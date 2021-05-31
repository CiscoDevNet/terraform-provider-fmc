package fmc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFtdDeploy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFtdDeployCreate,
		ReadContext:   resourceFtdDeployRead,
		UpdateContext: resourceFtdDeployCreate,
		DeleteContext: resourceFtdDeployDelete,
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

func resourceFtdDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	device_id := d.Get("device").(string)
	device, err := c.GetDeployableDevice(ctx, device_id)
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
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error in deployment",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(fmt.Sprintf("Deployment should now be in progress! Device Name: %s ID: %s", device.Name, device_id))
	return diags
}

func resourceFtdDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = m.(*Client)
	// Invalidate state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func resourceFtdDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = m.(*Client)
	// Invalidate state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
