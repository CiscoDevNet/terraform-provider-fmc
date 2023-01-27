package fmc

import (
	"context"
	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcDevices() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Access Control Policies in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_access_policies\" \"access_policy\" {\n" +
			"    name = \"Terraform Access Policy\"\n" +
			"    # default_action = \"block\" # Cannot have block with base IPS policy\n" +
			"    default_action = \"permit\"\n" +
			"    default_action_base_intrusion_policy_id = data.fmc_ips_policies.ips_policy.id\n" +
			"    default_action_send_events_to_fmc = \"true\"\n" +
			"    default_action_log_end = \"true\"\n" +
			"    default_action_syslog_config_id = data.fmc_syslog_alerts.syslog_alert.id\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcDeviceCreate,
		ReadContext:   resourceFmcDeviceRead,
		UpdateContext: resourceFmcDeviceUpdate,
		DeleteContext: resourceFmcDeviceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of FTD",
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The hostname of FTD",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of this resource",
			},
			"regkey": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Same regkey as entered in FTD",
			},
			"nat_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "NAT_ID is required if configured in FTD",
			},
			"license_caps": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "License caps for this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				  },
			},
			"access_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
					},
				},
				Description: "access policy for this resource",
			},
		},
	}
}

func resourceFmcDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	// var accp []AccessPolicyItem
	// if inputaccp, ok := d.GetOk("access_policy"); ok {
	// 	for _, acc := range inputaccp.([]interface{}) {
	// 		acci := acc.(map[string]interface{})
	// 		accp = append(accp, AccessPolicyItem{
	// 			ID:   acci["id"].(string),
	// 			Type: acci["type"].(string),
	// 		})
	// 	}
	//  }

	var accpolicy *AccessPolicyItem
	dynamicObjects2 := []**AccessPolicyItem{
		&accpolicy,
	}
	for i, objType := range []string{"access_policy"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects2[i] = &AccessPolicyItem{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}
	}

	lcap := []string{}
	for _, lic := range d.Get("license_caps").([]interface{}) {
		lcap = append(lcap, lic.(string))
	}
	
	res, err := c.CreateFmcDevice(ctx, &Device{
		Name:        d.Get("name").(string),
		HostName:     d.Get("hostname").(string),
		NatID:        d.Get("nat_id").(string),
		RegKey:       d.Get("regkey").(string),
		Type:        host_type,
		LicenseCaps: lcap,
		AccessPolicy:  accpolicy,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to add device",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcDeviceRead(ctx, d, m)
}

func resourceFmcDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcDevice(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("hostname", item.HostName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("nat_id", item.NatID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("regkey", item.RegKey); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}
	// accpolicy := make([]interface{}, 0, 1)
	// accpolicyObj := make(map[string]interface{})
	// accpolicyObj["id"] = item.AccessPolicy.ID
	// accpolicyObj["type"] = item.AccessPolicy.Type
	// accpolicy = append(accpolicy, accpolicyObj)

	// if err := d.Set("access_policy", accpolicy); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "unable to read policy devices assignment",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }
	// accesspolicy := make([]interface{}, 0, len(item.AccessPolicy))
	// for _, acc := range item.AccessPolicy{
	// 	acci := make(map[string]interface{})
	// 	acci["id"] = acc.ID
	// 	acci["type"] = acc.Type
	// 	accesspolicy = append(accesspolicy, acci)
	// }

	// if err := d.Set("access_policy", accesspolicy); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "unable to read network group object",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }
	if item.AccessPolicy != (AccessPolicyItem{}) {
		if err := d.Set("access_policy", convertTo1ListMapStringGeneric(item.AccessPolicy)); err != nil {
			return returnWithDiag(diags, err)
		}
	}

	return diags
}

func resourceFmcDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name", "hostname", "type", "license_caps", "access_policy") {
	var accpolicy *AccessPolicyItem
	dynamicObjects2 := []**AccessPolicyItem{
		&accpolicy,
	}
	for i, objType := range []string{"access_policy"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects2[i] = &AccessPolicyItem{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}
	}
	lcap := []string{}
	for _, lic := range d.Get("license_caps").([]interface{}) {
		lcap = append(lcap, lic.(string))
	}
	res, err := c.UpdateFmcDevice(ctx, d.Id(), &Device{
		Name:        d.Get("name").(string),
		HostName:     d.Get("hostname").(string),
		NatID:        d.Get("nat_id").(string),
		RegKey:       d.Get("regkey").(string),
		Type:        host_type,
		LicenseCaps: lcap,
		AccessPolicy:  accpolicy,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to add device",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	}
	return resourceFmcDeviceRead(ctx, d, m)
}

func resourceFmcDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcDevice(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete device",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
