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
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the physical interface this resource needs",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ID of the device this resource needs",
			},
			"ifname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of chosen interface",
			},
			"security_zone": {
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
				Description: "Security Zone for this resource",
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
				Optional: true,
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
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Dynamic IP of the interface",
									},
									"dhcp_route_metric": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: "DHCP of the interface",
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
	var patOptions *SubInterfaceIpv4
	// var szs SubInterfaceSecurityZone
	var securityZone *SubInterfaceSecurityZone
	dynamicObjects := []**SubInterfaceSecurityZone{
		&securityZone,
	}

	if pat_o := d.Get("ipv4").([]interface{}); len(pat_o) > 0 {
		pat_options := pat_o[0].(map[string]interface{})
		pat_pool_options := pat_options["dhcp"].([]interface{})[0].(map[string]interface{})
		pat_pool := SubInterfaceIpv4Dhcp{
			EnableDefaultRouteDHCP:   pat_pool_options["enable_default_route_dhcp"].(string),
			DhcpRouteMetric: pat_pool_options["dhcp_route_metric"].(int),
		}
		patOptions = &SubInterfaceIpv4{
			Dhcp: pat_pool,
		}
	}

	for i, objType := range []string{"security_zone"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects[i] = &SubInterfaceSecurityZone{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}
	}

	_, err := c.CreateFmcSubInterface(ctx, d.Get("device_id").(string), &SubInterface{
		Ifname:        d.Get("ifname").(string),
		Mode:          d.Get("mode").(string),
		Ipv4:          patOptions,
		VlanID: 	   d.Get("vlan_id").(int),
		SubInterfaceID: d.Get("subinterface_id").(int),
		Enabled:       d.Get("enabled").(bool),
		Security_Zone: securityZone,
		MTU: d.Get("mtu").(int),
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
		var securityZone *SubInterfaceSecurityZone
		dynamicObjects := []**SubInterfaceSecurityZone{
			&securityZone,
		}


		for i, objType := range []string{"security_zone"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entry := inputEntries.([]interface{})[0].(map[string]interface{})
				*dynamicObjects[i] = &SubInterfaceSecurityZone{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
				}
			}
		}

		_, err := c.UpdateFmcSubInterface(ctx, d.Get("device_id").(string), d.Get("subinterface_id").(string), &SubInterface{
			Ifname:        d.Get("ifname").(string),
			Mode:          d.Get("mode").(string),
			Ipv4:          ipv4,
			Security_Zone: securityZone,
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
