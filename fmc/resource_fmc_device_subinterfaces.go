package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcSubInterface() *schema.Resource {
	return &schema.Resource{
		// CreateContext: resourceFmcPhysicalInterfaceCreate,
		ReadContext:   resourceFmcSubInterfaceRead,
		UpdateContext: resourceFmcSubInterfaceUpdate,
		DeleteContext: resourceFmcSubInterfaceDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the device this resource needs",
			},
			"subinterface_id": {
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

func resourceFmcSubInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcSubInterface(ctx, d.Get("device_id").(string), id)
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
func resourceFmcSubInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics
	id := d.Id()
	var ipv4 *SubInterfaceIpv4
	var szs SubInterfaceSecurityZone

	if ip_v4 := d.Get("ipv4").([]interface{}); len(ip_v4) > 0 {
		ip_v42 := ip_v4[0].(map[string]interface{})
		ipv4_static_list := ip_v42["static"].([]interface{})
		var ipv4_static *SubInterfaceIpv4Static
		if len(ipv4_static_list) != 0 {
			ipv4_static_address := ipv4_static_list[0].(map[string]interface{})
			ipv4_static = &SubInterfaceIpv4Static{
				Address:  ipv4_static_address["address"].(string),
				Netmask: ipv4_static_address["netmask"].(string),
			}
		}
		ipv4_dhcp_list := ip_v42["dhcp"].([]interface{})
		var ipv4_dhcp *SubInterfaceIpv4Dhcp
		if len(ipv4_dhcp_list) != 0 {
			ipv4_dhcp_address := ipv4_dhcp_list[0].(map[string]interface{})
			ipv4_dhcp = &SubInterfaceIpv4Dhcp{
				EnableDefaultRouteDHCP:  ipv4_dhcp_address["enable_default_route_dhcp"].(string),
				DhcpRouteMetric: ipv4_dhcp_address["dhcp_route_metric"].(string),
			}
		}
		ipv4 = &SubInterfaceIpv4{
			Static: ipv4_static,
			Dhcp: ipv4_dhcp,
		}
	}

	if inputSzs, ok := d.GetOk("security_zone"); ok {
		for _, sz := range inputSzs.([]interface{}) {
			szi := sz.(map[string]interface{})
			szs = SubInterfaceSecurityZone{
				ID:   szi["id"].(string),
				Type: szi["type"].(string),
			}
		}
	}

	_, err := c.CreateFmcSubInterface(ctx, d.Get("device_id").(string), &SubInterface{
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
	return resourceFmcSubInterfaceRead(ctx, d, m)
}
func resourceFmcSubInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("ifname", "mode", "ipv4", "security_zone") {

		var ipv4 *SubInterfaceIpv4
		var szs SubInterfaceSecurityZone

		if ip_v4 := d.Get("ipv4").([]interface{}); len(ip_v4) > 0 {
			ip_v42 := ip_v4[0].(map[string]interface{})
			ipv4_static_list := ip_v42["static"].([]interface{})
			var ipv4_static *SubInterfaceIpv4Static
			if len(ipv4_static_list) != 0 {
				ipv4_static_address := ipv4_static_list[0].(map[string]interface{})
				ipv4_static = &SubInterfaceIpv4Static{
					Address:  ipv4_static_address["address"].(string),
					Netmask: ipv4_static_address["netmask"].(string),
				}
			}
			ipv4_dhcp_list := ip_v42["dhcp"].([]interface{})
			var ipv4_dhcp *SubInterfaceIpv4Dhcp
			if len(ipv4_dhcp_list) != 0 {
				ipv4_dhcp_address := ipv4_dhcp_list[0].(map[string]interface{})
				ipv4_dhcp = &SubInterfaceIpv4Dhcp{
					EnableDefaultRouteDHCP:  ipv4_dhcp_address["enable_default_route_dhcp"].(string),
					DhcpRouteMetric: ipv4_dhcp_address["dhcp_route_metric"].(string),
				}
			}
			ipv4 = &SubInterfaceIpv4{
				Static: ipv4_static,
				Dhcp: ipv4_dhcp,
			}
		}

		if inputSzs, ok := d.GetOk("security_zone"); ok {
			for _, sz := range inputSzs.([]interface{}) {
				szi := sz.(map[string]interface{})
				szs = SubInterfaceSecurityZone{
					ID:   szi["id"].(string),
					Type: szi["type"].(string),
				}
			}
		}

		_, err := c.UpdateFmcSubInterface(ctx, d.Get("device_id").(string), d.Get("subinterface_id").(string), &SubInterface{
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
	return resourceFmcSubInterfaceRead(ctx, d, m)
}

func resourceFmcSubInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
