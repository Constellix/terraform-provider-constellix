---
layout: "constellix"
page_title: "Constellix: constellix_hinfo_record"
sidebar_current: "docs-constellix-data-source-constellix_hinfo_record"
description: |-
  Data source for HINFO record
---

# constellix_hinfo_record #
Data source for HINFO record

## Example Usage ##

```hcl
data "constellix_hinfo_record" "hinfo" {
  domain_id   = "${data.constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "datahinforecord"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of HINFO record

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Optional) Set
* `roundrobin.cpu` - (Optional) A description of basic system hardware
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `roundrobin.os` - (Optional) A description of the operating system and version
* `name` - (Required) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type HINFO