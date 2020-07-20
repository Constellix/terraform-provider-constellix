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

			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"resolver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"check_sites": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verification_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expected_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

	if fqdn, ok := d.GetOk("fqdn"); ok {
		dnsAttr.FQDN = fqdn.(string)
	}

	if resolver, ok := d.GetOk("resolver"); ok {
		dnsAttr.Resolver = resolver.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		dnsAttr.CheckSites = checksites.([]interface{})
	}

	if interval, ok := d.GetOk("interval"); ok {
		dnsAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		dnsAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		dnsAttr.VerificationPolicy = verification_policy.(string)
	}

	if expected_response, ok := d.GetOk("expected_response"); ok {
		dnsAttr.ExpectedResponse = expected_response.(string)
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
	d.Set("fqdn", data["fqdn"])
	d.Set("resolver", data["resolver"])
	d.Set("check_sites", data["checkSites"])
	d.Set("interval", data["interval"])
	d.Set("interval_policy", data["monitorIntervalPolicy"])
	d.Set("verification_policy", data["verificationPolicy"])
	d.Set("expected_response", data["expectedResponse"])
	return nil
}

func resourceConstellixDNSCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	dnsAttr := models.DNSAttributes{}

	if name, ok := d.GetOk("name"); ok {
		dnsAttr.Name = name.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		dnsAttr.CheckSites = checksites.([]interface{})
	}

	if interval, ok := d.GetOk("interval"); ok {
		dnsAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		dnsAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		dnsAttr.VerificationPolicy = verification_policy.(string)
	}

	if expected_response, ok := d.GetOk("expected_response"); ok {
		dnsAttr.ExpectedResponse = expected_response.(string)
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
