package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcDevicePhysicalInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Physical interfaces in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_physical_interfaces\" \"acp\" {\n" +
			"	name = \"FTD ACP\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcDevicePhysicalInterfacesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this physical interfaces",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the physical interfaces",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this physical interfaces",
			},
		},
	}
}

func dataSourceFmcDevicePhysicalInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	physicalinterfaces, err := c.GetPhysicalInterfacesByName(ctx, d.Get("actId").(string), d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(physicalinterfaces.ID)

	if err := d.Set("name", physicalinterfaces.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", physicalinterfaces.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
