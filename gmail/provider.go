package gmail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOOGLE_CREDENTIALS", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gmail_filter": resourceFilter(),
		},
	}
}
