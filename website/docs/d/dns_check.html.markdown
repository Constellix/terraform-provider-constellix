---
layout: "constellix"
page_title: "Constellix: constellix_dns_check"
sidebar_current: "docs-constellix-data-source-constellix_dns_check"
description: |-
  Data source for DNS check resource
---

# constellix_dns_check #
Data source for DNS check resource

## Example Usage ##

```hcl
data "constellix_dns_check" "check" {
  name = "dns check"
}


```

## Argument Reference ##
* `name` - (Required) Name of resource. Name should be unique.

## Attribute Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `fqdn` - (Required) A website address. It can be set only once
* `resolver` - (Required) A website address. It can be set only once
* `check_sites` - (Required) Site ids to check.
* `interval` - (Optional) Check Interval. Allowed values are `THIRTYSECONDS`, `ONEMINUTE`, `TWOMINUTES`, `THREEMINUTES`, `FOURMINUTES`, `FIVEMINUTES`, `TENMINUTES`, `THIRTYMINUTES`, `HALFDAY` and `DAY`.
* `interval_policy` - (Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are `PARALLEL`, `ONCEPERSITE` and `ONCEPERREGION`.
* `verification_policy` - (Optional) Specifies how the check should be validated. Allowed values are `SIMPLE` and `MAJORITY`. This parameter will only work with the `interval_policy` set to `PARALLEL`.
* `expected_response` - (Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions.
