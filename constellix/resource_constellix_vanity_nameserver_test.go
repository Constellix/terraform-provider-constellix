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

func TestAccVanitynameserver_Basic(t *testing.T) {
	var VNS models.VanityNameserverAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixVanitynameserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixVanitynameserverConfig_basic("create"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixVanitynameserverExists("constellix_vanity_nameserver.one", &VNS),
					testAccCheckConstellixVanitynameserverAttributes("create", &VNS),
				),
			},
		},
	})
}

func TestAccConstellixVanitynameserver_Update(t *testing.T) {
	var VNS models.VanityNameserverAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixVanitynameserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixVanitynameserverConfig_basic("Name1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixVanitynameserverExists("constellix_vanity_nameserver.one", &VNS),
					testAccCheckConstellixVanitynameserverAttributes("Name1", &VNS),
				),
			},
			{
				Config: testAccCheckConstellixVanitynameserverConfig_basic("Name2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixVanitynameserverExists("constellix_vanity_nameserver.one", &VNS),
					testAccCheckConstellixVanitynameserverAttributes("Name2", &VNS),
				),
			},
		},
	})
}

func testAccCheckConstellixVanitynameserverConfig_basic(liststring string) string {
	return fmt.Sprintf(`
	resource "constellix_vanity_nameserver" "one"{
		name = "checkVNS"
		nameserver_group = 1
		nameserver_list_string = "%s"
	}
	`, liststring)
}

func testAccCheckConstellixVanitynameserverExists(VNSName string, model *models.VanityNameserverAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[VNSName]

		if !ok {
			return fmt.Errorf("Vanitynameserver %s not found", VNSName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vanitynameserver id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/vanityNameservers/" + rs.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := vnsfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixVanitynameserverAttributes(grpname string, model *models.VanityNameserverAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if grpname != model.NameserverListString {
			return fmt.Errorf("Bad Vanity Nameserver ListString %s", model.NameserverListString)
		}
		if "checkVNS" != model.Name {
			return fmt.Errorf("Bad Vanity Nameserver name %s", model.Name)
		}
		return nil
	}
}

func testAccCheckConstellixVanitynameserverDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_vanitynameserver" {
			_, err := client.GetbyId("v1/vanityNameservers/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Vanity Nameserver is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func vnsfromcontainer(resp *http.Response) (*models.VanityNameserverAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	model := models.VanityNameserverAttributes{}

	model.Name = fmt.Sprintf("%v", data["name"])
	model.NameserverListString = fmt.Sprintf("%v", data["nameserversListString"])

	return &model, nil
}
