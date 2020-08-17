package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixCaa() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixCaaCreate,
		Update: resourceConstellixCaaUpdate,
		Read:   resourceConstellixCaaRead,
		Delete: resourceConstellixCaaDelete,

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
						"caa_provider_id": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"data": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"flag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixCaaCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	caaAttr := models.CaaAttributes{}

	if name, ok := d.GetOk("name"); ok {
		caaAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		caaAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		caaAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		caaAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		caaAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		caaAttr.Type = tp.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["caaProviderId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["caa_provider_id"]))
			tpMap["tag"] = fmt.Sprintf("%v", inner["tag"])
			tpMap["data"] = fmt.Sprintf("%v", inner["data"])
			tpMap["flag"] = fmt.Sprintf("%v", inner["flag"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		caaAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	source := d.Get("source_type").(string)

	resp, err := client.Save(caaAttr, "v1/"+source+"/"+id+"/records/caa")
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
	return resourceConstellixCaaRead(d, m)
}

func resourceConstellixCaaUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	caaAttr := models.CaaAttributes{}

	if name, ok := d.GetOk("name"); ok {
		caaAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		caaAttr.TTL = ttl.(int)
	}

	if _, ok := d.GetOk("noanswer"); ok {
		caaAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		caaAttr.Note = note.(string)
	}

	if _, ok := d.GetOk("gtd_region"); ok {
		caaAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if _, ok := d.GetOk("type"); ok {
		caaAttr.Type = d.Get("type").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["caaProviderId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["caa_provider_id"]))
			tpMap["tag"] = fmt.Sprintf("%v", inner["tag"])
			tpMap["data"] = fmt.Sprintf("%v", inner["data"])
			tpMap["flag"] = fmt.Sprintf("%v", inner["flag"])
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			mapListRR = append(mapListRR, tpMap)
		}
		caaAttr.RoundRobin = mapListRR
	}

	domainid := d.Get("domain_id").(string)
	caaid := d.Id()
	source := d.Get("source_type").(string)
	_, err := client.UpdatebyID(caaAttr, "v1/"+source+"/"+domainid+"/records/caa/"+caaid)
	if err != nil {
		return err
	}
	return resourceConstellixCaaRead(d, m)
}

func resourceConstellixCaaRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	caaid := d.Id()
	source := d.Get("source_type").(string)

	resp, err := client.GetbyId("v1/" + source + "/" + domainid + "/records/caa/" + caaid)
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
		log.Println("RR are : ", val)
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["caa_provider_id"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["caaProviderId"]))
		tpMap["tag"] = fmt.Sprintf("%v", inner["tag"])
		tpMap["data"] = fmt.Sprintf("%v", inner["data"])
		tpMap["flag"] = fmt.Sprintf("%v", inner["flag"])
		tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixCaaDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainid := d.Get("domain_id").(string)
	dn := d.Id()
	source := d.Get("source_type").(string)

	err := client.DeletebyId("v1/" + source + "/" + domainid + "/records/caa/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
