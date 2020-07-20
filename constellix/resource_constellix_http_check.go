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
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"search_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expected_status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

	if interval, ok := d.GetOk("interval"); ok {
		httpcheckAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		httpcheckAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		httpcheckAttr.VerificationPolicy = verification_policy.(string)
	}

	if fqdn, ok := d.GetOk("fqdn"); ok {
		httpcheckAttr.FQDN = fqdn.(string)
	}
	if path, ok := d.GetOk("path"); ok {
		httpcheckAttr.PATH = path.(string)
	}
	if search_string, ok := d.GetOk("search_string"); ok {
		httpcheckAttr.SearchString = search_string.(string)
	}
	if expected_status_code, ok := d.GetOk("expected_status_code"); ok {
		httpcheckAttr.ExpectedStatus = expected_status_code.(int)
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
	if interval, ok := d.GetOk("interval"); ok {
		httpcheckAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		httpcheckAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		httpcheckAttr.VerificationPolicy = verification_policy.(string)
	}

	if fqdn, ok := d.GetOk("fqdn"); ok {
		httpcheckAttr.FQDN = fqdn.(string)
	}
	if path, ok := d.GetOk("path"); ok {
		httpcheckAttr.PATH = path.(string)
	}
	if search_string, ok := d.GetOk("search_string"); ok {
		httpcheckAttr.SearchString = search_string.(string)
	}
	if expected_status_code, ok := d.GetOk("expected_status_code"); ok {
		httpcheckAttr.ExpectedStatus = expected_status_code.(int)
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
	d.Set("interval", data["interval"])
	d.Set("interval_policy", data["monitorIntervalPolicy"])
	d.Set("verification_policy", data["verificationPolicy"])
	d.Set("fqdn", data["fqdn"])
	d.Set("path", data["path"])
	d.Set("search_string", data["searchString"])
	d.Set("expected_status_code", data["expectedStatusCode"])
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
