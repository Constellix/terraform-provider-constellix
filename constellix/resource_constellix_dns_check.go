package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixDNSCheck() *schema.Resource {
	return &schema.Resource{
		Create:        resourceConstellixDNSCheckCreate,
		Read:          resourceConstellixDNSCheckRead,
		Update:        resourceConstellixDNSCheckUpdate,
		Delete:        resourceConstellixDNSCheckDelete,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"protocol_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"resolver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"check_sites": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceConstellixDNSCheckCreate(d *schema.ResourceData, m interface{}) error {

	constellixConnect := m.(*client.Client)

	dnsAttr := models.DNSAttributes{}

	if name, ok := d.GetOk("name"); ok {
		dnsAttr.Name = name.(string)
	}

	if host, ok := d.GetOk("host"); ok {
		dnsAttr.Host = host.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		dnsAttr.Port = port.(int)
	}

	if ptype, ok := d.GetOk("protocol_type"); ok {
		dnsAttr.ProtocolType = ptype.(string)
	}

	if fqdn, ok := d.GetOk("fqdn"); ok {
		dnsAttr.FQDN = fqdn.(string)
	}

	if resolver, ok := d.GetOk("resolver"); ok {
		dnsAttr.Resolver = resolver.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		dnsAttr.CheckSites = checksites.([]interface{})
	}

	resp, err := constellixConnect.Save(dnsAttr, "https://api.sonar.constellix.com/rest/api/dns")
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	var location string
	var flag bool
	for k, v := range resp.Header {
		if k == "Location" {
			location = string(v[0])
			flag = true
		}
	}
	if flag == false {
		return fmt.Errorf("response contains empty location value")
	}

	locArr := strings.Split(location, "/")
	d.SetId(locArr[len(locArr)-1])

	return resourceConstellixDNSCheckRead(d, m)
}

func resourceConstellixDNSCheckRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	dnsid := d.Id()
	resp, err := constellixClient.GetbyId("https://api.sonar.constellix.com/rest/api/dns/" + dnsid)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)

	var data map[string]interface{}

	json.Unmarshal([]byte(bodystring), &data)
	d.Set("id", data["id"])
	d.Set("name", data["name"])
	d.Set("host", data["host"])
	d.Set("protocol_type", data["protocolType"])
	d.Set("port", data["port"])
	d.Set("fqdn", data["fqdn"])
	d.Set("resolver", data["resolver"])
	d.Set("check_sites", data["checkSites"])
	return nil
}

func resourceConstellixDNSCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	dnsAttr := models.DNSAttributes{}

	if name, ok := d.GetOk("name"); ok {
		dnsAttr.Name = name.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		dnsAttr.Port = port.(int)
	}

	if value, ok := d.GetOk("fqdn"); ok {
		dnsAttr.FQDN = value.(string)
	}

	if value, ok := d.GetOk("resolver"); ok {
		dnsAttr.Resolver = value.(string)
	}

	if ptype, ok := d.GetOk("protocol_type"); ok {
		dnsAttr.ProtocolType = ptype.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		dnsAttr.CheckSites = checksites.([]interface{})
	}

	dn := d.Id()
	resp, err := client.UpdatebyID(dnsAttr, "https://api.sonar.constellix.com/rest/api/dns/"+dn)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return resourceConstellixDNSCheckRead(d, m)
}

func resourceConstellixDNSCheckDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	dnsid := d.Id()

	err := constellixConnect.DeletebyId("https://api.sonar.constellix.com/rest/api/dns/" + dnsid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
