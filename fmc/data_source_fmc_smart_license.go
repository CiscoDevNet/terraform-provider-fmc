package fmc

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcSmartLicense() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Smart License in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_smart_license\" \"smart_license\" {}\n" +
			"```",
		ReadContext: dataSourceFmcSmartLicenseRead,
		Schema: map[string]*schema.Schema{
			"registration_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The registration status of FMC",
			},
			"evaluation_used": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Status of whether evaluation period used",
			},
			"eval_expires_in_days": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "If evaluation period is active, gives the remaining days",
			},
			"virtual_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Virtual account associated with the registered token",
			},
		},
	}
}

func dataSourceFmcSmartLicenseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	smartLicense, err := c.GetFmcSmartLicense(ctx)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if err := d.Set("registration_status", smartLicense.RegistrationStatus); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("evaluation_used", smartLicense.EvalUsed); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("eval_expires_in_days", smartLicense.EvalExpiresInDays); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("virtual_account", smartLicense.VirtualAccount); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
