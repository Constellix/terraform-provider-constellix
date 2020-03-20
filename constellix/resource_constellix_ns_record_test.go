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

func TestAccNs_Basic(t *testing.T) {
	var ns models.NSAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixNsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixNsConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNsExists("constellix_domain.domain1", "constellix_ns_record.ns1", &ns),
					testAccCheckConstellixNsAttributes(1800, &ns),
				),
			},
		},
	})
}

func TestAccConstellixNs_Update(t *testing.T) {
	var ns models.NSAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixNsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixNsConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNsExists("constellix_domain.domain1", "constellix_ns_record.ns1", &ns),
					testAccCheckConstellixNsAttributes(1800, &ns),
				),
			},
			{
				Config: testAccCheckConstellixNsConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNsExists("constellix_domain.domain1", "constellix_ns_record.ns1", &ns),
					testAccCheckConstellixNsAttributes(1900, &ns),
				),
			},
		},
	})
}

func testAccCheckConstellixNsConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkns.com"
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

	resource "constellix_ns_record" "ns1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempnsrecord"
		ttl = "%d"
		note = "Practice record naptr"
		roundrobin {
			       value = "f5."
			       disable_flag = "false"
			   }
	}
	`, ttl)
}

func testAccCheckConstellixNsExists(domainName string, nsName string, ns *models.NSAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[nsName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Ns record %s not found", nsName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Ns record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/ns/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := nsfromcontainer(resp)

		*ns = *tp
		return nil
	}
}

func testAccCheckConstellixNsAttributes(ttl interface{}, ns *models.NSAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempnsrecord" != ns.Name {
			return fmt.Errorf("Bad Ns record name %s", ns.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != ns.Ttl {
			return fmt.Errorf("Bad Ns record ttl %d", ns.Ttl)
		}
		return nil
	}
}

func testAccCheckConstellixNsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_ns_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/ns/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Ns record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func nsfromcontainer(resp *http.Response) (*models.NSAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	ns := models.NSAttributes{}

	ns.Name = fmt.Sprintf("%v", data["name"])
	ns.Ttl, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	ns.Note = fmt.Sprintf("%v", data["note"])

	return &ns, nil

}
