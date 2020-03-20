---
layout: "constellix"
page_title: "CONSTELLIX: constellix_spf_record"
sidebar_current: "docs-constellix-data-source-constellix_spf_record"
description: |-
  Data source for records of type SPF  for a specific domain.
---

# constellix_spf_record
Data source for records of type SPF for a specific domain.

## Example Usage ##

```hcl
data "constellix_spf_record" "spf1" {
  domain_id		= "${constellix_domain.domain1.id}"
  source_type = "domains"
  name 			  = "temp"
}

```

## Argument Reference
* `source_type` - (Required) Type of the SPF record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the SPF record.

## Attributes Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type A.
* `roundrobin` - (Optional) Object.
* `roundrobin.value` - (Optional) Value may contain multiple strings (each string enclosed in double quotes). Individual string length should not exceed 255 characters.
* `roundrobin.disable_flag` - (Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.

