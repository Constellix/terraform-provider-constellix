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

func resourceConstellixNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixNSCreate,
		Read:   resourceConstellixNSRead,
		Update: resourceConstellixNSUpdate,
		Delete: resourceConstellixNSDelete,

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
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"disable_flag": {
							Type:     schema.TypeString,
							Optional: true,
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

func resourceConstellixNSCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	nsAttr := models.NSAttributes{}

	if name, ok := d.GetOk("name"); ok {
		nsAttr.Name = name.(string)
	}
	if ttl, ok := d.GetOk("ttl"); ok {
		nsAttr.Ttl = ttl.(int)
	}
	if noAnswer, ok := d.GetOk("noanswer"); ok {
		nsAttr.NoAnswer = noAnswer.(bool)
	}
	if note, ok := d.GetOk("note"); ok {
		nsAttr.Note = note.(string)
	}
	if gtdRegion, ok := d.GetOk("gtd_region"); ok {
		nsAttr.GtdRegion = gtdRegion.(int)
	}
	if type1, ok := d.GetOk("type"); ok {
		nsAttr.Type = type1.(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			maplistrr = append(maplistrr, map1)
		}
		nsAttr.RoundRobin = maplistrr
	}
	resp, err := constellixConnect.Save(nsAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/ns")
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodystring := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring[1:len(bodystring)-1]), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixNSRead(d, m)
}

func resourceConstellixNSDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	dn := d.Id()
	err := constellixConnect.DeletebyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/ns" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}

func resourceConstellixNSUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	nsAttr := models.NSAttributes{}

	if name, ok := d.GetOk("name"); ok {
		nsAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		nsAttr.Ttl = ttl.(int)
	}

	if d.HasChange("noanswer") {
		nsAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		nsAttr.Note = note.(string)
	}

	if d.HasChange("gtd_region") {
		nsAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if d.HasChange("type") {
		nsAttr.Type = d.Get("type").(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			maplistrr = append(maplistrr, map1)
		}
		nsAttr.RoundRobin = maplistrr
	}

	nsRecord := d.Id()
	_, err := constellixClient.UpdatebyID(nsAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/ns/"+nsRecord)
	if err != nil {
		return err
	}
	return resourceConstellixNSRead(d, m)
}

func resourceConstellixNSRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	nsID := d.Id()

	resp, err := constellixClient.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/ns/" + nsID)

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
		map1["value"] = fmt.Sprintf("%v", val1["value"])
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
