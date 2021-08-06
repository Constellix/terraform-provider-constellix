---
layout: "constellix"
page_title: "CONSTELLIX: constellix_aaaa_record"
sidebar_current: "docs-constellix-data-source-constellix_aaaa_record"
description: |-
  Data source for AAAA record.
---

# constellix_aaaa_record
Data source for AAAA record.

## Example Usage ##

```hcl
data "constellix_aaaa_record" "firstrecord" {
  domain_id   = "${constellix_domain.first_domain.id}"
  source_type = "domains"
  name        = "firstrecord"
}


```

## Argument Reference
* `source_type` - (Required) Type of the AAAA record. The values which can be applied are "domains" or "templates".
* `name` - (Required) Name of record. Name should be unique.
* `domain_id` - (Required) Domain id of the AAAA record.

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
* `roundrobin` - (Optional) Object.
* `roundrobin.value` - (Optional) IPv6 address.
* `roundrobin.disable_flag` - (Optional) enable or disable the roundrobin object. Default is false. Atleast one roundrobin object should be false.
* `geo_location` - (Optional) Details of IP filter / Geo proximity to be applied. Default is null.
* `geo_location.geo_ip_user_region` - (Optional) For Geo proximity to be applied. geoipUserRegion should not be provided.
* `geo_location.drop` - (Optional) drop flag. Default is false.
* `geo_location.geo_ip_proximity` - (Optional) a valid geoipProximity id.
* `geo_location.geo_ip_user_region` - (Optional) For Geo IP Filter to be applied. geoipUserRegion should be [1].
* `geo_location.drop` - (Optional) drop flag. Default is false.
* `geo_location.geo_ip_failover` - (Optional) Flag to enable/disable Failover to nearest proximity when all the host fails. Works with the record type pools. It requires Geo Proximity to be enabled at the Domain level. Default is false. 
* `geo_location.geo_ip_proximity` - (Optional) for Geo IP Filter, geoipProximity must not be provided. please create an A record with "World (Default)" IP Filter first before a more specific IP Filter is applied. The "World (Default)" record would only be used if no matching Filter or Proximity records are found.
* `record_option` - (Optional) Type of record. "roundRobin" for Standard record (Default). "failover" for Failover. "pools" for Pools. "roundRobinFailover" for Round Robin with Failover.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active).
* `note` - (Optional)Record note.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `type` - (Optional) Record type AAAA.
* `contact_ids` - (Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.
* `pools` - (Optional) Ids of AAAApool.
* `roundrobin_failover` - (Optional) Set.
* `roundrobin_failover.value` - (Optional) IPv6 address.
* `roundrobin_failover.disable_flag` - (Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.
* `roundrobin_failover.sort_order` - (Optional) Integer value which decides in which order the roundrobinfailover should be sorted.
* `record_failover` - (Optional) To create a record failover object pass the following attributes.
* `record_failover_values` - (Optional) Set. 
* `record_failover_values.value` - (Optional) IPv6 address.
* `record_failover_values.check_id` - (Optional) Sonar check id.
* `record_failover_values.sort_order` - (Optional) Integer value which decides in which order the recordfailover should be sorted
* `record_failover_values.disable_flag` - (Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.
* `record_failover_failover_type` - (Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).
* `record_failover_disable_flag` - (Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.
