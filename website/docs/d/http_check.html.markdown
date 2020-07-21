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
* `interval` - (Optional) Check Interval. Allowed values are `THIRTYSECONDS`, `ONEMINUTE`, `TWOMINUTES`, `THREEMINUTES`, `FOURMINUTES`, `FIVEMINUTES`, `TENMINUTES`, `THIRTYMINUTES`, `HALFDAY` and `DAY`.
* `interval_policy` - (Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are `PARALLEL`, `ONCEPERSITE` and `ONCEPERREGION`.
* `validation_policy` - (Optional) Specifies how the check should be validated. Allowed values are `SIMPLE` and `MAJORITY`. This parameter will only work with the `interval_policy` set to `PARALLEL`.
* `fqdn` - (Optional) Fully qualified domain name of the URL should be checked.
* `path` - (Optional) In case of multi-page site, which path should be checked.
* `search_string` - (Optional) String to search in the first 2KB of resonse received.
* `expected_status_code` - (Optional) Expected HTTP status code for this check.