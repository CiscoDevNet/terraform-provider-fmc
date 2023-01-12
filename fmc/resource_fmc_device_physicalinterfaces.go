package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcPhysicalInterface() *schema.Resource {
	return &schema.Resource{
		// CreateContext: resourceFmcPhysicalInterfaceCreate,
		ReadContext:   resourceFmcPhysicalInterfaceRead,
		UpdateContext: resourceFmcPhysicalInterfaceUpdate,
		DeleteContext: resourceFmcPhysicalInterfaceDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the device this resource needs",
			},
			"physicalinterface_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the physical interface this resource needs",
			},
			"ifname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of chosen interface",
			},
			"security_zone": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of SZ",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "SecurityZone",
							Description: "The type of SZ",
						},
					},
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"mode": { // add method to verify
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "NONE",
				Description: "The mode of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the physcal interface",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1500,
				Description: "The type of this resource",
			},
			"ipv4": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP of the interface",
									},
									"netmask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Subnet mask of the interface",
									},
								},
							},
						},
						"dhcp": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_default_route_dhcp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Dynamic IP of the interface",
									},
									"dhcp_route_metric": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
								},
							},
						},
					},
				},
				Description: "IPV4 information",
			},
		},
	}
}

func resourceFmcPhysicalInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcPhysicalInterface(ctx, d.Get("device_id").(string), id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical interface",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("ifname", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read host object",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceFmcPhysicalInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("ifname", "mode", "ipv4", "security_zone") {

		var ipv4 PhysicalInterfaceIpv4
		var szs PhysicalInterfaceSecurityZone

		if inputIpv4 := d.Get("ipv4").([]interface{}); len(inputIpv4) > 0 {
			ipv4s := inputIpv4[0].(map[string]interface{})
			static_ipv4 := ipv4s["static"].([]interface{})[0].(map[string]interface{})
			static := PhysicalInterfaceIpv4Static{
				Address: static_ipv4["address"].(string),
				Netmask: static_ipv4["netmask"].(string),
			}

			dhcp_ipv4 := ipv4s["dhcp"].([]interface{})[0].(map[string]interface{})
			dhcp := PhysicalInterfaceIpv4Dhcp{
				EnableDefaultRouteDHCP: dhcp_ipv4["enableDefaultRouteDHCP"].(string),
				DhcpRouteMetric:        dhcp_ipv4["dhcpRouteMetric"].(string),
			}

			ipv4 = PhysicalInterfaceIpv4{
				Static: static,
				Dhcp:   dhcp,
			}
		}

		if inputSzs, ok := d.GetOk("security_zone"); ok {
			for _, sz := range inputSzs.([]interface{}) {
				szi := sz.(map[string]interface{})
				szs = PhysicalInterfaceSecurityZone{
					ID:   szi["id"].(string),
					Type: szi["type"].(string),
				}
			}
		}

		_, err := c.UpdateFmcPhysicalInterface(ctx, d.Get("device_id").(string), d.Get("physicalnterface_id").(string), &PhysicalInterface{
			Ifname:        d.Get("ifname").(string),
			Mode:          d.Get("mode").(string),
			Ipv4:          ipv4,
			Security_Zone: szs,
			ID:            id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update physical interface",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcPhysicalInterfaceRead(ctx, d, m)
}

func resourceFmcPhysicalInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}