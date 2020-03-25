---
layout: "constellix"
page_title: "Constellix: constellix_tcp_check"
sidebar_current: "docs-constellix-resource-constellix_tcp_check"
description: |-
    Manages one or more TCP check resource
---
# constellix_tcp_check #
Manages one or more TCP check resource.

# Example Usage #
```hcl
        
resource "constellix_tcp_check" "first" {
  name = "tcp check"
  host = "constellix.com"
  ip_version = "IPV4"
  port = 443
  check_sites = [1,2]
}

```

## Argument Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `host` - (Required) Host for the resource, for example "constellix.com". It can be set only once.
* `ip_version` - (Required) Specifies the version of IP. It can be set only once.
* `port` - (Required) Specifies the port number.
* `check_sites` - (Required) Site ids to check.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of TCP check resource.