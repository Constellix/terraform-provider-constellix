package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixNAPTR() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixNAPTRCreate,
		Read:   resourceConstellixNAPTRRead,
		Update: resourceConstellixNAPTRUpdate,
		Delete: resourceConstellixNAPTRDelete,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"roundrobin": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"order": {
							Type:     schema.TypeString,
							Required: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Required: true,
						},
						"flags": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
						},
						"regular_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"replacement": {
							Type:     schema.TypeString,
							Required: true,
						},
						"disable_flag": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Required: true,
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
		},
	}
}

func resourceConstellixNAPTRCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	naptrAttr := models.NAPTRAttributes{}

	if name, ok := d.GetOk("name"); ok {
		naptrAttr.Name = name.(string)
	}
	if ttl, ok := d.GetOk("ttl"); ok {
		naptrAttr.Ttl = ttl.(int)
	}
	if noAnswer, ok := d.GetOk("noanswer"); ok {
		naptrAttr.NoAnswer = noAnswer.(bool)
	}
	if note, ok := d.GetOk("note"); ok {
		naptrAttr.Note = note.(string)
	}
	if gtdRegion, ok := d.GetOk("gtd_region"); ok {
		naptrAttr.GtdRegion = gtdRegion.(int)
	}
	if type1, ok := d.GetOk("type"); ok {
		naptrAttr.Type = type1.(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["order"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["order"]))
			map1["preference"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["preference"]))
			map1["flags"] = fmt.Sprintf("%v", inner["flags"])
			map1["service"] = fmt.Sprintf("%v", inner["service"])
			map1["regularExpression"] = fmt.Sprintf("%v", inner["regular_expression"])
			map1["replacement"] = fmt.Sprintf("%v", inner["replacement"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			maplistrr = append(maplistrr, map1)
		}
		naptrAttr.RoundRobin = maplistrr
	}

	resp, err := constellixConnect.Save(naptrAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/naptr")
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString[1:len(bodyString)-1]), &data)
	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixNAPTRRead(d, m)
}

func resourceConstellixNAPTRDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	dn := d.Id()
	err := constellixConnect.DeletebyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/naptr" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}

func resourceConstellixNAPTRUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	naptrAttr := models.NAPTRAttributes{}

	if name, ok := d.GetOk("name"); ok {
		naptrAttr.Name = name.(string)
	}
	if ttl, ok := d.GetOk("ttl"); ok {
		naptrAttr.Ttl = ttl.(int)
	}

	if d.HasChange("noanswer") {

		naptrAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		naptrAttr.Note = note.(string)
	}
	if d.HasChange("gtd_region") {

		naptrAttr.GtdRegion = d.Get("gtd_region").(int)
	}
	if d.HasChange("type") {

		naptrAttr.Type = d.Get("type").(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["order"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["order"]))
			map1["preference"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["preference"]))
			map1["flags"] = fmt.Sprintf("%v", inner["flags"])
			map1["service"] = fmt.Sprintf("%v", inner["service"])
			map1["regularExpression"] = fmt.Sprintf("%v", inner["regular_expression"])
			map1["replacement"] = fmt.Sprintf("%v", inner["replacement"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			maplistrr = append(maplistrr, map1)
		}
		naptrAttr.RoundRobin = maplistrr
	}
	naptrRecord := d.Id()
	_, err := constellixClient.UpdatebyID(naptrAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/naptr/"+naptrRecord)
	if err != nil {
		return err
	}

	return resourceConstellixNAPTRRead(d, m)
}

func resourceConstellixNAPTRRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	naptrID := d.Id()

	resp, err := constellixClient.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/naptr/" + naptrID)
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	arecroundrobin := data["roundRobin"].([]interface{})
	rrlist := make([]interface{}, 0, 1)
	for _, valrrf := range arecroundrobin {
		map1 := make(map[string]interface{})
		val1 := valrrf.(map[string]interface{})
		map1["order"] = fmt.Sprintf("%v", val1["order"])
		map1["preference"] = fmt.Sprintf("%v", val1["preference"])
		map1["flags"] = fmt.Sprintf("%v", val1["flags"])
		map1["service"] = fmt.Sprintf("%v", val1["service"])
		map1["regular_expression"] = fmt.Sprintf("%v", val1["regularExpression"])
		map1["replacement"] = fmt.Sprintf("%v", val1["replacement"])
		map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])

		rrlist = append(rrlist, map1)
	}

	d.Set("id", data["id"])
	d.Set("roundrobin", rrlist)
	d.Set("name", data["name"])
	d.Set("ttl", data["ttl"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	return nil
}
