---
layout: "constellix"
page_title: "CONSTELLIX: constellix_cert_record"
sidebar_current: "docs-constellix-resource-constellix_cert_record"
description: |-
  Manages records of type CERT for a specific domain.
---

# constellix_cert_record
Manages records of type CERT for a specific domain.

## Example Usage ##

```hcl
resource "constellix_cert_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  type        = "CERT"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  roundrobin {
    certificate_type = 20
    key_tag          = 30
    certificate      = "certificate1"
    algorithm        = 100
    disable_flag     = "true"
  }
}

```

## Argument Reference ##
* `source_type` - (Required) Type of the CERT record. The values which can be applied are "domains" or "templates".
* `ttl` - (Required) TTL must be in between 0 and 2147483647.
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `roundrobin` - (Required) Object.
* `roundrobin.certificate_type` - (Required) certificateType 0 - 65,535
* `roundrobin.key_tag` - (Required) 0 - 65,535
* `roundrobin.disable_flag` - (Optional) disable flag. Default is false
* `roundrobin.certificate` - (Required) certificate.
* `roundrobin.algorithm` - (Required) 0-255.
* `type` - (Optional) Record type CERT.


## Attributes Reference
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of the CERT resource.