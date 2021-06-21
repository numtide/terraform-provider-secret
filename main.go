package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/numtide/terraform-provider-secret/secret"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: secret.Provider})
}
