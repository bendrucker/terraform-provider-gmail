package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/bendrucker/terraform-provider-gmail/gmail"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gmail.Provider,
	})
}
