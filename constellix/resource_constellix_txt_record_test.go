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

func TestAccTxt_Basic(t *testing.T) {
	var tx models.TxtAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTxtDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTxtConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTxtExists("constellix_domain.domain1", "constellix_txt_record.tx1", &tx),
					testAccCheckConstellixTxtAttributes(1800, &tx),
				),
			},
		},
	})
}

func TestAccConstellixTxt_Update(t *testing.T) {
	var tx models.TxtAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTxtDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTxtConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTxtExists("constellix_domain.domain1", "constellix_txt_record.tx1", &tx),
					testAccCheckConstellixTxtAttributes(1800, &tx),
				),
			},
			{
				Config: testAccCheckConstellixTxtConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTxtExists("constellix_domain.domain1", "constellix_txt_record.tx1", &tx),
					testAccCheckConstellixTxtAttributes(1900, &tx),
				),
			},
		},
	})
}

func testAccCheckConstellixTxtConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checktxt.com"
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

	resource "constellix_txt_record" "tx1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "temptxtrecord"
		ttl = "%d"
		roundrobin {
			value = "mail.com."
			disable_flag = "true"
		}
		roundrobin {
			value = "google.com."
			disable_flag = "false"
		}
	}
	`, ttl)
}

func testAccCheckConstellixTxtExists(domainName string, txtName string, model *models.TxtAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[txtName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Txt record %s not found", txtName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Txt record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/txt/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := txtfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func txtfromcontainer(resp *http.Response) (*models.TxtAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	tx := models.TxtAttributes{}
	tx.Name = fmt.Sprintf("%v", data["name"])
	tx.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	resrr := (data["roundRobin"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		log.Println("RR are : ", val)
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"] = fmt.Sprintf("%d", inner["value"])
		mapListRR = append(mapListRR, tpMap)
	}
	tx.RoundRobin = mapListRR

	return &tx, nil

}

func testAccCheckConstellixTxtDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_txt_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/txt/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Txt record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
func testAccCheckConstellixTxtAttributes(ttl interface{}, tx *models.TxtAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temptxtrecord" != tx.Name {
			return fmt.Errorf("Bad Txt record name %s", tx.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != tx.TTL {
			return fmt.Errorf("Bad Txt record ttl %d", tx.TTL)
		}
		return nil
	}
}
