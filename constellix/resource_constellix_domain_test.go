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

func TestAccDomain_Basic(t *testing.T) {
	var domain models.DomainAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixDomainConfig_basic("checkashu.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixDomainExists("constellix_domain.domain1", &domain),
					testAccCheckConstellixDomainAttributes("checkashu.com", &domain),
				),
			},
		},
	})
}

func TestAccConstellixDomain_Update(t *testing.T) {
	var domain models.DomainAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixDomainConfig_basic("shushu01.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixDomainExists("constellix_domain.domain1", &domain),
					testAccCheckConstellixDomainAttributes("shushu01.com", &domain),
				),
			},
			{
				Config: testAccCheckConstellixDomainConfig_basic("ashu70.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixDomainExists("constellix_domain.domain1", &domain),
					testAccCheckConstellixDomainAttributes("ashu70.com", &domain),
				),
			},
		},
	})
}

func testAccCheckConstellixDomainConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "%s"
		soa = {
			email = "dns.dnsmadeeasy.com."
			primary_nameserver = "ns41.constellix.com."
			ttl = 1800
			refresh = 48100
			retry = 7200
			expire = 1209
			negcache = 8000
		}
		note = "hello"
	}
	`, name)
}

func testAccCheckConstellixDomainExists(name string, domain *models.DomainAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Domain %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No domain id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := domainfromcontainer(resp)

		*domain = *tp
		return nil

	}
}

func domainfromcontainer(resp *http.Response) (*models.DomainAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	domain := models.DomainAttributes{}

	nameList := toStringList(data["name"])
	domain.Name = nameList
	domain.Note = data["note"].(string)

	return &domain, nil

}

func testAccCheckConstellixDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_domain" {
			_, err := client.GetbyId("v1/domains/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Domain is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckConstellixDomainAttributes(name string, domain *models.DomainAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != domain.Name[0] {
			return fmt.Errorf("Bad domain name %s", domain.Name)
		}
		if "hello" != domain.Note {
			return fmt.Errorf("Bad domain nameservergroup %s", domain.Note)
		}
		return nil
	}
}
