package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcNatPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Nat Policies in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_nat_policies\" \"natp\" {\n" +
			"	name = \"FTD NATP\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcNatPoliciesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the FTD natPolicy",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this resource",
			},
		},
	}
}

func dataSourceFmcNatPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	natPolicy, err := c.GetFmcNatPolicyByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get natPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(natPolicy.ID)

	if err := d.Set("name", natPolicy.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read natPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", natPolicy.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read natPolicy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
