package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcRavpn() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for RAVPN policy in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_ravpn_policy\" \"ravpn1\" {\n" +
			"	name = \"ravpn1\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcRavpnRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the FTD device",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this resource",
			},
		},
	}
}

func dataSourceFmcRavpnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ravpn, err := c.GetFmcRavpnByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get ravpn",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ravpn.ID)

	if err := d.Set("name", ravpn.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ravpn",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", ravpn.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ravpn",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
