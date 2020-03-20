package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixAnamerecord() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixAnamerecordRead,
		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
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

			"roundrobin": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disable_flag": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
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
						"checkidrcdf": &schema.Schema{
							Type:     schema.TypeString,
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
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"record_failover_disable_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixAnamerecordRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/aname")
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

			rrlist := make([]interface{}, 0, 1)
			if tp["roundRobin"] != nil {
				recroundrobin := tp["roundRobin"].([]interface{})
				rrlist := make([]interface{}, 0, 1)
				for _, valrrf := range recroundrobin {
					map1 := make(map[string]interface{})
					val1 := valrrf.(map[string]interface{})
					map1["value"] = fmt.Sprintf("%v", val1["value"])
					map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
					rrlist = append(rrlist, map1)
				}
			}

			rcdf := tp["recordFailover"]
			rcdf1 := rcdf.(map[string]interface{})
			rcdfset := make(map[string]interface{})
			rcdfset["record_failover_failover_type"] = fmt.Sprintf("%v", rcdf1["failoverType"])
			rcdfset["record_failover_disable_flag"] = fmt.Sprintf("%v", rcdf1["disabled"])

			rcdfvalues := rcdf1["values"].([]interface{})

			rcdflist := make([]interface{}, 0, 1)
			for _, valrcdf := range rcdfvalues {
				map1 := make(map[string]interface{})
				val1 := valrcdf.(map[string]interface{})
				map1["value"] = fmt.Sprintf("%v", val1["value"])
				map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
				map1["disable_flag"] = fmt.Sprintf("%v", val1["disabled"])
				rcdflist = append(rcdflist, map1)
			}

			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("ttl", tp["ttl"])
			d.Set("record_option", tp["recordOption"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])
			d.Set("contact_ids", tp["contactids"])
			d.Set("roundrobin", rrlist)
			d.Set("record_failover_values", rcdflist)
			d.Set("record_failover_failover_type", rcdfset["record_failover_failover_type"])
			d.Set("record_failover_disable_flag", rcdfset["record_failover_disable_flag"])
		}
	}
	if flag != true {
		return fmt.Errorf("ANAME record of specified name is not available")
	}
	return nil
}
