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
* `interval` - (Optional) Check Interval. Allowed values are `THIRTYSECONDS`, `ONEMINUTE`, `TWOMINUTES`, `THREEMINUTES`, `FOURMINUTES`, `FIVEMINUTES`, `TENMINUTES`, `THIRTYMINUTES`, `HALFDAY` and `DAY`.
* `interval_policy` - (Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are `PARALLEL`, `ONCEPERSITE` and `ONCEPERREGION`.
* `verification_policy` - (Optional) Specifies how the check should be validated. Allowed values are `SIMPLE` and `MAJORITY`. This parameter will only work with the `interval_policy` set to `PARALLEL`.
* `expected_response` - (Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of DNS check resource.

