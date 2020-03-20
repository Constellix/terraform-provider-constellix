---
layout: "constellix"
page_title: "Constellix: constellix_contact_lists"
sidebar_current: "docs-constellix-resource-constellix_contact_lists"
description: |-
    Manages Constellix DNS contact lists.
---
# constellix_contact_lists #
Manages Constellix DNS contact lists.

# Example Usage #
```hcl

resource "constellix_contact_lists" "contactlist1" {
  name = "Contacts"
  email_addresses = [
    "user1@example.com",
    "user2@example.com"
  ]
}

```

## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `email_addresses` - (Required) List of email addresses

## Attribute Reference ##
No attributes are exported