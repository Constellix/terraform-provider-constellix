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

func TestAccAAAARecordPool_Basic(t *testing.T) {
	var aaaarp models.AAAArecordPoolAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAAAARecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAAAARecordPoolConfig_basic(20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAAAARecordPoolExists("constellix_aaaa_record_pool.aaaap1", &aaaarp),
					testAccCheckConstellixAAAARecordPoolAttributes(20, &aaaarp),
				),
			},
		},
	})
}

func TestAccConstellixAAAARecordPool_Update(t *testing.T) {
	var aaaarp models.AAAArecordPoolAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAAAARecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAAAARecordPoolConfig_basic(20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAAAARecordPoolExists("constellix_aaaa_record_pool.aaaap1", &aaaarp),
					testAccCheckConstellixAAAARecordPoolAttributes(20, &aaaarp),
				),
			},
			{
				Config: testAccCheckConstellixAAAARecordPoolConfig_basic(30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAAAARecordPoolExists("constellix_aaaa_record_pool.aaaap1", &aaaarp),
					testAccCheckConstellixAAAARecordPoolAttributes(30, &aaaarp),
				),
			},
		},
	})
}

func testAccCheckConstellixAAAARecordPoolConfig_basic(numreturn int) string {
	return fmt.Sprintf(`
	resource "constellix_aaaa_record_pool" "aaaap1"{
		name = "tempaaaarecordpool"
		num_return = 1
		min_available_failover = 1
		values {
			value = "0:0:0:0:0:0:0:1"
			weight = "%d"
			policy = "followsonar"
		}
		values {
			value = "0:0:0:0:0:0:0:12"
			weight = 20
			policy = "followsonar"
		}
		note = "hello"
	}
	`, numreturn)
}

func testAccCheckConstellixAAAARecordPoolExists(arecordpoolName string, arp *models.AAAArecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs2, err2 := s.RootModule().Resources[arecordpoolName]

		if !err2 {
			return fmt.Errorf("AAAARecordPool record %s not found", arecordpoolName)
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No AAAARecordPool record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/pools/AAAA/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := aaaarecordpoolfromcontainer(resp)

		*arp = *tp
		return nil
	}
}

func aaaarecordpoolfromcontainer(resp *http.Response) (*models.AAAArecordPoolAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	arp := models.AAAArecordPoolAttributes{}
	arp.Name = fmt.Sprintf("%v", data["name"])
	arp.NumReturn, _ = strconv.Atoi(fmt.Sprintf("%v", data["numReturn"]))
	arp.MinavailFailover, _ = strconv.Atoi(fmt.Sprintf("%v", data["minAvailableFailover"]))
	arp.Note = fmt.Sprintf("%v", data["note"])

	return &arp, nil

}

func testAccCheckConstellixAAAARecordPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_arecordpool" {
			_, err := client.GetbyId("v1/pools/AAAA/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("AAAARecordPool record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckConstellixAAAARecordPoolAttributes(numreturn interface{}, arp *models.AAAArecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempaaaarecordpool" != arp.Name {
			return fmt.Errorf("Bad AAAA pool name %s", arp.Name)
		}
		if "hello" != arp.Note {
			return fmt.Errorf("Bad AAAARecordPool note %s", arp.Note)
		}
		return nil
	}
}
