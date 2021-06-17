package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFilePolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for File Policies in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_file_policies\" \"file_policy\" {\n" +
			"	name = \"AMP Policy\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFilePoliciesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the file policy",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func dataSourceFilePoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	filePolicy, err := c.GetFmcFilePolicyByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get file policy",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(filePolicy.ID)

	if err := d.Set("name", filePolicy.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read file policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", filePolicy.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read file policy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
