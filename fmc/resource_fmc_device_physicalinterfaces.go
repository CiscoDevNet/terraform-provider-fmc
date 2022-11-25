package fmc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var physical_interface_type string = "PhysicalInterface"

func resourceFmcDevicePhysicalInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for access policy category in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_access_policies_category\" \"category\" {\n" +
			"    name        		  = \"test-time-range\"\n" +
			"    access_policy_id     = \"BB62F664-7168-4C8E-B4CE-F70D522889D2\"\n" +
			"}\n" +
			"```",
		UpdateContext: resourceFmcDevicePhysicalInterfacesUpdate,
		ReadContext:   resourceFmcDevicePhysicalInterfacesRead,
		//ReadAllContext: resourceFmcDevicePhysicalInterfacesReadAll,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of this physical interfaces",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of physical interfaces",
			},
		},
	}
}

func resourceFmcDevicePhysicalInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	acpId := d.Get("acpId").(string)
	item, err := c.GetPhysicalInterfacesByName(ctx, id, acpId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("acpId", acpId); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcDevicePhysicalInterfacesReadAll(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	items, err := c.GetPhysicalInterfaces(ctx, id)
	fmt.Println(items)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physicalinterfaces",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcDevicePhysicalInterfacesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("hardware", "LLDP", "name", "managementOnly", "MTU", "mode") {
		var hardware HardWare
		var lldp LLDP

		if inputObjs, ok := d.GetOk("hardware"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			hardware = HardWare{
				Duplex:          obj["duplex"].(string),
				Speed:           obj["speed"].(string),
				AutoNegState:    obj["autoNegState"].(string),
				FecMode:         obj["fecMode"].(string),
				FlowControlSend: obj["flowControlSend"].(string),
			}
		}

		if inputObjs, ok := d.GetOk("LLDP"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			lldp = LLDP{
				Transmit: obj["transmit"].(bool),
				Receive:  obj["receive"].(bool),
			}
		}

		res, err := c.UpdatePhysicalInterfaces(ctx, d.Get("acp").(string), d.Id(), &Physicalinterfaces{
			ID:                 d.Id(),
			Name:               d.Get("name").(string),
			Type:               physical_interface_type,
			ManagementOnly:     d.Get("managementOnly").(bool),
			MTU:                d.Get("MTU").(int),
			Mode:               d.Get("mode").(string),
			Enabled:            d.Get("enabled").(bool),
			NveOnly:            d.Get("nveOnly").(bool),
			EnableSGTPropagate: d.Get("enableSGTPropagate").(bool),
			HardWare:           hardware,
			LLDP:               lldp,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceFmcDevicePhysicalInterfacesRead(ctx, d, m)
}
