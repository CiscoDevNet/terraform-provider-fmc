package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcDeviceCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for FTD Device Cluster in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_cluster\" \"cluster\" {\n" +
			"	name = \"ftd-cluster1\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcDeviceClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cluster",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this resource",
			},
		},
	}
}

func dataSourceFmcDeviceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	device, err := c.GetFmcDeviceClusterByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get device cluster",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(device.ID)

	if err := d.Set("name", device.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device cluster",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", device.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device cluster",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}
