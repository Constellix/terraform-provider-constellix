output "a_pool_info" {
  value = {
    for key, pool in constellix_a_record_pool.this : key => {
      pool_id = pool.id,
      ips     = [for value in pool.values : value.value]
    }
  }
}

output "cname_pool_info" {
  value = {
    for key, pool in constellix_cname_record_pool.this : key => {
      pool_id = pool.id,
      ips     = [for value in pool.values : value.value]
    }
  }
}

output "aaaa_pool_info" {
  value = {
    for key, pool in constellix_aaaa_record_pool.this : key => {
      pool_id = pool.id,
      ips     = [for value in pool.values : value.value]
    }
  }
}