package constellix

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixCertCreate,
		Update: resourceConstellixCertUpdate,
		Read:   resourceConstellixCertRead,
		Delete: resourceConstellixCertDelete,

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

			"parentid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"roundrobin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_type": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"key_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"algorithm": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixCertCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	CertAttr := models.CertAttributes{}
	if name, ok := d.GetOk("name"); ok {
		CertAttr.Name = name.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		CertAttr.TTL = ttl.(int)
	}

	if noans, ok := d.GetOk("noanswer"); ok {
		CertAttr.NoAnswer = noans.(bool)
	}

	if note, ok := d.GetOk("note"); ok {
		CertAttr.Note = note.(string)
	}

	if gtdr, ok := d.GetOk("gtd_region"); ok {
		CertAttr.GtdRegion = gtdr.(int)
	}

	if tp, ok := d.GetOk("type"); ok {
		CertAttr.Type = tp.(string)
	}

	if pid, ok := d.GetOk("parentid"); ok {
		CertAttr.ParentID = pid.(int)
	}

	if p, ok := d.GetOk("parent"); ok {
		CertAttr.Parent = p.(string)
	}

	if source, ok := d.GetOk("source"); ok {
		CertAttr.Source = source.(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["certificateType"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["certificate_type"]))
			tpMap["keyTag"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["key_tag"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			tpMap["algorithm"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["algorithm"]))
			sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", inner["certificate"])))
			tpMap["certificate"] = sEnc
			mapListRR = append(mapListRR, tpMap)
		}
		CertAttr.RoundRobin = mapListRR
	}

	id := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)

	resp, err := client.Save(CertAttr, "v1/"+stid+"/"+id+"/records/cert")
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
	return resourceConstellixCertRead(d, m)
}

func resourceConstellixCertUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	CertAttr := models.CertAttributes{}

	CertAttr.Name = d.Get("name").(string)

	CertAttr.TTL = d.Get("ttl").(int)

	CertAttr.NoAnswer = d.Get("noanswer").(bool)

	CertAttr.Note = d.Get("note").(string)

	CertAttr.GtdRegion = d.Get("gtd_region").(int)

	if d.HasChange("type") {
		CertAttr.Type = d.Get("type").(string)
	}

	CertAttr.ParentID = d.Get("parentid").(int)

	if d.HasChange("parent") {
		CertAttr.Parent = d.Get("parent").(string)
	}

	if d.HasChange("source") {
		CertAttr.Source = d.Get("source").(string)
	}

	if rr, ok := d.GetOk("roundrobin"); ok {
		mapListRR := make([]interface{}, 0, 1)
		tp := rr.(*schema.Set).List()
		for _, val := range tp {
			tpMap := make(map[string]interface{})
			inner := val.(map[string]interface{})
			tpMap["certificateType"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["certificate_type"]))
			tpMap["keyTag"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["key_tag"]))
			tpMap["disableFlag"] = fmt.Sprintf("%v", inner["disable_flag"])
			tpMap["algorithm"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["algorithm"]))
			sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", inner["certificate"])))
			tpMap["certificate"] = sEnc
			mapListRR = append(mapListRR, tpMap)
		}
		CertAttr.RoundRobin = mapListRR
	}

	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	certid := d.Id()
	_, err := client.UpdatebyID(CertAttr, "v1/"+stid+"/"+domainID+"/records/cert/"+certid)
	if err != nil {
		return err
	}
	return resourceConstellixCertRead(d, m)
}

func resourceConstellixCertRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	certid := d.Id()

	resp, err := client.GetbyId("v1/" + stid + "/" + domainID + "/records/cert/" + certid)
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
		tpMap["certificate_type"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["certificateType"]))
		tpMap["key_tag"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["keyTag"]))
		tpMap["disable_flag"] = fmt.Sprintf("%v", inner["disableFlag"])
		tpMap["algorithm"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["algorithm"]))
		sEnc, _ := b64.StdEncoding.DecodeString(fmt.Sprintf("%v", inner["certificate"]))
		tpMap["certificate"] = sEnc
		mapListRR = append(mapListRR, tpMap)
	}

	d.Set("roundrobin", mapListRR)
	return nil
}

func resourceConstellixCertDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)
	domainID := d.Get("domain_id").(string)
	stid := d.Get("source_type").(string)
	dn := d.Id()

	err := constellixConnect.DeletebyId("v1/" + stid + "/" + domainID + "/records/cert/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
