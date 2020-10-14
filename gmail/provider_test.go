package gmail

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var testAccProviders map[string]*schema.Provider

func init() {
	testAccProviders = map[string]*schema.Provider{
		"gmail": Provider(),
	}
}
