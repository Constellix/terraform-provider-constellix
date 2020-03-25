package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixDNSCheck() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixDNSCheckRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"protocol_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"resolver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"check_sites": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixDNSCheckRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := constellixClient.GetbyId("https://api.sonar.constellix.com/rest/api/dns/")
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)
	var data []interface{}
	var flag bool
	var tp map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	for _, val := range data {
		tp = val.(map[string]interface{})
		if tp["name"].(string) == name {
			flag = true
			d.Set("id", tp["id"])
			d.Set("name", tp["name"])
			d.Set("fqdn", tp["fqdn"])
			d.Set("host", tp["host"])
			d.Set("protocol_type", tp["protocolType"])
			d.Set("port", tp["port"])
			d.Set("resolver", tp["resolver"])
			d.Set("check_sites", tp["checkSites"])
		}
	}
	if flag == false {
		return (fmt.Errorf("DNS Check with specified name is not present"))
	}
	return nil
}
