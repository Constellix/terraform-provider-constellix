package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixTags() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixTagsCreate,
		Update: resourceConstellixTagsUpdate,
		Read:   resourceConstellixTagsRead,
		Delete: resourceConstellixTagsDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConstellixTagsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	model := models.Tags{}

	if name, ok := d.GetOk("name"); ok {
		model.Name = name.(string)
	}

	resp, err := client.Save(model, "v2/tags")
	if err != nil {
		return err
	}

	bodybtes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybtes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)

	idstruct := data["successTags"].([]interface{})[0]
	var idStruct map[string]interface{}

	idStruct = idstruct.(map[string]interface{})
	log.Println("Struct for Hinfo :", idStruct["id"])

	d.SetId(fmt.Sprintf("%.0f", idStruct["id"]))
	return resourceConstellixTagsRead(d, m)
}

func resourceConstellixTagsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	model := models.Tags{}

	if d.HasChange("name") {
		model.Name = d.Get("name").(string)
	}

	dn := d.Id()

	_, err := client.UpdatebyID(model, "v2/tags/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixTagsRead(d, m)
}

func resourceConstellixTagsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()
	resp, err := client.GetbyId("v2/tags/" + dn)
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybytes)

	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)
	d.Set("id", data["id"])
	d.Set("name", data["name"])
	return nil
}

func resourceConstellixTagsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v2/tags/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
