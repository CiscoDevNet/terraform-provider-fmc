package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcPhysicalInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Physical Interfaces in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_physical_interfaces\" \"test-phy-interfaces\" {\n" +
			"	name = \"TEST-PHY\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourcePhysicalInterfaceRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Physical Interface",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Physical Interface",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of Physical Interface",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Physical Interface description",
			},
			"if_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Physical Interface logical name(ifname)",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Physical Interface MTU",
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Physical Interface Mode",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Physical Interface Security Zone ID",
			},
		},
	}
}

func dataSourcePhysicalInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	ifc, err := c.GetFmcPhysicalInterface(ctx, d.Get("device_id").(string), d.Get("name").(string))

	log.Printf("ID=%s Name=%s IFName=%s Type=%s", ifc.ID, ifc.Name, ifc.Ifname, ifc.Type)
	log.Printf("SecurityZone=%s", ifc.SecurityZone.ID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get physical interface",
			Detail:   err.Error(),
		})
		return diags
	}
	log.Printf("Description=%s", ifc.Name)

	d.SetId(ifc.ID)

	d.Set("type", ifc.Type)
	d.Set("description", ifc.Description)
	d.Set("if_name", ifc.Ifname)
	d.Set("mtu", ifc.MTU)
	d.Set("mode", ifc.Mode)
	d.Set("security_zone_id", ifc.SecurityZone.ID)

	if err := d.Set("name", ifc.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical Interface",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
