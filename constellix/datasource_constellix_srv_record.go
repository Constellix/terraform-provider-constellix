package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixSRV() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixSRVRead,

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

			"roundrobin": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"weight": {
							Type:     schema.TypeInt,
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

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixSRVRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/srv")
	if err != nil {
		return err
	}
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
			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("ttl", tp["ttl"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])

			srvroundrobin := tp["roundRobin"].([]interface{})
			rrlist := make([]interface{}, 0, 1)
			for _, valrrf := range srvroundrobin {
				map1 := make(map[string]interface{})
				val1 := valrrf.(map[string]interface{})
				map1["value"] = fmt.Sprintf("%v", val1["value"])
				map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
				map1["port"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["port"]))
				map1["priority"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["priority"]))
				map1["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["weight"]))
				rrlist = append(rrlist, map1)
			}
			d.Set("roundrobin", rrlist)
		}
	}

	if flag != true {
		return fmt.Errorf("SRV record of specified name is not available")
	}
	return nil
}
