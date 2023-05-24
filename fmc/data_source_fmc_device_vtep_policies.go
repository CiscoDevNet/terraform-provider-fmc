package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcVTEPPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for VTEP policy on FMC\n\n" +
            "An example is shown below: \n" +
            "```hcl\n" +
            "data \"fmc_device_vtep_policies\" \"vtep\" {\n" +
            "   device_id = \"<device ID>\"\n" +
            "}\n" +
            "```",
		ReadContext: dataSourceFmcVTEPPoliciesRead,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of this VTEP Policies",
			},

			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this VTEP Policies",
			},

			"nveenable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "NveEnable of this VTEP Policies",
			},
		},
	}
}

func dataSourceFmcVTEPPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	resp, err := c.GetVTEPPolicies(ctx, d.Get("device_id").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VTEP Policies",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VTEP Details ID=%v Type=%v", resp.ID, resp.Type)

	d.SetId(resp.ID)

	if err := d.Set("type", resp.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read VTEP Policies",
			Detail:   err.Error(),
		})
		return diags
	}

	vtepdetails, err := c.GetFmcVTEPDetails(ctx, d.Get("device_id").(string), resp.ID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get VTEP",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VTEP Details ID=%s Type=%s NVEEnabled=%v", vtepdetails.ID, vtepdetails.Type, vtepdetails.NveEnable)

	d.SetId(vtepdetails.ID)
	d.Set("type", vtepdetails.Type)
	d.Set("nveenable", vtepdetails.NveEnable)

	return diags
}
