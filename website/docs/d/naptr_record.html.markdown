---
layout: "constellix"
page_title: "CONSTELLIX: constellix_naptr_record"
sidebar_current: "docs-constellix-data-source-constellix_naptr_record"
description: |-
  Data source for records of type NAPTR for a specific domain.
---

# constellix_naptr_record
 Data source for records of type NAPTR for a specific domain.

## Example Usage ##

```hcl
data "constellix_naptr_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  name        = "firstrecord"
}

```

## Argument Reference
* `source_type` - (Required) Type of the NAPTR record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the NAPTR record.

## Attributes Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type Naptr.
* `roundrobin` - (Optional) Set.
* `roundrobin.order` - (Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.
* `roundrobin.preference` - (Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal). A 16-bit value ranging from 0 to 63535, the lowest number having the highest order.
* `roundrobin.flags` - (Optional) A Flag is a single character from the set A-Z and 0-9, defined to be application specific, such that each application may define a specific use of the flag or which flags are valid. The flag is enclosed in quotes (“”). Currently defined values are: 
U – a terminal condition – the result of the regexp is a valid URI.
S – a terminal condition – the replace field contains the FQDN of an SRV record.
A – a terminal condition – the replace field contains the FQDN of an A or AAAA record.
P – a non-terminal condition – the protocol/services part of the params field determines the application specific behavior and subsequent processing is external to the record.
“” (empty string) – a non-terminal condition to indicate that regexp is empty and the replace field contains the FQDN of a further NAPTR record.
* `roundrobin.service` - (Optional) Defines the application specific service parameters. The generic format is: protocol+rs. Where “protocol” defines the protocol used by the application and “rs” is the resolution service. There may be 0 or more resolution services each separated by +.
* `roundrobin.regular_expression` - (Optional) A 16-bit value ranging from 0 to 63535, the lowest number having the highest order. For example, an order of 10 is of more importance (has a higher order value) than an order of 50.
* `roundrobin.replacement` - (Optional) Preference is used only when two NAPTR records with the same name also have the same order and is used to indicate preference (all other things being equal).
* `roundrobin.disable_flag` - (Optional) disable flag. Default is false

