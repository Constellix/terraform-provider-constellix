package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixTemplateCreate,
		Read:   resourceConstellixTemplateRead,
		Update: resourceConstellixTemplateUpdate,
		Delete: resourceConstellixTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceConstellixTemplateCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	templateAttr := models.TemplateAttributes{}

	if name, ok := d.GetOk("name"); ok {
		templateAttr.Name = toStringList(name)
	}
	if domain, ok := d.GetOk("domain"); ok {
		templateAttr.Domain = domain.(int)
	}

	if hasgtdRegion, ok := d.GetOk("has_gtd_regions"); ok {
		templateAttr.HasGtdRegions = hasgtdRegion.(bool)
	}
	if hasGeoIP, ok := d.GetOk("has_geoip"); ok {
		templateAttr.HasGeoIP = hasGeoIP.(bool)
	}
	if version, ok := d.GetOk("version"); ok {
		templateAttr.Version = version.(int)
	}

	resp, err := constellixConnect.Save(templateAttr, "v1/templates/")
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString[1:len(bodyString)-1]), &data)
	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixTemplateRead(d, m)
}

func resourceConstellixTemplateRead(d *schema.ResourceData, m interface{}) error {
	constellixclient := m.(*client.Client)
	dn := d.Id()
	resp, err := constellixclient.GetbyId("v1/templates/" + dn)
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	d.Set("domain", data["Domain"])
	d.Set("name", data["name"])
	d.Set("has_geoip", data["hasGeoIP"])
	d.Set("has_gtd_regions", data["hasGtdRegions"])
	d.Set("version", data["version"])

	return nil

}

func resourceConstellixTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	templateAttr := models.TemplateAttributes{}

	if domain, ok := d.GetOk("domain"); ok {
		templateAttr.Domain = domain.(int)
	}

	if hasgtdRegion, ok := d.GetOk("has_gtd_regions"); ok {
		templateAttr.HasGtdRegions = hasgtdRegion.(bool)
	}
	if hasGeoIP, ok := d.GetOk("has_geoip"); ok {
		templateAttr.HasGeoIP = hasGeoIP.(bool)
	}
	if version, ok := d.GetOk("version"); ok {
		templateAttr.Version = version.(int)
	}

	dn := d.Id()

	_, err := constellixClient.UpdatebyID(templateAttr, "v1/templates/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixTemplateRead(d, m)
}

func resourceConstellixTemplateDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	dn := d.Id()

	err := constellixConnect.DeletebyId("v1/templates/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}
