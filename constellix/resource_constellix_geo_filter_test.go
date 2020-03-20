package constellix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/Constellix/constellix-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIpFilter_Basic(t *testing.T) {
	var htp models.IPFilterAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixIpFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixIpFilterConfig_basic("1.1.1.0/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixIpFilterExists("constellix_geo_filter.ip1", &htp),
					testAccCheckConstellixIpFilterAttributes(&htp),
				),
			},
		},
	})
}

func TestAccConstellixIpFilter_Update(t *testing.T) {
	var model models.IPFilterAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixIpFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixIpFilterConfig_basic("1.1.1.0/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixIpFilterExists("constellix_geo_filter.ip1", &model),
					testAccCheckConstellixIpFilterAttributes(&model),
				),
			},
			{
				Config: testAccCheckConstellixIpFilterConfig_basic("1.2.0.0/32"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixIpFilterExists("constellix_geo_filter.ip1", &model),
					testAccCheckConstellixIpFilterAttributes(&model),
				),
			},
		},
	})
}

func testAccCheckConstellixIpFilterConfig_basic(tp string) string {
	return fmt.Sprintf(`

	resource "constellix_geo_filter" "ip1" {
	
		name ="ipfilter3"
  		geoip_continents = ["AS"]
  		geoip_countries = ["IN", "PK"]
  		geoip_regions = ["IN/BR", "IN/MP"]
  		asn = [1,2]
  		ipv4 = ["%v"]
  		ipv6 = ["2:0:0:2:0:0:1:abc/128"]
		
	}
	`, tp)
}

func testAccCheckConstellixIpFilterExists(hinfoName string, model *models.IPFilterAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs2, err2 := s.RootModule().Resources[hinfoName]
		if !err2 {
			return fmt.Errorf("IpFilter record %s not found", hinfoName)
		}

		if rs2.Primary.ID == "" {
			return fmt.Errorf("No ipfilter record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/geoFilters/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := ipfilterfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixIpFilterAttributes(model *models.IPFilterAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "ipfilter3" != model.Name {
			return fmt.Errorf("Bad IpFilter record name %s", model.Name)
		}

		return nil
	}
}

func testAccCheckConstellixIpFilterDestroy(s *terraform.State) error {
	return nil
}

func ipfilterfromcontainer(resp *http.Response) (*models.IPFilterAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	model := models.IPFilterAttributes{}

	model.Name = fmt.Sprintf("%v", data["name"])

	return &model, nil

}
