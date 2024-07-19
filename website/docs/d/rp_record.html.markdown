---
layout: "constellix"
page_title: "Constellix: constellix_rp_record"
sidebar_current: "docs-constellix-data-source-constellix_rp_record"
description: |-
  Data source for RP record
---

# constellix_rp_record #
Data source for RP record

## Example Usage ##

```hcl
data "constellix_rp_record" "datarp" {
  name        = "rpdatasource"
  source_type = "domains"
  domain_id   = "${data.constellix_domain.first.id}"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of RP record

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Optional) Set
* `roundrobin.mailbox` - (Optional) A mailbox for the responsible person of the domain
* `roundrobin.txt` - (Optional) A hostname for the responsible person of the domain
* `roundrobin.disable_flag` - (Optional) Enable or Disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `name` - (Required) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type RP