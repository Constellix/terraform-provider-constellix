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

func resourceConstellixTCPCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixTCPCheckCreate,
		Update: resourceConstellixTCPCheckUpdate,
		Read:   resourceConstellixTCPCheckRead,
		Delete: resourceConstellixTCPCheckDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"check_sites": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Required: true,
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
			"string_to_send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"string_to_receive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConstellixTCPCheckCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*client.Client)

	tcpcheckAttr := models.TCPCheckAttributes{}

	if name, ok := d.GetOk("name"); ok {
		tcpcheckAttr.Name = name.(string)
	}

	if host, ok := d.GetOk("host"); ok {
		tcpcheckAttr.Host = host.(string)
	}

	if ipv, ok := d.GetOk("ip_version"); ok {
		tcpcheckAttr.Ipversion = ipv.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		tcpcheckAttr.Port = port.(int)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		tcpcheckAttr.Checksites = checksites.([]interface{})
	}

	if interval, ok := d.GetOk("interval"); ok {
		tcpcheckAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		tcpcheckAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		tcpcheckAttr.VerificationPolicy = verification_policy.(string)
	}

	if string_to_send, ok := d.GetOk("string_to_send"); ok {
		tcpcheckAttr.StringToSend = string_to_send.(string)
	}
	if string_to_receive, ok := d.GetOk("string_to_receive"); ok {
		tcpcheckAttr.StringToReceive = string_to_receive.(string)
	}

	resp, err := client.Save(tcpcheckAttr, "https://api.sonar.constellix.com/rest/api/tcp")
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
		return fmt.Errorf("Response contains empty location value")
	}

	locArr := strings.Split(location, "/")
	d.SetId(locArr[len(locArr)-1])
	return resourceConstellixTCPCheckRead(d, m)
}

func resourceConstellixTCPCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	tcpcheckAttr := models.TCPCheckAttributes{}

	if name, ok := d.GetOk("name"); ok {
		tcpcheckAttr.Name = name.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		tcpcheckAttr.Port = port.(int)
	}

	if checksites, ok := d.GetOk("check_sites"); ok {
		tcpcheckAttr.Checksites = checksites.([]interface{})
	}

	if interval, ok := d.GetOk("interval"); ok {
		tcpcheckAttr.Interval = interval.(string)
	}

	if interval_policy, ok := d.GetOk("interval_policy"); ok {
		tcpcheckAttr.IntervalPolicy = interval_policy.(string)
	}

	if verification_policy, ok := d.GetOk("verification_policy"); ok {
		tcpcheckAttr.VerificationPolicy = verification_policy.(string)
	}

	if string_to_send, ok := d.GetOk("string_to_send"); ok {
		tcpcheckAttr.StringToSend = string_to_send.(string)
	}
	if string_to_receive, ok := d.GetOk("string_to_receive"); ok {
		tcpcheckAttr.StringToReceive = string_to_receive.(string)
	}

	dn := d.Id()
	_, err := client.UpdatebyID(tcpcheckAttr, "https://api.sonar.constellix.com/rest/api/tcp/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixTCPCheckRead(d, m)
}

func resourceConstellixTCPCheckRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	resp, err := client.GetbyId("https://api.sonar.constellix.com/rest/api/tcp/" + dn)
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
	d.Set("ip_version", data["ipVersion"])
	d.Set("port", data["port"])
	d.Set("check_sites", data["checkSites"])
	d.Set("interval", data["interval"])
	d.Set("interval_policy", data["monitorIntervalPolicy"])
	d.Set("string_to_send", data["stringToSend"])
	d.Set("string_to_receive", data["stringToReceive"])
	return nil
}

func resourceConstellixTCPCheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("https://api.sonar.constellix.com/rest/api/tcp/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
