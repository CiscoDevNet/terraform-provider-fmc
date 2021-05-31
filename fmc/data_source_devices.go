package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDevices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDevicesRead,
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

func dataSourceDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	device, err := c.GetDeviceByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get device",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(device.ID)

	if err := d.Set("name", device.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", device.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
