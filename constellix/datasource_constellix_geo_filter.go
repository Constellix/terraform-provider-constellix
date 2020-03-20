package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixIPFilter() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixIPFilterRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"geoipcontinents": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"geoipregions": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"geoipcountries": &schema.Schema{
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
			"regions": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"continentcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"countrycode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"regioncode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"filterruleslimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixIPFilterRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	name1 := d.Get("name").(string)

	resp, err := constellixClient.GetbyId("v1/geoFilters")
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data []interface{}
	var temp map[string]interface{}
	var flag bool
	json.Unmarshal([]byte(bodyString), &data)
	for _, val := range data {
		temp = val.(map[string]interface{})
		if temp["name"].(string) == name1 {
			flag = true
			resrr := temp["regions"].([]interface{})

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

			ipaddr := temp["ipAddresses"].([]interface{})

			tp1 := ipaddr[0].(map[string]interface{})
			tp2 := ipaddr[1].(map[string]interface{})

			ipv4s := tp1["ipv4Addresses"].([]interface{})
			ipv6s := tp2["ipv6Addresses"].([]interface{})

			ipv4List := make([]string, 0, 1)
			ipv6List := make([]string, 0, 1)

			for _, val := range ipv4s {
				temp := val.(map[string]interface{})["ipv4"]
				ipv4List = append(ipv4List, temp.(string))
			}

			for _, val := range ipv6s {
				temp := val.(map[string]interface{})["ipv6"]
				ipv6List = append(ipv6List, temp.(string))
			}

			log.Println("Hllllllllll :", ipv4List)
			log.Println("HLLLLLLL :", ipv6List)

			d.SetId(fmt.Sprintf("%v", temp["id"]))
			d.Set("ipv4", ipv4List)
			d.Set("ipv6", ipv6List)
			d.Set("geoipregions", geoipregionsList)
			d.Set("name", temp["name"])
			d.Set("geoipcontinents", temp["geoipContinents"])
			d.Set("geoipcountries", temp["geoipCountries"])
			d.Set("asn", temp["asn"])
			d.Set("filterruleslimit", temp["filterRulesLimit"])
		}
	}

	if flag == false {
		return fmt.Errorf("The ipfilter with the name %v is not present", name1)
	}

	return nil
}
