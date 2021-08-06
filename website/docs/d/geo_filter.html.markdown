---
layout: "constellix"
page_title: "CONSTELLIX: constellix_geo_filter"
sidebar_current: "docs-constellix-data-source-constellix_geo_filter"
description: |-
  Data source for Geofilters for A, AAAA, CNAME or ANAME records. 
---

# constellix_geo_filter
 Data source for Geofilters for A, AAAA, CNAME or ANAME records. 

## Example Usage ##

```hcl
data "constellix_geo_filter" "firstgeofilter" {
  name = "firstfilter"
}

```

## Argument Reference
* `name` - (Required) Geo Filter name should be unique.

## Attribute Reference ##
* `geoip_continents` - (Optional) Two digit Continents Code.
* `geoip_countries` - (Optional)Two digit Countries Code.
* `geoip_regions` - (Optional) Two digit country code followed by "/" followed by two digit region code. 
* `asn` - (Optional) Autonomous System Number (ASN). ASN code should be a number between 0 and 4294967295.
* `ipv4` - (Optional) IPV4 Address.
* `ipv6` - (Optional) IPV6 Address.
* `filter_rules_limit` - (Optional) Default is 100. For more than 100 rules, parameter should be set explicitly for ADD and Update API calls. Value should be in mulitple of 100 like 200, 300 ...upto the quota limit assigned to the account. Check quota details for IP Filter Rule Limit.

