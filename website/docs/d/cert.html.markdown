---
layout: "constellix"
page_title: "CONSTELLIX: constellix_cert_record"
sidebar_current: "docs-constellix-data-source-constellix_cert_record"
description: |-
  Data source for records of type CERT for a specific domain.
---

# constellix_cert_record
Data source for records of type CERT for a specific domain.

## Example Usage ##

```hcl
data "constellix_cert_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  name        = "firstrecord"
}



```

## Argument Reference
* `source_type` - (Required) Type of the CERT record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the CERT record.

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `roundrobin` - (Optional) Object.
* `roundrobin.certificate_type` - (Optional) certificateType 0 - 65,535
* `roundrobin.key_tag` - (Optional) 0 - 65,535
* `roundrobin.disable_flag` - (Optional) disable flag. Default is false
* `roundrobin.certificate` - (Optional) certificate.
* `roundrobin.algorithm` - (Optional) 0-255.
* `type` - (Optional) Record type CERT.


