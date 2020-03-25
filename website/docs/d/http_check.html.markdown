---
layout: "constellix"
page_title: "Constellix: constellix_http_check"
sidebar_current: "docs-constellix-data-source-constellix_http_check"
description: |-
  Data source for HTTP check resource
---

# constellix_HTTP_Check_resource #
Data source for HTTP check resource

## Example Usage ##

```hcl
data "constellix_http_check" "check" {
  name        = "http check"
}

```

## Argument Reference ##
* `name` - (Required) Name of resource. Name should be unique.

## Attribute Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `host` - (Optional) Host for the resource, for example "constellix.com". It can be set only once.
* `ip_version` - (Optional) Specifies the version of IP. It can be set only once.
* `port` - (Optional) Specifies the port number.
* `protocol_type` - (Optional) Specifies upper layer protocol like HTTP, HTTPs, etc.
* `check_sites` - (Optional) Site ids to check.