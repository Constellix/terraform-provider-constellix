package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixRP() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixRPCreate,
		Update: resourceConstellixRPUpdate,
		Read:   resourceConstellixRPRead,
		Delete: resourceConstellixRPDelete,

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
						"mailbox": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"txt": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceConstellixRPCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	rpAttr := models.RPAttributes{}

	if name, ok := d.GetOk("name"); ok {
		rpAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		rpAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		rpAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		rpAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		rpAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		rpAttr.Type = tp.(string)
	}

	if pid, ok := d.GetOk("parentid"); ok {
		rpAttr.ParentId = pid.(int)
	}

	if p, ok := d.GetOk("parent"); ok {
		rpAttr.Parent = p.(string)
	}

	if source, ok := d.GetOk("source"); ok {
		rpAttr.Source = source.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["mailbox"] = fmt.Sprintf("%v", inner["mailbox"])
			tpMap["txt"] = fmt.Sprintf("%v", inner["txt"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		rpAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	source := d.Get("source_type").(string)

	resp, err := client.Save(rpAttr, "v1/"+source+"/"+id+"/records/rp")
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
	return resourceConstellixRPRead(d, m)
}

func resourceConstellixRPUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	rpAttr := models.RPAttributes{}

	if name, ok := d.GetOk("name"); ok {
		rpAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		rpAttr.TTL = ttl.(int)
	}

	if d.HasChange("noanswer") {
		rpAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if d.HasChange("note") {
		rpAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("gtd_region") {
		rpAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if d.HasChange("type") {
		rpAttr.Type = d.Get("type").(string)
	}

	if d.HasChange("parentid") {
		rpAttr.ParentId = d.Get("parentid").(int)
	}

	if d.HasChange("parent") {
		rpAttr.Parent = d.Get("parent").(string)
	}

	if d.HasChange("source") {
		rpAttr.Source = d.Get("source").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["mailbox"] = fmt.Sprintf("%v", inner["mailbox"])
			tpMap["txt"] = fmt.Sprintf("%v", inner["txt"])
			tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
			mapListRR = append(mapListRR, tpMap)
		}
		rpAttr.RoundRobin = mapListRR
	}

	domainID := d.Get("domain_id").(string)
	rpid := d.Id()
	source := d.Get("source_type").(string)
	_, err := client.UpdatebyID(rpAttr, "v1/"+source+"/"+domainID+"/records/rp/"+rpid)
	if err != nil {
		return err
	}
	return resourceConstellixRPRead(d, m)
}

func resourceConstellixRPRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	rpid := d.Id()
	source := d.Get("source_type").(string)

	resp, err := client.GetbyId("v1/" + source + "/" + domainID + "/records/rp/" + rpid)
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
	d.Set("id", data["id"])
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
	length := len(resrr)
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["mailbox"] = fmt.Sprintf("%v", inner["mailbox"])
		tpMap["txt"] = fmt.Sprintf("%v", inner["txt"])
		if length > 1 {
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disableFlag"])
		}
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixRPDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	dn := d.Id()
	source := d.Get("source_type").(string)

	err := client.DeletebyId("v1/" + source + "/" + domainID + "/records/rp/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
