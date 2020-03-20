package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixVanityNameserver() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConstellixVanityNameserverRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"is_public": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"nameserver_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"nameserver_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nameserver_list_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixVanityNameserverRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)

	resp, err := client.GetbyId("v1/vanityNameservers")
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
			d.SetId(fmt.Sprintf("%v", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("is_default", tp["isDefault"])
			d.Set("is_public", tp["isPublic"])
			d.Set("nameserver_group", tp["nameserverGroup"])
			d.Set("nameserver_group_name", tp["nameserverGroupName"])
			d.Set("nameserver_list_string", tp["nameserversListString"])
		}
	}

	if flag != true {
		return fmt.Errorf("Vanity Nameserver of specified name is not available")
	}
	return nil
}
