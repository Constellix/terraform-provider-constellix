---
layout: "constellix"
page_title: "Provider: Constellix"
sidebar_current: "docs-constellix-index"
description: |-
  The Constellix provider is used to manage various DNS Objects supported by Constellix DNS platform. The provider needs to be configured with the proper credentials before it can be used.
---
Constellix Provider
------------
Constellix is a leading DNS service provider with a feature rich DNS services which includes, various kinds of dns records such as Aname record, Cname record, HTTPredirection, MX record, a granular filtering method using Geofilter and Geoproximity. The Constellix provider is used to manage various DNS Objects supported by Constellix DNS platform. The provider needs to be configured with the proper credentials before it can be used.

Authentication
--------------
The Provider supports authentication with Constellix using API-key and SECRET-key. 

 1. Authentication with user-id and password.  
 example:  

 ```hcl
provider "constellix" {
  # constellix Api key
  apikey    = "apikey"
  # cosntellix secret key
  secretkey = "secretkey"
  insecure  = true
  proxy_url = "https://proxy_server:proxy_port"
}
 ```

Example Usage
------------
```hcl
provider "constellix" {
  # constellix Api key
  apikey    = "apikey"
  # cosntellix secret key
  secretkey = "secretkey"
  insecure  = true
  proxy_url = "https://proxy_server:proxy_port"
}

resource "constellix_domain" "domain1" {
  name = "domain1.com"
  soa = {
    primary_nameserver = "ns41.constellix.com."
    ttl                = 1800
    refresh            = 48100
    retry              = 7200
    expire             = 1209
    negcache           = 8000
  }
}
```

Argument Reference
------------------
Following arguments are supported with Constellix terraform provider.

 * `apikey` - (Required) API key of a user which has the access to perform CRUD operations on all the DNS objects of Constellix platform.
 * `secretkey` - (Required) Secret key of a user which has the access to perform CRUD operations on all the DNS objects of Constellix platform.
 * `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.  
 * `proxy_url` - (Optional) A proxy server URL when configured, all the requests to Constellix platform will be passed through the proxy-server configured.
