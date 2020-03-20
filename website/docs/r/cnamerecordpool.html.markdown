---
layout: "constellix"
page_title: "CONSTELLIX: constellix_cname_record_pool"
sidebar_current: "docs-constellix-resource-constellix_cname_record_pool"
description: |-
    Manages the pools of CNAME records.
---

# constellix_cname_record_pool
  Manages the pools of CNAME records.

## Example Usage ##

```hcl
resource "constellix_cname_record_pool" "firstrecord" {
  name                   = "firstrecord"
  num_return             = "10"
  min_available_failover = 1
  values {
    value        = "8.1.1.1"
    weight       = 20
    policy       = "followsonar"
    disable_flag = false
  }
  failed_flag  = false
  disable_flag = false
  note         = "First record"
}


```

## Argument Reference ##
* `name` - (Required) Pool name should be unique.
* `num_return` - (Required) minimum number of value object to return. Value must be in between 0 and 64.
* `min_available_failover` - (Required) minimum number of Available Failover . Value must be in between 0 and 64.
* `version` - (Optional) System generated version number.
* `failed_flag` - (Optional) failed flag. Default is false.
* `disable_flag` - (Optional) Enable or disable pool values. Default is false.
* `values` - (Required) Object Number of IP/Hosts in a pool values cannot be less than the "Num Return" and "Min Available" values
* `values.value` - (Required) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.
* `values.weight` - (Required)weight number to sort the priorty. Weight must be in between 1 and 1000000
* `values.disable_flag` - (Optional) Enable or disable pool values. Default is false.
* `values.check_id` - (Optional) Sonar check id is required when you want to apply the ITO feature on a pool.
* `values.policy` - (Required) "followsonar" for Follow sonar. "alwaysoff" for Always off. "alwayson" for Always on. "offonfailure" for Off on Failure.
* `note` - (Optional) Description.

## Attributes Reference
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of the cname record pool resource.