<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_constellix"></a> [constellix](#requirement\_constellix) | >= 0.4.5 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_constellix"></a> [constellix](#provider\_constellix) | >= 0.4.5 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [constellix_a_record_pool.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/a_record_pool) | resource |
| [constellix_aaaa_record_pool.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/aaaa_record_pool) | resource |
| [constellix_cname_record_pool.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/cname_record_pool) | resource |
| [constellix_http_check.this_a](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/http_check) | resource |
| [constellix_http_check.this_aaaa](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/http_check) | resource |
| [constellix_http_check.this_cname](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/http_check) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_note"></a> [note](#input\_note) | Note to add to records | `string` | n/a | yes |
| <a name="input_pools"></a> [pools](#input\_pools) | Map of pools to have created in Constellix | `map(any)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_a_pool_info"></a> [a\_pool\_info](#output\_a\_pool\_info) | n/a |
| <a name="output_aaaa_pool_info"></a> [aaaa\_pool\_info](#output\_aaaa\_pool\_info) | n/a |
| <a name="output_cname_pool_info"></a> [cname\_pool\_info](#output\_cname\_pool\_info) | n/a |
<!-- END_TF_DOCS -->