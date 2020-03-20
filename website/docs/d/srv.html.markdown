---
layout: "constellix"
page_title: "Constellix: constellix_srv_record"
sidebar_current: "docs-constellix-data-source-constellix_srv_record"
description: |-
  Data source for SRV record
---

# constellix_srv_record #
Data source for SRV record

## Example Usage ##

```hcl
data "constellix_srv_record" "datasrv" {
  name        = "srvdatasource"
  source_type = "domains"
  domain_id   = "${data.constellix_domain.first.id}"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of SRV record

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Optional) Set
* `roundrobin.value` - (Optional) The system that will receive the service.
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `roundrobin.port` - (Optional) The port of the service offered.
* `roundrobin.priority` - (Optional) The lower the number in the priority field, the higher the preference of the associated target. 0 is the highest priority (lowest number).
* `roundrobin.weight` - (Optional) The weight of the record allows an administrator to distribute load to multiple targets (load balance).
* `name` - (Required) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type SRV