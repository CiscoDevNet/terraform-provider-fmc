package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataFmcIseSGTObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for ise security group tags object from pxgrid in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_ise_sgt_objects\" \"workstation\" {\n" +
			"	name = \"workstation\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcIseSGTObjectsRead,
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

func dataSourceFmcIseSGTObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	iseSGTObject, err := c.GetFmcIseSGTObjectsByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ise security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(iseSGTObject.ID)

	if err := d.Set("name", iseSGTObject.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ise security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", iseSGTObject.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ise security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
