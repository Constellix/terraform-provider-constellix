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
* `host` - (Required) Host for the resource, for example "constellix.com". It can be set only once.
* `fqdn` - (Required) A website address.
* `resolver` - (Required) A website address.
* `port` - (Required) Specifies the port number.
* `protocol_type` - (Required) Specifies upper layer protocol like HTTP, HTTPs, etc.
* `check_sites` - (Required) Site ids to check.
