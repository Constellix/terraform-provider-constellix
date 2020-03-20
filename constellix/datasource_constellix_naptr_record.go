package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixNAPTR() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixNAPTRRead,

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
						"order": {
							Type:     schema.TypeString,
							Required: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Required: true,
						},
						"flags": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
						},
						"regular_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"replacement": {
							Type:     schema.TypeString,
							Required: true,
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

func datasourceConstellixNAPTRRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	name1 := d.Get("name").(string)
	domainID := d.Get("domain_id").(string)
	sid := d.Get("source_type").(string)

	resp, err := constellixClient.GetbyId("v1/" + sid + "/" + domainID + "/records/naptr")
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
			resrr := tp["roundRobin"].([]interface{})
			rrlist := make([]interface{}, 0, 1)
			for _, valrrf := range resrr {
				map1 := make(map[string]interface{})
				val1 := valrrf.(map[string]interface{})
				map1["order"] = fmt.Sprintf("%v", val1["order"])
				map1["preference"] = fmt.Sprintf("%v", val1["preference"])
				map1["flags"] = fmt.Sprintf("%v", val1["flags"])
				map1["service"] = fmt.Sprintf("%v", val1["service"])
				map1["regular_expression"] = fmt.Sprintf("%v", val1["regularExpression"])
				map1["replacement"] = fmt.Sprintf("%v", val1["replacement"])
				map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])

				rrlist = append(rrlist, map1)
			}

			d.Set("roundrobin", rrlist)
			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("ttl", tp["ttl"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])
			d.Set("parentid", tp["parentId"])
			d.Set("parent", tp["parent"])
			d.Set("source", tp["source"])

		}
	}

	if flag == false {
		return (fmt.Errorf("Pointer record with name:%v is not present", name1))
	}
	return nil

}
