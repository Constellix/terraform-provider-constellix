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
