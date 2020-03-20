---
layout: "constellix"
page_title: "Constellix: constellix_http_redirection_record"
sidebar_current: "docs-constellix-data-source-constellix_http_redirection_record"
description: |-
  Data source for HTTP Redirection record
---

# constellix_http_redirection_record #
Data source for HTTP Redirection record


## Example Usage ##

```hcl
data "constellix_http_redirection_record" "datahttpredirection" {
  name        = "httpredirectiondatasource"
  source_type = "domains"
  domain_id   = "${data.constellix_domain.first.id}"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of HTTP Redirection record

## Attribute Reference ##
* `ttl` - (Required) TTL must be in between 0 and 2147483647
* `url` - (Required) URL link to redirect
* `redirect_type_id` - (Required) 1 for Hidden Frame Masked, 2 for Standard - 301, 3 for Standard - 302
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `hardlink_flag` - (Optional) Hardlink flag. Default is false
* `description` - (Optional) Description
* `title` - (Optional) Title of redirection
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created, 1 for World (Default), 2 for Europe, 3 for US East, 4 for US West, 5 for Asia Pacific, 6 for Oceania, note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type HTTP Redirection