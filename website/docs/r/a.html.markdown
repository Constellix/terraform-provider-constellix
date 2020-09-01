---
layout: "constellix"
page_title: "CONSTELLIX: constellix_a_record"
sidebar_current: "docs-constellix-resource-constellix_a_record"
description: |-
  Manages one or more A records within the specified domain.
---

# constellix_a_record
Manages one or more A records within the specified domain.

## Example Usage ##

```hcl
resource "constellix_a_record" "firstrecord" {
  domain_id     = "${constellix_domain.first_domain.id}"
  source_type   = "domains"
  record_option = "roundRobinFailover"
  ttl           = 100
  name          = "firstrecord"
  geo_location = {
    geo_ip_user_region = 1
    drop               = "false"
  }
  pools       = [123]
  contact_ids = [1234]
  type        = "A"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  roundrobin {
    value        = "5.45.25.35"
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "5.45.2.35"
    sort_order   = 1
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "5.45.25.3"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5.45.25.5"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5.45.25.5"
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_failover_type = 2
  record_failover_disable_flag  = "false"
}

```

## Argument Reference ##
* `source_type` - (Required) Type of the A record. The values which can be applied are "domains" or "templates".
* `ttl` - (Required) TTL must be in between 0 and 2147483647.
* `roundrobin` - (Required) Object.
* `roundrobin.value` - (Required) IPv4 address.
* `roundrobin.disable_flag` - (Required) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `name` - (Optional) Name of record. Name should be unique.
* `geo_location` - (Optional) Details of IP filter / Geo proximity to be applied. Default is null.
* `geo_location.geo_ip_user_region` - (Optional) For Geo proximity to be applied. geoipUserRegion should not be provided.
* `geo_location.drop` - (Optional) drop flag. Default is false.
* `geo_location.geo_ip_proximity` - (Optional) a valid geoipProximity id.
* `geo_location.geo_ip_user_region` - (Optional) For Geo IP Filter to be applied. geoipUserRegion should be [1].
* `geo_location.drop` - (Optional) drop flag. Default is false.
* `geo_location.geo_ip_proximity` - (Optional) for Geo IP Filter, geoipProximity must not be provided. please create an A record with "World (Default)" IP Filter first before a more specific IP Filter is applied. The "World (Default)" record would only be used if no matching Filter or Proximity records are found.
* `record_option` - (Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover. "pools" for Pools. "roundRobinFailover" for Round Robin with Failover.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type A.
* `contact_ids` - (Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.
* `pools` - (Optional) Ids of Apool.
* `roundrobin_failover` - (Optional) Set.
* `roundrobin_failover.value` - (Required) IPv4 address.
* `roundrobin_failover.disable_flag` - (Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.
* `roundrobin_failover.sort_order` - (Required) Integer value which decides in which order the rounrobinfailover should be sorted.
* `record_failover` - (Optional) To create a record failover object pass the following attributes.
* `record_failover_values` - (Required) Set. 
* `record_failover_values.value` - (Required) IPv4 address.
* `record_failover_values.check_id` - (Optional) Sonar check id.
* `record_failover_values.sort_order` - (Required) Integer value which decides in which order the recordfailover should be sorted.
* `record_failover_values.disable_flag` - (Required) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.
* `record_failover_failover_type` - (Required) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).
* `record_failover_disable_flag` - (Required) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.

## Attributes Reference
The only attribute that this resource exports is the `id`, which is set to the constellix calculated id of the A resource.

## Importing ##

An existing Record can be [imported][docs-import] into this resource using its Id, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import constellix_a_record.example <source>:<parent-id>:<record-id>
```

Where source can be either domains or templates; parent-id is domain-id or template-id based on the source provided and record-id is the Id of record calculated via Constellix API.