package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcVNI() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for VNI Interface in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_device_vni\" \"test-vni\" {\n" +
			"	name = \"TEST-VNI\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceVNIRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the VNI",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the VNI",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the VNI",
			},
			"segment_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SegmentId of the VNI",
			},
			"vnid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "VNID of the VNI",
			},
			"enable_proxy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "EnableProxy of the VNI",
			},
			"multicast_groupaddress": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "MulticastGroupAddress of the VNI",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Priority of the VNI",
			},
			"if_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ifname of the VNI",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "securityZone of the VNI",
			},
		},
	}
}

func dataSourceVNIRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VNI Datasource")

	c := m.(*Client)

	var diags diag.Diagnostics
	vni, err := c.GetFmcVNI(ctx, d.Get("device_id").(string), d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("ID=%s Name=%s", vni.ID, vni.Name)

	d.SetId(vni.ID)
	d.Set("name", vni.Name)

	if err := d.Set("name", vni.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	vnidetail, err := c.GetFmcVNIDetails(ctx, d.Get("device_id").(string), vni.ID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VNI Details ID=%s Name=%s Description=%s", vnidetail.ID, vnidetail.Name, vnidetail.Description)

	d.SetId(vnidetail.ID)
	d.Set("device_id", d.Get("device_id").(string))
	d.Set("name", vnidetail.Name)
	d.Set("description", vnidetail.Description)
	d.Set("segment_id", vnidetail.SegmentId)
	d.Set("vnid", vnidetail.VNIID)
	d.Set("enable_proxy", vnidetail.EnableProxy)
	d.Set("multicast_groupaddress", vnidetail.MulticastGroupAddress)
	d.Set("priority", vnidetail.Priority)
	d.Set("if_name", vnidetail.Ifname)
	d.Set("security_zone_id", vnidetail.SecurityZone.ID)

	return diags
}
