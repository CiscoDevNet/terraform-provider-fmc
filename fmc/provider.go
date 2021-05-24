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
	host := d.Get("fmc_host").(string)
	insecureSkipVerify := d.Get("fmc_insecure_skip_verify").(bool)
	var diags diag.Diagnostics

	if username != "" && password != "" && host != "" {
		client := NewClient(username, password, host, insecureSkipVerify)
		err := client.Login()
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
	return nil, diag.FromErr(errors.New("missing fmc username, password or base url"))
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
			"fmc_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FMC_BASE_URL", nil),
			},
			"fmc_insecure_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				DefaultFunc: schema.EnvDefaultFunc("FMC_INSECURE_SKIP_VERIFY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fmc_url_objects":           resourceURLObjects(),
			"fmc_url_object_group":      resourceURLObjectGroup(),
			"fmc_port_objects":          resourcePortObjects(),
			"fmc_network_objects":       resourceNetworkObjects(),
			"fmc_host_objects":          resourceHostObjects(),
			"fmc_range_objects":         resourceRangeObjects(),
			"fmc_fqdn_objects":          resourceFQDNObjects(),
			"fmc_icmpv4_objects":        resourceICMPV4Objects(),
			"fmc_access_rules":          resourceAccessRules(),
			"fmc_access_policies":       resourceAccessPolicies(),
			"fmc_network_group_objects": resourceNetworkGroupObjects(),
			"fmc_port_group_objects":    resourcePortGroupObjects(),
			"fmc_ftd_nat_policies":      resourceNatPolicies(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fmc_devices":         dataSourceDevices(),
			"fmc_ips_policies":    dataSourceIPSPolicies(),
			"fmc_file_policies":   dataSourceFilePolicies(),
			"fmc_syslog_alerts":   dataSourceSyslogAlerts(),
			"fmc_security_zones":  dataSourceSecurityZones(),
			"fmc_network_objects": dataSourceNetworkObjects(),
			"fmc_url_objects":     dataSourceURLObjects(),
			"fmc_port_objects":    dataSourcePortObjects(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
