package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcAccessPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Access Policies in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_access_policies\" \"acp\" {\n" +
			"	name = \"FTD ACP\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcAccessPoliciesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the FTD accessPolicy",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this resource",
			},
		},
	}
}

func dataSourceFmcAccessPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	version := c.GetVersions(ctx)
	accessPolicy, err := c.GetFmcAccessPolicyByName(ctx, d.Get("name").(string), version)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get accessPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(accessPolicy.ID)

	if err := d.Set("name", accessPolicy.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read accessPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", accessPolicy.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read accessPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
