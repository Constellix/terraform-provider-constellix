package main

import (
	"github.com/Constellix/terraform-provider-constellix/constellix"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: constellix.Provider})
}
