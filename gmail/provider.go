package gmail

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOOGLE_CLIENT_CREDENTIALS", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOOGLE_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gmail_filter": resourceFilter(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cfg := &Config{}

	client := d.Get("client").(string)
	if _, err := os.Stat(client); err == nil {
		bytes, err := ioutil.ReadFile(client)
		if err != nil {
			return nil, diag.Errorf("error reading client credentials file: %s", err)
		}

		if err := cfg.Client(bytes); err != nil {
			return nil, diag.Errorf("error parsing client credentials JSON: %s", err)
		}
	}

	token := d.Get("token").(string)
	if _, err := os.Stat(token); err == nil {
		bytes, err := ioutil.ReadFile(token)
		if err != nil {
			return nil, diag.Errorf("error reading token file: %s", err)
		}

		if err := cfg.Token(bytes); err != nil {
			return nil, diag.Errorf("error parsing token JSON: %s", err)
		}
	}

	return cfg, nil
}
