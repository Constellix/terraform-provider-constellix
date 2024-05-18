resource "constellix_a_record" "this" {
  for_each      = var.records.a
  domain_id     = var.domain_id
  source_type   = "domains"
  record_option = "roundRobin"
  ttl           = 100
  name          = each.key == "root" ? "" : each.key
  dynamic "roundrobin" {
    for_each = each.value
    content {
      value        = roundrobin.value
      disable_flag = false
    }
  }
  note = var.note
}


resource "constellix_aaaa_record" "this" {
  for_each      = var.records.aaaa
  domain_id     = var.domain_id
  source_type   = "domains"
  record_option = "roundRobin"
  ttl           = 100
  name          = each.key == "root" ? "" : each.key
  type          = "AAAA"
  gtd_region    = 1

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value        = roundrobin.value
      disable_flag = false
    }
  }

  note = var.note
}

resource "constellix_aname_record" "this" {
  for_each      = var.records.aname
  domain_id     = var.domain_id
  source_type   = "domains"
  record_option = "roundRobin"
  ttl           = 100
  name          = each.key == "root" ? "" : each.key
  type          = "ANAME"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value = roundrobin.value
    }
  }

  note = var.note
}

resource "constellix_cname_record" "this" {
  for_each      = var.records.cname
  domain_id     = var.domain_id
  source_type   = "domains"
  record_option = "roundRobin"
  ttl           = 100
  name          = each.key
  host          = each.value[0]
  type          = "CNAME"
  note          = var.note
}

resource "constellix_hinfo_record" "this" {
  for_each    = var.records.srv
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "HINFO"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      cpu = split(",", roundrobin.value)[0]
      os  = split(",", roundrobin.value)[1]
    }
  }
}

resource "constellix_http_redirection_record" "this" {
  for_each         = var.records.httpredirection
  domain_id        = var.domain_id
  source_type      = "domains"
  name             = each.key == "root" ? "" : each.key
  ttl              = 1800
  redirect_type_id = 1
  url              = each.value[0]
  type             = "HTTPRedirection"
}

resource "constellix_mx_record" "this" {
  for_each    = var.records.mx
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "MX"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      level = split(",", roundrobin.value)[0]
      value = split(",", roundrobin.value)[1]
    }
  }
}


resource "constellix_naptr_record" "this" {
  for_each    = var.records.naptr
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "NAPTR"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      order              = split(",", roundrobin.value)[0]
      preference         = split(",", roundrobin.value)[1]
      flags              = split(",", roundrobin.value)[2]
      service            = split(",", roundrobin.value)[3]
      regular_expression = split(",", roundrobin.value)[4]
      replacement        = split(",", roundrobin.value)[5]
      disable_flag       = false
    }
  }
}

resource "constellix_caa_record" "this" {
  for_each    = var.records.caa
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "CAA"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      caa_provider_id = split(",", roundrobin.value)[0]
      tag             = split(",", roundrobin.value)[1]
      data            = split(",", roundrobin.value)[2]
      flag            = split(",", roundrobin.value)[3]
      disable_flag    = false
    }
  }
}

resource "constellix_cert_record" "this" {
  for_each    = var.records.cert
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "CERT"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      certificate_type = split(",", roundrobin.value)[0]
      key_tag          = split(",", roundrobin.value)[1]
      certificate      = split(",", roundrobin.value)[2]
      algorithm        = split(",", roundrobin.value)[3]
      disable_flag     = false
    }
  }
}

resource "constellix_ptr_record" "this" {
  for_each    = var.records.ptr
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "PTR"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value        = roundrobin.value
      disable_flag = false
    }
  }
}

resource "constellix_rp_record" "this" {
  for_each    = var.records.rp
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  source_type = "domains"
  type        = "RP"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      mailbox      = split(",", roundrobin.value)[0]
      txt          = split(",", roundrobin.value)[1]
      disable_flag = ""
    }
  }
}

resource "constellix_spf_record" "this" {
  for_each    = var.records.spf
  domain_id   = var.domain_id
  source_type = "domains"
  ttl         = 100
  name        = each.key == "root" ? "" : each.key

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value = roundrobin.value
    }
  }

  type = "SPF"
  note = var.note
}

resource "constellix_srv_record" "this" {
  for_each    = var.records.srv
  domain_id   = var.domain_id
  ttl         = 1800
  name        = each.key == "root" ? "" : each.key
  type        = "SRV"
  source_type = "domains"

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value    = split(",", roundrobin.value)[0]
      port     = split(",", roundrobin.value)[1]
      priority = split(",", roundrobin.value)[2]
      weight   = split(",", roundrobin.value)[3]
    }
  }
}

resource "constellix_txt_record" "this" {
  for_each    = var.records.txt
  domain_id   = var.domain_id
  source_type = "domains"
  ttl         = 100
  name        = each.key == "root" ? "" : each.key

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value = roundrobin.value
    }
  }

  type = "TXT"
  note = var.note
}

resource "constellix_ns_record" "this" {
  for_each    = var.records.ns
  domain_id   = var.domain_id
  source_type = "domains"
  ttl         = 100
  name        = each.key == "root" ? "" : each.key

  dynamic "roundrobin" {
    for_each = each.value
    content {
      value        = roundrobin.value
      disable_flag = "false"
    }
  }

  type = "NS"
  note = var.note
}
