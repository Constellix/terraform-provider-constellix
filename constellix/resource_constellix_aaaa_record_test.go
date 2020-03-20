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

func TestAccAaaa_Basic(t *testing.T) {
	var aaaa models.AAAARecordAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAaaaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAaaaConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAaaaExists("constellix_domain.domain1", "constellix_aaaa_record.aaaa1", &aaaa),
					testAccCheckConstellixAaaaAttributes("1800", &aaaa),
				),
			},
		},
	})
}

func TestAccConstellixAaaa_Update(t *testing.T) {
	var model models.AAAARecordAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAaaaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAaaaConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAaaaExists("constellix_domain.domain1", "constellix_aaaa_record.aaaa1", &model),
					testAccCheckConstellixAaaaAttributes("1800", &model),
				),
			},
			{
				Config: testAccCheckConstellixAaaaConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAaaaExists("constellix_domain.domain1", "constellix_aaaa_record.aaaa1", &model),
					testAccCheckConstellixAaaaAttributes("1900", &model),
				),
			},
		},
	})
}

func testAccCheckConstellixAaaaConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkaaaa.com"
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

	resource "constellix_aaaa_record" "aaaa1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempaaaarecord"
		ttl = "%d"
		
		roundrobin{
			    value       = "5:0:0:0:0:0:0:6"
			    disable_flag = "false"
				}
		roundrobin{
				value = "6:0:0:0:0:0:0:8"
				disable_flag = "true"
				}
	}
	`, ttl)
}

func testAccCheckConstellixAaaaExists(domainName string, aaaaName string, model *models.AAAARecordAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[aaaaName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Aaaa record %s not found", aaaaName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Aaaa record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/aaaa/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := aaaarecordfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixAaaaAttributes(ttl interface{}, model *models.AAAARecordAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempaaaarecord" != model.Name {
			return fmt.Errorf("Bad Aaaa record name %s", model.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != model.TTL {
			return fmt.Errorf("Bad Aaaa record ttl value %d", model.TTL)
		}

		return nil
	}
}

func testAccCheckConstellixAaaaDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_aaaa_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/aaaa/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Aaaa record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func aaaarecordfromcontainer(resp *http.Response) (*models.AAAARecordAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	model := models.AAAARecordAttributes{}
	model.Name = fmt.Sprintf("%v", data["name"])
	model.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	model.Note = fmt.Sprintf("%v", data["note"])

	return &model, nil

}
