---
layout: "constellix"
page_title: "Constellix: constellix_srv_record"
sidebar_current: "docs-constellix-resource-constellix_srv_record"
description: |-
    Manages records of type SRV  for a specific domain.
---
# constellix_srv_record #
Manages records of type SRV  for a specific domain.

# Example Usage #
```hcl
resource "constellix_srv_record" "srvrecord1" {
  domain_id   = "${constellix_domain.domain1.id}"
  ttl         = 1800
  name        = "srvrecord"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "SRV"
  source_type = "domains"
  roundrobin {
    value        = "www.google.com"
    port         = 8888
    priority     = 65
    weight       = 20
    disable_flag = false
  }
}



```

## Argument Reference ##
* `ttl` - (Required) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Required) Set
* `roundrobin.value` - (Required) The system that will receive the service.
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `roundrobin.port` - (Required) The port of the service offered. Value should be between 0 and 65535.
* `roundrobin.priority` - (Required) The lower the number in the priority field, the higher the preference of the associated target. 0 is the highest priority (lowest number). Value should be between 0 and 65535.
* `roundrobin.weight` - (Required) The weight of the record allows an administrator to distribute load to multiple targets (load balance). Value should be between 0 and 65535.
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type SRV

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of srv resource.

## Importing ##

An existing Record can be [imported][docs-import] into this resource using its Id, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import constellix_srv_record.example <source>:<parent-id>:<record-id>
```

Where source can be either domains or templates; parent-id is domain-id or template-id based on the source provided and record-id is the Id of record calculated via Constellix API.