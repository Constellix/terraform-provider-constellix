---
layout: "constellix"
page_title: "Constellix: constellix_tags"
sidebar_current: "docs-constellix-resource-constellix_tags"
description: |-
    Manages tags for Constellix DNS.
---
# constellix_tags #
Manages tags for Constellix DNS.

# Example Usage #
```hcl
        
resource "constellix_tags" "tags1" {
  name = "tagsdns"
}

```

## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.


## Attribute Reference ##
No attributes are exported