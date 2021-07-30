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
* `notification_groups` - (Optional) List of group IDs for the notification group of TCP Check.
* `interval` - (Optional) Check Interval. Allowed values are `THIRTYSECONDS`, `ONEMINUTE`, `TWOMINUTES`, `THREEMINUTES`, `FOURMINUTES`, `FIVEMINUTES`, `TENMINUTES`, `THIRTYMINUTES`, `HALFDAY` and `DAY`.
* `interval_policy` - (Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are `PARALLEL`, `ONCEPERSITE` and `ONCEPERREGION`.
* `verification_policy` - (Optional) Specifies how the check should be validated. Allowed values are `SIMPLE` and `MAJORITY`. This parameter will only work with the `interval_policy` set to `PARALLEL`.
* `string_to_send` - (Optional) String to send along with the check. It can be any parameter to the endpoint.
* `string_to_receive` - (Optional) String which should be received as a result of TCP check.