package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixMX() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixMXCreate,
		Update: resourceConstellixMXUpdate,
		Read:   resourceConstellixMXRead,
		Delete: resourceConstellixMXDelete,

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

			"roundrobin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"level": &schema.Schema{
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

func resourceConstellixMXCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	mxAttr := models.MXAttributes{}

	if name, ok := d.GetOk("name"); ok {
		mxAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		mxAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		mxAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		mxAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		mxAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		mxAttr.Type = tp.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"] = fmt.Sprintf("%v", inner["value"])
			tpMap["level"] = fmt.Sprintf("%v", inner["level"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		mxAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	source := d.Get("source_type").(string)

	resp, err := client.Save(mxAttr, "v1/"+source+"/"+id+"/records/mx")
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
	return resourceConstellixMXRead(d, m)
}

func resourceConstellixMXUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	mxAttr := models.MXAttributes{}

	if name, ok := d.GetOk("name"); ok {
		mxAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		mxAttr.TTL = ttl.(int)
	}

	if d.HasChange("noanswer") {
		mxAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if d.HasChange("note") {
		mxAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("gtd_region") {
		mxAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if d.HasChange("type") {
		mxAttr.Type = d.Get("type").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"] = fmt.Sprintf("%v", inner["value"])
			tpMap["level"] = fmt.Sprintf("%v", inner["level"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		mxAttr.RoundRobin = mapListRR
	}

	domainid := d.Get("domain_id").(string)
	mxid := d.Id()
	source := d.Get("source_type").(string)

	_, err := client.UpdatebyID(mxAttr, "v1/"+source+"/"+domainid+"/records/mx/"+mxid)
	if err != nil {
		return err
	}
	return resourceConstellixMXRead(d, m)
}

func resourceConstellixMXRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	mxid := d.Id()
	source := d.Get("source_type").(string)

	resp, err := client.GetbyId("v1/" + source + "/" + domainid + "/records/mx/" + mxid)
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
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"] = fmt.Sprintf("%v", inner["value"])
		tpMap["level"] = fmt.Sprintf("%v", inner["level"])
		tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixMXDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	mxid := d.Id()
	source := d.Get("source_type").(string)

	err := constellixConnect.DeletebyId("v1/" + source + "/" + domainid + "/records/mx/" + mxid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
