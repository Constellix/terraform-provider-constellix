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

func resourceConstellixContactList() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixContactListCreate,
		Read:   resourceConstellixContactListRead,
		Update: resourceConstellixcontactListUpdate,
		Delete: resourceConstellixContactListDelete,

		Importer: &schema.ResourceImporter{
			State: resourceConstellixContactListImport,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceConstellixContactListImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	constellixClient := m.(*client.Client)
	cid := d.Id()
	resp, err := constellixClient.GetbyId("v2/contactLists/" + cid)
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil, err
		}
		return nil, err
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodystring := string(bodybytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	d.Set("name", data["name"])
	d.Set("email_addresses", data["emailAddresses"])
	log.Printf("[DEBUG] %s finished import", d.Id())
	return []*schema.ResourceData{d}, nil
}

func resourceConstellixContactListCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	contactlistAttr := models.ContactListAttributes{}
	if name, ok := d.GetOk("name"); ok {
		contactlistAttr.Name = name.(string)
	}

	if contactemails, ok := d.GetOk("email_addresses"); ok {
		contactemailsList := toListOfString(contactemails)
		contactlistAttr.EmailAddresses = contactemailsList
	}

	resp, err := client.Save(contactlistAttr, "v2/contactLists")
	if err != nil {
		return err
	}

	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodystring := string(bodybytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring), &data)
	idstruct := data["successContactLists"].([]interface{})[0]
	var idStruct map[string]interface{}

	idStruct = idstruct.(map[string]interface{})

	d.SetId(fmt.Sprintf("%.0f", idStruct["id"]))
	return resourceConstellixContactListRead(d, m)
}

func resourceConstellixcontactListUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	contactlistAttr := models.ContactListAttributes{}

	contactlistAttr.Name = d.Get("name").(string)

	if _, ok := d.GetOk("email_addresses"); ok {
		contactemaillist := toListOfString(d.Get("email_addresses"))
		contactlistAttr.EmailAddresses = contactemaillist

	}

	cid := d.Id()
	_, err := client.UpdatebyID(contactlistAttr, "v2/contactLists/"+cid)
	if err != nil {
		return err
	}

	return resourceConstellixContactListRead(d, m)
}

func resourceConstellixContactListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	cid := d.Id()
	resp, err := client.GetbyId("v2/contactLists/" + cid)
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

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	d.Set("name", data["name"])
	d.Set("email_addresses", data["emailAddresses"])
	return nil
}

func resourceConstellixContactListDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	cid := d.Id()

	err := client.DeletebyId("v2/contactLists/" + cid)
	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
