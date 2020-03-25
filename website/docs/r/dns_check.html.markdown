---
layout: "constellix"
page_title: "Constellix: constellix_dns_check"
sidebar_current: "docs-constellix-resource-constellix_dns_check"
description: |-
    Manages one or more DNS check resource
---
# constellix_dns_check #
Manages one or more DNS check resource.

# Example Usage #
```hcl
        
resource "constellix_dns_check" "first" {
  name          = "dns check"
  host          = "constellix.com"
  fqdn          = "google.co.in"
  resolver      = "google.co.in"
  port          = 443
  protocol_type = "HTTPS"
  check_sites   = [1, 2]
}


```

## Argument Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `host` - (Required) Host for the resource, for example "constellix.com". It can be set only once.
* `fqdn` - (Required) A website address.
* `resolver` - (Required) A website address.
* `port` - (Required) Specifies the port number.
* `protocol_type` - (Required) Specifies upper layer protocol like HTTP, HTTPs, etc.
* `check_sites` - (Required) Site ids to check.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of DNS check resource.

