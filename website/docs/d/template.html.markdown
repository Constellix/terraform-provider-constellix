---
layout: "constellix"
page_title: "CONSTELLIX: constellix_template"
sidebar_current: "docs-constellix-data-source-constellix_template"
description: |-
   Data source for the DNS record templates.
---

# constellix_template
  Data source for the DNS record templates.

## Example Usage ##

```hcl
data "constellix_template" "firsttemplate" {
  name = "sample"
}


```

## Argument Reference
* `name` - (Required) Template names. e.g "sampletemplate".

## Attributes Reference ##
* `domain` - (Optional) Id of domain to be applied.
* `has_gtd_region` - (Optional) Enable/Disable GTD Region of the domain. The Default value is false.
* `has_geoip` - (Optional) Enable/Disable GEO IP. The Default value is false.
* `version` - (Optional) System generated template history version.

