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

func TestAccCaa_Basic(t *testing.T) {
	var caa models.CaaAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCaaRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCaaRecordConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCaaRecordExists("constellix_domain.domain1", "constellix_caa_record.caa1", &caa),
					testAccCheckConstellixCaaRecordAttributes(1800, &caa),
				),
			},
		},
	})
}

func TestAccConstellixCaaRecord_Update(t *testing.T) {
	var caa models.CaaAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCaaRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCaaRecordConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCaaRecordExists("constellix_domain.domain1", "constellix_caa_record.caa1", &caa),
					testAccCheckConstellixCaaRecordAttributes(1800, &caa),
				),
			},
			{
				Config: testAccCheckConstellixCaaRecordConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCaaRecordExists("constellix_domain.domain1", "constellix_caa_record.caa1", &caa),
					testAccCheckConstellixCaaRecordAttributes(1900, &caa),
				),
			},
		},
	})
}

func testAccCheckConstellixCaaRecordConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkcaarecord.com"
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

	resource "constellix_caa_record" "caa1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempcaarecord"
		ttl = "%d"
		roundrobin{
			caa_provider_id = 3
			tag = "issue"
			data = "como.com"
			flag = "0"
			disable_flag = "false"
		}
		roundrobin{
			caa_provider_id = 4
			tag = "issue"
			data = "como01.com"
			flag = "1"
			disable_flag = "true"
		}
	}
	`, ttl)
}

func testAccCheckConstellixCaaRecordExists(domainName string, caaName string, caa *models.CaaAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[caaName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Caa record %s not found", caaName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Caa record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/caa/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := caarecordfromcontainer(resp)

		*caa = *tp
		return nil
	}
}

func testAccCheckConstellixCaaRecordAttributes(ttl interface{}, caa *models.CaaAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempcaarecord" != caa.Name {
			return fmt.Errorf("Bad Caa record name %s", caa.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != caa.TTL {
			return fmt.Errorf("Bad Caa record ttl %d", caa.TTL)
		}

		return nil
	}
}

func testAccCheckConstellixCaaRecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_caa_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/caa/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Caa record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func caarecordfromcontainer(resp *http.Response) (*models.CaaAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	caa := models.CaaAttributes{}

	caa.Name = fmt.Sprintf("%v", data["name"])
	caa.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	caa.Note = fmt.Sprintf("%v", data["note"])

	return &caa, nil

}
