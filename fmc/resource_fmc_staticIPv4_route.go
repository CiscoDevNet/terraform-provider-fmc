package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var static_route_type string = "IPv4StaticRoute"

func resourceFmcStaticIPv4Route() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Static IPv4 rotues in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_staticIPv4_route\" \"route1\" {\n" +
			"    device_id = data.fmc_devices.device.id\n" +
			"    interface_name = \"outside\"\n" +
			"    is_tunneled = false\n" +
			"    metric_value = 22\n" +
			"    route_tracking {\n" +
			"        id = <ID of SLA>\n" +
			"        type = \"SLAMonitor\"\n" +
			"        name = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    selected_networks {\n" +
			"        id = data.fmc_network_objects.nw.id\n" +
			"        type = data.fmc_network_objects.nw.type\n" +
			"        name = data.fmc_network_objects.nw.name\n" +
			"    gateway {\n" +
			"        object {\n" +
			"            id = data.fmc_host_objects.igw.id\n" +
			"            type = data.fmc_host_objects.igw.type\n" +
			"        	 name = data.fmc_host_objects.igw.name\n" +
			"        }\n" +
			"	}\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcStaticIPv4RouteCreate,
		ReadContext:   resourceFmcStaticIPv4RouteRead,
		UpdateContext: resourceFmcStaticIPv4RouteUpdate,
		DeleteContext: resourceFmcStaticIPv4RouteDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of this resource",
			},
			"interface_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the interface",
			},
			"metric_value": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The metric value to send",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"is_tunneled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If the route is tunneled or not",
			},
			"route_tracking": {
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
				Description: "Route tracking information",
			},
			"selected_networks": {
				Type:     schema.TypeList,
				Required: true,
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
							Optional:    true,
							Description: "The type of this resource",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of this resource",
						},
					},
				},
				Description: "Route tracking information",
			},
			"gateway": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ID of this resource",
									},
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of this resource",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of this resource",
									},
								},
							},
						},
					},
				},
				Description: "The gateway for this resource",
			},
		},
	}
}

func resourceFmcStaticIPv4RouteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var gateway *ObjectItems
	if gate_o := d.Get("gateway").([]interface{}); len(gate_o) > 0 {
		gateway_options := gate_o[0].(map[string]interface{})
		gateway_list := gateway_options["object"].([]interface{})
		var gw *ObjectItem
		if len(gateway_list) != 0 {
			gw_object := gateway_list[0].(map[string]interface{})
			gw = &ObjectItem{
				ID:   gw_object["id"].(string),
				Type: gw_object["type"].(string),
				Name: gw_object["name"].(string),
			}
		}
		gateway = &ObjectItems{
			Objects: gw,
		}
	}
	var selectedNetworks []ObjectItem
	dynamicObjects2 := []*[]ObjectItem{
		&selectedNetworks,
	}
	for i, objType := range []string{"selected_networks"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects2[i] = []ObjectItem{
				{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
					Name: entry["name"].(string),
				},
			}
		}
	}
	var routeTracking *ObjectItem
	dynamicObjects3 := []**ObjectItem{
		&routeTracking,
	}
	for i, objType := range []string{"route_tracking"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects3[i] = &ObjectItem{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
				Name: entry["name"].(string),
			}
		}
	}
	res, err := c.CreateFmcStaticIPv4Route(ctx, d.Get("device_id").(string), &StaticIpv4Route{
		InterfaceName:    d.Get("interface_name").(string),
		Type:             static_route_type,
		Tunneled:         d.Get("is_tunneled").(bool),
		MetricValue:      d.Get("metric_value").(int),
		SelectedNetworks: selectedNetworks,
		Gateway:          gateway,
		RouteTracking:    routeTracking,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to add route",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcStaticIPv4RouteRead(ctx, d, m)
}

func resourceFmcStaticIPv4RouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcStaticIPv4Route(ctx, d.Get("device_id").(string), d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read route",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("interface_name", item.InterfaceName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read route",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read route",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("metric_value", item.MetricValue); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read route",
			Detail:   err.Error(),
		})
		return diags
	}

	if item.Gateway != (ObjectItems{}) {
		gateway_options := make(map[string]interface{})
		if item.Gateway.Objects != nil && *item.Gateway.Objects != (ObjectItem{}) {
			gateway_options["object"] = convertTo1ListMapStringGeneric(item.Gateway.Objects)
		}
		if err := d.Set("gateway", convertTo1ListGeneric(gateway_options)); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read route",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	if item.RouteTracking != (ObjectItem{}) {
		if err := d.Set("route_tracking", convertTo1ListMapStringGeneric(item.RouteTracking)); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read route",
				Detail:   err.Error(),
			})
			return diags
		}
	}

	return diags
}

func resourceFmcStaticIPv4RouteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("interface_name", "is_tunneled", "metric_value", "gateway", "selected_networks") {

		var gateway *ObjectItems
		if gate_o := d.Get("gateway").([]interface{}); len(gate_o) > 0 {
			gateway_options := gate_o[0].(map[string]interface{})
			gateway_list := gateway_options["object"].([]interface{})
			var gw *ObjectItem
			if len(gateway_list) != 0 {
				gw_object := gateway_list[0].(map[string]interface{})
				gw = &ObjectItem{
					ID:   gw_object["id"].(string),
					Type: gw_object["type"].(string),
					Name: gw_object["name"].(string),
				}
			}
			gateway = &ObjectItems{
				Objects: gw,
			}
		}
		var selectedNetworks []ObjectItem
		dynamicObjects2 := []*[]ObjectItem{
			&selectedNetworks,
		}
		for i, objType := range []string{"selected_networks"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entry := inputEntries.([]interface{})[0].(map[string]interface{})
				*dynamicObjects2[i] = []ObjectItem{
					{
						ID:   entry["id"].(string),
						Type: entry["type"].(string),
						Name: entry["name"].(string),
					},
				}
			}
		}
		var routeTracking *ObjectItem
		dynamicObjects3 := []**ObjectItem{
			&routeTracking,
		}
		for i, objType := range []string{"route_tracking"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entry := inputEntries.([]interface{})[0].(map[string]interface{})
				*dynamicObjects3[i] = &ObjectItem{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
					Name: entry["name"].(string),
				}
			}
		}
		res, err := c.UpdateFmcStaticIPv4Response(ctx, d.Get("device_id").(string), d.Id(), &StaticIpv4Route{
			ID:               d.Id(),
			InterfaceName:    d.Get("interface_name").(string),
			Type:             static_route_type,
			Tunneled:         d.Get("is_tunneled").(bool),
			MetricValue:      d.Get("metric_value").(int),
			SelectedNetworks: selectedNetworks,
			Gateway:          gateway,
			RouteTracking:    routeTracking,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update route",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcStaticIPv4RouteRead(ctx, d, m)
}

func resourceFmcStaticIPv4RouteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcStaticIPv4Response(ctx, d.Get("device_id").(string), d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete route",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
