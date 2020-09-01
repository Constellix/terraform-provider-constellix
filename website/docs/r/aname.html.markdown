---
layout: "constellix"
page_title: "Constellix: constellix_aname_record"
sidebar_current: "docs-constellix-resource-constellix_aname_record"
description: |-
    Manages one or more domain ANAME record.
---
# constellix_aname_record #
Manages one or more domain ANAME record.

# Example Usage #
```hcl
        
resource "constellix_aname_record" "aname_record1" {
  domain_id   = "${constellix_domain.domain1.id}"
  ttl         = 18
  name        = "anamerecord"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "ANAME"
  contact_ids = [123]
  source_type = "domains"
  roundrobin {
    value        = "www.whatsapp.com."
    disable_flag = false
  }
  roundrobin {
    value        = "www.info.com."
    disable_flag = false
  }
  record_option = "failover"
  record_failover_values {
    value        = "www.w3schools.com."
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "www.messenger.com."
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_values {
    value        = "www.gmail.com."
    sort_order   = 3
    disable_flag = "false"
  }
  record_failover_failover_type = 1
  record_failover_disable_flag  = "false"
}

```

## Argument Reference ##
* `ttl` - (Required) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Required) Set
* `roundrobin.value` - (Required) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.
* `roundrobin.disable_flag` - (Required) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `name` - (Optional) Name of record. Name should be unique.
* `record_option` - (Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type ANAME
* `contact_ids` - (Optional) Applied contact list id. Only applicable to record with type failover.
* `record_failover` - (Optional) Set
* `record_failover_values` - (Required) Set
* `record_failover_values.value` - (Required) Host name
* `record_failover_values.check_id` - (Optional) Sonar check id
* `record_failover_values.disable_flag` - (Required) Enable or Disable the recordfailover values object. Default is false. Atleast one object should be false.
* `record_failover_values.sort_order` - (Required) Integer value which decides in which order recordfailover should be sorted.
* `record_failover_failover_type` - (Optional) 1 for Normal (always lowest level), 2 for Off on any Failover event, 3 for One Way (move to higher level)
* `record_failover_disable_flag` - (Optional) Enable or Disable the recordfailover object. Default is false. Atleast one recordfailover object should be false.

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of aname resource.

## Importing ##

An existing Record can be [imported][docs-import] into this resource using its Id, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import constellix_aname_record.example <source>:<parent-id>:<record-id>
```

Where source can be either domains or templates; parent-id is domain-id or template-id based on the source provided and record-id is the Id of record calculated via Constellix API.