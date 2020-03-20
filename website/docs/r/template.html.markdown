---
layout: "constellix"
page_title: "CONSTELLIX: constellix_template"
sidebar_current: "docs-constellix-resource-constellix_template"
description: |-
   Manages the DNS record templates.
---

# constellix_template
  Manages the DNS record templates.

## Example Usage ##

```hcl
resource "constellix_template" "firsttemplate" {
  name            = "sample"
  has_gtd_regions = "true"
  has_geoip       = "false"
}



```

## Argument Reference ##
* `name` - (Required) Template names. e.g "sampletemplate".
* `domain` - (Optional) Id of domain to be applied.
* `has_gtd_regions` - (Optional) Enable/Disable GTD Region of the domain. The Default value is false.
* `has_geoip` - (Optional) Enable/Disable GEO IP. The Default value is false.
* `version` - (Optional) System generated template history version.

## Attributes Reference
No attributes are exported.