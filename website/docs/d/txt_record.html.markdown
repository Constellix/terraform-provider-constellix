---
layout: "constellix"
page_title: "Constellix: constellix_txt_record"
sidebar_current: "docs-constellix-data-source-constellix_txt_record"
description: |-
  Data source for TXT record
---

# constellix_txt_record #
Data source for TXT record


## Example Usage ##

```hcl
data "constellix_txt_record" "datatxt" {
  name        = "txtdatasource"
  source_type = "domains"
  domain_id   = "${data.constellix_domain.first.id}"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of TXT record

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Optional) Set
* `roundrobin.value` - (Optional) Free form text data of any type which may be no longer than 255 characters unless divided into multiple strings with sets of quotation marks..
* `roundrobin.disable_flag` - (Optional) Disable flag. Default is false
* `name` - (Required) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type TXT