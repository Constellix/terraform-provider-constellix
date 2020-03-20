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

func TestAccSRV_Basic(t *testing.T) {
	var srv models.SRVAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixSRVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixSRVConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSRVExists("constellix_domain.domain1", "constellix_srv_record.srvrecord1", &srv),
					testAccCheckConstellixSRVAttributes("1800", &srv),
				),
			},
		},
	})
}

func TestAccConstellixSRV_Update(t *testing.T) {
	var model models.SRVAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixSRVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixSRVConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSRVExists("constellix_domain.domain1", "constellix_srv_record.srvrecord1", &model),
					testAccCheckConstellixSRVAttributes("1800", &model),
				),
			},
			{
				Config: testAccCheckConstellixSRVConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixSRVExists("constellix_domain.domain1", "constellix_srv_record.srvrecord1", &model),
					testAccCheckConstellixSRVAttributes("1900", &model),
				),
			},
		},
	})

}

func testAccCheckConstellixSRVConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "domaintest1.com"
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

	resource "constellix_srv_record" "srvrecord1" {
		domain_id="${constellix_domain.domain1.id}"
		ttl="%d"
		source_type = "domains"
		roundrobin {
			value="www.abc.com"
			port=8888
			priority=345
			weight=2
			disable_flag=false
		}
	}
	`, ttl)
}

func testAccCheckConstellixSRVExists(domainName string, srvName string, model *models.SRVAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[srvName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}
		if !err2 {
			return fmt.Errorf("Aname record %s not found", srvName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No srv record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/srv/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := srvfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixSRVAttributes(ttl interface{}, model *models.SRVAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println()
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != model.TTL {
			return fmt.Errorf("Bad srv record %d", model.TTL)
		}

		return nil
	}
}

func testAccCheckConstellixSRVDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_srv_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/srv/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("srv record still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func srvfromcontainer(resp *http.Response) (*models.SRVAttributes, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body string : %v", bodyString)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	log.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD: ", &data)
	model := models.SRVAttributes{}

	model.Name = fmt.Sprintf("%v", data["name"])
	model.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	model.Note = fmt.Sprintf("%v", data["note"])
	return &model, nil
}
