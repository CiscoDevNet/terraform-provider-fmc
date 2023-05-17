package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcStandardAcl() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Standard accesslist in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_standard_acl\" \"std-acl\" {\n" +
			"	name = \"STD-ACL\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcStandardAclRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Standard Acl resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this resource",
			},
		},
	}
}

func dataSourceFmcStandardAclRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	acl, err := c.GetFmcStandardAclByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get Standard Acl",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(acl.ID)

	if err := d.Set("name", acl.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read standard Acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", acl.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read standard Acl",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
