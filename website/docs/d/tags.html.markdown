---
layout: "constellix"
page_title: "Constellix: constellix_tags"
sidebar_current: "docs-constellix-data-source-constellix_tags"
description: |-
  Data source for Tags
---

# constellix_tags #
Data source for Tags


## Example Usage ##

```hcl
data "constellix_tags" "datatags" {
  name = "tagsdatasource"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of record. Name should be unique.
