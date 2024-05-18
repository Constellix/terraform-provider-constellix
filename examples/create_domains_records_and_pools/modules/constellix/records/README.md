# constellix-ops
Testing the constellix terraform repo

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
| [constellix_a_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/a_record) | resource |
| [constellix_aaaa_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/aaaa_record) | resource |
| [constellix_aname_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/aname_record) | resource |
| [constellix_caa_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/caa_record) | resource |
| [constellix_cert_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/cert_record) | resource |
| [constellix_cname_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/cname_record) | resource |
| [constellix_hinfo_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/hinfo_record) | resource |
| [constellix_http_redirection_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/http_redirection_record) | resource |
| [constellix_mx_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/mx_record) | resource |
| [constellix_naptr_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/naptr_record) | resource |
| [constellix_ns_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/ns_record) | resource |
| [constellix_ptr_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/ptr_record) | resource |
| [constellix_rp_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/rp_record) | resource |
| [constellix_spf_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/spf_record) | resource |
| [constellix_srv_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/srv_record) | resource |
| [constellix_txt_record.this](https://registry.terraform.io/providers/Constellix/constellix/latest/docs/resources/txt_record) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_a_pools"></a> [a\_pools](#input\_a\_pools) | Map of A record pools created in Constellix | `map(any)` | n/a | yes |
| <a name="input_aaaa_pools"></a> [aaaa\_pools](#input\_aaaa\_pools) | Map of AAAA record pools created in Constellix | `map(any)` | n/a | yes |
| <a name="input_cname_pools"></a> [cname\_pools](#input\_cname\_pools) | Map of CNAME record pools created in Constellix | `map(any)` | n/a | yes |
| <a name="input_domain_id"></a> [domain\_id](#input\_domain\_id) | ID of the Constellix Domain to have resources created in | `string` | n/a | yes |
| <a name="input_note"></a> [note](#input\_note) | Note to add to records | `string` | n/a | yes |
| <a name="input_records"></a> [records](#input\_records) | Map of records to have created in Constellix | `map(any)` | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->