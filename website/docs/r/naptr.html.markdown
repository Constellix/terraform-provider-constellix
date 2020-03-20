---
layout: "constellix"
page_title: "CONSTELLIX: constellix_naptr_record"
sidebar_current: "docs-constellix-resource-constellix_naptr_record"
description: |-
  Manages records of type NAPTR for a specific domain.
---

# constellix_naptr_record
 Manages records of type NAPTR for a specific domain.

## Example Usage ##

```hcl
resource "constellix_naptr_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  roundrobin {
    order              = 10
    preference         = 100
    flags              = "s"
    service            = "SIP+D2U"
    regular_expression = "hello"
    replacement        = "foobar.example.com."
    disable_flag       = "true"
  }
  type       = "NAPTR"
  gtd_region = 1
  note       = "First record"
  noanswer   = false

}

```

## Argument Reference ##
* `source_type` - (Required) Type of the Naptr record. The values which can be applied are "domains" or "templates".
* `ttl` - (Required) TTL must be in between 0 and 2147483647.
* `name` - (Optional) Name of record. Name should be unique.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type Naptr.
* `roundrobin` - (Required) Set.
* `roundrobin.order` - (Required) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.
* `roundrobin.preference` - (Required) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal). A 16-bit value ranging from 0 to 63535, the lowest number having the highest order.
* `roundrobin.flags` - (Required) A Flag is a single character from the set A-Z and 0-9, defined to be application specific, such that each application may define a specific use of the flag or which flags are valid. The flag is enclosed in quotes (“”). Currently defined values are: 
U – a terminal condition – the result of the regexp is a valid URI.
S – a terminal condition – the replace field contains the FQDN of an SRV record.
A – a terminal condition – the replace field contains the FQDN of an A or AAAA record.
P – a non-terminal condition – the protocol/services part of the params field determines the application specific behavior and subsequent processing is external to the record.
“” (empty string) – a non-terminal condition to indicate that regexp is empty and the replace field contains the FQDN of a further NAPTR record.
* `roundrobin.service` - (Required) Defines the application specific service parameters. The generic format is: protocol+rs. Where “protocol” defines the protocol used by the application and “rs” is the resolution service. There may be 0 or more resolution services each separated by +.
* `roundrobin.regular_expression` - (Required) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.
* `roundrobin.replacement` - (Required) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal).
* `roundrobin.disable_flag` - (Required) disable flag. Default is false

## Attributes Reference
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of the NAPTR resource.