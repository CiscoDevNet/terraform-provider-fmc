package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcIPSPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Intrusion Policies in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_ips_policies\" \"ips_policy\" {\n" +
			"    name = \"test\"\n" +
			"    inspection_mode = \"DETECTION\"\n" +
			"    basepolicy_id = <basepolicy-id>\n" +
			"}\n" +
			"```",

		CreateContext: resourceFmcIPSPoliciesCreate,
		ReadContext:   resourceFmcIPSPoliciesRead,
		UpdateContext: resourceFmcIPSPoliciesUpdate,
		DeleteContext: resourceFmcIPSPoliciesDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},

			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"inspection_mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The inspection mode of this policy",
			},
			"basepolicy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the base policy",
			},
			"basepolicy_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the base policy",
			},
		},
	}
}

func resourceFmcIPSPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcIPS(ctx, &IPSPolicy{
		Name:           d.Get("name").(string),
		Type:           "IntrusionPolicy",
		InspectionMode: d.Get("inspection_mode").(string),
		BasePolicy: Base_policy{
			ID:   d.Get("basepolicy_id").(string),
			Name: d.Get("basepolicy_name").(string),
			Type: "IntrusionPolicy",
		},
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create IPS policy",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcIPSPoliciesRead(ctx, d, m)

}

func resourceFmcIPSPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcIPSPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read IPS policy",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read IPS policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read IPS policy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcIPSPoliciesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name", "type", "inspection_mode", "basepolicy_id", "basepolicy_name") {

		res, err := c.UpdateFmcIPSPolicy(ctx, d.Id(), &IPSPolicy{
			ID:             d.Id(),
			Name:           d.Get("name").(string),
			Type:           "IntrusionPolicy",
			InspectionMode: d.Get("inspection_mode").(string),
			BasePolicy: Base_policy{
				ID:   d.Get("basepolicy_id").(string),
				Name: d.Get("basepolicy_name").(string),
				Type: "IntrusionPolicy",
			},
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to create IPS policy",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcIPSPoliciesRead(ctx, d, m)
}

func resourceFmcIPSPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcIPSPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete IPS policy",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
