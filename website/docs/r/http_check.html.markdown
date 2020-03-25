---
layout: "constellix"
page_title: "Constellix: constellix_http_check"
sidebar_current: "docs-constellix-resource-constellix_http_check"
description: |-
    Manages one or more HTTP check resource
---
# constellix_HTTP_Check_resource #
Manages one or more HTTP check resource.

# Example Usage #
```hcl
        
resource "constellix_http_check" "first" {
  name = "http check"
  host = "constellix.com"
  ip_version = "IPV4"
  port = 443
  protocol_type = "HTTPS"
  check_sites = [1,2]
}

```

## Argument Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `host` - (Required) Host for the resource, for example "constellix.com". It can be set only once.
* `ip_version` - (Required) Specifies the version of IP. It can be set only once.
* `port` - (Required) Specifies the port number.
* `protocol_type` - (Required) Specifies upper layer protocol like HTTP, HTTPs, etc.
* `check_sites` - (Required) Site ids to check.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of HTTP check resource.