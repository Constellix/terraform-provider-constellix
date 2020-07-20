resource "constellix_http_check" "first" {
  name = "http check"
  host = "constellix.com"
  ip_version = "IPV4"
  port = 443
  protocol_type = "HTTPS"
  check_sites = [1,2]
  fqdn = "test.com"
  path = "root"
  interval = "ONEMINUTE"
}

provider "constellix" {
  apikey    = ""
  secretkey = ""
}