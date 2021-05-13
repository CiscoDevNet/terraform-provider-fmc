package fmc

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("fmc_username").(string)
	password := d.Get("fmc_password").(string)
	baseURL := d.Get("fmc_base_url").(string)
	var diags diag.Diagnostics

	if username != "" && password != "" && baseURL != "" {
		client := NewClient(username, password, baseURL)
		err := client.Login()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
	return nil, diag.FromErr(errors.New("Missing FMC username, password or base url"))
}

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"fmc_username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_USERNAME", nil),
			},
			"fmc_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_PASSWORD", nil),
			},
			"fmc_base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_BASE_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fmc_network_objects": resourceNetworkObjects(),
			"fmc_url_objects": resourceURLObjects(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
