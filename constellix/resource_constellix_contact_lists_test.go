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

func TestAccContactList_Basic(t *testing.T) {
	var ctl models.ContactListAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixContactListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixContactListConfig_basic("abc@yahoo.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixContactListExists("constellix_contact_lists.ctl1", &ctl),
					testAccCheckConstellixContactListAttributes(&ctl),
				),
			},
		},
	})
}

func TestAccConstellixContactList_Update(t *testing.T) {
	var ctl models.ContactListAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixContactListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixContactListConfig_basic("abc@yahoo.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixContactListExists("constellix_contact_lists.ctl1", &ctl),
					testAccCheckConstellixContactListAttributes(&ctl),
				),
			},
			{
				Config: testAccCheckConstellixContactListConfig_basic("shaival@yahoo.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixContactListExists("constellix_contact_lists.ctl1", &ctl),
					testAccCheckConstellixContactListAttributes(&ctl),
				),
			},
		},
	})
}

func testAccCheckConstellixContactListConfig_basic(tp string) string {
	return fmt.Sprintf(`

	resource "constellix_contact_lists" "ctl1" {
	
		name = "tempcontactlist"
		email_addresses = ["%s"]
		
	}
	`, tp)
}

func testAccCheckConstellixContactListExists(hinfoName string, ctl *models.ContactListAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs2, err2 := s.RootModule().Resources[hinfoName]
		if !err2 {
			return fmt.Errorf("ContactList record %s not found", hinfoName)
		}

		if rs2.Primary.ID == "" {
			return fmt.Errorf("No contactlist record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v2/contactLists/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := contactlistfromcontainer(resp)

		*ctl = *tp
		return nil
	}
}

func testAccCheckConstellixContactListAttributes(ctl *models.ContactListAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempcontactlist" != ctl.Name {
			return fmt.Errorf("Bad ContactList record name %s", ctl.Name)
		}

		return nil
	}
}

func testAccCheckConstellixContactListDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_contact_lists_record" {
			_, err := client.GetbyId("v2/contactLists/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("ContactList record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func contactlistfromcontainer(resp *http.Response) (*models.ContactListAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)

	ctl := models.ContactListAttributes{}

	ctl.Name = fmt.Sprintf("%v", data["name"])

	return &ctl, nil

}
