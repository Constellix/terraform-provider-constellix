---
layout: "constellix"
page_title: "CONSTELLIX: constellix_cname_record"
sidebar_current: "docs-constellix-data-source-constellix_cname_record"
description: |-
  Data source for one or more domain CNAME records.
---

# constellix_cname_record
Data source one or more domain CNAME records.

## Example Usage ##

```hcl
data "constellix_cname_record" "firstrecord" {
  domain_id	 	= "${constellix_domain.first_domain.id}"
  source_type 	= "domains"
  name     		= "firstrecord"
}

```

## Argument Reference
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) Type of the CName record. The values which can be applied are "domains" or "templates".
* `domain_id` - (Required) Domain id of the CName record.


## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647.
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
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created. 1 for World (Default). 2 for Europe. 3 for US East. 4 for US West. 5 for Asia Pacific. 6 for Oceania. note: "gtdRegion" from 2 to 6 will be applied only when GTD region is enabled on domain.
* `type` - (Optional) Record type A.
* `contact_ids` - (Optional) Applied contact list id. Only applicable to record with type roundRobin with failover and failover.
* `pools` - (Optional) Ids of CNamepool.
* `record_failover` - (Optional) To create a record failover object pass the following attributes.
* `record_failover_values` - (Optional) Set. 
* `record_failover_values.value` - (Optional) Host name.
* `record_failover_values.checkid` - (Optional) Sonar check id.
* `record_failover_values.sort_order` - (Optional) Integer value which decides in which order the recordfailover should be sorted.
* `record_failover_values.disable_flag` - (Optional) enable or disable the recordFailover value object. Default is false (Active). Atleast one recordFailover value object should be false.
* `record_failover_failover_type` - (Optional) 1 for Normal (always lowest level). 2 for Off on any Failover event. 3 for One Way (move to higher level).
* `record_failover_disable_flag` - (Optional) enable or disable the recordFailover object. Default is false (Active). Atleast one recordFailover object should be false.

