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

func resourceConstellixARecordPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixARecordPoolCreate,
		Update: resourceConstellixARecordPoolUpdate,
		Read:   resourceConstellixARecordPoolRead,
		Delete: resourceConstellixARecordPoolDelete,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"num_return": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"min_available_failover": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failed_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"disable_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"values": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixARecordPoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	arecordpoolAttr := models.ARecordPoolAttributes{}

	if name, ok := d.GetOk("name"); ok {
		arecordpoolAttr.Name = name.(string)
	}

	if nr, ok := d.GetOk("num_return"); ok {
		arecordpoolAttr.NumReturn = nr.(int)
	}

	if minaf, ok := d.GetOk("min_available_failover"); ok {
		arecordpoolAttr.MinAvailableFailover = minaf.(int)
	}

	if note, ok := d.GetOk("note"); ok {
		arecordpoolAttr.Note = note.(string)
	}

	if vr, ok := d.GetOk("version"); ok {
		arecordpoolAttr.Version = vr.(int)
	}

	if ff, ok := d.GetOk("failed_flag"); ok {
		arecordpoolAttr.FailedFlag = ff.(string)
	}

	if df1, ok := d.GetOk("disable_flag"); ok {
		arecordpoolAttr.DisableFlag1 = df1.(string)
	}

	if rr, ok := d.GetOk("values"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"] = fmt.Sprintf("%v", inner["value"])
			tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			tpMap["checkId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["check_id"]))
			tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])

			mapListRR = append(mapListRR, tpMap)
		}
		arecordpoolAttr.Values = mapListRR
	}

	resp, err := client.Save(arecordpoolAttr, "v1/pools/A")
	if err != nil {
		return err
	}

	bodybtes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodystring := string(bodybtes)
	log.Println("Body String of ARecordPool Respince :", bodystring)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodystring[1:len(bodystring)-1]), &data)

	d.SetId(fmt.Sprintf("%.0f", data["id"]))
	return resourceConstellixARecordPoolRead(d, m)
}

func resourceConstellixARecordPoolUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	arecordpoolAttr := models.ARecordPoolAttributes{}

	arecordpoolAttr.Name = d.Get("name").(string)

	arecordpoolAttr.NumReturn = d.Get("num_return").(int)

	arecordpoolAttr.MinAvailableFailover = d.Get("min_available_failover").(int)

	if d.HasChange("failed_flag") {
		arecordpoolAttr.FailedFlag = d.Get("failed_flag").(string)
	}

	if d.HasChange("disable_flag") {
		arecordpoolAttr.DisableFlag1 = d.Get("disable_flag").(string)
	}

	if d.HasChange("note") {
		arecordpoolAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("version") {
		arecordpoolAttr.Version = d.Get("version").(int)
	}

	if rr, ok := d.GetOk("values"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"] = fmt.Sprintf("%v", inner["value"])
			tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			tpMap["checkId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["check_id"]))
			tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])

			mapListRR = append(mapListRR, tpMap)
		}
		arecordpoolAttr.Values = mapListRR
	}

	arecordpoolid := d.Id()
	_, err := client.UpdatebyID(arecordpoolAttr, "v1/pools/A/"+arecordpoolid)
	if err != nil {
		return err
	}
	return resourceConstellixARecordPoolRead(d, m)
}

func resourceConstellixARecordPoolRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	arecordpoolid := d.Id()

	resp, err := client.GetbyId("v1/pools/A/" + arecordpoolid)
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
	d.Set("num_return", data["numReturn"])
	d.Set("min_available_failover", data["minAvailableFailover"])
	d.Set("note", data["note"])
	d.Set("version", data["version"])
	d.Set("failed_flag", data["failedFlag"])
	d.Set("disable_flag", data["disableFlag"])
	resrr := (data["values"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"] = fmt.Sprintf("%v", inner["value"])
		tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
		tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
		tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])
		tpMap["check_id"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["checkId"]))

		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("values", mapListRR)
	return nil
}

func resourceConstellixARecordPoolDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v1/pools/A/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
