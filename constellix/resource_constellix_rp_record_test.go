package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRP_Basic(t *testing.T) {
	var rp models.RPAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixRPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixRPConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixRPExists("constellix_domain.domain1", "constellix_rp_record.rp1", &rp),
					testAccCheckConstellixRPAttributes(1800, &rp),
				),
			},
		},
	})
}

func TestAccConstellixRP_Update(t *testing.T) {
	var rp models.RPAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixRPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixRPConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixRPExists("constellix_domain.domain1", "constellix_rp_record.rp1", &rp),
					testAccCheckConstellixRPAttributes(1800, &rp),
				),
			},
			{
				Config: testAccCheckConstellixRPConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixRPExists("constellix_domain.domain1", "constellix_rp_record.rp1", &rp),
					testAccCheckConstellixRPAttributes(1900, &rp),
				),
			},
		},
	})
}

func testAccCheckConstellixRPConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkrp.com"
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

	resource "constellix_rp_record" "rp1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "temprprecord"
		ttl = "%d"
		
		roundrobin {
			mailbox = "mx.com"
			txt = "hello"
		}
	}
	`, ttl)
}

func testAccCheckConstellixRPExists(domainName string, RPName string, rp *models.RPAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[RPName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("RP record %s not found", RPName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No RP record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/rp/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := rpfromcontainer(resp)

		*rp = *tp
		return nil
	}
}

func testAccCheckConstellixRPAttributes(ttl interface{}, rp *models.RPAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temprprecord" != rp.Name {
			return fmt.Errorf("Bad RP record name %s", rp.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != rp.TTL {
			return fmt.Errorf("Bad RP record ttl %d", rp.TTL)
		}
		return nil
	}
}

func testAccCheckConstellixRPDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_rp_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/rp/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("RP record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func rpfromcontainer(resp *http.Response) (*models.RPAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body string : %v", bodyString)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	rp := models.RPAttributes{}

	rp.Name = fmt.Sprintf("%v", data["name"])
	rp.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	return &rp, nil

}
