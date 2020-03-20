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

func resourceConstellixSRVRecord() *schema.Resource {
	return &schema.Resource{
		Create:        resourceConstellixSRVRecordCreate,
		Read:          resourceConstellixSRVRecordRead,
		Update:        resourceConstellixSRVRecordUpdate,
		Delete:        resourceConstellixSRVRecordDelete,
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

			"roundrobin": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConstellixSRVRecordCreate(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	srvAttr := models.SRVAttributes{}

	if name, ok := d.GetOk("name"); ok {
		srvAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		srvAttr.TTL = ttl.(int)
	}
	if noanswer, ok := d.GetOk("noanswer"); ok {
		srvAttr.NoAnswer = noanswer.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		srvAttr.Note = note.(string)
	}

	if gtdregion, ok := d.GetOk("gtd_region"); ok {
		srvAttr.GtdRegion = gtdregion.(int)
	}

	if types, ok := d.GetOk("type"); ok {
		srvAttr.Type = types.(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["port"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["port"]))
			map1["priority"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["priority"]))
			map1["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
			maplistrr = append(maplistrr, map1)
		}
		srvAttr.RoundRobin = maplistrr
	}

	resp, err := constellixConnect.Save(srvAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/srv")
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

	return resourceConstellixSRVRecordRead(d, m)
}

func resourceConstellixSRVRecordRead(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	srvid := d.Id()

	resp, err := constellixClient.GetbyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/srv/" + srvid)
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
		map1["port"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["port"]))
		map1["priority"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["priority"]))
		map1["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", val1["weight"]))
		rrlist = append(rrlist, map1)
	}

	d.Set("id", data["id"])
	d.Set("name", data["name"])
	d.Set("ttl", data["ttl"])
	d.Set("noanswer", data["noAnswer"])
	d.Set("note", data["note"])
	d.Set("gtd_region", data["gtdRegion"])
	d.Set("type", data["type"])
	d.Set("roundrobin", rrlist)
	return nil
}

func resourceConstellixSRVRecordUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	srvAttr := models.SRVAttributes{}

	if d.HasChange("name") {
		srvAttr.Name = d.Get("name").(string)
	}

	if d.HasChange("ttl") {
		srvAttr.TTL = d.Get("ttl").(int)
	}

	if d.HasChange("noanswer") {
		srvAttr.NoAnswer = d.Get("noanswer").(bool)
	}

	if d.HasChange("note") {
		srvAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("gtd_region") {
		srvAttr.GtdRegion = d.Get("gtd_region").(int)
	}

	if d.HasChange("type") {
		srvAttr.Type = d.Get("type").(string)
	}

	maplistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("roundrobin"); ok {
		tp := val.(*schema.Set).List()
		for _, val := range tp {
			map1 := make(map[string]interface{})
			inner := val.(map[string]interface{})
			map1["value"] = fmt.Sprintf("%v", inner["value"])
			map1["disableFlag"], _ = strconv.ParseBool(fmt.Sprintf("%v", inner["disable_flag"]))
			map1["port"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["port"]))
			map1["priority"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["priority"]))
			map1["weight"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["weight"]))
			maplistrr = append(maplistrr, map1)
		}
		srvAttr.RoundRobin = maplistrr
	}

	srvid := d.Id()
	_, err := constellixClient.UpdatebyID(srvAttr, "v1/"+d.Get("source_type").(string)+"/"+d.Get("domain_id").(string)+"/records/srv/"+srvid)
	if err != nil {
		return err
	}
	return resourceConstellixSRVRecordRead(d, m)

}

func resourceConstellixSRVRecordDelete(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)
	srvid := d.Id()

	err := constellixClient.DeletebyId("v1/" + d.Get("source_type").(string) + "/" + d.Get("domain_id").(string) + "/records/srv/" + srvid)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}
