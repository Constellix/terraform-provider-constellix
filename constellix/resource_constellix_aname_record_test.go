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

func TestAccAname_Basic(t *testing.T) {
	var aname models.AnameAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAnameConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAnameExists("constellix_domain.domain1", "constellix_aname_record.aname_record1", &aname),
					testAccCheckConstellixAnameAttributes("1800", &aname),
				),
			},
		},
	})
}

func TestAccConstellixAname_Update(t *testing.T) {
	var model models.AnameAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixAnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixAnameConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAnameExists("constellix_domain.domain1", "constellix_aname_record.aname_record1", &model),
					testAccCheckConstellixAnameAttributes("1800", &model),
				),
			},
			{
				Config: testAccCheckConstellixAnameConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixAnameExists("constellix_domain.domain1", "constellix_aname_record.aname_record1", &model),
					testAccCheckConstellixAnameAttributes("1900", &model),
				),
			},
		},
	})

}

func testAccCheckConstellixAnameConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "domaintest1.com"
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

	resource "constellix_aname_record" "aname_record1" {
			domain_id="${constellix_domain.domain1.id}"
			ttl="%d"
			name="aname1"
			source_type = "domains"
			roundrobin {
				value="www.whatsapp.com."
				disable_flag=false
			}
			roundrobin {
				value="www.info.com."
				disable_flag=false
			}
		}
	`, ttl)
}

func testAccCheckConstellixAnameExists(domainName string, anameName string, model *models.AnameAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[anameName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}
		if !err2 {
			return fmt.Errorf("Aname record %s not found", anameName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No aname record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/aname/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := anamefromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckConstellixAnameAttributes(ttl interface{}, model *models.AnameAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "aname1" != model.Name {
			return fmt.Errorf("Bad aname record name %s", model.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != model.TTL {
			return fmt.Errorf("Bad aname record %d", model.TTL)
		}

		return nil
	}
}

func testAccCheckConstellixAnameDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_aname_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/aname/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("aname record still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func anamefromcontainer(resp *http.Response) (*models.AnameAttributes, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	model := models.AnameAttributes{}

	model.Name = fmt.Sprintf("%v", data["name"])
	model.TTL, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	model.Note = fmt.Sprintf("%v", data["note"])
	return &model, nil
}
