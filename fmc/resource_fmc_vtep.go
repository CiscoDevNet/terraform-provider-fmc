package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVTEP() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for configuring VTEP\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_device_vtep\" \"my_fmc_device_vtep\" {\n" +
			"    device_id = data.fmc_devices.device.id\n" +
			"    nve_enabled = true\n" +
			"    nve_vtep_id = 1\n" +
			"    nve_destination_port = 6081\n" +
			"    nve_encapsulation_type = \"GENEVE\"\n" +
			"    source_interface_id = data.fmc_device_physical_interfaces.physical_interface1.id\n" +
			"}\n" +
			"```\n" +
			"**Note:** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.\n",
		CreateContext: resourceVTEPCreate,
		ReadContext:   resourceVTEPRead,
		UpdateContext: resourceVTEPUpdate,
		DeleteContext: resourceVTEPDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Device Id of VTEP",
			},
			"nve_enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "NVE Enabled",
			},
			"source_interface_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source Interface Id",
			},
			"nve_vtep_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "NVE VTEP Id",
			},
			"nve_encapsulation_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "NVE Encapsulation type",
			},
			"nve_destination_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "NVE Destination port",
			},
			"nve_neighbor_discovery_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "NVE Discovery type",
			},
			"neighbor_addr_object_overridable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If Object Overridable",
			},
			"neighbor_addr_object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Neighbor Address Object ID",
			},
			"neighbor_addr_literal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Neighbor Address literal type",
			},
			"neighbor_addr_literal_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Neighbor Address literal value",
			},
		},
	}
}

func resourceVTEPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	log.Printf("VTEP: Fetching VTEP details")

	if d.Get("device_id") != nil && d.Get("id") != nil {

		deviceId := d.Get("device_id").(string)
		id := d.Get("id").(string)

		log.Printf("VTEP: DeviceId=%s, Id=%s", deviceId, id)

		c := m.(*Client)

		vtep, err := c.GetFmcVTEPDetails(ctx, deviceId, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read vtep",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("VTEP: VTEP Details Id=%v Type=%v NVEEnabled=%v", vtep.ID, vtep.Type, vtep.NveEnable)

		d.SetId(vtep.ID)

	} else {
		log.Printf("VTEP: No Proper Data Of Device Id=%s, Id=%s", d.Get("device_id"), d.Get("id"))
	}

	return diags
}

func resourceVTEPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VTEP: Creating VTEP")

	device_id := d.Get("device_id").(string)
	nve_enabled := d.Get("nve_enabled").(bool)
	source_interface_id := d.Get("source_interface_id").(string)
	nve_vtep_id := d.Get("nve_vtep_id").(int)
	nve_encapsulation_type := d.Get("nve_encapsulation_type").(string)
	nve_destination_port := d.Get("nve_destination_port").(int)
	nve_neighbor_discovery_type := d.Get("nve_neighbor_discovery_type").(string)
	objectOverridable := d.Get("neighbor_addr_object_overridable").(bool)
	neighbor_addr_object_id := d.Get("neighbor_addr_object_id").(string)
	neighbor_addr_literal_type := d.Get("neighbor_addr_literal_type").(string)
	neighbor_addr_literal_value := d.Get("neighbor_addr_literal_value").(string)
	log.Printf("device_id=%v nve_enabled=%v, source_interface_id=%v, nve_vtep_id=%v, nve_encapsulation_type=%v nve_destination_port=%v", device_id, nve_enabled, source_interface_id, nve_vtep_id, nve_encapsulation_type, nve_destination_port)

	c := m.(*Client)

	var diags diag.Diagnostics

	var sourceInterface = SourceInterface{
		ID: source_interface_id,
	}
	
	var objects = Object{
		Overridable:objectOverridable,
		ID: neighbor_addr_object_id,
	}
	var literals = Literal{
		Type: neighbor_addr_literal_type,
		Value: neighbor_addr_literal_value,
	}
	var NveNeighborAddress = NveNeighborAddress{
		Object: objects,
		Literal: literals,
	}
	var vtepEntry [1]VTEPEntry

	vtepEntry[0] = VTEPEntry{
		NveVtepId:            nve_vtep_id,
		NveEncapsulationType: nve_encapsulation_type,
		NveDestinationPort:   nve_destination_port,
		SourceInterface:      sourceInterface,
		NveNeighborDiscoveryType: nve_neighbor_discovery_type,
		NveNeighborAddress: NveNeighborAddress,
	}
	vtepResp, err := c.CreateFmcVTEP(ctx, device_id, &VTEPPolicy{
		NveEnable:   nve_enabled,
		VTEPEntries: vtepEntry[:],
	})

	if err != nil {
		log.Printf("VTEP: err=%s", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to crete VTEP",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VTEP Id=%s", vtepResp.ID)

	d.SetId(vtepResp.ID)
	return diags
}

func resourceVTEPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VTEP: Updating VTEP")
	id := d.Id()
	var diags diag.Diagnostics

	if len(id) > 0 {

		device_id := d.Get("device_id").(string)
		nve_enabled := d.Get("nve_enabled").(bool)
		source_interface_id := d.Get("source_interface_id").(string)
		nve_vtep_id := d.Get("nve_vtep_id").(int)
		nve_encapsulation_type := d.Get("nve_encapsulation_type").(string)
		nve_destination_port := d.Get("nve_destination_port").(int)
		nve_neighbor_discovery_type := d.Get("nve_neighbor_discovery_type").(string)
		objectOverridable := d.Get("neighbor_addr_object_overridable").(bool)
		neighbor_addr_object_id := d.Get("neighbor_addr_object_id").(string)
		neighbor_addr_literal_type := d.Get("neighbor_addr_literal_type").(string)
		neighbor_addr_literal_value := d.Get("neighbor_addr_literal_value").(string)

		log.Printf("id=%v device_id=%v nve_enabled=%v, source_interface_id=%v, nve_vtep_id=%v, nve_encapsulation_type=%v nve_destination_port=%v", id, device_id, nve_enabled, source_interface_id, nve_vtep_id, nve_encapsulation_type, nve_destination_port)

		c := m.(*Client)

		var objects = Object{
			Overridable:objectOverridable,
			ID: neighbor_addr_object_id,
		}
	
		var literals = Literal{
			Type: neighbor_addr_literal_type,
			Value: neighbor_addr_literal_value,
		}
	
		var NveNeighborAddress = NveNeighborAddress{
			Object: objects,
			Literal: literals,
		}

		var sourceInterface = SourceInterface{
			ID: source_interface_id,
		}
		var vtepEntry [1]VTEPEntry

		vtepEntry[0] = VTEPEntry{
			NveVtepId:            nve_vtep_id,
			NveEncapsulationType: nve_encapsulation_type,
			NveDestinationPort:   nve_destination_port,
			SourceInterface:      sourceInterface,
			NveNeighborDiscoveryType: nve_neighbor_discovery_type,
			NveNeighborAddress: NveNeighborAddress,
		}
		vtepResp, err := c.UpdateFmcVTEP(ctx, device_id, id, &VTEPPolicyUpdate{
			ID:          id,
			NveEnable:   nve_enabled,
			VTEPEntries: vtepEntry[:],
		})

		if err != nil {
			log.Printf("VTEP: err=%s", err)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to crete VTEP",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("VTEP Id=%s", vtepResp.ID)

		d.SetId(vtepResp.ID)
	} else {
		log.Printf("Update VTEP ID is null, so record cannot be updated")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update vtep",
			Detail:   " VTEP ID is null, so record cannot be updated",
		})
	}

	return diags
}

func resourceVTEPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VTEP: Deleting VTEP")

	var diags diag.Diagnostics

	id := d.Id()

	log.Printf("VTEP Delete Id=%v", id)

	if len(id) > 0 && len(d.Get("device_id").(string)) > 0 {
		device_id := d.Get("device_id").(string)

		log.Printf("VTEP Delete Id=%v and Device Id=%v", id, device_id)

		c := m.(*Client)

		vtepResp, err := c.DeleteFmcVTEPDetails(ctx, device_id, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to delete VTEP as not able to get VTEP details based on id",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Deleted VTEP Id=%s", vtepResp.ID)

		d.SetId(vtepResp.ID)
	} else {
		log.Printf("Unable to delete the VTEP as device_id or id is not present")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete the VTEP as device_id or id is not present",
			Detail:   "ID is null, so record cannot be deleted",
		})
	}

	return diags
}
