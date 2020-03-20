package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixSPF() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixSPFRead,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
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

			"roundrobin": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceConstellixSPFRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/" + stid + "/" + domainid + "/records/spf/")
	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)

	var data []interface{}
	json.Unmarshal([]byte(bodystring), &data)

	var tp map[string]interface{}
	var flag bool
	for _, val := range data {
		tp = val.(map[string]interface{})
		if name == tp["name"].(string) {
			flag = true

			d.SetId(fmt.Sprintf("%v", (tp["id"])))
			d.Set("name", tp["name"])
			d.Set("ttl", tp["ttl"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])
			d.Set("parentid", tp["parentId"])
			d.Set("parent", tp["parent"])
			d.Set("source", tp["source"])
			resrr := (tp["roundRobin"]).([]interface{})
			mapListRR := make([]interface{}, 0, 1)
			length := len(resrr)
			for _, val := range resrr {
				tpMap := make(map[string]interface{})
				inner := val.(map[string]interface{})
				tpMap["mailbox"] = fmt.Sprintf("%v", inner["mailbox"])
				tpMap["txt"] = fmt.Sprintf("%v", inner["txt"])
				if length > 1 {
					tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
				}
				mapListRR = append(mapListRR, tpMap)
			}

			d.Set("roundrobin", mapListRR)
		}
	}

	if flag != true {
		return fmt.Errorf("SPF record of specified name is not available")
	}
	return nil
}
