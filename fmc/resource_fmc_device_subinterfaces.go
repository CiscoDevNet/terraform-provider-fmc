package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcSubInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for SubInterfaces in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_device_subinterfaces\" \"sub\" {\n" +
			"    name = \"GigabitEthernet0/1\"\n" +
			"	 device_id = \"<ID of the ftd>\"\n" +
			"	 subinterface_id = 1\n" +
			"	 vlan_id = 30\n" +
			"	 security_zone_id = \"<ID of the security zone>\"\n" +
			"	 if_name = \"Inside\"\n" +
			"	 mtu = 1700\n" +
			"	 mode = \"NONE\"\n" +
			"	 enabled =   true\n" +
			"	 ipv4_static_address = \"10.20.220.45\"\n" +
			"	 ipv4_static_netmask = 24\n" +
			"	 ipv4_dhcp_enabled = false\n" +
			"	 ipv4_dhcp_route_metric = 1\n" +
			"	 ipv6_address = \"2001:10:240:ac::a\"\n" +
			"	 ipv6_prefix = 124\n" +
			"	 ipv6_enforce_eui =   false\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcSubInterfaceCreate,
		ReadContext:   resourceFmcSubInterfaceRead,
		UpdateContext: resourceFmcSubInterfaceUpdate,
		DeleteContext: resourceFmcSubInterfaceDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the device this resource needs",
			},
			"subinterface_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID this resource needs",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Vlan ID needed for this resource",
			},
			"ifname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of chosen interface",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable this resource",
			},
			"management_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Managenment only or not",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"mode": { 
				Type:        schema.TypeString,
				Optional:    true,
				Default: 	 "NONE",
				Description: "The mode of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the physical interface",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1500,
				Description: "The mtu value",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default: 	 0,
				Description: "The type of this resource",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Security Zone ID",
			},
			"ipv4_static_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 Static address",
			},
			"ipv4_static_netmask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv4 Static address netmask",
			},
			"ipv4_dhcp_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "IPv4 DHCP enabled",
			},
			"ipv4_dhcp_route_metric": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv4 DHCP Route Metric",
			},
			"enable_ipv6": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "IPv6 address enabled",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address",
			},
			"ipv6_prefix": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv6 netmask",
			},
			"ipv6_enforce_eui": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "IPv6 EnforceEUI64",
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
	if err := d.Set("ifname", item.Ifname); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical interface",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
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


	ipv4StaticAddress := d.Get("ipv4_static_address").(string)
	ipv4StaticNetmask := d.Get("ipv4_static_netmask").(int)
	ipv4DhcpEnabled := d.Get("ipv4_dhcp_enabled").(bool)
	ipv4DhcpRouteMetric := d.Get("ipv4_dhcp_route_metric").(int)
	
	securityZoneId := d.Get("security_zone_id").(string)
	
	ipv6Address := d.Get("ipv6_address").(string)
	ipv6Prefix := d.Get("ipv6_prefix").(int)
	ipv6EnforceEUI := d.Get("ipv6_enforce_eui").(bool)
	enable_ipv6 := d.Get("enable_ipv6").(bool)
	var IPv6Add []IPv6Address

	if len(ipv6Address) > 0 {
		IPv6Add = append(IPv6Add, IPv6Address{
			Address:      ipv6Address,
			Prefix:       ipv6Prefix,
			EnforceEUI64: ipv6EnforceEUI,
		})
	}

	var IPv6 = IPv6{EnableIPv6:enable_ipv6 ,Addresses: IPv6Add}

	var SubInterfaceSecurityZone = PhysicalInterfaceSecurityZone{
		ID:   securityZoneId,
		Type: "SecurityZone",
	}
	var IPv4Static = IPv4Static{
		Address: ipv4StaticAddress,
		Netmask: ipv4StaticNetmask,
	}
	var IPv4DHCP = IPv4DHCP{
		Enable:      ipv4DhcpEnabled,
		RouteMetric: ipv4DhcpRouteMetric,
	}

	var IPv4 = IPv4{}

	if ipv4DhcpEnabled {
		IPv4.DHCP = &IPv4DHCP
	} else if len(ipv4StaticAddress) > 0 {
		IPv4.Static = &IPv4Static
	}
	res, err := c.CreateFmcSubInterface(ctx, d.Get("device_id").(string), &SubInterface{
		Ifname:         d.Get("ifname").(string),
		Mode:           d.Get("mode").(string),
		Name:           d.Get("name").(string),
		IPv4:           IPv4,
		VlanID: 	    d.Get("vlan_id").(int),
		SubInterfaceID: d.Get("subinterface_id").(int),
		Enabled:        d.Get("enabled").(bool),
		SecurityZone:   SubInterfaceSecurityZone,
		MgmntOnly:      d.Get("management_only").(bool),
		Priority:       d.Get("priority").(int),
		MTU:            d.Get("mtu").(int),
		IPv6: IPv6,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(res.ID)
	return resourceFmcSubInterfaceRead(ctx, d, m)
}
func resourceFmcSubInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("ipv6_enforce_eui","ipv6_prefix","ipv6_address","vlan_id","management_only","ifname","name", "mode", "ipv4_static_address", "security_zone_id", "ipv4_dhcp_enabled","ipv4_dhcp_route_metric", "priority", "enabled") {

		ipv4StaticAddress := d.Get("ipv4_static_address").(string)
		ipv4StaticNetmask := d.Get("ipv4_static_netmask").(int)
		ipv4DhcpEnabled := d.Get("ipv4_dhcp_enabled").(bool)
		ipv4DhcpRouteMetric := d.Get("ipv4_dhcp_route_metric").(int)
		securityZoneId := d.Get("security_zone_id").(string)
		enable_ipv6 := d.Get("enable_ipv6").(bool)

		var SubInterfaceSecurityZone = PhysicalInterfaceSecurityZone{
			ID:   securityZoneId,
			Type: "SecurityZone",
		}
		var IPv4Static = IPv4Static{
			Address: ipv4StaticAddress,
			Netmask: ipv4StaticNetmask,
		}
		var IPv4DHCP = IPv4DHCP{
			Enable:      ipv4DhcpEnabled,
			RouteMetric: ipv4DhcpRouteMetric,
		}

		var IPv4 = IPv4{}

		if ipv4DhcpEnabled {
			IPv4.DHCP = &IPv4DHCP
		} else if len(ipv4StaticAddress) > 0 {
			IPv4.Static = &IPv4Static
		}
		ipv6Address := d.Get("ipv6_address").(string)
		ipv6Prefix := d.Get("ipv6_prefix").(int)
		ipv6EnforceEUI := d.Get("ipv6_enforce_eui").(bool)

		var IPv6Add []IPv6Address

		if len(ipv6Address) > 0 {
			IPv6Add = append(IPv6Add, IPv6Address{
				Address:      ipv6Address,
				Prefix:       ipv6Prefix,
				EnforceEUI64: ipv6EnforceEUI,
			})
		}

		var IPv6 = IPv6{EnableIPv6:enable_ipv6 ,Addresses: IPv6Add}

		_, err := c.UpdateFmcSubInterface(ctx, d.Get("device_id").(string), id, &SubInterface{
			ID:             id,
			Ifname:         d.Get("ifname").(string),
			Mode:           d.Get("mode").(string),
			Name:           d.Get("name").(string),
			IPv4:           IPv4,
			VlanID: 	    d.Get("vlan_id").(int),
			Enabled:        d.Get("enabled").(bool),
			SecurityZone:   SubInterfaceSecurityZone,
			SubInterfaceID: d.Get("subinterface_id").(int),
			MgmntOnly:      d.Get("management_only").(bool),
			Priority:       d.Get("priority").(int),
			MTU:            d.Get("mtu").(int),
			IPv6:          IPv6,
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
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcSubInterface(ctx, d.Get("device_id").(string) ,id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete subinterface",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
