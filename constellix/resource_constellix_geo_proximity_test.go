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

func TestAccGeoProximity_Basic(t *testing.T) {
	var gp models.GeoProximityAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixGeoProximityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixGeoProximityConfig_basic(273890),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixGeoProximityExists("constellix_geo_proximity.gp1", &gp),
					testAccCheckConstellixGeoProximityAttributes(&gp),
				),
			},
		},
	})
}

func TestAccConstellixGeoProximity_Update(t *testing.T) {
	var gp models.GeoProximityAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixGeoProximityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixGeoProximityConfig_basic(0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixGeoProximityExists("constellix_geo_proximity.gp1", &gp),
					testAccCheckConstellixGeoProximityAttributes(&gp),
				),
			},
			{
				Config: testAccCheckConstellixGeoProximityConfig_basic(273890),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixGeoProximityExists("constellix_geo_proximity.gp1", &gp),
					testAccCheckConstellixGeoProximityAttributes(&gp),
				),
			},
		},
	})
}

func testAccCheckConstellixGeoProximityConfig_basic(tp int) string {
	return fmt.Sprintf(`

	resource "constellix_geo_proximity" "gp1" {
	
		name = "tempgeoproximityrecordtestcheck"
		country = "OM"
		city = "%v"
	}
	`, tp)
}

func testAccCheckConstellixGeoProximityExists(gpName string, gp *models.GeoProximityAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs2, err2 := s.RootModule().Resources[gpName]
		if !err2 {
			return fmt.Errorf("GeoProximity record %s not found", gpName)
		}

		if rs2.Primary.ID == "" {
			return fmt.Errorf("No geoproximity record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/geoProximities/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := geoproximityfromcontainer(resp)

		*gp = *tp
		return nil
	}
}

func testAccCheckConstellixGeoProximityAttributes(gp *models.GeoProximityAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempgeoproximityrecordtestcheck" != gp.Name {
			return fmt.Errorf("Bad GeoProximity record name %s", gp.Name)
		}
		return nil
	}
}

func testAccCheckConstellixGeoProximityDestroy(s *terraform.State) error {
	return nil
}

func geoproximityfromcontainer(resp *http.Response) (*models.GeoProximityAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	bodyString := string(bodyBytes)
	json.Unmarshal([]byte(bodyString), &data)

	gp := models.GeoProximityAttributes{}
	gp.Name = fmt.Sprintf("%v", data["name"])
	gp.City, _ = strconv.Atoi(fmt.Sprintf("%v", data["city"]))
	gp.Country = fmt.Sprintf("%v", data["country"])

	return &gp, nil

}
