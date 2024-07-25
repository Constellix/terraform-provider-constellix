locals {
  falttened_cname_values = flatten([
    for pool_name, pool in var.pools.cname : [
      for value in pool.values : {
        pool_name    = pool_name
        value        = value.value
        fqdn         = value.fqdn
        weight       = value.weight
        policy       = value.policy
        disable_flag = value.disable_flag
      }
    ]
  ])
}

resource "constellix_http_check" "this_cname" {
  for_each = { for value in local.falttened_cname_values : "${value.pool_name}_${value.value}" => value }

  name                = each.value.pool_name
  host                = each.value.value
  fqdn                = each.value.fqdn
  ip_version          = "IPV4"
  port                = 443
  protocol_type       = "HTTPS"
  check_sites         = [1, 2]
  interval            = "ONEMINUTE"
  interval_policy     = "ONCEPERSITE"
  verification_policy = "SIMPLE"
}


resource "constellix_cname_record_pool" "this" {
  for_each               = var.pools.cname
  name                   = each.key
  num_return             = "1"
  min_available_failover = 1

  dynamic "values" {
    for_each = each.value.values
    content {
      value        = values.value.value
      weight       = values.value.weight
      policy       = values.value.policy
      disable_flag = values.value.disable_flag
      check_id     = resource.constellix_http_check.this_cname["${each.key}_${values.value.value}"].id
    }
  }

  failed_flag  = false
  disable_flag = false
  note         = var.note
}