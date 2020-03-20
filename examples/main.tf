provider "constellix" {
  apikey    = ""
  secretkey = ""
}
//DATASOURCES
# data "constellix_domain" "domain1" {
#   name = "datasourcedomain.com"
# }

# data "constellix_aname_record" "dataanamerecord" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name        = "anamerecorddatasource"
# }

# data "constellix_caa_record" "datacaarecord" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   name        = "anamerecorddatasource"
#   source_type = "domains"
# }

# data "constellix_contact_lists" "contactlist" {
#   name = "contactlist1"
# }

# data "constellix_hinfo_record" "hinfo" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name        = "datahinforecord"
# }

# data "constellix_http_redirection_record" "datahttpredirection" {
#   name        = "httpredirectiondatasource"
#   source_type = "domains"
#   domain_id   = "${data.constellix_domain.domain1.id}"
# }

# data "constellix_mx_record" "datamx" {
#   name        = "mxdatasource"
#   source_type = "domains"
#   domain_id   = "${data.constellix_domain.domain1.id}"
# }

# data "constellix_rp_record" "datarp" {
#   name        = "rpdatasource"
#   source_type = "domains"
#   domain_id   = "${data.constellix_domain.domain1.id}"
# }

# data "constellix_srv_record" "datasrv" {
#   name        = "srvdatasource"
#   source_type = "domains"
#   domain_id   = "${data.constellix_domain.domain1.id}"
# }

# data "constellix_tags" "datatags" {
#   name = "tagsdatasource"
# }

# data "constellix_txt_record" "datatxt" {
#   name        = "txtdatasource"
#   source_type = "domains"
#   domain_id   = "${data.constellix_domain.domain1.id}"
# }

# data "constellix_vanity_nameserver" "datavanitynameserver" {
#   name = "vanitynameserverdatasource"
# }

# data "constellix_a_record" "firstrecord" {
#   domain_id     = "${data.constellix_domain.domain1.id}"
#   source_type   = "domains"
#   name        = "firstrecord"
# }

# data "constellix_aaaa_record" "firstrecord" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name        = "firstrecord"
# }

# data "constellix_aaaa_record_pool" "prac" {
#   name = "firstrecord"
# }

# data "constellix_a_record_pool" "prac" {
#   name = "firstrecord"
# }

# data "constellix_cert_record" "firstrecord" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name        = "firstrecord"
# }

# data "constellix_cname_record" "firstrecord" {
#   domain_id     = "${data.constellix_domain.domain1.id}"
#   source_type   = "domains"
#   name          = "firstrecord"
# }

# data "constellix_cname_record_pool" "prac" {
#   name = "firstrecord"
# }

# data "constellix_geo_filter" "firstgeofilter" {
#   name = "firstfilter"
# }

# data "constellix_geo_proximity" "firstgeoproximity" {
#   name = "practice"
# }

# data "constellix_naptr_record" "firstrecord" {
#   domain_id   = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name        = "firstrecord"
# }

# data "constellix_ns_record" "firstrecord" {
#   domain_id       = "${data.constellix_domain.domain1.id}"
#   source_type   = "domains"
#   name            = "firstrecord" 
# }

# data "constellix_ptr_record" "ptr1" {
#   domain_id     = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name            = "pointer1"
# }

# data "constellix_spf_record" "spf1" {
#   domain_id     = "${data.constellix_domain.domain1.id}"
#   source_type = "domains"
#   name            = "temp"
# }

# data "constellix_template" "firsttemplate" {
#   name = "sample"
# }

//RESOURCES
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

resource "constellix_aname_record" "aname_record1" {
  domain_id   = "${constellix_domain.domain1.id}"
  ttl         = 18
  name        = "anamerecord"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "ANAME"
  contact_ids = [123]
  source_type = "domains"
  roundrobin {
    value        = "www.whatsapp.com."
    disable_flag = false
  }
  roundrobin {
    value        = "www.info.com."
    disable_flag = false
  }
  record_option = "failover"
  record_failover_values {
    value        = "www.w3schools.com."
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "www.messenger.com."
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_values {
    value        = "www.gmail.com."
    sort_order   = 3
    disable_flag = "false"
  }
  record_failover_failover_type = 1
  record_failover_disable_flag  = "false"
}

resource "constellix_srv_record" "srvrecord1" {
  domain_id   = "${constellix_domain.domain1.id}"
  ttl         = 1800
  name        = "srvrecord"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "SRV"
  source_type = "domains"
  roundrobin {
    value        = "www.google.com"
    port         = 8888
    priority     = 65
    weight       = 20
    disable_flag = false
  }
}

resource "constellix_hinfo_record" "hinfo1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "hinforecord"
  ttl         = "1900"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "HINFO"
  roundrobin {
    cpu          = "quard core"
    os           = "linux2"
    disable_flag = "false"
  }
  roundrobin {
    cpu          = "abc"
    os           = "winddows"
    disable_flag = "true"
  }
}

resource "constellix_http_redirection_record" "http1" {
  domain_id        = "${constellix_domain.domain1.id}"
  source_type      = "domains"
  name             = "redirectionrecord"
  ttl              = 1800
  redirect_type_id = 1
  url              = "https://www.google.com"
  noanswer         = false
  note             = ""
  gtd_region       = 1
  type             = "HTTPRedirection"
  hardlink_flag    = false
  description      = false
  title            = ""
}

resource "constellix_caa_record" "caacheck" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "caarecord"
  ttl         = 1900
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "CAA"
  roundrobin {
    caa_provider_id = 3
    tag             = "issue"
    data            = "como.com"
    flag            = "0"
    disable_flag    = "false"
  }
  roundrobin {
    caa_provider_id = 4
    tag             = "issue"
    data            = "como01.com"
    flag            = "1"
    disable_flag    = "true"
  }
}

resource "constellix_mx_record" "mx1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "mxrecord"
  ttl         = "1900"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "MX"
  roundrobin {
    value        = "abc"
    level        = "100"
    disable_flag = "false"
  }
  roundrobin {
    value        = "dce"
    level        = "200"
    disable_flag = "true"
  }
}

resource "constellix_rp_record" "rp1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "rprecord"
  ttl         = "1900"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "RP"
  roundrobin {
    mailbox      = "one.com"
    txt          = "domain.com"
    disable_flag = "false"
  }
  roundrobin {
    mailbox      = "second.com"
    txt          = "two.com"
    disable_flag = "true"
  }
}


resource "constellix_tags" "tags1" {
  name = "tagsdns"
}

resource "constellix_vanity_nameserver" "vanitynameserver1" {
  name                   = "vanitynameserverrecord"
  nameserver_group       = 1
  nameserver_list_string = "www.google.com,\nwww.facebook.com,\nwww.instegram.com"
  is_default             = false
  is_public              = false
  nameserver_group_name  = "NS user group 1"
}
resource "constellix_txt_record" "txtrecord1" {
  domain_id   = "${constellix_domain.domain1.id}"
  ttl         = 1800
  name        = "txtrecord"
  noanswer    = false
  note        = ""
  gtd_region  = 1
  type        = "TXT"
  source_type = "domains"
  roundrobin {
    value        = "\"{\\\"cfg\\\":[{\\\"useAS\\\":0}]}\""
    disable_flag = false
  }
}

resource "constellix_contact_lists" "contactlist1" {
  name = "Contacts"
  email_addresses = [
    "user1@example.com",
    "user2@example.com"
  ]
}

resource "constellix_a_record" "firstrecord" {
  domain_id     = "${constellix_domain.domain1.id}"
  source_type   = "domains"
  record_option = "roundRobinFailover"
  ttl           = 100
  name          = "firstrecord"
  geo_location = {
    geo_ip_user_region = 1
    drop               = "false"
  }
  pools       = [123]
  contact_ids = [1234]
  type        = "A"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  roundrobin {
    value        = "5.45.25.35"
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "5.45.2.35"
    sort_order   = 1
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "5.45.25.3"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5.45.25.5"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5.45.25.5"
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_failover_type = 2
  record_failover_disable_flag  = "false"
}

resource "constellix_aaaa_record" "firstrecord" {
  domain_id     = "${constellix_domain.domain1.id}"
  source_type   = "domains"
  record_option = "roundRobinFailover"
  ttl           = 100
  name          = "firstrecord"
  geo_location = {
    geo_ip_user_region = 1
    drop               = "false"
  }
  pools       = [123]
  contact_ids = [1234]
  type        = "AAAA"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  roundrobin {
    value        = "5:0:0:0:0:0:0:6"
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "4:0:0:0:0:0:0:6"
    sort_order   = 1
    disable_flag = "false"
  }
  roundrobin_failover {
    value        = "3:0:0:0:0:0:0:6"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5:0:0:0:0:0:1:6"
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "5:0:0:0:1:0:0:6"
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_failover_type = 2
  record_failover_disable_flag  = "false"
}

resource "constellix_cname_record" "firstrecord" {
  domain_id     = "${constellix_domain.domain1.id}"
  source_type   = "domains"
  record_option = "failover"
  ttl           = 100
  name          = "arecordname350"
  host          = "abcd.com."
  geo_location = {
    geo_ip_user_region = 1
    drop               = "false"
  }
  pools       = [123]
  contact_ids = [1234]
  type        = "CNAME"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  record_failover_values {
    value        = "abc.com."
    sort_order   = 1
    disable_flag = "false"
  }
  record_failover_values {
    value        = "ab.com."
    sort_order   = 2
    disable_flag = "false"
  }
  record_failover_failover_type = 2
  record_failover_disable_flag  = "false"
}

resource "constellix_naptr_record" "firstrecord" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  roundrobin {
    order              = 10
    preference         = 100
    flags              = "s"
    service            = "SIP+D2U"
    regular_expression = "hello"
    replacement        = "foobar.example.com."
    disable_flag       = "true"
  }
  type       = "NAPTR"
  gtd_region = 1
  note       = "First record"
  noanswer   = false

}

resource "constellix_ns_record" "firstrecord" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  roundrobin {
    value        = "prac."
    disable_flag = "false"
  }
  type       = "NS"
  gtd_region = 1
  note       = "First record"
  noanswer   = false

}

resource "constellix_template" "firsttemplate" {
  name            = "sample"
  has_gtd_regions = "true"
  has_geoip       = "false"
}

resource "constellix_geo_filter" "ipfilter1" {
  name               = "first135"
  geoip_continents   = ["AS"]
  geoip_countries    = ["IN", "PK"]
  geoip_regions      = ["IN/BR", "IN/MP"]
  asn                = [1, 2]
  ipv4               = ["1.1.1.0/32", "1.1.2.2/32"]
  ipv6               = ["2:0:0:2:0:0:1:abc/128"]
  filter_rules_limit = 100
}

resource "constellix_cert_record" "firstrecord" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  ttl         = 100
  name        = "firstrecord"
  type        = "CERT"
  gtd_region  = 1
  note        = "First record"
  noanswer    = false
  roundrobin {
    certificate_type = 20
    key_tag          = 30
    certificate      = "certificate1"
    algorithm        = 100
    disable_flag     = "true"
  }
}

resource "constellix_a_record_pool" "firstrecord" {
  name                   = "firstrecord"
  num_return             = "10"
  min_available_failover = 1
  values {
    value        = "8.1.1.1"
    weight       = 20
    policy       = "followsonar"
    disable_flag = false
  }
  failed_flag  = false
  disable_flag = false
  note         = "First record"
}

resource "constellix_cname_record_pool" "firstrecord" {
  name                   = "firstrecord"
  num_return             = "10"
  min_available_failover = 1
  values {
    value        = "8.1.1.1"
    weight       = 20
    policy       = "followsonar"
    disable_flag = false
  }
  failed_flag  = false
  disable_flag = false
  note         = "First record"
}

resource "constellix_ptr_record" "ptr1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "pointer1"
  ttl         = "10"
  note        = "Practice record"
  noanswer    = false
  gtd_region  = 1
  type        = "PTR"
  roundrobin {
    value        = 13
    disable_flag = "true"
  }
}

resource "constellix_spf_record" "spf1" {
  domain_id   = "${constellix_domain.domain1.id}"
  source_type = "domains"
  name        = "temp"
  ttl         = 10
  noanswer    = false
  gtd_region  = 1
  type        = "SPF"
  note        = "Practice record"
  roundrobin {
    value        = "124.56.8.1"
    disable_flag = "false"
  }

}

resource "constellix_geo_proximity" "firstgeoproximity" {
  name      = "practice"
  latitude  = "22.7"
  longitude = "56.8333"
  region    = "05"
  city      = "273890"
  country   = "OM"
}

resource "constellix_aaaa_record_pool" "firstrecord" {
  name                   = "firstrecord"
  num_return             = "10"
  min_available_failover = 1
  values {
    value        = "0:0:0:0:0:0:0:12"
    weight       = 20
    policy       = "followsonar"
    disable_flag = false
  }
  failed_flag  = false
  disable_flag = false
  note         = "First record"
}
