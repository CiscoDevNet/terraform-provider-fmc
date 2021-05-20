package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFilePolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFilePoliciesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFilePoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	filePolicy, err := c.GetFilePolicyByName(ctx, d.Get("name").(string))

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
