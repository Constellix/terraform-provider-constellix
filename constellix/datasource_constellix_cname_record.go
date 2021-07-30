package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixCNameRecord() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixCNameRecordRead,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"geo_location": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"geo_ip_user_region": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"geo_ip_proximity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"geo_ip_failover": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},

			"record_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"noanswer": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"gtd_region": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"record_failover_values": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sort_order": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},

			"record_failover_failover_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"record_failover_disable_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func datasourceConstellixCNameRecordRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	name1 := d.Get("name").(string)
	domainID := d.Get("domain_id").(string)
	sid := d.Get("source_type").(string)

	resp, err := constellixClient.GetbyId("v1/" + sid + "/" + domainID + "/records/cname/")
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
		if tp["name"].(string) == name1 {
			flag = true

			geoloc1 := tp["geolocation"]
			log.Println("GEOLOC VALUE INSIDE READ :", geoloc1)

			geoLocMap := make(map[string]interface{})
			if geoloc1 != nil {
				geoloc := geoloc1.(map[string]interface{})
				if geoloc["geoipFilter"] != nil {
					geoLocMap["geo_ip_user_region"] = fmt.Sprintf("%v", geoloc["geoipFilter"])
				}
				if geoloc["drop"] != nil {
					geoLocMap["drop"] = fmt.Sprintf("%v", geoloc["drop"])
				}
				if geoloc["geoipFailover"] != nil {
					geoLocMap["geo_ip_failover"] = fmt.Sprintf("%v", geoloc["geoipFailover"])
				}
				if geoloc["geoipProximity"] != nil {
					geoLocMap["geo_ip_proximity"] = fmt.Sprintf("%v", geoloc["geoipProximity"])
				}
				d.Set("geo_location", geoLocMap)
			} else {
				d.Set("geo_location", geoLocMap)
			}

			rcdf := tp["recordFailover"]
			rcdfset := make(map[string]interface{})
			rcdflist := make([]interface{}, 0, 1)
			if rcdf != nil {
				rcdf1 := rcdf.(map[string]interface{})
				rcdfset["record_failover_failover_type"] = fmt.Sprintf("%v", rcdf1["failoverType"])
				rcdfset["record_failover_disable_flag"] = fmt.Sprintf("%v", rcdf1["disabled"])

				rcdfvalues := rcdf1["values"].([]interface{})

				for _, valrcdf := range rcdfvalues {
					map1 := make(map[string]interface{})
					val1 := valrcdf.(map[string]interface{})
					map1["value"] = fmt.Sprintf("%v", val1["value"])
					map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
					map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
					rcdflist = append(rcdflist, map1)
				}
			}

			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("host", tp["host"])
			d.Set("name", tp["name"])
			d.Set("ttl", tp["ttl"])
			d.Set("record_option", tp["recordOption"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])
			d.Set("pools", tp["pools"])
			d.Set("contact_ids", tp["contactId"])
			d.Set("record_failover_values", rcdflist)
			d.Set("record_failover_failover_type", rcdfset["record_failover_failover_type"])
			d.Set("record_failover_disable_flag", rcdfset["record_failover_disable_flag"])
		}
	}

	if flag == false {
		return (fmt.Errorf("Pointer record with name:%v is not present", name1))
	}
	return nil
}
