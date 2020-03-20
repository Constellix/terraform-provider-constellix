---
layout: "constellix"
page_title: "Constellix: constellix_aname_record"
sidebar_current: "docs-constellix-data-source-constellix_aname_record"
description: |-
  Data source for Aname record
---

# constellix_aname_record #
Data source for Aname record


## Example Usage ##

```hcl
data "constellix_aname_record" "dataanamerecord" {
  domain_id   = "${data.constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "anamerecorddatasource"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of Aname record

## Attribute Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `roundrobin` - (Optional) Set
* `roundrobin.value` - (Optional) Host name. If "Host" value does not end in a dot, your domain name will be appended to it.
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `record_option` - (Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type ANAME
* `contact_ids` - (Optional) Applied contact list id. Only applicable to record with type failover.
* `record_failover` - (Optional) Set
* `record_failover_values` - (Required) Set
* `record_failover_values.value` - (Optional) Host name
* `record_failover_values.check_id` - (Optional) Sonar check id
* `record_failover_values.disable_flag` - (Optional) Enable or Disable the recordfailover values object. Default is false. Atleast one object should be false.
* `record_failover_values.sort_order` - (Optional) Integer value which decides in which order recordfailover should be sorted.
* `record_failover_failover_type` - (Optional) 1 for Normal (always lowest level), 2 for Off on any Failover event, 3 for One Way (move to higher level)
* `record_failover_disable_flag` - (Optional) Enable or Disable the recordfailover object. Default is false. Atleast one recordfailover object should be false.
