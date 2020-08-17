package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixVanityNameserver() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixVanityNameserverCreate,
		Update: resourceConstellixVanityNameserverUpdate,
		Read:   resourceConstellixVanityNameserverRead,
		Delete: resourceConstellixVanityNameserverDelete,

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
				Required: true,
			},

			"nameserver_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nameserver_list_string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConstellixVanityNameserverCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	model := models.VanityNameserverAttributes{}

	if name, ok := d.GetOk("name"); ok {
		model.Name = name.(string)
	}

	if isdef, ok := d.GetOk("is_default"); ok {
		model.IsDefault = isdef.(bool)
	}

	if ispub, ok := d.GetOk("is_public"); ok {
		model.IsDefault = ispub.(bool)
	}

	if group, ok := d.GetOk("nameserver_group"); ok {
		model.NameserverGroup = group.(int)
	}

	if gname, ok := d.GetOk("nameserver_group_name"); ok {
		model.NameserverGroupName = gname.(string)
	}

	if liststr, ok := d.GetOk("nameserver_list_string"); ok {
		model.NameserverListString = liststr.(string)
	}

	resp, err := client.Save(model, "v1/vanityNameservers/")
	if err != nil {
		return err
	}

	//Managing response and extracting id of resource
	bodybtes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybtes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixVanityNameserverRead(d, m)
}

func resourceConstellixVanityNameserverUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	model := models.VanityNameserverAttributes{}

	if name, ok := d.GetOk("name"); ok {
		model.Name = name.(string)
	}

	if _, ok := d.GetOk("is_default"); ok {
		model.IsDefault = d.Get("is_default").(bool)
	}

	if _, ok := d.GetOk("is_public"); ok {
		model.IsDefault = d.Get("is_public").(bool)
	}

	if group, ok := d.GetOk("nameserver_group"); ok {
		model.NameserverGroup = group.(int)
	}

	if _, ok := d.GetOk("nameserver_group_name"); ok {
		model.NameserverGroupName = d.Get("nameserver_group_name").(string)
	}

	if liststr, ok := d.GetOk("nameserver_list_string"); ok {
		model.NameserverListString = liststr.(string)
	}

	dn := d.Id()

	_, err := client.UpdatebyID(model, "v1/vanityNameservers/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixVanityNameserverRead(d, m)
}

func resourceConstellixVanityNameserverRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	resp, err := client.GetbyId("v1/vanityNameservers/" + dn)
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
	d.Set("is_default", data["isDefault"])
	d.Set("is_public", data["isPublic"])
	d.Set("nameserver_group", data["nameserverGroup"])
	d.Set("nameserver_group_name", data["nameserverGroupName"])
	d.Set("nameserver_list_string", data["nameserversListString"])
	return nil
}

func resourceConstellixVanityNameserverDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v1/vanityNameservers/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
