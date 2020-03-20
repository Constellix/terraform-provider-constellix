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

func TestAccPtr_Basic(t *testing.T) {
	var ptr models.PtrAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixPtrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixPtrConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixPtrExists("constellix_domain.domain1", "constellix_ptr_record.ptr1", &ptr),
					testAccCheckConstellixPtrAttributes(1800, &ptr),
				),
			},
		},
	})
}

func TestAccConstellixPtr_Update(t *testing.T) {
	var ptr models.PtrAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixPtrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixPtrConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixPtrExists("constellix_domain.domain1", "constellix_ptr_record.ptr1", &ptr),
					testAccCheckConstellixPtrAttributes(1800, &ptr),
				),
			},
			{
				Config: testAccCheckConstellixPtrConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixPtrExists("constellix_domain.domain1", "constellix_ptr_record.ptr1", &ptr),
					testAccCheckConstellixPtrAttributes(1900, &ptr),
				),
			},
		},
	})
}

func testAccCheckConstellixPtrConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkptr.com"
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

	resource "constellix_ptr_record" "ptr1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempptrrecord"
		ttl = "%d"
		roundrobin {
			value = 13
			disable_flag = "true"
		}
		roundrobin {
			value = 14
			disable_flag = "false"
		}
	}
	`, ttl)
}

func testAccCheckConstellixPtrExists(domainName string, ptrName string, ptr *models.PtrAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[ptrName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Ptr record %s not found", ptrName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Ptr record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/ptr/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := ptrfromcontainer(resp)

		*ptr = *tp
		return nil
	}
}

func ptrfromcontainer(resp *http.Response) (*models.PtrAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	ptr := models.PtrAttributes{}

	ptr.Name = fmt.Sprintf("%v", data["name"])
	ptr.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	resrr := (data["roundRobin"]).([]interface{})
	mapListRR := make([]interface{}, 0, 1)
	for _, val := range resrr {
		log.Println("RR are : ", val)
		tpMap := make(map[string]interface{})
		inner := val.(map[string]interface{})
		tpMap["value"], _ = strconv.Atoi(fmt.Sprintf("%d", inner["value"]))
		mapListRR = append(mapListRR, tpMap)
	}
	ptr.RoundRobin = mapListRR

	return &ptr, nil

}

func testAccCheckConstellixPtrDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_ptr_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/ptr/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Ptr record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
func testAccCheckConstellixPtrAttributes(ttl interface{}, ptr *models.PtrAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempptrrecord" != ptr.Name {
			return fmt.Errorf("Bad Ptr record name %s", ptr.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != ptr.TTL {
			return fmt.Errorf("Bad Ptr record ttl %d", ptr.TTL)
		}
		return nil
	}
}
