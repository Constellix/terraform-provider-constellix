---
layout: "constellix"
page_title: "CONSTELLIX: constellix_ptr_record"
sidebar_current: "docs-constellix-data-source-constellix_ptr_record"
description: |-
  Data source for records of type PTR  for a specific domain.
---

# constellix_ptr_record
Data source for records of type PTR for a specific domain.

## Example Usage ##

```hcl
data "constellix_ptr_record" "ptr1"{
  domain_id		= "${constellix_domain.domain1.id}"
  source_type = "domains"
  name 			  = "pointer1"
}
```

## Argument Reference
* `source_type` - (Required) Type of the PTR record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the PTR record.

## Attributes Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type A.
* `roundrobin` - (Optional) Object.
* `roundrobin.value` - (Optional) This will be the host name of the computer or server the IP resolves to, for example mail.example.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).
* `roundrobin.disable_flag` - (Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.

