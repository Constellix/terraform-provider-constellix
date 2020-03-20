package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"has_gtd_regions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"has_geoip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixTemplateRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/templates")
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)

	var data []interface{}
	json.Unmarshal([]byte(bodyString), &data)

	var tp map[string]interface{}
	var flag bool
	for _, val := range data {
		tp = val.(map[string]interface{})
		if name == tp["name"].(string) {
			flag = true
			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("domain", tp["Domain"])
			d.Set("name", tp["name"])
			d.Set("has_geoip", tp["hasGeoIP"])
			d.Set("has_gtd_regions", tp["hasGtdRegions"])
			d.Set("version", tp["version"])
		}
	}
	if flag != true {
		return fmt.Errorf("Template of specified name is not available")
	}
	return nil
}
