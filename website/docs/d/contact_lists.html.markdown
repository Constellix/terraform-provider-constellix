---
layout: "constellix"
page_title: "Constellix: constellix_contact_lists"
sidebar_current: "docs-constellix-data-source-constellix_contact_lists"
description: |-
  Data source for Contact List
---

# constellix_contact_lists #
Data source for Contact List


## Example Usage ##

```hcl
data "constellix_contact_lists" "contactlist" {
  name = "contactlist1"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `email_addresses` - (Optional) List of email addresses
