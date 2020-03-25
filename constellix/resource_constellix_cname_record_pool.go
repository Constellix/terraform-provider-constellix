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

func resourceConstellixCnameRecordPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixCnameRecordPoolCreate,
		Update: resourceConstellixCnameRecordPoolUpdate,
		Read:   resourceConstellixCnameRecordPoolRead,
		Delete: resourceConstellixCnameRecordPoolDelete,

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

func resourceConstellixCnameRecordPoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	cnamerecordpoolAttr := models.CnameRecordPoolAttributes{}

	if name, ok := d.GetOk("name"); ok {
		cnamerecordpoolAttr.Name = name.(string)
	}

	if nr, ok := d.GetOk("num_return"); ok {
		cnamerecordpoolAttr.NumReturn = nr.(int)
	}

	if minaf, ok := d.GetOk("min_available_failover"); ok {
		cnamerecordpoolAttr.MinAvailableFailover = minaf.(int)
	}

	if note, ok := d.GetOk("note"); ok {
		cnamerecordpoolAttr.Note = note.(string)
	}

	if vr, ok := d.GetOk("version"); ok {
		cnamerecordpoolAttr.Version = vr.(int)
	}

	if ff, ok := d.GetOk("failed_flag"); ok {
		cnamerecordpoolAttr.FailedFlag = ff.(string)
	}

	if df1, ok := d.GetOk("disableflag1"); ok {
		cnamerecordpoolAttr.DisableFlag1 = df1.(string)
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
		cnamerecordpoolAttr.ValuesCname = mapListRR
	}

	resp, err := client.Save(cnamerecordpoolAttr, "v1/pools/CNAME")
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
	return resourceConstellixCnameRecordPoolRead(d, m)
}

func resourceConstellixCnameRecordPoolUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	cnamerecordpoolAttr := models.CnameRecordPoolAttributes{}

	cnamerecordpoolAttr.Name = d.Get("name").(string)

	cnamerecordpoolAttr.NumReturn = d.Get("num_return").(int)

	cnamerecordpoolAttr.MinAvailableFailover = d.Get("min_available_failover").(int)

	if d.HasChange("failed_flag") {
		cnamerecordpoolAttr.FailedFlag = d.Get("failed_flag").(string)
	}

	if d.HasChange("disableflag1") {
		cnamerecordpoolAttr.DisableFlag1 = d.Get("disableflag1").(string)
	}

	if d.HasChange("note") {
		cnamerecordpoolAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("version") {
		cnamerecordpoolAttr.Version = d.Get("version").(int)
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
		cnamerecordpoolAttr.ValuesCname = mapListRR
	}

	cnamerecordpoolid := d.Id()
	_, err := client.UpdatebyID(cnamerecordpoolAttr, "v1/pools/CNAME/"+cnamerecordpoolid)
	if err != nil {
		return err
	}
	return resourceConstellixCnameRecordPoolRead(d, m)
}

func resourceConstellixCnameRecordPoolRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	cnamerecordpoolid := d.Id()

	resp, err := client.GetbyId("v1/pools/CNAME/" + cnamerecordpoolid)
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
	d.Set("num_return", data["numReturn"])
	d.Set("min_available_failover", data["minAvailableFailover"])
	d.Set("note", data["note"])
	d.Set("version", data["version"])
	d.Set("failed_flag", data["failedFlag"])
	d.Set("disable_flag", data["disableFlag"])
	resrr := (data["values"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		log.Println("RR are : ", val)
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"] = fmt.Sprintf("%v", inner["value"])
		tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
		tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disableFlag"])
		tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])
		tpMap["checkId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["checkId"]))

		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("values", mapListRR)
	return nil
}

func resourceConstellixCnameRecordPoolDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v1/pools/CNAME/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
