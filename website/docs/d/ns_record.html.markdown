---
layout: "constellix"
page_title: "CONSTELLIX: constellix_ns_record"
sidebar_current: "docs-constellix-data-source-constellix_ns_record"
description: |-
  Data source for records of type NS  for a specific domain.
---

# constellix_ns_record
 Data source for records of type NS  for a specific domain.

## Example Usage ##

```hcl
data "constellix_ns_record" "firstrecord" {
  domain_id	 	  = "${constellix_domain.first_domain.id}"
  source_type 	= "domains"
  name     		  = "firstrecord" 
}

```

## Argument Reference
* `source_type` - (Required) Type of the NS record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the NS record.

## Attributes Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type NS.
* `roundrobin` - (Optional) Set.
* `roundrobin.value` - (Optional) This will be the host name for the name server, for example ns0.nameserver.com. It is important to note, the domain name is automatically appended to the end of this field unless it ends with a dot (.).
* `roundrobin.disable_flag` - (Optional) disable flag. Default is false

