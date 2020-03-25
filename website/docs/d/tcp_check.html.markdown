---
layout: "constellix"
page_title: "Constellix: constellix_tcp_check"
sidebar_current: "docs-constellix-data-source-constellix_tcp_check"
description: |-
  Data source for TCP check resource
---

# constellix_tcp_check #
Data source for TCP check resource

## Example Usage ##

```hcl
data "constellix_tcp_check" "check" {
  name        = "tcp check"
}

```

## Argument Reference ##
* `name` - (Required) Name of resource. Name should be unique.

## Attribute Reference ##
* `name` - (Required) name of the resource. Name should be unique.
* `host` - (Optional) Host for the resource, for example "constellix.com". It can be set only once.
* `ip_version` - (Optional) Specifies the version of IP. It can be set only once.
* `port` - (Optional) Specifies the port number.
* `check_sites` - (Optional) Site ids to check.