package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPSPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPSPoliciesRead,
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

func dataSourceIPSPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ipsPolicy, err := c.GetIPSPolicyByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get ips policy",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ipsPolicy.ID)

	if err := d.Set("name", ipsPolicy.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ips policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", ipsPolicy.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ips policy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
