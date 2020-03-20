package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixTxt() *schema.Resource {
	return &schema.Resource{

		Read: datasourceConstellixTxtRead,

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

			"parentid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"source": &schema.Schema{
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
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceConstellixTxtRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	name1 := d.Get("name").(string)

	resp, err := client.GetbyId("v1/" + stid + "/" + domainID + "/records/txt")
	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)

	var data []interface{}
	var flag bool
	var temp map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)
	for _, val := range data {
		temp = val.(map[string]interface{})
		if temp["name"].(string) == name1 {
			flag = true
			d.SetId(fmt.Sprintf("%v", temp["id"]))
			d.Set("name", temp["name"])
			d.Set("ttl", temp["ttl"])
			d.Set("noanswer", temp["noAnswer"])
			d.Set("note", temp["note"])
			d.Set("gtd_region", temp["gtdRegion"])
			d.Set("type", temp["type"])
			d.Set("parentid", temp["parentId"])
			d.Set("parent", temp["parent"])
			d.Set("source", temp["source"])

			resrr := (temp["roundRobin"]).([]interface{})
			mapListRR := make([]interface{}, 0, 1)
			for _, val := range resrr {
				tpMap := make(map[string]interface{})
				inner := val.(map[string]interface{})
				tpMap["value"] = fmt.Sprintf("%v", inner["value"])
				tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
				mapListRR = append(mapListRR, tpMap)
			}

			d.Set("roundrobin", mapListRR)

		}
	}

	if flag == false {
		return (fmt.Errorf("Pointer record with name:%v is not present", name1))
	}
	return nil

}
