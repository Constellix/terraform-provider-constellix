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

func TestAccTags_Basic(t *testing.T) {
	var tags models.Tags
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTagsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTagsConfig_basic("checkTag"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTagsExists("constellix_tags.tag1", &tags),
					testAccCheckConstellixTagsAttributes("checkTag", &tags),
				),
			},
		},
	})
}

func TestAccConstellixTags_Update(t *testing.T) {
	var tags models.Tags

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixTagsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixTagsConfig_basic("Name1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTagsExists("constellix_tags.tag1", &tags),
					testAccCheckConstellixTagsAttributes("Name1", &tags),
				),
			},
			{
				Config: testAccCheckConstellixTagsConfig_basic("Name2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixTagsExists("constellix_tags.tag1", &tags),
					testAccCheckConstellixTagsAttributes("Name2", &tags),
				),
			},
		},
	})
}

func testAccCheckConstellixTagsConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "constellix_tags" "tag1"{
		name = "%s"
	}
	`, name)
}

func testAccCheckConstellixTagsExists(TagName string, model *models.Tags) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[TagName]

		if !ok {
			return fmt.Errorf("Tags %s not found", TagName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Tags id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v2/tags/" + rs.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := tagsfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixTagsAttributes(name string, model *models.Tags) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != model.Name {
			return fmt.Errorf("Bad Tags name %s", model.Name)
		}
		return nil
	}
}

func testAccCheckConstellixTagsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_tags" {
			_, err := client.GetbyId("v2/tags/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Tag still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func tagsfromcontainer(resp *http.Response) (*models.Tags, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	model := models.Tags{}

	model.Name = fmt.Sprintf("%v", data["name"])

	return &model, nil
}
