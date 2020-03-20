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

func TestAccSPF_Basic(t *testing.T) {
	var spf models.SpfAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixSPFDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixSPFConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSPFExists("constellix_domain.domainSPF", "constellix_spf_record.spf1", &spf),
					testAccCheckConstellixSPFAttributes(1800, &spf),
				),
			},
		},
	})
}

func TestAccConstellixSPF_Update(t *testing.T) {
	var spf models.SpfAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixSPFDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixSPFConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSPFExists("constellix_domain.domainSPF", "constellix_spf_record.spf1", &spf),
					testAccCheckConstellixSPFAttributes(1800, &spf),
				),
			},
			{
				Config: testAccCheckConstellixSPFConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSPFExists("constellix_domain.domainSPF", "constellix_spf_record.spf1", &spf),
					testAccCheckConstellixSPFAttributes(1900, &spf),
				),
			},
		},
	})
}

func testAccCheckConstellixSPFConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domainSPF" {
		name = "checkspf.com"
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

	resource "constellix_spf_record" "spf1" {
		domain_id = "${constellix_domain.domainSPF.id}"
		source_type = "domains"
		name = "tempspfrecord"
		ttl = %d
		roundrobin{
		  value = "1.2.3.5"
		  disable_flag = "false"
		}
		roundrobin{
		  value = "124.56.8.1"
		  disable_flag = "true"
		}
	  }
	`, ttl)
}

func testAccCheckConstellixSPFExists(domainName string, SPFName string, model *models.SpfAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[SPFName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("SPF record %s not found", SPFName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No SPF record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/spf/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := spffromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixSPFAttributes(ttl interface{}, model *models.SpfAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempspfrecord" != model.Name {
			return fmt.Errorf("Bad SPF record name %s", model.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != model.TTL {
			return fmt.Errorf("Bad SPF record ttl %d", model.TTL)
		}

		return nil
	}
}

func testAccCheckConstellixSPFDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domainSPF"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domainSPF")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_spf_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/spf/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("SPF record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func spffromcontainer(resp *http.Response) (*models.SpfAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	model := models.SpfAttributes{}

	model.Name = fmt.Sprintf("%v", data["name"])
	model.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))

	return &model, nil
}
