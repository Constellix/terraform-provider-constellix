package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixGeoProximity() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixGeoProximityCreate,
		Update: resourceConstellixGeoProximityUpdate,
		Read:   resourceConstellixGeoProximityRead,
		Delete: resourceConstellixGeoProximityDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == new {
						return true
					}
					s, err := strconv.ParseFloat(new, 64)
					if err != nil {
						return false
					}
					val := fmt.Sprintf("%.2f", s)
					if val == old {
						return true
					}
					return false
				},
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == new {
						return true
					}
					s, err := strconv.ParseFloat(new, 64)
					if err != nil {
						return false
					}
					val := fmt.Sprintf("%.2f", s)
					if val == old {
						return true
					}
					return false
				},
			},
		},
	}
}

func resourceConstellixGeoProximityCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	GeoProximityAttr := models.GeoProximityAttributes{}

	if nm, ok := d.GetOk("name"); ok {
		GeoProximityAttr.Name = nm.(string)
	}

	if ct, ok := d.GetOk("country"); ok {
		GeoProximityAttr.Country = ct.(string)
	}

	if rsp, ok := d.GetOk("region"); ok {
		GeoProximityAttr.RegionStateProvince = rsp.(string)
	}

	if city, ok := d.GetOk("city"); ok {
		GeoProximityAttr.City = city.(int)
	}

	if lat, ok := d.GetOk("latitude"); ok {
		GeoProximityAttr.Latitude = lat.(float64)
	}

	if long, ok := d.GetOk("longitude"); ok {
		GeoProximityAttr.Longitude = long.(float64)
	}

	resp, err := client.Save(GeoProximityAttr, "v1/geoProximities/")
	if err != nil {
		return err
	}

	//Managing response and extracting id of resource
	bodybtes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybtes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring[1:len(bodystring)-1]), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixGeoProximityRead(d, m)
}

func resourceConstellixGeoProximityUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	geoproximityAttr := models.GeoProximityAttributes{}

	geoproximityAttr.Name = d.Get("name").(string)

	geoproximityAttr.Country = d.Get("country").(string)

	geoproximityAttr.RegionStateProvince = d.Get("region").(string)

	geoproximityAttr.City = d.Get("city").(int)

	geoproximityAttr.Latitude = d.Get("latitude").(float64)

	geoproximityAttr.Longitude = d.Get("longitude").(float64)

	geoproximityid := d.Id()
	_, err := client.UpdatebyID(geoproximityAttr, "v1/geoProximities/"+geoproximityid)
	if err != nil {
		return err
	}
	return resourceConstellixGeoProximityRead(d, m)
}

func resourceConstellixGeoProximityRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	geoproximityid := d.Id()

	resp, err := client.GetbyId("v1/geoProximities/" + geoproximityid)
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
	d.Set("country", data["country"])
	d.Set("region", data["region"])
	d.Set("latitude", data["latitude"])
	d.Set("longitude", data["longitude"])

	return nil
}

func resourceConstellixGeoProximityDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v1/geoProximities/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
