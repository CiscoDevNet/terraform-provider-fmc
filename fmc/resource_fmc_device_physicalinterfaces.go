package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var physical_interface_type string = "PhysicalInterface"

func resourceFmcDevicePhysicalInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Physical Interfaces in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_physical_interfaces\" \"physical_interfaces_1\" {\n" +
			"    id = fpphysicalinterfaceUUID1\n" +
			"    name = \"s1p1\"\n" +
			"    type = \"FPPhysicalInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"resource \"fmc_physical_interfaces\" \"physical_interfaces_2\" {\n" +
			"    id = fpphysicalinterfaceUUID2\n" +
			"    name = \"s1p2\"\n" +
			"    type = \"FPPhysicalInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"\n" +
			"}\n" +
			"```\n",
		UpdateContext: resourceFmcDevicePhysicalInterfacesUpdate,
		ReadContext:   resourceFmcDevicePhysicalInterfacesRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this physical interfaces",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of physical interfaces",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of this physical interfaces",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this physical interfaces",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enabled of this physical interfaces",
			},
			"MTU": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "MTU of this physical interfaces",
			},
			"mode": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "mode of this physical interfaces",
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ifname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ifname of this physical interfaces",
			},
			"nveOnly": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "ifname of this physical interfaces",
			},
			"managementOnly": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "managementOnly of this physical interfaces",
			},
			"securityZone": {
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
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of this resource",
						},
					},
				},
				Description: "Destination zones for this resource",
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
