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

	c := m.(*Client)

	var diags diag.Diagnostics

	var PhysicalInterfaceSecurityZone = PhysicalInterfaceSecurityZone{
		ID:   securityZoneId,
		Type: "SecurityZone",
	}

	log.Printf("PhysicalInterfaceSecurityZone=%s", PhysicalInterfaceSecurityZone)

	physicalInterfaceResponse, err := c.UpdateFmcPhysicalInterface(ctx, deviceId, physicalInterfaceId, &PhysicalInterfaceRequest{
		ID:           physicalInterfaceId,
		Ifname:       iFName,
		Mode:         mode,
		Name:         name,
		Description:  description,
		MTU:          mtu,
		SecurityZone: PhysicalInterfaceSecurityZone,
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
