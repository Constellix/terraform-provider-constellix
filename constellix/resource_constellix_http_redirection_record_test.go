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

func TestAccHTTPRedirection_Basic(t *testing.T) {
	var htp models.HTTPRedirectionAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHTTPRedirectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHTTPRedirectionConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPRedirectionExists("constellix_domain.domain1", "constellix_http_redirection_record.http1", &htp),
					testAccCheckConstellixHTTPRedirectionAttributes(1800, &htp),
				),
			},
		},
	})
}

func TestAccConstellixHTTPRedirection_Update(t *testing.T) {
	var htp models.HTTPRedirectionAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixHTTPRedirectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixHTTPRedirectionConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPRedirectionExists("constellix_domain.domain1", "constellix_http_redirection_record.http1", &htp),
					testAccCheckConstellixHTTPRedirectionAttributes(1800, &htp),
				),
			},
			{
				Config: testAccCheckConstellixHTTPRedirectionConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixHTTPRedirectionExists("constellix_domain.domain1", "constellix_http_redirection_record.http1", &htp),
					testAccCheckConstellixHTTPRedirectionAttributes(1900, &htp),
				),
			},
		},
	})
}

func testAccCheckConstellixHTTPRedirectionConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checkhttp.com"
		soa = {
			email = "com.com."
			primary_nameserver = "ns41.constellix.com."
			ttl = 1900
			refresh = 48100
			retry = 7200
			expire = 1209
			negcache = 8000
		}
	}

	resource "constellix_http_redirection_record" "http1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "temphttpredirectionrecord"
		redirect_type_id = 1
		ttl = "%d"
		url = "https://www.google.com"
	}
		
		
	`, ttl)
}

func testAccCheckConstellixHTTPRedirectionExists(domainName string, httpredirectionName string, htp *models.HTTPRedirectionAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[httpredirectionName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Http redirection record %s not found", httpredirectionName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Httpredirection record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/httpredirection/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := httpfromcontainer(resp)

		*htp = *tp
		return nil
	}
}

func httpfromcontainer(resp *http.Response) (*models.HTTPRedirectionAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	htp := models.HTTPRedirectionAttributes{}

	htp.Name = fmt.Sprintf("%v", data["name"])
	htp.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	htp.URL = fmt.Sprintf("%v", data["url"])
	htp.RedirectTypeID, _ = strconv.Atoi(fmt.Sprintf("%v", data["redirectTypeId"]))

	return &htp, nil

}

func testAccCheckConstellixHTTPRedirectionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_http_redirection_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/httpredirection/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Httpredirection record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
func testAccCheckConstellixHTTPRedirectionAttributes(ttl interface{}, htp *models.HTTPRedirectionAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temphttpredirectionrecord" != htp.Name {
			return fmt.Errorf("Bad Httpredirection record name %s", htp.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != htp.TTL {
			return fmt.Errorf("Bad Httpredirection record ttl %d", htp.TTL)
		}
		return nil
	}
}
