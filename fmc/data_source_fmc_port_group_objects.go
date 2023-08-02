package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcPortGroupObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Port Group Objects in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_port_group_objects\" \"test\" {\n" +
			"	name = \"test-object-group\"\n" +
			"}\n" +
			"```",

		ReadContext: dataSourcePortGroupObjectsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the port group object",
			},
		},
	}
}

func dataSourcePortGroupObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ifc, err := c.GetFmcPortGroupObjectByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ifc.ID)

	if err := d.Set("name", ifc.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
