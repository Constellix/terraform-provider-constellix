resource "constellix_tcp_check" "first" {
  name = "tcp check"
  host = "constellix.com"
  ip_version = "IPV4"
  port = 443
  check_sites = [1,2]
  verification_policy = "MAJORITY"
  string_to_send = "tp"
  string_to_receive = "rc"
  interval_policy = "ONCEPERSITE"
}

provider "constellix" {
  apikey    = ""
  secretkey = ""
}