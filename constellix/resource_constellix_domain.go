package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/Jeffail/gabs"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConstellixDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceConstellixDNSCreate,
		Update: resourceConstellixDNSUpdate,
		Read:   resourceConstellixDNSRead,
		Delete: resourceConstellixDNSDelete,

		Importer: &schema.ResourceImporter{
			State: resourceConstellixDNSImport,
		},

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

			"vanity_nameserver": &schema.Schema{
				Type:     schema.TypeString,
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
				Computed: true,
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
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"retry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"expire": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"negcache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceConstellixDNSImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	constellixClient := m.(*client.Client)
	resp, err := constellixClient.GetbyId("v1/domains/" + d.Id())
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil, err
		}
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj, err := gabs.ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}

	soaset := make(map[string]interface{})
	soaset["primary_nameserver"] = stripQuotes(obj.S("soa", "primaryNameserver").String())
	soaset["ttl"] = stripQuotes(obj.S("soa", "ttl").String())
	if value, ok := d.GetOk("soa"); ok {
		tp := value.(map[string]interface{})
		if tp["email"] != nil {
			soaset["email"] = stripQuotes(obj.S("soa", "email").String())
		}
	}
	soaset["refresh"] = stripQuotes(obj.S("soa", "refresh").String())
	soaset["expire"] = stripQuotes(obj.S("soa", "expire").String())
	soaset["retry"] = stripQuotes(obj.S("soa", "retry").String())
	soaset["negcache"] = stripQuotes(obj.S("soa", "negCache").String())

	d.Set("id", stripQuotes(obj.S("id").String()))
	d.Set("name", stripQuotes(obj.S("name").String()))
	d.Set("soa", soaset)
	if hasGeoIP, err := strconv.ParseBool(stripQuotes(obj.S("hasGeoIP").String())); err == nil {
		d.Set("has_geoip", hasGeoIP)
	}
	if hasGTDRegion, err := strconv.ParseBool(stripQuotes(obj.S("hasGtdRegions").String())); err == nil {
		d.Set("has_gtd_regions", hasGTDRegion)
	}
	if obj.Exists("vanityNameServer") {
		d.Set("vanity_nameserver", stripQuotes(obj.S("vanityNameServer").String()))
	}
	if obj.Exists("nameserverGroup") {
		d.Set("nameserver_group", stripQuotes(obj.S("nameserverGroup").String()))
	}
	d.Set("note", stripQuotes(obj.S("note").String()))

	if obj.S("tags").Data() != nil {
		d.Set("tags", toListOfString(obj.S("tags").Data()))
	} else {
		d.Set("tags", make([]string, 0, 1))
	}

	log.Printf("[DEBUG] %s finished import", d.Id())
	return []*schema.ResourceData{d}, nil
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

	if vanityNS, ok := d.GetOk("vanity_nameserver"); ok {
		domainAttr.VanityNameServer = vanityNS.(string)
	}

	if nsg, ok := d.GetOk("nameserver_group"); ok {
		domainAttr.NameserverGroup = nsg.(string)
	}

	if note, ok := d.GetOk("note"); ok {
		domainAttr.Note = note.(string)
	}

	if tg, ok := d.GetOk("tags"); ok {
		tagsList := tg.([]interface{})
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
			soaAttr.TTL = tp["ttl"].(string)
		}
		if tp["expire"] != nil {
			soaAttr.Expire = tp["expire"].(string)
		}
		if tp["negcache"] != nil {
			soaAttr.NegCache = tp["negcache"].(string)
		}
		if tp["refresh"] != nil {
			soaAttr.Refresh = tp["refresh"].(string)
		}
		if tp["retry"] != nil {
			soaAttr.Retry = tp["retry"].(string)
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

	obj, err := gabs.ParseJSON(bodyBytes)
	if err != nil {
		return err
	}

	soaset := make(map[string]interface{})
	soaset["primary_nameserver"] = stripQuotes(obj.S("soa", "primaryNameserver").String())
	soaset["ttl"] = stripQuotes(obj.S("soa", "ttl").String())
	if value, ok := d.GetOk("soa"); ok {
		tp := value.(map[string]interface{})
		if tp["email"] != nil {
			soaset["email"] = stripQuotes(obj.S("soa", "email").String())
		}
	}
	soaset["refresh"] = stripQuotes(obj.S("soa", "refresh").String())
	soaset["expire"] = stripQuotes(obj.S("soa", "expire").String())
	soaset["retry"] = stripQuotes(obj.S("soa", "retry").String())
	soaset["negcache"] = stripQuotes(obj.S("soa", "negCache").String())

	d.Set("id", stripQuotes(obj.S("id").String()))
	d.Set("name", stripQuotes(obj.S("name").String()))
	d.Set("soa", soaset)
	if hasGeoIP, err := strconv.ParseBool(stripQuotes(obj.S("hasGeoIP").String())); err == nil {
		d.Set("has_geoip", hasGeoIP)
	}
	if hasGTDRegion, err := strconv.ParseBool(stripQuotes(obj.S("hasGtdRegions").String())); err == nil {
		d.Set("has_gtd_regions", hasGTDRegion)
	}
	if obj.Exists("vanityNameServer") {
		d.Set("vanity_nameserver", stripQuotes(obj.S("vanityNameServer").String()))
	}
	if obj.Exists("nameserverGroup") {
		d.Set("nameserver_group", stripQuotes(obj.S("nameserverGroup").String()))
	}
	d.Set("note", stripQuotes(obj.S("note").String()))

	if obj.S("tags").Data() != nil {
		d.Set("tags", toListOfString(obj.S("tags").Data()))
	} else {
		d.Set("tags", make([]string, 0, 1))
	}

	return nil
}

func resourceConstellixDNSUpdate(d *schema.ResourceData, m interface{}) error {
	constellixClient := m.(*client.Client)

	domainAttr := models.DomainAttributes{}

	domainAttr.HasGtdRegions = d.Get("has_gtd_regions").(bool)

	domainAttr.HasGeoIP = d.Get("has_geoip").(bool)

	if vanityNS, ok := d.GetOk("vanity_nameserver"); ok {
		domainAttr.VanityNameServer = vanityNS.(string)
	}

	if _, ok := d.GetOk("nameserver_group"); ok {
		domainAttr.NameserverGroup = d.Get("nameserver_group").(string)
	}

	if _, ok := d.GetOk("note"); ok {
		domainAttr.Note = d.Get("note").(string)
	}

	if _, ok := d.GetOk("tags"); ok {
		tagsList := d.Get("tags").([]interface{})
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
		soaAttr.TTL = tp["ttl"].(string)
	}
	if tp["expire"] != nil {
		soaAttr.Expire = tp["expire"].(string)
	}
	if tp["negcache"] != nil {
		soaAttr.NegCache = tp["negcache"].(string)
	}
	if tp["refresh"] != nil {
		soaAttr.Refresh = tp["refresh"].(string)
	}
	if tp["retry"] != nil {
		soaAttr.Retry = tp["retry"].(string)
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
