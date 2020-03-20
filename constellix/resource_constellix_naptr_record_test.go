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

func TestAccNaptr_Basic(t *testing.T) {
	var naptr models.NAPTRAttributes
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixNaptrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixNaptrConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNaptrExists("constellix_domain.domain1", "constellix_naptr_record.naptr1", &naptr),
					testAccCheckConstellixNaptrAttributes(1800, &naptr),
				),
			},
		},
	})
}

func TestAccConstellixNaptr_Update(t *testing.T) {
	var naptr models.NAPTRAttributes

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConstellixNaptrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConstellixNaptrConfig_basic(1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNaptrExists("constellix_domain.domain1", "constellix_naptr_record.naptr1", &naptr),
					testAccCheckConstellixNaptrAttributes(1800, &naptr),
				),
			},
			{
				Config: testAccCheckConstellixNaptrConfig_basic(1900),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConstellixNaptrExists("constellix_domain.domain1", "constellix_naptr_record.naptr1", &naptr),
					testAccCheckConstellixNaptrAttributes(1900, &naptr),
				),
			},
		},
	})
}

func testAccCheckConstellixNaptrConfig_basic(ttl int) string {
	return fmt.Sprintf(`
	resource "constellix_domain" "domain1" {
		name = "checknaptr.com"
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

	resource "constellix_naptr_record" "naptr1"{
		domain_id = "${constellix_domain.domain1.id}"
		source_type = "domains"
		name = "tempnaptrrecord"
		ttl = "%d"
		note = "Practice record naptr"
	    roundrobin {
			       order = 10
			       preference = 100
			       flags = "s"
			       service = "SIP+D2U"
			       regular_expression = "hello"
			       replacement = "foobar.example.com."
			       disable_flag = "false"
			     }
	}
	`, ttl)
}

func testAccCheckConstellixNaptrExists(domainName string, naptrName string, naptr *models.NAPTRAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[naptrName]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Naptr record %s not found", naptrName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Naptr record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("v1/domains/" + rs1.Primary.ID + "/records/naptr/" + rs2.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := naptrfromcontainer(resp)

		*naptr = *tp
		return nil
	}
}

func testAccCheckConstellixNaptrAttributes(ttl interface{}, naptr *models.NAPTRAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "tempnaptrrecord" != naptr.Name {
			return fmt.Errorf("Bad Naptr record name %s", naptr.Name)
		}
		ttl, _ := strconv.Atoi(fmt.Sprintf("%v", ttl))
		if ttl != naptr.Ttl {
			return fmt.Errorf("Bad Naptr record ttl %d", naptr.Ttl)
		}
		return nil
	}
}

func testAccCheckConstellixNaptrDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["constellix_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "constellix_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "constellix_naptr_record" {
			_, err := client.GetbyId("v1/domains/" + domainid + "/records/naptr/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Naptr record still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func naptrfromcontainer(resp *http.Response) (*models.NAPTRAttributes, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	var data map[string]interface{}
	json.Unmarshal([]byte(bodyString), &data)
	naptr := models.NAPTRAttributes{}

	naptr.Name = fmt.Sprintf("%v", data["name"])
	naptr.Ttl, _ = strconv.Atoi(fmt.Sprintf("%v", data["ttl"]))
	naptr.Note = fmt.Sprintf("%v", data["note"])

	return &naptr, nil

}
