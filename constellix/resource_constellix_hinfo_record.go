package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixHinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixHinfoCreate,
		Update: resourceConstellixHinfoUpdate,
		Read:   resourceConstellixHinfoRead,
		Delete: resourceConstellixHinfoDelete,

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
						"cpu": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"os": &schema.Schema{
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

func resourceConstellixHinfoCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	hinfoAttr := models.HinfoAttributes{}

	if name, ok := d.GetOk("name"); ok {
		hinfoAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		hinfoAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		hinfoAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		hinfoAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		hinfoAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		hinfoAttr.Type = tp.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["cpu"] = fmt.Sprintf("%v", inner["cpu"])
			tpMap["os"] = fmt.Sprintf("%v", inner["os"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		hinfoAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	source := d.Get("source_type").(string)

	resp, err := client.Save(hinfoAttr, "v1/"+source+"/"+id+"/records/hinfo")
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
	return resourceConstellixHinfoRead(d, m)
}

func resourceConstellixHinfoUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	hinfoAttr := models.HinfoAttributes{}

	if name, ok := d.GetOk("name"); ok {
		hinfoAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		hinfoAttr.TTL = ttl.(int)
	}

	if _, ok := d.GetOk("noanswer"); ok {
		hinfoAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if _, ok := d.GetOk("note"); ok {
		hinfoAttr.Note = d.Get("note").(string)
	}

	if _, ok := d.GetOk("gtd_region"); ok {
		hinfoAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if _, ok := d.GetOk("type"); ok {
		hinfoAttr.Type = d.Get("type").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["cpu"] = fmt.Sprintf("%v", inner["cpu"])
			tpMap["os"] = fmt.Sprintf("%v", inner["os"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		hinfoAttr.RoundRobin = mapListRR
	}

	domainid := d.Get("domain_id").(string)
	hinfoid := d.Id()
	source := d.Get("source_type").(string)

	_, err := client.UpdatebyID(hinfoAttr, "v1/"+source+"/"+domainid+"/records/hinfo/"+hinfoid)
	if err != nil {
		return err
	}
	return resourceConstellixHinfoRead(d, m)
}

func resourceConstellixHinfoRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	hinfoid := d.Id()
	source := d.Get("source_type").(string)

	resp, err := client.GetbyId("v1/" + source + "/" + domainid + "/records/hinfo/" + hinfoid)
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
		tpMap["cpu"] = fmt.Sprintf("%v", inner["cpu"])
		tpMap["os"] = fmt.Sprintf("%v", inner["os"])
		tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixHinfoDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	dn := d.Id()
	source := d.Get("source_type").(string)

	err := client.DeletebyId("v1/" + source + "/" + domainid + "/records/hinfo/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
