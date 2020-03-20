package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceConstellixContactList() *schema.Resource {
	return &schema.Resource{

		Read: datasourceConstellixContactListRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func datasourceConstellixContactListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	name := d.Get("name").(string)
	resp, err := client.GetbyId("v2/contactLists")
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
			d.SetId(fmt.Sprintf("%.0f", tp["id"]))
			d.Set("name", tp["name"])
			d.Set("email_addresses", tp["emailAddresses"])
		}
	}

	if flag == false {
		return (fmt.Errorf("Cert record with name:%v,is not present", name))
	}
	return nil

}
