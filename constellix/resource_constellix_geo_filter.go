package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixIPFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixIPFilterCreate,
		Read:   resourceConstellixIPFilterRead,
		Update: resourceConstellixIPFilterUpdate,
		Delete: resourceConstellixIPFilterDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"geoip_continents": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"geoip_regions": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"geoip_countries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"asn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"ipv4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_rules_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConstellixIPFilterCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	ipfilterattr := models.IPFilterAttributes{}

	if name, ok := d.GetOk("name"); ok {
		ipfilterattr.Name = fmt.Sprintf("%v", name)
	}
	if geoipcontinents, ok := d.GetOk("geoip_continents"); ok {
		ipfilterattr.GeoIPContinents = toListOfString(geoipcontinents)
	}
	if geoipcountries, ok := d.GetOk("geoip_countries"); ok {
		ipfilterattr.GeoIPCountries = toListOfString(geoipcountries)
	}
	if asn, ok := d.GetOk("asn"); ok {
		ipfilterattr.Asn = toListOfInt(asn)
	}
	if geoipregions, ok := d.GetOk("geoip_regions"); ok {
		ipfilterattr.GeoIPRegions = toListOfString(geoipregions)
	}
	if filterruleslimit, ok := d.GetOk("filter_rules_limit"); ok {
		ipfilterattr.FilterRulesLimit = filterruleslimit.(int)
	}

	var count1, count2 int
	mainList := make([]interface{}, 0, 1)
	tp01 := make([]map[string]interface{}, 0, 1)
	inner1 := make(map[string]interface{}, 1)
	inner2 := make(map[string]interface{}, 1)
	if ipv4, ok := d.GetOk("ipv4"); ok {
		values := ipv4.([]interface{})
		count1 = 1
		for _, val := range values {
			temp := make(map[string]interface{}, 1)
			temp["ipv4"] = val.(string)
			tp01 = append(tp01, temp)
		}
		inner1["ipv4Addresses"] = tp01
	}

	tp02 := make([]map[string]interface{}, 0, 1)
	if ipv6, ok := d.GetOk("ipv6"); ok {
		values := ipv6.([]interface{})
		count2 = 1
		for _, val := range values {
			temp := make(map[string]interface{}, 1)
			temp["ipv6"] = val.(string)
			tp02 = append(tp02, temp)
		}
		inner2["ipv6Addresses"] = tp02
	}
	if count1 == 1 {
		mainList = append(mainList, inner1)
	}
	if count2 == 1 {
		mainList = append(mainList, inner2)
	}

	ipfilterattr.IPAddresses = mainList

	resp, err := constellixConnect.Save(ipfilterattr, "v1/geoFilters")
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
	return resourceConstellixIPFilterRead(d, m)
}

func resourceConstellixIPFilterRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	nsid := d.Id()

	resp, err := constellixClient.GetbyId("v1/geoFilters/" + nsid)
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

	resrr := data["regions"].([]interface{})

	geoipregionsList := make([]string, 0)
	for _, value := range resrr {
		tp := value.(map[string]interface{})
		str := fmt.Sprintf("%v", tp["countryCode"])
		str1 := fmt.Sprintf("%v", tp["regionCode"])
		if str != "" && str1 != "" {
			geoip := str + "/" + str1
			geoipregionsList = append(geoipregionsList, geoip)
		}
	}

	ipaddr1 := data["ipAddresses"]

	if ipaddr1 != nil {
		ipaddr := ipaddr1.([]interface{})
		for _, val := range ipaddr {

			tp := val.(map[string]interface{})
			if tp["ipv4Addresses"] != nil {
				ipv4s := tp["ipv4Addresses"].([]interface{})
				ipv4List := make([]string, 0, 1)
				for _, val := range ipv4s {
					temp := val.(map[string]interface{})["ipv4"]
					ipv4List = append(ipv4List, temp.(string))
				}
				d.Set("ipv4", ipv4List)

			} else {
				ipv6s := tp["ipv6Addresses"].([]interface{})
				ipv6List := make([]string, 0, 1)
				for _, val := range ipv6s {
					temp := val.(map[string]interface{})["ipv6"]
					ipv6List = append(ipv6List, temp.(string))
				}
				d.Set("ipv6", ipv6List)

			}
		}
	}

	d.Set("geoip_regions", geoipregionsList)
	d.Set("name", data["name"])
	d.Set("geoip_continents", data["geoipContinents"])
	d.Set("geoip_countries", data["geoipCountries"])
	d.Set("asn", data["asn"])
	d.Set("filter_rules_limit", data["filterRulesLimit"])

	return nil
}

func resourceConstellixIPFilterDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	dn := d.Id()
	err := constellixConnect.DeletebyId("v1/geoFilters/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}

func resourceConstellixIPFilterUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	ipfilterattr := models.IPFilterAttributes{}

	if name, ok := d.GetOk("name"); ok {
		ipfilterattr.Name = fmt.Sprintf("%v", name)
	}
	if geoipcontinents, ok := d.GetOk("geoip_continents"); ok {
		ipfilterattr.GeoIPContinents = toListOfString(geoipcontinents)
	}
	if geoipcountries, ok := d.GetOk("geoip_countries"); ok {
		ipfilterattr.GeoIPCountries = toListOfString(geoipcountries)
	}
	if geoipregions, ok := d.GetOk("geoip_regions"); ok {
		ipfilterattr.GeoIPRegions = toListOfString(geoipregions)
	}
	if asn, ok := d.GetOk("asn"); ok {
		ipfilterattr.Asn = toListOfInt(asn)
	}
	if _, ok := d.GetOk("filter_rules_limit"); ok {

		ipfilterattr.FilterRulesLimit = d.Get("filter_rules_limit").(int)
	}

	var count1, count2 int
	mainList := make([]interface{}, 0, 1)
	tp01 := make([]map[string]interface{}, 0, 1)
	inner1 := make(map[string]interface{}, 1)
	inner2 := make(map[string]interface{}, 1)
	if ipv4, ok := d.GetOk("ipv4"); ok {
		values := ipv4.([]interface{})
		count1 = 1
		for _, val := range values {
			temp := make(map[string]interface{}, 1)
			temp["ipv4"] = val.(string)
			tp01 = append(tp01, temp)
		}
		inner1["ipv4Addresses"] = tp01
	}

	tp02 := make([]map[string]interface{}, 0, 1)
	if ipv6, ok := d.GetOk("ipv6"); ok {
		values := ipv6.([]interface{})
		count2 = 1
		for _, val := range values {
			temp := make(map[string]interface{}, 1)
			temp["ipv6"] = val.(string)
			tp02 = append(tp02, temp)
		}
		inner2["ipv6Addresses"] = tp02
	}
	if count1 == 1 {
		mainList = append(mainList, inner1)
	}
	if count2 == 1 {
		mainList = append(mainList, inner2)
	}

	ipfilterattr.IPAddresses = mainList

	ipfilterattr.IPAddresses = mainList
	nsRecord := d.Id()
	_, err := constellixClient.UpdatebyID(ipfilterattr, "v1/geoFilters/"+nsRecord)
	if err != nil {
		return err
	}
	return resourceConstellixIPFilterRead(d, m)
}
