package constellix

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixCert() *schema.Resource {
	return &schema.Resource{

		Read: datasourceConstellixCertRead,

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
						"certificate_type": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"keytag": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"algorithm": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func datasourceConstellixCertRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)

	resp, err := client.GetbyId("v1/" + stid + "/" + domainID + "/records/cert")
	if err != nil {
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)
	var flag bool
	var tp map[string]interface{}
	var data []interface{}
	json.Unmarshal([]byte(bodystring), &data)
	for _, val := range data {
		tp = val.(map[string]interface{})
		if tp["name"] == name {
			flag = true
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
			resrr := (tp["roundRobin"]).([]interface{})
			mapListRR := make([]interface{}, 0, 1)
			for _, val := range resrr {
				tpMap := make(map[string]interface{})
				inner := val.(map[string]interface{})
				tpMap["certificate_type"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["certificateType"]))
				tpMap["keytag"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["keyTag"]))
				tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
				tpMap["algorithm"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["algorithm"]))
				sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", inner["certificate"])))
				tpMap["certificate"] = sEnc
				mapListRR = append(mapListRR, tpMap)
			}

			d.Set("roundrobin", mapListRR)

		}
	}

	if flag == false {
		return (fmt.Errorf("Cert record with name:%v,is not present", name))
	}
	return nil

}
