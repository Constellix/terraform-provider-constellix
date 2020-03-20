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

func TestAccCNameRecordPool_Basic(t *testing.T) {
	var crp models.CnameRecordPoolAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCNameRecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCNameRecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCNameRecordPoolExists("constellix_cname_record_pool.cp1", &crp),
					testAccCheckConstellixCNameRecordPoolAttributes(1, &crp),
				),
			},
		},
	})
}

func TestAccConstellixCNameRecordPool_Update(t *testing.T) {
	var crp models.CnameRecordPoolAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixCNameRecordPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixCNameRecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCNameRecordPoolExists("constellix_cname_record_pool.cp1", &crp),
					testAccCheckConstellixCNameRecordPoolAttributes(1, &crp),
				),
			},
			{
				Config: testAccCheckConstellixCNameRecordPoolConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixCNameRecordPoolExists("constellix_cname_record_pool.cp1", &crp),
					testAccCheckConstellixCNameRecordPoolAttributes(1, &crp),
				),
			},
		},
	})
}

func testAccCheckConstellixCNameRecordPoolConfig_basic(numreturn int) string {
	return fmt.Sprintf(`
	resource "constellix_cname_record_pool" "cp1"{
		name = "tempcnamerecordpool"
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

func testAccCheckConstellixCNameRecordPoolExists(arecordpoolName string, crp *models.CnameRecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs2, err2 := s.RootModule().Resources[arecordpoolName]

		if !err2 {
			return fmt.Errorf("CNameRecordPool record %s not found", arecordpoolName)
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No CNameRecordPool record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/pools/CNAME/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := cnamerecordpoolfromcontainer(resp)

		*crp = *tp
		return nil
	}
}

func cnamerecordpoolfromcontainer(resp *http.Response) (*models.CnameRecordPoolAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	crp := models.CnameRecordPoolAttributes{}

	crp.Name = fmt.Sprintf("%v", data["name"])
	crp.NumReturn, _ = strconv.Atoi(fmt.Sprintf("%v", data["numReturn"]))
	crp.MinAvailableFailover, _ = strconv.Atoi(fmt.Sprintf("%v", data["minAvailableFailover"]))

	return &crp, nil

}

func testAccCheckConstellixCNameRecordPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_cname_record_pool" {
			_, err := client.GetbyId("v1/pools/CNAME/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("CNameRecordPool record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
func testAccCheckConstellixCNameRecordPoolAttributes(numreturn interface{}, crp *models.CnameRecordPoolAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempcnamerecordpool" != crp.Name {
			return fmt.Errorf("Bad cname record pool name %s", crp.Name)
		}
		numreturn, _ := strconv.Atoi(fmt.Sprintf("%v", numreturn))
		if numreturn != crp.NumReturn {
			return fmt.Errorf("Bad CNameRecordPool record Numreturn %d", crp.NumReturn)
		}
		return nil
	}
}
