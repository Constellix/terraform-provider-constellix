# constellix-ops
Testing the constellix terraform repo

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_constellix"></a> [constellix](#requirement\_constellix) | >= 0.4.5 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_constellix_domain"></a> [constellix\_domain](#module\_constellix\_domain) | ./modules/constellix/domains | n/a |
| <a name="module_constellix_pools"></a> [constellix\_pools](#module\_constellix\_pools) | ./modules/constellix/pools | n/a |
| <a name="module_constellix_records"></a> [constellix\_records](#module\_constellix\_records) | ./modules/constellix/records | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_apikey"></a> [apikey](#input\_apikey) | constellix api key | `string` | n/a | yes |
| <a name="input_secretkey"></a> [secretkey](#input\_secretkey) | constellix secret key | `string` | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->