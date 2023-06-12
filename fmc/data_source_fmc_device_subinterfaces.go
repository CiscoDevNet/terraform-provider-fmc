package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcSubInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Sub-Interfaces in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_subinterfaces\" \"sub_interfaces\" {\n" +
			"	device_id = \"<ID of the ftd>\"\n" +	
			"	subinterface_id = 1\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceSubInterfacesRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the device this resource needs",
			},
			"subinterface_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of sub interface",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SubInterface type",
			},
			"if_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SubInterface logical name",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The Vlan ID needed for this resource",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SubInterface MTU",
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mode of this resource",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security Zone ID",
			},
			"ipv4_static_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 Static address",
			},
		},
	}
}

func dataSourceSubInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ifc, err := c.GetFmcSubInterfaceByName(ctx, d.Get("device_id").(string), d.Get("subinterface_id").(int))

	d.Set("type", ifc.Type)
	d.Set("if_name", ifc.Ifname)
	d.Set("mtu", ifc.MTU)
	d.Set("mode", ifc.Mode)
	d.Set("security_zone_id", ifc.SecurityZone.ID)
	d.Set("vlan_id", ifc.VlanID)
	d.Set("ipv4_static_address", ifc.IPv4.Static.Address)
	
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get subinterface",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(ifc.ID)

	if err := d.Set("subinterface_id", ifc.SubInterfaceID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read subinterface",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
