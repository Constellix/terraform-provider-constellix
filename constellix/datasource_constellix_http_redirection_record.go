package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixHTTPRedirection() *schema.Resource {
	return &schema.Resource{

		Read: datasourceConstellixHTTPRedirectionRead,

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
				ForceNew: true,
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

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"keywords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
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

			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"hardlink_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"redirect_type_id": &schema.Schema{
				Type:     schema.TypeInt,
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

func datasourceConstellixHTTPRedirectionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/" + stid + "/" + domainID + "/records/httpredirection")
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
	var flag bool
	var tp map[string]interface{}
	for _, val := range data {
		tp = val.(map[string]interface{})
		if tp["name"] == name {
			flag = true
			d.Set("name", tp["name"])
			d.SetId(fmt.Sprintf("%v", tp["id"]))

			d.Set("ttl", tp["ttl"])
			d.Set("noanswer", tp["noAnswer"])
			d.Set("note", tp["note"])
			d.Set("gtd_region", tp["gtdRegion"])
			d.Set("type", tp["type"])
			d.Set("parentid", tp["parentId"])
			d.Set("parent", tp["parent"])
			d.Set("source", tp["source"])
			d.Set("title", tp["title"])
			d.Set("keywords", tp["keywords"])
			d.Set("description", tp["description"])
			d.Set("url", tp["url"])
			d.Set("hardlink_flag", tp["hardlinkFlag"])
			d.Set("redirect_type_id", tp["redirectTypeId"])
		}
	}

	if flag == false {
		return (fmt.Errorf("Cert record with name:%v,is not present", name))
	}
	return nil
}
