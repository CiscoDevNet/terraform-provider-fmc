package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcNetworkGroupObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Network Group Objects in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_network_group_objects\" \"PrivateGroup\" {\n" +
			"  name = \"PrivateGroup\"\n" +
			"  description = \"Terraform private group\"\n" +
			"  objects {\n" +
			"      id = data.fmc_network_objects.PrivateVLAN.id\n" +
			"      type = data.fmc_network_objects.PrivateVLAN.type\n" +
			"  }\n" +
			"```",

		ReadContext: dataSourceNetworkGroupObjectsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the network group object",
			},
			// added code here
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of this resource",
			},	

			// added code here
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},			
		},
	}
}

func dataSourceNetworkGroupObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ifc, err := c.GetFmcNetworkGroupObjectByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get network group object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ifc.ID)

	if err := d.Set("name", ifc.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get network group object",
			Detail:   err.Error(),
		})
		return diags
	}
	// added code here
	if err := d.Set("type", ifc.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network group object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
