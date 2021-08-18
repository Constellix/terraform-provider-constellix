package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixPtr() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixPtrCreate,
		Update: resourceConstellixPtrUpdate,
		Read:   resourceConstellixPtrRead,
		Delete: resourceConstellixPtrDelete,

		Importer: &schema.ResourceImporter{
			State: resourceConstellixPtrImport,
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
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixPtrImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	constellixClient := m.(*client.Client)
	params := strings.Split(d.Id(), ":")
	resp, err := constellixClient.GetbyId("v1/" + params[0] + "/" + params[1] + "/records/ptr/" + params[2])
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
	resrr := (data["roundRobin"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["value"]))
		tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disableFlag"])
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	d.Set("domain_id", params[1])
	d.Set("source_type", params[0])
	log.Printf("[DEBUG] %s finished import", d.Id())
	return []*schema.ResourceData{d}, nil
}
func resourceConstellixPtrCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	PtrAttr := models.PtrAttributes{}

	if nm, ok := d.GetOk("name"); ok {
		PtrAttr.Name = nm.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		PtrAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		PtrAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		PtrAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		PtrAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		PtrAttr.Type = tp.(string)
	}

	if pid, ok := d.GetOk("parentid"); ok {
		PtrAttr.ParentID = pid.(int)
	}

	if p, ok := d.GetOk("parent"); ok {
		PtrAttr.Parent = p.(string)
	}

	if source, ok := d.GetOk("source"); ok {
		PtrAttr.Source = source.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["value"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])

			mapListRR = append(mapListRR, tpMap)
		}
		PtrAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)

	resp, err := client.Save(PtrAttr, "v1/"+stid+"/"+id+"/records/ptr")
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
	return resourceConstellixPtrRead(d, m)
}

func resourceConstellixPtrUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	PtrAttr := models.PtrAttributes{}

	if nm, ok := d.GetOk("name"); ok {
		PtrAttr.Name = nm.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		PtrAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		PtrAttr.NoAnswer = noans.(bool)
	}

	if _, ok := d.GetOk("note"); ok {
		PtrAttr.Note = d.Get("note").(string)
	}

	if _, ok := d.GetOk("gtd_region"); ok {
		PtrAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if _, ok := d.GetOk("type"); ok {
		PtrAttr.Type = d.Get("type").(string)
	}

	if _, ok := d.GetOk("parentid"); ok {
		PtrAttr.ParentID = d.Get("parentid").(int)
	}

	if _, ok := d.GetOk("parent"); ok {
		PtrAttr.Parent = d.Get("parent").(string)
	}

	if _, ok := d.GetOk("source"); ok {
		PtrAttr.Source = d.Get("source").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["value"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		PtrAttr.RoundRobin = mapListRR
	}

	domainid := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	ptrid := d.Id()
	_, err := client.UpdatebyID(PtrAttr, "v1/"+stid+"/"+domainid+"/records/ptr/"+ptrid)
	if err != nil {
		return err
	}
	return resourceConstellixPtrRead(d, m)
}

func resourceConstellixPtrRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	ptrid := d.Id()

	resp, err := client.GetbyId("v1/" + stid + "/" + domainid + "/records/ptr/" + ptrid)
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
	resrr := (data["roundRobin"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["value"]))
		tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disableFlag"])
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixPtrDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	dn := d.Id()

	err := constellixConnect.DeletebyId("v1/" + stid + "/" + domainid + "/records/ptr/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
