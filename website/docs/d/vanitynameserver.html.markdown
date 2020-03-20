---
layout: "constellix"
page_title: "Constellix: constellix_vanity_nameserver"
sidebar_current: "docs-constellix-data-source-constellix_vanity_nameserver"
description: |-
  Data source for Vanitynameserver record
---

# constellix_vanity_nameserver #
Data source for Vanitynameserver record


## Example Usage ##

```hcl
data "constellix_vanity_nameserver" "datavanitynameserver" {
  name = "vanitynameserverdatasource"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Vanity nameserver name should be unique.
* `nameserver_group` - (Optional) Name server group id. 1 .. Available nameserver groups
* `nameserver_list_string` - (Optional) Comma separedted name servers list
* `is_default` - (Optional) Default flag. Default is false.
* `is_public` - (Optional) isPublic flag. Default is false
* `nameserver_group_name` - (Optional) Name server group name