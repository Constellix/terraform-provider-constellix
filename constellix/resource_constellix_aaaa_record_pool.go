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

func resourceConstellixAAAArecordPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixAAAAPoolCreate,
		Update: resourceConstellixAAAAPoolUpdate,
		Read:   resourceConstellixAAAAPoolRead,
		Delete: resourceConstellixAAAAPoolDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"num_return": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"min_available_failover": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"failed_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"disable_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
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

						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"check_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConstellixAAAAPoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	aaaapoolAttr := models.AAAArecordPoolAttributes{}

	if name, ok := d.GetOk("name"); ok {
		aaaapoolAttr.Name = name.(string)
	}

	if numr, ok := d.GetOk("num_return"); ok {
		aaaapoolAttr.NumReturn = numr.(int)
	}

	if minaf, ok := d.GetOk("min_available_failover"); ok {
		aaaapoolAttr.MinavailFailover = minaf.(int)
	}

	if fflag, ok := d.GetOk("failed_flag"); ok {
		aaaapoolAttr.Failedflag = fflag.(bool)
	}

	if dflag, ok := d.GetOk("disable_flag"); ok {
		aaaapoolAttr.Disableflag = dflag.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		aaaapoolAttr.Note = note.(string)
	}

	if rr, ok := d.GetOk("values"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["value"] = fmt.Sprintf("%v", inner["value"])
			tpMap["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
			tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disable_flag"])
			tpMap["check_id"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["check_id"]))
			tpMap["policy"] = fmt.Sprintf("%v", inner["policy"])

			mapListRR = append(mapListRR, tpMap)
		}
		aaaapoolAttr.Values = mapListRR
	}

	resp, err := client.Save(aaaapoolAttr, "v1/pools/AAAA")
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
	return resourceConstellixAAAAPoolRead(d, m)
}

func resourceConstellixAAAAPoolUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	aaaapoolAttr := models.AAAArecordPoolAttributes{}

	if d.HasChange("name") {
		aaaapoolAttr.Name = d.Get("name").(string)
	}

	if d.HasChange("num_return") {
		aaaapoolAttr.NumReturn = d.Get("num_return").(int)
	}

	if minaf, ok := d.GetOk("min_available_failover"); ok {
		aaaapoolAttr.MinavailFailover = minaf.(int)
	}

	if d.HasChange("failed_flag") {
		aaaapoolAttr.Failedflag = d.Get("failed_flag").(bool)
	}

	if d.HasChange("disable_flag") {
		aaaapoolAttr.Disableflag = d.Get("disable_flag").(bool)
	}

	if d.HasChange("note") {
		aaaapoolAttr.Note = d.Get("note").(string)
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
		aaaapoolAttr.Values = mapListRR
	}

	dn := d.Id()

	_, err := client.UpdatebyID(aaaapoolAttr, "v1/pools/AAAA/"+dn)
	if err != nil {
		return err
	}

	//Managing response and extracting id of resource
	return resourceConstellixAAAAPoolRead(d, m)
}

func resourceConstellixAAAAPoolRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	resp, err := client.GetbyId("v1/pools/AAAA/" + dn)
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
	d.Set("failed_flag", data["failedFlag"])
	d.Set("disable_flag", data["disableFlag"])
	d.Set("note", data["note"])

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

func resourceConstellixAAAAPoolDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	dn := d.Id()

	err := client.DeletebyId("v1/pools/AAAA/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
