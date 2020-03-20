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

func TestAccCert_Basic(t *testing.T) {
	var ct models.CertAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCertConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCertExists("constellix_domain.domain1", "constellix_cert_record.cert1", &ct),
					testAccCheckConstellixCertAttributes(1800, &ct),
				),
			},
		},
	})
}

func TestAccConstellixCert_Update(t *testing.T) {
	var ct models.CertAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCertConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCertExists("constellix_domain.domain1", "constellix_cert_record.cert1", &ct),
					testAccCheckConstellixCertAttributes(1800, &ct),
				),
			},
			{
				Config: testAccCheckConstellixCertConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCertExists("constellix_domain.domain1", "constellix_cert_record.cert1", &ct),
					testAccCheckConstellixCertAttributes(1900, &ct),
				),
			},
		},
	})
}

func testAccCheckConstellixCertConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkcert.com"
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

	resource "constellix_cert_record" "cert1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempcert"
		ttl = "%d"
		roundrobin {
			certificate_type = 20
			key_tag = 30
			algorithm = 100
			disable_flag = "true"
			certificate = "certificate1"
		}
		roundrobin {
			certificate_type = 40
			key_tag = 62
			certificate = "certificate1"
			algorithm = 45
			disable_flag = "false"
		}
	}
	`, ttl)
}

func testAccCheckConstellixCertExists(domainName string, certName string, ct *models.CertAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[certName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Cert record %s not found", certName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Cert record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/cert/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := certfromcontainer(resp)

		*ct = *tp
		return nil
	}
}

func certfromcontainer(resp *http.Response) (*models.CertAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	ct := models.CertAttributes{}

	ct.Name = fmt.Sprintf("%v", data["name"])
	ct.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	ct.Note = fmt.Sprintf("%v", data["note"])
	resrr := (data["roundRobin"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["certificatetype"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["certificateType"]))
		mapListRR = append(mapListRR, tpMap)
	}
	ct.RoundRobin = mapListRR

	return &ct, nil

}

func testAccCheckConstellixCertDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_cert_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/cert/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Cert record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
func testAccCheckConstellixCertAttributes(ttl interface{}, ct *models.CertAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempcert" != ct.Name {
			return fmt.Errorf("Bad Cert record name %s", ct.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != ct.TTL {
			return fmt.Errorf("Bad Cert record ttl %d", ct.TTL)
		}

		return nil
	}
}
