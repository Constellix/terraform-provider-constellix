variable "pools" {
  description = "Map of pools to have created in Constellix"
  type        = map(any)
}

variable "note" {
  description = "Note to add to records"
  type        = string
}
