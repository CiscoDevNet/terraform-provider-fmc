package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcDynamicObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Dynamic Objects in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_dynamic_object\" \"dyobj\" {\n" +
			"	name = \"Dynamic Object 1\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcDynamicObjectsRead,
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

func dataSourceFmcDynamicObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	dynamicobject, err := c.GetFmcDynamicObjectByName(ctx, d.Get("name").(string))
	//    dynamicobject, err = c.GetFmcDynamicObject(ctx, idInput.(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(dynamicobject.ID)

	if err := d.Set("name", dynamicobject.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read Dynamic Object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", dynamicobject.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read Dynamic Object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
