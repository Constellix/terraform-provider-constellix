package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixNS() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixNSRead,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"noanswer": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"roundrobin": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"disable_flag": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
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
		},
	}
}

func datasourceConstellixNSRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	name1 := d.Get("name").(string)
	sid := d.Get("source_type").(string)
	domainID := d.Get("domain_id").(string)

	resp, err := constellixClient.GetbyId("v1/" + sid + "/" + domainID + "/records/ns")
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
			resrr := temp["roundRobin"].([]interface{})
			rrlist := make([]interface{}, 0, 1)
			for _, valrrf := range resrr {
				map1 := make(map[string]interface{})
				val1 := valrrf.(map[string]interface{})
				map1["value"] = fmt.Sprintf("%v", val1["value"])
				map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])

				rrlist = append(rrlist, map1)
			}

			d.SetId(fmt.Sprintf("%v", temp["id"]))
			d.Set("roundrobin", rrlist)
			d.Set("name", temp["name"])
			d.Set("ttl", temp["ttl"])
			d.Set("noanswer", temp["noAnswer"])
			d.Set("note", temp["note"])
			d.Set("gtd_region", temp["gtdRegion"])
			d.Set("type", temp["type"])
			d.Set("parentid", temp["parentId"])
			d.Set("parent", temp["parent"])
			d.Set("source", temp["source"])
		}
	}

	if flag == false {
		return (fmt.Errorf("Pointer record with name:%v is not present", name1))
	}
	return nil

}
