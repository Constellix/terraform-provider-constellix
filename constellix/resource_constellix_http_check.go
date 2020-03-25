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

func resourceConstellixHTTPCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixHTTPCheckCreate,
		Read:   resourceConstellixHTTPCheckRead,
		Update: resourceConstellixHTTPCheckUpdate,
		Delete: resourceConstellixHTTPCheckDelete,

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

			"ip_version": &schema.Schema{
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

			"check_sites": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceConstellixHTTPCheckCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	httpcheckAttr := models.HttpcheckAttr{}

	if name, ok := d.GetOk("name"); ok {
		httpcheckAttr.Name = name.(string)
	}

	if host, ok := d.GetOk("host"); ok {
		httpcheckAttr.Host = host.(string)
	}

	if ipv, ok := d.GetOk("ip_version"); ok {
		httpcheckAttr.Ipversion = ipv.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		httpcheckAttr.Port = port.(int)
	}

	if ptype, ok := d.GetOk("protocol_type"); ok {
		httpcheckAttr.ProtoType = ptype.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		httpcheckAttr.Checksites = checksites.([]interface{})
	}

	resp, err := client.Save(httpcheckAttr, "https://api.sonar.constellix.com/rest/api/http")
	if err != nil {
		return nil
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
	return resourceConstellixHTTPCheckRead(d, m)
}

func resourceConstellixHTTPCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	httpcheckAttr := models.HttpcheckAttr{}

	if name, ok := d.GetOk("name"); ok {
		httpcheckAttr.Name = name.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		httpcheckAttr.Port = port.(int)
	}

	if ptype, ok := d.GetOk("protocol_type"); ok {
		httpcheckAttr.ProtoType = ptype.(string)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		httpcheckAttr.Checksites = checksites.([]interface{})
	}

	dn := d.Id()
	_, err := client.UpdatebyID(httpcheckAttr, "https://api.sonar.constellix.com/rest/api/http/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixHTTPCheckRead(d, m)
}

func resourceConstellixHTTPCheckRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	resp, err := client.GetbyId("https://api.sonar.constellix.com/rest/api/http/" + dn)
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
	d.Set("ip_version", data["ipVersion"])
	d.Set("port", data["port"])
	d.Set("check_sites", data["checkSites"])
	return nil
}

func resourceConstellixHTTPCheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("https://api.sonar.constellix.com/rest/api/http/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
