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

func TestAccHTTPCheck_Basic(t *testing.T) {
	var httpCheck models.HttpcheckAttr
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHTTPCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHTTPCheckConfig_basic(443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPCheckExists("constellix_http_check.http1", &httpCheck),
					testAccCheckConstellixHTTPCheckAttributes(443, &httpCheck),
				),
			},
		},
	})
}

func TestAccConstellixHTTPCheck_Update(t *testing.T) {
	var httpCheck models.HttpcheckAttr

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHTTPCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHTTPCheckConfig_basic(443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPCheckExists("constellix_http_check.http1", &httpCheck),
					testAccCheckConstellixHTTPCheckAttributes(443, &httpCheck),
				),
			},
			{
				Config: testAccCheckConstellixHTTPCheckConfig_basic(80),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPCheckExists("constellix_http_check.http1", &httpCheck),
					testAccCheckConstellixHTTPCheckAttributes(80, &httpCheck),
				),
			},
		},
	})
}

func testAccCheckConstellixHTTPCheckConfig_basic(port int) string {
	return fmt.Sprintf(`
	resource "constellix_http_check" "http1"{
		name = "http check"
  		host = "constellix.com"
  		ip_version = "IPV4"
  		port = %d
  		protocol_type = "HTTPS"
  		check_sites = [1,2]
	}
	`, port)
}

func testAccCheckConstellixHTTPCheckExists(httpName string, http *models.HttpcheckAttr) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, err := s.RootModule().Resources[httpName]

		if !err {
			return fmt.Errorf("HTTP check %s not found", httpName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No HTTP check id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("https://api.sonar.constellix.com/rest/api/http/" + rs.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := httpCheckfromcontainer(resp)

		*http = *tp
		return nil
	}
}

func testAccCheckConstellixHTTPCheckAttributes(port interface{}, http *models.HttpcheckAttr) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "http check" != http.Name {
			return fmt.Errorf("Bad HTTP check resource name %s", http.Name)
		}
		port, _ := strconv.Atoi(fmt.Sprintf("%v", port))
		if port != http.Port {
			return fmt.Errorf("Bad HTTP check resource port %d", http.Port)
		}
		return nil
	}
}

func testAccCheckConstellixHTTPCheckDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_http_check" {
			_, err := client.GetbyId("https://api.sonar.constellix.com/rest/api/http/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("HTTP check resource still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func httpCheckfromcontainer(resp *http.Response) (*models.HttpcheckAttr, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	http := models.HttpcheckAttr{}
	http.Name = fmt.Sprintf("%v", data["name"])
	http.Port, _ = strconv.Atoi(fmt.Sprintf("%v", data["port"]))
	http.Host = fmt.Sprintf("%v", data["host"])

	return &http, nil

}
