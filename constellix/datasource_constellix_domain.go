package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixDomain() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixDomainRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"has_gtd_regions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"has_geoip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"nameserver_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"soa": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_nameserver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"refresh": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"retry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"expire": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"negcache": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceConstellixDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/domains")
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
	for _, data := range data {
		tp = data.(map[string]interface{})
		if name == tp["name"].(string) {
			flag = true
			recsoa := tp["soa"].(map[string]interface{})

			soaSet := make(map[string]interface{})
			soaSet["primary_nameserver"] = recsoa["primaryNameserver"]
			soaSet["ttl"] = fmt.Sprintf("%v", recsoa["ttl"])
			soaSet["email"] = recsoa["email"]
			soaSet["refresh"] = fmt.Sprintf("%v", recsoa["refresh"])
			soaSet["expire"] = fmt.Sprintf("%v", recsoa["expire"])
			soaSet["retry"] = fmt.Sprintf("%v", recsoa["retry"])
			soaSet["negcache"] = fmt.Sprintf("%v", recsoa["negCache"])

			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("soa", soaSet)
			d.Set("typeid", tp["typeId"])
			d.Set("has_geoip", tp["hasGeoIP"])
			d.Set("has_gtd_regions", tp["hasGtdRegions"])
			d.Set("nameserver_group", tp["nameserverGroup"])
			d.Set("nameservers", tp["nameservers"])
			d.Set("note", tp["note"])
			d.Set("version", tp["version"])
			d.Set("status", tp["status"])
			d.Set("tags", tp["tags"])
		}
	}

	if flag != true {
		return fmt.Errorf("Domain of specified name is not available")
	}
	return nil
}
