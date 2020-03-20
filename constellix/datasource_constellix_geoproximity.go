package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixGeoProximity() *schema.Resource {
	return &schema.Resource{

		Read: datasourceConstellixGeoProximityRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"longitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},

			"city": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"latitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixGeoProximityRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/geoProximities")
	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)

	var data []interface{}
	var tp map[string]interface{}
	var flag bool
	json.Unmarshal([]byte(bodystring), &data)

	for _, val := range data {
		tp = val.(map[string]interface{})
		if tp["name"] == name {
			flag = true
			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("country", tp["country"])
			d.Set("region", tp["region"])
			d.Set("latitude", tp["latitude"])
			d.Set("longitude", tp["longitude"])
		}
	}

	if flag == false {
		return (fmt.Errorf("Cert record with name:%v,is not present", name))
	}
	return nil
}
