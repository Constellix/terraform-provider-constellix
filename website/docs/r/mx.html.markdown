---
layout: "constellix"
page_title: "Constellix: constellix_mx_record"
sidebar_current: "docs-constellix-resource-constellix_mx_record"
description: |-
    Manages records of type MX for a specific domain.
---
# constellix_mx_record #
Manages records of type MX for a specific domain.

# Example Usage #
```hcl

resource "constellix_mx_record" "mx1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "mxrecord"
  ttl         = "1900"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "MX"
  roundrobin {
    value        = "abc"
    level        = "100"
    disable_flag = "false"
  }
  roundrobin {
    value        = "dce"
    level        = "200"
    disable_flag = "true"
  }
}

```

## Argument Reference ##
* `ttl` - (Required) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Required) Set
* `roundrobin.value` - (Required) The mail server that will accept mail for the host that is specified in the name field. Your domain name is automatically appended to your value if it does not end it a dot.
* `roundrobin.level` - (Required) Level must be in between 0 and 65535. The MX level determines the order (by priority) that remote mail servers will attempt to deliver email. The mail server with the lowest MX level will be the first priority.
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type MX

## Attribute Reference ##
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of mx resource.