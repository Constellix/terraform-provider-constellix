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

func TestAccMX_Basic(t *testing.T) {
	var mx models.MXAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixMXDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixMXConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixMXExists("constellix_domain.domain1", "constellix_mx_record.mx1", &mx),
					testAccCheckConstellixMXAttributes(1800, &mx),
				),
			},
		},
	})
}

func TestAccConstellixMx_Update(t *testing.T) {
	var mx models.MXAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixMXDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixMXConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixMXExists("constellix_domain.domain1", "constellix_mx_record.mx1", &mx),
					testAccCheckConstellixMXAttributes(1800, &mx),
				),
			},
			{
				Config: testAccCheckConstellixMXConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixMXExists("constellix_domain.domain1", "constellix_mx_record.mx1", &mx),
					testAccCheckConstellixMXAttributes(1900, &mx),
				),
			},
		},
	})
}

func testAccCheckConstellixMXConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkmx.com"
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

	resource "constellix_mx_record" "mx1"{
		domain_id = "${constellix_domain.domain1.id}"
	 source_type = "domains"
		name = "tempmxrecord"
		ttl = "%d"
		
		roundrobin {
			value = "abc"
			level = "100"
			disable_flag = false
		}
	}
	`, ttl)
}

func testAccCheckConstellixMXExists(domainName string, MXName string, mx *models.MXAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[MXName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("MX record %s not found", MXName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No MX record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/mx/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := mxfromcontainer(resp)

		*mx = *tp
		return nil
	}
}

func testAccCheckConstellixMXAttributes(ttl interface{}, mx *models.MXAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempmxrecord" != mx.Name {
			return fmt.Errorf("Bad MX record name %s", mx.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != mx.TTL {
			return fmt.Errorf("Bad MX record ttl %d", mx.TTL)
		}
		return nil
	}
}

func testAccCheckConstellixMXDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_mx_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/mx/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("MX record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func mxfromcontainer(resp *http.Response) (*models.MXAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	mx := models.MXAttributes{}

	mx.Name = fmt.Sprintf("%v", data["name"])
	mx.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))

	return &mx, nil

}
