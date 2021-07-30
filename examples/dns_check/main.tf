resource "constellix_dns_check" "first" {
  name          = "dns check"
  fqdn          = "google.co.in"
  resolver      = "google.co.in"
  check_sites   = [1, 2]
  notification_groups = [874, 875, 876]
  expected_response = "2.2.2.2"
}

provider "constellix" {
  apikey    = ""
  secretkey = ""
}