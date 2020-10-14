package gmail

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
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
