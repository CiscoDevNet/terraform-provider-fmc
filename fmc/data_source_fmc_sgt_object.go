package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataFmcSGTObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for security group tags object from pxgrid in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_sgt_objects\" \"workstation\" {\n" +
			"	name = \"workstation\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcSGTObjectsRead,
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

func dataSourceFmcSGTObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	SGTObject, err := c.GetFmcSGTObjectsByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(SGTObject.ID)

	if err := d.Set("name", SGTObject.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ise security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", SGTObject.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read ise security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	// if err := d.Set("tag", securityGroupTag.Tag); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "unable to tag of security group tag",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	return diags
}
