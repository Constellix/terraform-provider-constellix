---
layout: "constellix"
page_title: "CONSTELLIX: constellix_ns_record"
sidebar_current: "docs-constellix-resource-constellix_ns_record"
description: |-
  Manages records of type NS  for a specific domain.
---

# constellix_ns_record
 Manages records of type NS  for a specific domain.

## Example Usage ##

```hcl
resource "constellix_ns_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  roundrobin {
    value        = "prac."
    disable_flag = "false"
  }
  type       = "NS"
  gtd_region = 1
  note       = "First record"
  noanswer   = false

}


```

## Argument Reference ##
* `source_type` - (Required) Type of the NS record. The values which can be applied are "domains" or "templates".
* `ttl` - (Required) TTL must be in between 0 and 2147483647.
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type NS.
* `roundrobin` - (Required) Set.
* `roundrobin.value` - (Required) This will be the host name for the name server, for example ns0.nameserver.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).
* `roundrobin.disable_flag` - (Required) disable flag. Default is false

## Attributes Reference
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of the NS resource.