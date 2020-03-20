---
layout: "constellix"
page_title: "CONSTELLIX: constellix_geo_proximity"
sidebar_current: "docs-constellix-data-source-constellix_geo_proximity"
description: |-
  Data source for Geoproximity for A, AAAA, CNAME or ANAME records. 
---

# constellix_geo_proximity
 Data source for Geoproximity for A, AAAA, CNAME or ANAME records. 

## Example Usage ##

```hcl
data "constellix_geo_proximity" "firstgeoproximity" {
  name = "practice"
}

```

## Argument Reference
* `name` - (Required) Geo Proximity name should be unique.

## Attribute Reference ##
* `country` - (Optional) Country code. Default is null.
* `region` - (Optional)Region or state or province code. Default is null.
* `latitude` - (Optional) Latitude value.
* `longitude` - (Optional) Longitude value.
* `city` - (Optional)City code. Default is null.

