package fmc

import (
	"context"
	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
var device_type string = "Device"
func resourceFmcDevices() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for adding device in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_devicese\" \"device1\" {\n" +
			"    name = \"ftd\"\n" +
			"    hostname = \"<IP ADDR OF HOST>\"\n" +
			"    regkey = \"<Reg key used in FTD>\"\n" +
			"    metric_value = 22\n" +
			"    type = \"Device\"\n" +
			"    license_caps = [\n" +
			"		\"MALWARE\"\n" +
			"    ]\n" +
			"    access_policy {\n" +
			"        id = data.fmc_access_policies.access_policy.id\n" +
			"        type = data.fmc_access_policies.access_policy.type\n" +
			"	}\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
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
	
	_, err := c.CreateFmcDevice(ctx, &Device{
		Name:        d.Get("name").(string),
		HostName:     d.Get("hostname").(string),
		NatID:        d.Get("nat_id").(string),
		RegKey:       d.Get("regkey").(string),
		Type:        device_type,
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
	//ADD CODE TO GET ID
	device, err := c.GetFmcDeviceByName(ctx, d.Get("name").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get device",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(device.ID)

	//Code ENDED
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
	if d.HasChanges("name", "hostname", "type", "license_caps") {
	// var accpolicy *AccessPolicyItem
	// dynamicObjects2 := []**AccessPolicyItem{
	// 	&accpolicy,
	// }
	// for i, objType := range []string{"access_policy"} {
	// 	if inputEntries, ok := d.GetOk(objType); ok {
	// 		entry := inputEntries.([]interface{})[0].(map[string]interface{})
	// 		*dynamicObjects2[i] = &AccessPolicyItem{
	// 			ID:   entry["id"].(string),
	// 			Type: entry["type"].(string),
	// 		}
	// 	}
	// }
	lcap := []string{}
	for _, lic := range d.Get("license_caps").([]interface{}) {
		lcap = append(lcap, lic.(string))
	}
	res, err := c.UpdateFmcDevice(ctx, d.Id(), &Device{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		HostName:     d.Get("hostname").(string),
		NatID:        d.Get("nat_id").(string),
		RegKey:       d.Get("regkey").(string),
		Type:        device_type,
		LicenseCaps: lcap,
		// AccessPolicy:  accpolicy,
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
