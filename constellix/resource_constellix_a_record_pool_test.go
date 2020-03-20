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

func TestAccARecordPool_Basic(t *testing.T) {
	var arp models.ARecordPoolAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixARecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixARecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixARecordPoolExists("constellix_a_record_pool.ap1", &arp),
					testAccCheckConstellixARecordPoolAttributes(1, &arp),
				),
			},
		},
	})
}

func TestAccConstellixARecordPool_Update(t *testing.T) {
	var arp models.ARecordPoolAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixARecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixARecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixARecordPoolExists("constellix_a_record_pool.ap1", &arp),
					testAccCheckConstellixARecordPoolAttributes(1, &arp),
				),
			},
			{
				Config: testAccCheckConstellixARecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixARecordPoolExists("constellix_a_record_pool.ap1", &arp),
					testAccCheckConstellixARecordPoolAttributes(1, &arp),
				),
			},
		},
	})
}

func testAccCheckConstellixARecordPoolConfig_basic(numreturn int) string {
	return fmt.Sprintf(`
	resource "constellix_a_record_pool" "ap1"{
		name = "temparecordpool"
		num_return = "%d"
		min_available_failover = 1
		values {
			value = "8.1.1.1"
			weight = 20
			policy = "followsonar"
		}
		values {
			value = "8.2.1.1"
			weight = 20
			policy = "followsonar"
		}
	}
	`, numreturn)
}

func testAccCheckConstellixARecordPoolExists(arecordpoolName string, arp *models.ARecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs2, err2 := s.RootModule().Resources[arecordpoolName]

		if !err2 {
			return fmt.Errorf("ARecordPool record %s not found", arecordpoolName)
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No ARecordPool record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/pools/A/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := arecordpoolfromcontainer(resp)

		*arp = *tp
		return nil
	}
}

func arecordpoolfromcontainer(resp *http.Response) (*models.ARecordPoolAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	arp := models.ARecordPoolAttributes{}
	arp.Name = fmt.Sprintf("%v", data["name"])
	arp.NumReturn, _ = strconv.Atoi(fmt.Sprintf("%v", data["numReturn"]))
	arp.MinAvailableFailover, _ = strconv.Atoi(fmt.Sprintf("%v", data["minAvailableFailover"]))

	return &arp, nil

}

func testAccCheckConstellixARecordPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_arecordpool" {
			_, err := client.GetbyId("v1/pools/A/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("ARecordPool record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckConstellixARecordPoolAttributes(numreturn interface{}, arp *models.ARecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temparecordpool" != arp.Name {
			return fmt.Errorf("Bad Hinfo record name %s", arp.Name)
		}
		numreturn, _ := strconv.Atoi(fmt.Sprintf("%v", numreturn))
		if numreturn != arp.NumReturn {
			return fmt.Errorf("Bad ARecordPool record Numreturn %d", arp.NumReturn)
		}
		return nil
	}
}
