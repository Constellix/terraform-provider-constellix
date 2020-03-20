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

func resourceConstellixDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixDNSCreate,
		Update: resourceConstellixDNSUpdate,
		Read:   resourceConstellixDNSRead,
		Delete: resourceConstellixDNSDelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"has_gtd_regions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"has_geoip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"nameserver_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"soa": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_nameserver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"refresh": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"serial": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"retry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"expire": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"negcache": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixDNSCreate(d *schema.ResourceData, m interface{}) error {

	constellixConnect := m.(*client.Client)

	domainAttr := models.DomainAttributes{}

	if name, ok := d.GetOk("name"); ok {
		nameList := toStringList(name)
		domainAttr.Name = nameList
	}

	if hasgtdregions, ok := d.GetOk("has_gtd_regions"); ok {
		domainAttr.HasGtdRegions = hasgtdregions.(bool)
	}

	if hasgeoip, ok := d.GetOk("has_geoip"); ok {
		domainAttr.HasGeoIP = hasgeoip.(bool)
	}

	if nsg, ok := d.GetOk("nameserver_group"); ok {
		domainAttr.NameserverGroup = nsg.(string)
	}

	if note, ok := d.GetOk("note"); ok {
		domainAttr.Note = note.(string)
	}

	if tg, ok := d.GetOk("tags"); ok {
		tagsList := toStringList(tg.([]interface{}))
		domainAttr.Tags = tagsList
	}

	var soaAttr *models.Soa
	if value, ok := d.GetOk("soa"); ok {
		soaAttr = &models.Soa{}
		tp := value.(map[string]interface{})
		if tp["primary_nameserver"] != nil {
			soaAttr.PrimaryNameServer = fmt.Sprintf("%v", tp["primary_nameserver"])
		}
		if tp["email"] != nil {
			soaAttr.Email = fmt.Sprintf("%v", tp["email"])
		}
		if tp["ttl"] != nil {
			var1 := fmt.Sprintf("%v", tp["ttl"])
			soaAttr.TTL, _ = strconv.Atoi(var1)
		}
		if tp["expire"] != nil {
			var2 := fmt.Sprintf("%v", tp["expire"])
			soaAttr.Expire, _ = strconv.Atoi(var2)
		}
		if tp["negcache"] != nil {
			var3 := fmt.Sprintf("%v", tp["negcache"])
			soaAttr.NegCache, _ = strconv.Atoi(var3)
		}
		if tp["refresh"] != nil {
			var4 := fmt.Sprintf("%v", tp["refresh"])
			soaAttr.Refresh, _ = strconv.Atoi(var4)
		}
		if tp["retry"] != nil {
			var5 := fmt.Sprintf("%v", tp["retry"])
			soaAttr.Retry, _ = strconv.Atoi(var5)
		}
	}

	domainAttr.Soa = soaAttr

	resp, err := constellixConnect.Save(domainAttr, "v1/domains")

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
	return resourceConstellixDNSRead(d, m)
}
func resourceConstellixDNSRead(d *schema.ResourceData, m interface{}) error {
	constellixclient := m.(*client.Client)
	dn := d.Id()
	resp, err := constellixclient.GetbyId("v1/domains/" + dn)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	recsoa := data["soa"].(map[string]interface{})

	soaset := make(map[string]interface{})
	soaset["primary_nameserver"] = recsoa["primaryNameserver"]
	soaset["ttl"] = fmt.Sprintf("%v", recsoa["ttl"])
	if value, ok := d.GetOk("soa"); ok {
		tp := value.(map[string]interface{})
		if tp["email"] != nil {
			soaset["email"] = recsoa["email"]
		}
	}
	soaset["refresh"] = fmt.Sprintf("%v", recsoa["refresh"])
	soaset["expire"] = fmt.Sprintf("%v", recsoa["expire"])
	soaset["retry"] = fmt.Sprintf("%v", recsoa["retry"])
	soaset["negcache"] = fmt.Sprintf("%v", recsoa["negCache"])

	d.Set("id", data["id"])
	d.Set("name", data["name"])
	d.Set("soa", soaset)
	d.Set("typeid", data["typeId"])
	d.Set("has_geoip", data["hasGeoIP"])
	d.Set("has_gtd_regions", data["hasGtdRegions"])
	d.Set("nameserver_group", data["nameserverGroup"])
	d.Set("note", data["note"])
	d.Set("version", data["version"])
	d.Set("status", data["status"])
	d.Set("tags", data["tags"])

	return nil

}

func resourceConstellixDNSUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	domainAttr := models.DomainAttributes{}

	domainAttr.HasGtdRegions = d.Get("has_gtd_regions").(bool)

	domainAttr.HasGeoIP = d.Get("has_geoip").(bool)

	if d.HasChange("nameserver_group") {
		domainAttr.NameserverGroup = d.Get("nameserver_group").(string)
	}

	if d.HasChange("note") {
		domainAttr.Note = d.Get("note").(string)
	}

	if d.HasChange("tags") {
		tagsList := toStringList(d.Get("tags").([]interface{}))
		domainAttr.Tags = tagsList
	}

	var soaAttr *models.Soa

	value := d.Get("soa")
	soaAttr = &models.Soa{}
	tp := value.(map[string]interface{})
	if tp["primary_nameserver"] != nil {
		soaAttr.PrimaryNameServer = fmt.Sprintf("%v", tp["primary_nameserver"])
	}
	if tp["email"] != nil {
		soaAttr.Email = fmt.Sprintf("%v", tp["email"])
	}
	if tp["ttl"] != nil {
		var1 := fmt.Sprintf("%v", tp["ttl"])
		soaAttr.TTL, _ = strconv.Atoi(var1)
	}
	if tp["expire"] != nil {
		var2 := fmt.Sprintf("%v", tp["expire"])
		soaAttr.Expire, _ = strconv.Atoi(var2)
	}
	if tp["negcache"] != nil {
		var3 := fmt.Sprintf("%v", tp["negcache"])
		soaAttr.NegCache, _ = strconv.Atoi(var3)
	}
	if tp["refresh"] != nil {
		var4 := fmt.Sprintf("%v", tp["refresh"])
		soaAttr.Refresh, _ = strconv.Atoi(var4)
	}
	if tp["retry"] != nil {
		var5 := fmt.Sprintf("%v", tp["retry"])
		soaAttr.Retry, _ = strconv.Atoi(var5)
	}

	domainAttr.Soa = soaAttr

	dn := d.Id()

	_, err := constellixClient.UpdatebyID(domainAttr, "v1/domains/"+dn)
	if err != nil {
		return err
	}
	return resourceConstellixDNSRead(d, m)

}

func resourceConstellixDNSDelete(d *schema.ResourceData, m interface{}) error {
	constellixConnect := m.(*client.Client)

	dn := d.Id()

	err := constellixConnect.DeletebyId("v1/domains/" + dn)
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}
