module "constellix_domain" {
  source   = "./modules/constellix/domains"
  for_each = toset(local.domains)
  domain   = each.key
}

module "constellix_pools" {
  source = "./modules/constellix/pools"
  note   = local.note
  pools  = local.pools
}

module "constellix_records" {
  source      = "./modules/constellix/records"
  for_each    = module.constellix_domain
  records     = local.records
  note        = local.note
  a_pools     = module.constellix_pools.a_pool_info
  cname_pools = module.constellix_pools.cname_pool_info
  aaaa_pools  = module.constellix_pools.aaaa_pool_info
  domain_id   = each.value.domain_id
}