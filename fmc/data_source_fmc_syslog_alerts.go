package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcSyslogAlerts() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Syslog Alert Configuration in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_security_zones\" \"inside\" {\n" +
			"	name = \"inside\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcSyslogAlertsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func dataSourceFmcSyslogAlertsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	syslogAlert, err := c.GetFmcSyslogAlertByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get syslog alert",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(syslogAlert.ID)

	if err := d.Set("name", syslogAlert.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read syslog alert",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", syslogAlert.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read syslog alert",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
