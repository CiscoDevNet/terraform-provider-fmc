package fmc

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourcePhyInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhyInterfaceCreate,
		ReadContext:   resourcePhyInterfaceRead,
		UpdateContext: resourcePhyInterfaceUpdate,
		DeleteContext: resourcePhyInterfaceDelete,
		Schema: map[string]*schema.Schema{
			"physical_interface_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the physical interface this resource needs",
			},
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Physical Interface",
			},

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of chosen interface",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Physical Interface description",
			},
			"if_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of chosen interface",
			},

			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Physical Interface MTU",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Physical Interface Mode",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Physical Interface Security Zone",
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

func resourcePhyInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("FPC: Creating physical interface, this is mock method. In actual physical interface is not going to get created, instead actual physical interface will be fetched and store terraform state")

	return resourcePhyInterfaceUpdate(ctx, d, m)
}

func resourcePhyInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Printf("FPR: Fetching physical interface details")

	deviceId := d.Get("device_id").(string)
	physicalInterfaceId := d.Get("physical_interface_id").(string)

	log.Printf("FPR: DeviceId=%s, PhysicalInterfaceId=%s", deviceId, physicalInterfaceId)

	c := m.(*Client)

	var diags diag.Diagnostics

	physicalInterfaces, err := c.GetFmcPhysicalInterfaceByID(ctx, deviceId, physicalInterfaceId)

	log.Printf("FPR: Physical Interface Details PhysicalInterfaceId=%s Name=%s IFName=%s Type=%s", physicalInterfaces.ID, physicalInterfaces.Name, physicalInterfaces.Ifname, physicalInterfaces.Type)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(physicalInterfaces.ID)

	return diags
}

func resourcePhyInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("FPU: Updating physical interface details")

	deviceId := d.Get("device_id").(string)
	physicalInterfaceId := d.Get("physical_interface_id").(string)
	name := d.Get("name").(string)
	iFName := d.Get("if_name").(string)
	description := d.Get("description").(string)
	mtu := d.Get("mtu").(int)
	mode := d.Get("mode").(string)
	securityZoneId := d.Get("security_zone_id").(string)

	log.Printf("FPU: DeviceId=%s, PhysicalInterfaceId=%s, IFName=%s Name=%s, Description=%s, security_zone_id=%s", deviceId, physicalInterfaceId, iFName, name, description, securityZoneId)

	ipv4StaticAddress := d.Get("ipv4_static_address").(string)
	ipv4StaticNetmask := d.Get("ipv4_static_netmask").(int)
	ipv4DhcpEnabled := d.Get("ipv4_dhcp_enabled").(bool)
	ipv4DhcpRouteMetric := d.Get("ipv4_dhcp_route_metric").(int)

	log.Printf("ipv4_static_address=%s, ipv4_static_netmask=%s, ipv4_dhcp_enabled=%s, ipv4_dhcp_route_metric=%s", ipv4StaticAddress, ipv4StaticNetmask, ipv4DhcpEnabled, ipv4DhcpRouteMetric)

	ipv6Address := d.Get("ipv6_address").(string)
	ipv6Prefix := d.Get("ipv6_prefix").(int)
	ipv6EnforceEUI := d.Get("ipv6_enforce_eui").(bool)

	log.Printf("ipv6_address=%s, ipv6_prefix=%s, ipv6_enforceEUI64=%s", ipv6Address, ipv6Prefix, ipv6EnforceEUI)

	c := m.(*Client)

	var diags diag.Diagnostics

	var PhysicalInterfaceSecurityZone = PhysicalInterfaceSecurityZone{
		ID:   securityZoneId,
		Type: "SecurityZone",
	}
	log.Printf("PhysicalInterfaceSecurityZone=%s", PhysicalInterfaceSecurityZone)

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

	log.Printf("IPv4=%s", IPv4)

	var IPv6Add []IPv6Address

	if len(ipv6Address) > 0 {
		IPv6Add = append(IPv6Add, IPv6Address{
			Address:      ipv6Address,
			Prefix:       ipv6Prefix,
			EnforceEUI64: ipv6EnforceEUI,
		})
	}

	var IPv6 = IPv6{Addresses: IPv6Add}
	log.Printf("IPv6Address=%s", IPv6Add)

	physicalInterfaceResponse, err := c.UpdateFmcPhysicalInterface(ctx, deviceId, physicalInterfaceId, &PhysicalInterfaceRequest{
		ID:           physicalInterfaceId,
		Ifname:       iFName,
		Mode:         mode,
		Name:         name,
		Description:  description,
		MTU:          mtu,
		SecurityZone: PhysicalInterfaceSecurityZone,
		IPv4:         IPv4,
		IPv6:         IPv6,
	})

	if err != nil {
		log.Printf("FPU: err=%s", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("FPU: Updated physical interface=%s", physicalInterfaceResponse)

	return resourcePhyInterfaceRead(ctx, d, m)
}

func resourcePhyInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("FPD: Deleting physical interface details, Delete will reset the physical interface")

	deviceId := d.Get("device_id").(string)
	physicalInterfaceId := d.Get("physical_interface_id").(string)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	c := m.(*Client)

	var PhysicalInterfaceSecurityZone = PhysicalInterfaceSecurityZone{
		ID:   "",
		Type: "SecurityZone",
	}
	physicalInterfaceResponse, err := c.UpdateFmcPhysicalInterface(ctx, deviceId, physicalInterfaceId, &PhysicalInterfaceRequest{
		ID:           physicalInterfaceId,
		Ifname:       "-",
		Mode:         "NONE",
		Name:         name,
		Description:  "",
		SecurityZone: PhysicalInterfaceSecurityZone,
	})

	if err != nil {
		log.Printf("FPU: err=%s", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to reset physical interface",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("FPD: Physical interface reseted with value=%s", physicalInterfaceResponse)

	return resourcePhyInterfaceRead(ctx, d, m)
}
