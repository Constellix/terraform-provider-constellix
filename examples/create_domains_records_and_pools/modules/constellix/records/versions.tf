terraform {
  required_version = ">= 1.0.0"

  required_providers {
    constellix = {
      source  = "Constellix/constellix"
      version = ">= 0.4.5"
    }
  }
}