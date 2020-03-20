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

func resourceConstellixANAMERecord() *schema.Resource {
	return &schema.Resource{
		Create:        resourceConstellixANAMERecordCreate,
		Read:          resourceConstellixANAMERecordRead,
		Update:        resourceConstellixANAMERecordUpdate,
		Delete:        resourceConstellixANAMERecordDelete,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"record_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

			"contact_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
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
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
				Required: true,
			},

			"record_failover_values": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"check_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sort_order": {
							Type:     schema.TypeString,
							Required: true,
						},
						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},
			"record_failover_failover_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"record_failover_disable_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConstellixANAMERecordCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	anameAttr := models.AnameAttributes{}

	if name, ok := d.GetOk("name"); ok {
		anameAttr.Name = name.(string)
	}

	if value, ok := d.GetOk("ttl"); ok {
		anameAttr.TTL = value.(int)
	}

	if recordoption, ok := d.GetOk("record_option"); ok {
		anameAttr.RecordOption = recordoption.(string)
	}

	if noanswer, ok := d.GetOk("noanswer"); ok {
		anameAttr.NoAnswer = noanswer.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		anameAttr.Note = note.(string)
	}

	if gtdregion, ok := d.GetOk("gtd_region"); ok {
		anameAttr.GtdRegion = gtdregion.(int)
	}

	if types, ok := d.GetOk("type"); ok {
		anameAttr.Type = types.(string)
	}

	if contactids, ok := d.GetOk("contact_ids"); ok {
		contactidList := toStringList(contactids.([]interface{}))
		var intlist []int
		intlist = make([]int, len(contactidList))
		for _, i := range contactidList {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			intlist = append(intlist, j)
		}

		anameAttr.ContactIDs = intlist
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
		anameAttr.RoundRobin = maplistrr
	}

	var valuesrcdf *models.ValuesAname
	var rcdfaname *models.RCDFAname
	valueslist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("record_failover_values"); ok {
		rcdfaname = &models.RCDFAname{}
		valuesrcdf = &models.ValuesAname{}
		tp := value.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["checkid"] = fmt.Sprintf("%v", inner["check_id"])
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			valueslist = append(valueslist, map1)
		}
		rcdfaname.Values = valueslist
	}

	if failtype, ok := d.GetOk("record_failover_failover_type"); ok {
		rcdfaname.FailoverTypeAname, _ = strconv.Atoi(fmt.Sprintf("%v", failtype))
	}

	if disflag, ok := d.GetOk("record_failover_disable_flag"); ok {
		rcdfaname.DisableFlagAname, _ = strconv.ParseBool(fmt.Sprintf("%v", disflag))
	}

	if valuesrcdf != nil {
		rcdfaname.Values = valueslist
		anameAttr.RecordFailoverAname = rcdfaname
	} else {
		anameAttr.RecordFailoverAname = nil
	}

	resp, err := constellixConnect.Save(anameAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/aname")

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

	return resourceConstellixANAMERecordRead(d, m)
}

func resourceConstellixANAMERecordRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	anameid := d.Id()

	resp, err := constellixClient.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/aname/" + anameid)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(bodyString), &data)
	if err != nil {
		return err
	}

	arecroundrobin := data["roundRobin"].([]interface{})
	rrlist := make([]interface{}, 0, 1)
	for _, valrrf := range arecroundrobin {
		map1 := make(map[string]interface{})
		val1 := valrrf.(map[string]interface{})
		map1["value"] = fmt.Sprintf("%v", val1["value"])
		map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
		rrlist = append(rrlist, map1)
	}

	rcdf := data["recordFailover"]
	rcdfSet := make(map[string]interface{})
	rcdflist := make([]interface{}, 0, 1)
	if rcdf != nil {
		rcdf1 := rcdf.(map[string]interface{})
		rcdfSet["record_failover_failover_type"] = fmt.Sprintf("%v", rcdf1["failoverType"])
		rcdfSet["record_failover_disable_flag"] = fmt.Sprintf("%v", rcdf1["disabled"])

		rcdfValues := rcdf1["values"].([]interface{})

		for _, valrcdf := range rcdfValues {
			map1 := make(map[string]interface{})
			val1 := valrcdf.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", val1["value"])
			map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
			map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
			rcdflist = append(rcdflist, map1)
		}
	}

	d.Set("id", data["id"])
	d.Set("name", data["name"])
	d.Set("ttl", data["ttl"])
	d.Set("record_option", data["recordOption"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	d.Set("contact_ids", data["contactids"])
	d.Set("roundrobin", rrlist)
	d.Set("record_failover_values", rcdflist)
	d.Set("record_failover_failover_type", rcdfSet["record_failover_failover_type"])
	d.Set("record_failover_disable_flag", rcdfSet["record_failover_disable_flag"])

	return nil

}

func resourceConstellixANAMERecordUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	anameAttr := models.AnameAttributes{}

	if value, ok := d.GetOk("ttl"); ok {
		anameAttr.TTL = value.(int)
	}

	if d.HasChange("record_option") {
		anameAttr.RecordOption = d.Get("record_option").(string)
	}

	if name, ok := d.GetOk("name"); ok {
		anameAttr.Name = name.(string)
	}

	if d.HasChange("noanswer") {
		anameAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if d.HasChange("note") {
		anameAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("gtd_region") {
		anameAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if d.HasChange("type") {
		anameAttr.Type = d.Get("type").(string)
	}

	if d.HasChange("contact_ids") {

		contactidslist := toStringList(d.Get("contact_ids").([]interface{}))
		var intlist []int
		intlist = make([]int, len(contactidslist))
		for _, i := range contactidslist {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			intlist = append(intlist, j)
		}
		anameAttr.ContactIDs = intlist
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
		anameAttr.RoundRobin = maplistrr
	}

	var valuesrcdf *models.ValuesAname
	var rcdfaname *models.RCDFAname
	valueslist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("record_failover_values"); ok {
		rcdfaname = &models.RCDFAname{}
		valuesrcdf = &models.ValuesAname{}
		tp := value.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["checkid"] = fmt.Sprintf("%v", inner["check_id"])
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			valueslist = append(valueslist, map1)
		}
		rcdfaname.Values = valueslist
	}

	if failtype, ok := d.GetOk("record_failover_failover_type"); ok {
		rcdfaname.FailoverTypeAname, _ = strconv.Atoi(fmt.Sprintf("%v", failtype))
	}

	if disflag, ok := d.GetOk("record_failover_disable_flag"); ok {
		rcdfaname.DisableFlagAname, _ = strconv.ParseBool(fmt.Sprintf("%v", disflag))
	}

	if valuesrcdf != nil {
		rcdfaname.Values = valueslist
		anameAttr.RecordFailoverAname = rcdfaname
	} else {
		anameAttr.RecordFailoverAname = nil
	}

	anamerecordid := d.Id()

	_, err := constellixClient.UpdatebyID(anameAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/aname/"+anamerecordid)
	if err != nil {
		return err
	}
	return resourceConstellixANAMERecordRead(d, m)
}

func resourceConstellixANAMERecordDelete(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	anamerecordid := d.Id()

	err := constellixClient.DeletebyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/aname/" + anamerecordid)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}
