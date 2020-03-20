package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixARecordPool() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceARecordPoolRead,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"num_return": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"min_available_failover": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failed_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"disable_flag_1": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"values": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceARecordPoolRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/pools/A/")
	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)
	var flag bool
	var data []interface{}
	var tp map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)
	for _, val := range data {
		tp = val.(map[string]interface{})
		if tp["name"] == name {
			flag = true

			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("num_return", tp["numReturn"])
			d.Set("min_available_failover", tp["minAvailableFailover"])
			d.Set("note", tp["note"])
			d.Set("version", tp["version"])
			d.Set("failed_flag", tp["failedFlag"])
			d.Set("disable_flag_1", tp["disableFlag"])
			resrr := (tp["values"]).([]interface{})
			mapListRR := make([]interface{}, 0, 1)
			for _, val := range resrr {
				tpMap := make(map[string]interface{})
				inner := val.(map[string]interface{})
				tpMap["value"] = fmt.Sprintf("%v", inner["value"])
				tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
				tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
				tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])
				tpMap["check_id"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["checkId"]))

				mapListRR = append(mapListRR, tpMap)
			}
			d.Set("values", mapListRR)
		}
	}

	if flag == false {
		return fmt.Errorf("ARecord Pool named:%v is not available", name)
	}
	return nil

}
