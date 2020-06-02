package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-constellix/constellix"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: constellix.Provider})
}
