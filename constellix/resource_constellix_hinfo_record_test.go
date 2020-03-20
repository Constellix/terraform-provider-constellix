package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHinfo_Basic(t *testing.T) {
	var hinfo models.HinfoAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHinfoConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHinfoExists("constellix_domain.domain1", "constellix_hinfo_record.hinfo1", &hinfo),
					testAccCheckConstellixHinfoAttributes(1800, &hinfo),
				),
			},
		},
	})
}

func TestAccConstellixHinfo_Update(t *testing.T) {
	var hinfo models.HinfoAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHinfoConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHinfoExists("constellix_domain.domain1", "constellix_hinfo_record.hinfo1", &hinfo),
					testAccCheckConstellixHinfoAttributes(1800, &hinfo),
				),
			},
			{
				Config: testAccCheckConstellixHinfoConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHinfoExists("constellix_domain.domain1", "constellix_hinfo_record.hinfo1", &hinfo),
					testAccCheckConstellixHinfoAttributes(1900, &hinfo),
				),
			},
		},
	})
}

func testAccCheckConstellixHinfoConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkhinfo.com"
		soa = {
			email = "com.com."
			primary_nameserver = "ns41.constellix.com."
			ttl = 1900
			refresh = 48100
			retry = 7200
			expire = 1209
			negcache = 8000
		}
	}

	resource "constellix_hinfo_record" "hinfo1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "temphinforecord"
		ttl = "%d"
		
		roundrobin {
			cpu = "quard core"
			os = "linux2"
			disable_flag = "false"
		}
		roundrobin{
			cpu = "abc"
			os = "windows"
			disable_flag = "true"
		}
	}
	`, ttl)
}

func testAccCheckConstellixHinfoExists(domainName string, hinfoName string, hinfo *models.HinfoAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[hinfoName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Hinfo record %s not found", hinfoName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Hinfo record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/hinfo/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := hinfofromcontainer(resp)

		*hinfo = *tp
		return nil
	}
}

func testAccCheckConstellixHinfoAttributes(ttl interface{}, hinfo *models.HinfoAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temphinforecord" != hinfo.Name {
			return fmt.Errorf("Bad Hinfo record name %s", hinfo.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != hinfo.TTL {
			return fmt.Errorf("Bad Hinfo record ttl %d", hinfo.TTL)
		}
		return nil
	}
}

func testAccCheckConstellixHinfoDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_hinfo_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/hinfo/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Hinfo record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func hinfofromcontainer(resp *http.Response) (*models.HinfoAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	hinfo := models.HinfoAttributes{}
	hinfo.Name = fmt.Sprintf("%v", data["name"])
	hinfo.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	hinfo.Note = fmt.Sprintf("%v", data["note"])

	return &hinfo, nil

}
