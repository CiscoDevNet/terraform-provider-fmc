package fmc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcFtdDeploy() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for deploying changes to FTD in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_ftd_deploy\" \"ftd\" {\n" +
			"    device = data.fmc_devices.ftd.id\n" +
			"    ignore_warning = false\n" +
			"    force_deploy = false\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcFtdDeployCreate,
		ReadContext:   resourceFmcFtdDeployRead,
		UpdateContext: resourceFmcFtdDeployCreate,
		DeleteContext: resourceFmcFtdDeployDelete,
		Schema: map[string]*schema.Schema{
			"device": {
				Type:     schema.TypeString,
				Required: true,
			},
			"force_deploy": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_warning": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceFmcFtdDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	device_id := d.Get("device").(string)
	device, err := c.GetFmcDeployableDevice(ctx, device_id)
	if err != nil {
		d.SetId(fmt.Sprintf("Device not in deployable state! No devices found for deployment with ID: %s", device_id))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Device not in deployable state!",
			Detail:   err.Error(),
		})
		return diags
	}
	object := FtdDeploy{
		Type:          deployment_type,
		Version:       device.Version,
		Forcedeploy:   d.Get("force_deploy").(bool),
		Ignorewarning: d.Get("ignore_warning").(bool),
		Devicelist:    []string{device_id},
	}
	if err := c.DeployToFTD(ctx, object); err != nil {
		d.SetId(fmt.Sprintf("Error in deployment, there might be another deployment in progress for device Name: %s ID: %s", device.Name, device_id))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Error in deployment, there might be another deployment in progress!",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(fmt.Sprintf("Deployment should now be in progress! Device Name: %s ID: %s", device.Name, device_id))
	return diags
}

func resourceFmcFtdDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = m.(*Client)
	// Invalidate state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func resourceFmcFtdDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = m.(*Client)
	// Invalidate state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
