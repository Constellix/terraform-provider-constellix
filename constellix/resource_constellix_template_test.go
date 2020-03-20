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

func TestAccTemplate_Basic(t *testing.T) {
	var template models.TemplateAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTemplateConfig_basic("checktemplate"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTemplateExists("constellix_template.template1", &template),
					testAccCheckConstellixTemplateAttributes("checktemplate", &template),
				),
			},
		},
	})
}

func TestAccConstellixTemplate_Update(t *testing.T) {
	var template models.TemplateAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTemplateConfig_basic("checkone"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTemplateExists("constellix_template.template1", &template),
					testAccCheckConstellixTemplateAttributes("checkone", &template),
				),
			},
			{
				Config: testAccCheckConstellixTemplateConfig_basic("checktwo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTemplateExists("constellix_template.template1", &template),
					testAccCheckConstellixTemplateAttributes("checktwo", &template),
				),
			},
		},
	})
}

func testAccCheckConstellixTemplateConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "constellix_template" "template1" {
		name = "%s"
		has_geoip = "true"
	    has_gtd_regions = "false"
	}
	`, name)
}

func testAccCheckConstellixTemplateExists(name string, template *models.TemplateAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Template %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No template id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/templates/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := templatefromcontainer(resp)

		*template = *tp
		return nil

	}
}

func templatefromcontainer(resp *http.Response) (*models.TemplateAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	template := models.TemplateAttributes{}
	nameList := toStringList(data["name"])
	template.Name = nameList
	template.HasGeoIP, _ = strconv.ParseBool(fmt.Sprintf("%v", data["hasGeoIP"]))
	template.HasGtdRegions, _ = strconv.ParseBool(fmt.Sprintf("%v", data["hasGtdRegions"]))
	return &template, nil

}

func testAccCheckConstellixTemplateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_template" {
			_, err := client.GetbyId("v1/templates/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Template still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckConstellixTemplateAttributes(name string, template *models.TemplateAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != template.Name[0] {
			return fmt.Errorf("Bad template name %s", template.Name)
		}
		v1 := fmt.Sprintf("%v", template.HasGeoIP)
		if "true" != v1 {
			return fmt.Errorf("Bad value of hasgeoip %s", v1)
		}
		v2 := fmt.Sprintf("%v", template.HasGtdRegions)
		if "false" != v2 {
			return fmt.Errorf("Bad value of hasgtdregions %s", v2)
		}

		return nil
	}
}
