---
layout: "constellix"
page_title: "CONSTELLIX: constellix_aaaa_record_pool"
sidebar_current: "docs-constellix-data-source-constellix_aaaa_record_pool"
description: |-
   Data source for the pools of AAAA records.
---

# constellix_aaaa_record_pool
 Data source for the pools of AAAA records.

## Example Usage ##

```hcl
data "constellix_aaaa_record_pool" "prac" {
  name = "firstrecord"
}

```

## Attributes Reference
* `name` - (Required) Pool name should be unique.

## Attribute Reference ##
* `num_return` - (Optional) minimum number of value object to return. Value must be in between 0 and 64.
* `min_available_failover` - (Optional)minimum number of Available Failover . Value must be in between 0 and 64.
* `failed_flag` - (Optional) failed flag. Default is false.
* `disable_flag` - (Optional) Enable or disable pool values. Default is false.
* `values` - (Optional) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values
* `values.value` - (Optional) IPv6 address.
* `values.weight` - (Optional) weight number to sort the priorty. Weight must be in between 1 and 1000000.
* `values.disable_flag` - (Optional) Enable or disable pool values. Default is false.
* `values.check_id` - (Optional) Sonar check id is Optional when you want to apply the ITO feature on a pool.
* `values.policy` - (Optional) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.
* `note` - (Optional) Description.
