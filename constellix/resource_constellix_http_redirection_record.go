package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixHTTPRedirection() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixHTTPRedirectionCreate,
		Read:   resourceConstellixHTTPRedirectionRead,
		Update: resourceConstellixHTTPRedirectionUpdate,
		Delete: resourceConstellixHTTPRedirectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceConstellixHTTPRedirectionImport,
		},

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
				Optional: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
				Required: true,
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
				Required: true,
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

func resourceConstellixHTTPRedirectionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	constellixClient := m.(*client.Client)
	params := strings.Split(d.Id(), ":")
	resp, err := constellixClient.GetbyId("v1/" + params[0] + "/" + params[1] + "/records/httpredirection/" + params[2])
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
	d.Set("ttl", data["ttl"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	d.Set("parentid", data["parentId"])
	d.Set("parent", data["parent"])
	d.Set("source", data["source"])
	d.Set("title", data["title"])
	d.Set("keywords", data["keywords"])
	d.Set("description", data["description"])
	d.Set("url", data["url"])
	d.Set("hardlink_flag", data["hardlinkflag"])
	d.Set("redirect_type_id", data["redirectTypeId"])
	d.Set("domain_id", params[1])
	d.Set("source_type", params[0])
	log.Printf("[DEBUG] %s finished import", d.Id())
	return []*schema.ResourceData{d}, nil
}

func resourceConstellixHTTPRedirectionCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	httpAttr := models.HTTPRedirectionAttributes{}
	if name, ok := d.GetOk("name"); ok {
		httpAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		httpAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		httpAttr.NoAnswer = noans.(bool)
	}

	if title, ok := d.GetOk("title"); ok {
		httpAttr.Title = title.(string)
	}

	if kwrds, ok := d.GetOk("keywords"); ok {
		httpAttr.Keywords = kwrds.(string)
	}

	if desc, ok := d.GetOk("description"); ok {
		httpAttr.Description = desc.(string)
	}

	if note, ok := d.GetOk("note"); ok {
		httpAttr.Note = note.(string)
	}

	if gtd, ok := d.GetOk("gtd_region"); ok {
		httpAttr.GtdRegion = gtd.(int)
	}

	if url, ok := d.GetOk("url"); ok {
		httpAttr.URL = url.(string)
	}
	if tp, ok := d.GetOk("type"); ok {
		httpAttr.Type = tp.(string)
	}
	if hlinkflag, ok := d.GetOk("hardlink_flag"); ok {
		httpAttr.Hardlinkflag = hlinkflag.(bool)
	}
	if redtpid, ok := d.GetOk("redirect_type_id"); ok {
		httpAttr.RedirectTypeID = redtpid.(int)
	}
	if pid, ok := d.GetOk("parentid"); ok {
		httpAttr.ParentID = pid.(int)
	}
	if p, ok := d.GetOk("parent"); ok {
		httpAttr.Parent = p.(string)
	}
	if sc, ok := d.GetOk("source"); ok {
		httpAttr.Source = sc.(string)
	}

	id := d.Get("domain_id").(string)
	stype := d.Get("source_type").(string)

	resp, err := constellixConnect.Save(httpAttr, "v1/"+stype+"/"+id+"/records/httpredirection")
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
	json.Unmarshal([]byte(bodystring[1:len(bodystring)-1]), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixHTTPRedirectionRead(d, m)

}

func resourceConstellixHTTPRedirectionUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	httpAttr := models.HTTPRedirectionAttributes{}

	httpAttr.Name = d.Get("name").(string)

	httpAttr.TTL = d.Get("ttl").(int)

	httpAttr.NoAnswer = d.Get("noanswer").(bool)

	if _, ok := d.GetOk("title"); ok {
		httpAttr.Title = d.Get("title").(string)
	}

	if _, ok := d.GetOk("keywords"); ok {
		httpAttr.Keywords = d.Get("keywords").(string)
	}

	if _, ok := d.GetOk("description"); ok {
		httpAttr.Description = d.Get("description").(string)
	}

	if _, ok := d.GetOk("note"); ok {
		httpAttr.Note = d.Get("note").(string)
	}

	if gtd, ok := d.GetOk("gtd_region"); ok {
		httpAttr.GtdRegion = gtd.(int)
	}

	if url, ok := d.GetOk("url"); ok {
		httpAttr.URL = url.(string)
	}

	if _, ok := d.GetOk("type"); ok {
		httpAttr.Type = d.Get("type").(string)
	}

	if hlinkflag, ok := d.GetOk("hardlink_flag"); ok {
		httpAttr.Hardlinkflag = hlinkflag.(bool)
	}
	if redtpid, ok := d.GetOk("redirect_type_id"); ok {
		httpAttr.RedirectTypeID = redtpid.(int)
	}
	if pid, ok := d.GetOk("parentid"); ok {
		httpAttr.ParentID = pid.(int)
	}

	if _, ok := d.GetOk("parent"); ok {
		httpAttr.Parent = d.Get("parent").(string)
	}

	if _, ok := d.GetOk("source"); ok {
		httpAttr.Source = d.Get("source").(string)
	}

	domainID := d.Get("domain_id").(string)
	stype := d.Get("source_type").(string)
	httpid := d.Id()
	_, err := client.UpdatebyID(httpAttr, "v1/"+stype+"/"+domainID+"/records/httpredirection/"+httpid)
	if err != nil {
		return err
	}
	return resourceConstellixHTTPRedirectionRead(d, m)
}

func resourceConstellixHTTPRedirectionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	httpid := d.Id()

	resp, err := client.GetbyId("v1/" + stid + "/" + domainID + "/records/httpredirection/" + httpid)
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
	d.Set("ttl", data["ttl"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	d.Set("parentid", data["parentId"])
	d.Set("parent", data["parent"])
	d.Set("source", data["source"])
	d.Set("title", data["title"])
	d.Set("keywords", data["keywords"])
	d.Set("description", data["description"])
	d.Set("url", data["url"])
	d.Set("hardlink_flag", data["hardlinkflag"])
	d.Set("redirect_type_id", data["redirectTypeId"])
	log.Println("Data  : ", data)

	return nil
}

func resourceConstellixHTTPRedirectionDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	dn := d.Id()

	err := constellixConnect.DeletebyId("v1/" + stid + "/" + domainID + "/records/httpredirection/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
