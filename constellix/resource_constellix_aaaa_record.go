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

func resourceConstellixAAAARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixAAAARecordCreate,
		Update: resourceConstellixAAAARecordUpdate,
		Read:   resourceConstellixAAAARecordRead,
		Delete: resourceConstellixAAAARecordDelete,

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
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"geo_location": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"geo_ip_user_region": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"geo_ip_proximity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
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
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Required: true,
			},
			"roundrobin_failover": &schema.Schema{
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
							Computed: true,
						},
						"sort_order": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
				Computed: true,
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
			},

			"record_failover_failover_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"record_failover_disable_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceConstellixAAAARecordCreate(d *schema.ResourceData, m interface{}) error {

	constellixConnect := m.(*client.Client)

	aAttr := models.AAAARecordAttributes{}

	if name, ok := d.GetOk("name"); ok {
		aAttr.Name = name.(string)
	}
	if TTL, ok := d.GetOk("ttl"); ok {
		aAttr.TTL = TTL.(int)
	}
	if RecordOption, ok := d.GetOk("record_option"); ok {
		aAttr.RecordOption = RecordOption.(string)
	}
	if NoAnswer, ok := d.GetOk("noanswer"); ok {
		aAttr.NoAnswer = NoAnswer.(bool)
	}
	if Note, ok := d.GetOk("note"); ok {
		aAttr.Note = Note.(string)
	}
	if GtdRegion, ok := d.GetOk("gtd_region"); ok {
		aAttr.GtdRegion = GtdRegion.(int)
	}
	if Type, ok := d.GetOk("type"); ok {
		aAttr.Type = Type.(string)
	}
	if contactid, ok := d.GetOk("contact_ids"); ok {
		aAttr.ContactId = toListOfInt(contactid)
	}
	if pools, ok := d.GetOk("pools"); ok {
		aAttr.Pools = toListOfInt(pools)
	}

	var geoloc *models.Geolocation
	if geoipuserregion, ok := d.GetOk("geo_location"); ok {
		geoloc = &models.Geolocation{}
		geouserlist := make([]int, 0)
		tp := geoipuserregion.(map[string]interface{})
		var1, _ := strconv.Atoi(fmt.Sprintf("%v", tp["geo_ip_user_region"]))

		if tp["geo_ip_user_region"] != nil {
			geouserlist = append(geouserlist, var1)
			geoloc.GeoIpUserRegion = geouserlist
		}
		geoloc.Drop, _ = strconv.ParseBool(fmt.Sprintf("%v", tp["drop"]))
		geoloc.GeoIpProximity, _ = strconv.Atoi(fmt.Sprintf("%v", tp["geo_ip_proximity"]))

		if geoloc != nil {
			aAttr.GeoLocation = geoloc
		} else {
			aAttr.GeoLocation = nil
		}
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			maplistrr = append(maplistrr, map1)
		}
		aAttr.RoundRobin = maplistrr
	}

	maplist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("roundrobin_failover"); ok {
		tp := value.(*schema.Set).List()

		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			maplist = append(maplist, map1)
		}
		aAttr.RoundRobinFailoverA = maplist
	}

	var valuesrcdf *models.ValuesRCDFA
	var rcdfa *models.RCDFA //added
	valueslist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("record_failover_values"); ok {
		rcdfa = &models.RCDFA{} //added
		valuesrcdf = &models.ValuesRCDFA{}
		tp := value.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["checkId"] = fmt.Sprintf("%v", inner["check_id"])
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			valueslist = append(valueslist, map1)
		}
		rcdfa.Values = valueslist
	}

	if failovertype, ok := d.GetOk("record_failover_failover_type"); ok {
		rcdfa.FailoverTypeRCDFA, _ = strconv.Atoi(fmt.Sprintf("%v", failovertype)) //added
	}

	if disableflag, ok := d.GetOk("record_failover_disable_flag"); ok {
		rcdfa.DisableFlagRCDFA, _ = strconv.ParseBool(fmt.Sprintf("%v", disableflag)) //added
	}

	if valuesrcdf != nil {
		rcdfa.Values = valueslist     //added
		aAttr.RecordFailoverA = rcdfa //added
	} else {
		aAttr.RecordFailoverA = nil
	}

	resp, err := constellixConnect.Save(aAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/aaaa")

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

	return resourceConstellixAAAARecordRead(d, m)

}

func resourceConstellixAAAARecordRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	arecordid := d.Id()

	resp, err := constellixClient.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/aaaa/" + arecordid)
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
	err = json.Unmarshal([]byte(bodyString), &data)
	if err != nil {
		return err
	}

	geoloc1 := data["geo_location"]
	geoset := make(map[string]interface{})
	if geoloc1 != nil {
		geoloc := geoloc1.(map[string]interface{})
		if geoloc["geoipFilter"] != nil {
			geoset["geo_ip_user_region"] = fmt.Sprintf("%v", geoloc["geoipFilter"])
			geoset["drop"] = fmt.Sprintf("%v", geoloc["drop"])
		} else {
			geoset["geo_ip_proximity"] = fmt.Sprintf("%v", geoloc["geoipProximity"])
		}
	} else {
		geoset = nil
	}

	arecroundrobin := data["roundRobin"].([]interface{})
	rrlist := make([]interface{}, 0, 1)
	for _, valrrf := range arecroundrobin {
		map1 := make(map[string]interface{})
		val1 := valrrf.(map[string]interface{})
		map1["value"] = fmt.Sprintf("%v", val1["value"])
		map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
		map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])

		rrlist = append(rrlist, map1)
	}

	arecroundrobinfailover := data["roundRobinFailover"].([]interface{})

	rrflist := make([]interface{}, 0, 1)
	for _, valrrf := range arecroundrobinfailover {
		map1 := make(map[string]interface{})
		val1 := valrrf.(map[string]interface{})
		map1["value"] = fmt.Sprintf("%v", val1["value"])
		map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
		map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])

		rrflist = append(rrflist, map1)
	}

	rcdf := data["recordFailover"]
	rcdfset := make(map[string]interface{})
	rcdflist := make([]interface{}, 0, 1)
	if rcdf != nil {
		rcdf1 := rcdf.(map[string]interface{})
		rcdfset["record_failover_failover_type"] = fmt.Sprintf("%v", rcdf1["failoverType"])

		rcdfset["record_failover_disable_flag"] = fmt.Sprintf("%v", rcdf1["disabled"])

		rcdfvalues := rcdf1["values"].([]interface{})

		for _, valrcdf := range rcdfvalues {
			map1 := make(map[string]interface{})
			val1 := valrcdf.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", val1["value"])
			map1["sort_order"] = fmt.Sprintf("%v", val1["sortOrder"])
			map1["disable_flag"] = fmt.Sprintf("%v", val1["disableFlag"])
			map1["check_id"] = fmt.Sprintf("%v", val1["checkId"])
			rcdflist = append(rcdflist, map1)
		}
	}

	d.Set("id", data["id"])
	d.Set("name", data["name"])
	d.Set("ttl", data["ttl"])
	d.Set("geo_location", geoset)
	d.Set("record_option", data["recordOption"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	d.Set("pools", data["pools"])
	d.Set("contact_ids", data["contactId"])
	d.Set("roundrobin", rrlist)
	d.Set("roundrobin_failover", rrflist)
	d.Set("record_failover_values", rcdflist)
	d.Set("record_failover_failover_type", rcdfset["record_failover_failover_type"])
	d.Set("record_failover_disable_flag", rcdfset["record_failover_disable_flag"])
	return nil
}

func resourceConstellixAAAARecordUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	aAttr := models.AAAARecordAttributes{}

	if ttl, ok := d.GetOk("ttl"); ok {
		aAttr.TTL = ttl.(int)
	}
	if name, ok := d.GetOk("name"); ok {
		aAttr.Name = name.(string)
	}

	if recordoption, ok := d.GetOk("record_option"); ok {
		aAttr.RecordOption = recordoption.(string)
	}

	if _, ok := d.GetOk("noanswer"); ok {
		aAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		aAttr.Note = note.(string)
	}

	if _, ok := d.GetOk("gtd_region"); ok {
		aAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if _, ok := d.GetOk("type"); ok {
		aAttr.Type = d.Get("type").(string)
	}
	if contactid, ok := d.GetOk("contact_ids"); ok {
		aAttr.ContactId = toListOfInt(contactid)
	}
	if pools, ok := d.GetOk("pools"); ok {
		aAttr.Pools = toListOfInt(pools)
	}

	var geoloc *models.Geolocation
	if geoipuserregion, ok := d.GetOk("geo_location"); ok {
		geoloc = &models.Geolocation{}
		geouserlist := make([]int, 0)
		tp := geoipuserregion.(map[string]interface{})
		var1, _ := strconv.Atoi(fmt.Sprintf("%v", tp["geo_ip_user_region"]))
		if tp["geo_ip_user_region"] != nil {
			geouserlist = append(geouserlist, var1)
			geoloc.GeoIpUserRegion = geouserlist
		}
		geoloc.Drop, _ = strconv.ParseBool(fmt.Sprintf("%v", tp["drop"]))
		geoloc.GeoIpProximity, _ = strconv.Atoi(fmt.Sprintf("%v", tp["geo_ip_proximity"]))

		if geoloc != nil {
			aAttr.GeoLocation = geoloc
		} else {
			aAttr.GeoLocation = nil
		}
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			maplistrr = append(maplistrr, map1)
		}
		aAttr.RoundRobin = maplistrr
	}

	maplist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("roundrobin_failover"); ok {
		tp := value.(*schema.Set).List()

		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			maplist = append(maplist, map1)
		}
		aAttr.RoundRobinFailoverA = maplist
	}

	var valuesrcdf *models.ValuesRCDFA
	var rcdfa *models.RCDFA //added
	valueslist := make([]interface{}, 0, 1)
	if value, ok := d.GetOk("record_failover_values"); ok {
		rcdfa = &models.RCDFA{} //added
		valuesrcdf = &models.ValuesRCDFA{}
		tp := value.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["checkId"] = fmt.Sprintf("%v", inner["check_id"])
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["sortOrder"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["sort_order"]))
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			valueslist = append(valueslist, map1)
		}
		rcdfa.Values = valueslist
	}

	if failovertype, ok := d.GetOk("record_failover_failover_type"); ok {
		rcdfa.FailoverTypeRCDFA, _ = strconv.Atoi(fmt.Sprintf("%v", failovertype)) //added
	}

	if disableflag, ok := d.GetOk("record_failover_disable_flag"); ok {
		rcdfa.DisableFlagRCDFA, _ = strconv.ParseBool(fmt.Sprintf("%v", disableflag)) //added
	}

	if valuesrcdf != nil {
		rcdfa.Values = valueslist     //added
		aAttr.RecordFailoverA = rcdfa //added
	} else {
		aAttr.RecordFailoverA = nil
	}

	arecordid := d.Id()

	_, err := constellixClient.UpdatebyID(aAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/aaaa/"+arecordid)
	if err != nil {
		return err
	}
	return resourceConstellixAAAARecordRead(d, m)
}

func resourceConstellixAAAARecordDelete(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	arecordid := d.Id()

	err := constellixClient.DeletebyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/aaaa/" + arecordid)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}
