variable "records" {
  description = "Map of records to have created in Constellix"
  type        = map(any)
}

variable "a_pools" {
  description = "Map of A record pools created in Constellix"
  type        = map(any)
}

variable "cname_pools" {
  description = "Map of CNAME record pools created in Constellix"
  type        = map(any)
}

variable "aaaa_pools" {
  description = "Map of AAAA record pools created in Constellix"
  type        = map(any)
}

variable "note" {
  description = "Note to add to records"
  type        = string
}

variable "domain_id" {
  description = "ID of the Constellix Domain to have resources created in"
  type        = string
}