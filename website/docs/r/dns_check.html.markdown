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
  fqdn          = "google.co.in"
  resolver      = "google.co.in"
  check_sites   = [1, 2]
}


```

## Argument Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `fqdn` - (Required) A website address. It can be set only once
* `resolver` - (Required) A website address. It can be set only once
* `check_sites` - (Required) Site ids to check.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of DNS check resource.

