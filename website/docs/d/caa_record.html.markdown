---
layout: "constellix"
page_title: "Constellix: constellix_caa_record"
sidebar_current: "docs-constellix-data-source-constellix_caa_record"
description: |-
  Data source for CAA record
---

# constellix_caa_record #
Data source for CAA record

## Example Usage ##

```hcl
data "constellix_caa_record" "datacaarecord" {
  domain_id   = "${constellix_domain.domain1.id}"
  name        = "anamerecorddatasource"
  source_type = "domains"
}

```
## Argument Reference ##
* `name` - (Required) Name of record. Name should be unique.
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `domain_id` - (Required) Record id of CAA record

## Attribute Reference ##
* `ttl` - (Optional) TTL must be in between 0 and 2147483647
* `source_type` - (Required) "domains" for Domain records and "template" for Template records
* `roundrobin` - (Optional) Set
* `roundrobin.caa_provider_id` - (Optional) 1 for [ Custom ], 2 for [ No Provider ], 3 for Comodo, 4 for Digicert, 5 for Entrust, 6 for GeoTrust, 7 for Izenpe, 8 for Lets Encrypt, 9 for Symantec, 10 for Thawte

* `roundrobin.tag` - (Optional) "issue" for Issue, "IssueWild" for IssueWild, "iodef" for iodef. Type allows you to choose how you want certificates to be issued by the CA. Each CAA record can contain only one tag-value pair. Options:
issue: Explicitly authorizes a single certificate authority to issue a certificate (any type) for the hostname.

issuewild: Authorization to issue certificates that specify a wildcard domain. Please note: issuewild properties take precedence over issue properties when specified.

iodef: (Incident Description Exchange Format) Specifies a means of reporting certificate issue requests or cases of certificate issue for the corresponding domain that violate the security policy of the issuer or the domain name holder.

* `roundrobin.data` - (Optional) "" for [ Custom ] if CAA provider Id is 1, ";" for [ No Provider ], "comodoca.com" for Comodo, "digicert.com" for Digicert, "entrust.net" for Entrust, "geotrust.com" for GeoTrust, "izenpe.com" for Izenpe, "letsencrypt.org" for Lets Encrypt, "symantec.com" for Symantec, "thawte.com" for Thawte

* `roundrobin.flag` - (Optional) roundRobin.Issuer Critical

There is currently only one flag defined, “issuer critical” at a value of 1. If a CA does not understand the flag value for an issuer critical record, then the CA will return with “no issue” for the certification.

All records will have the default issuer critical value of 0, which means they are “not critical”. Issuer Critical Value should be between 0 to 255.

* `roundrobin.disable_flag` - (Optional) Disable flag. Default is false

* `name` - (Required) Name of record. Name should be unique.
* `gtd_region` - (Optional) Shows id of GTD region in which record is to be created.
* `noanswer` - (Optional) Shows if record is enabled or disabled. Default is false (Active)
* `note` - (Optional) Record note
* `type` - (Optional) Record type CAA
