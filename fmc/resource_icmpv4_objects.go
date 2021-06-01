package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var icmpv4_type string = "ICMPV4Object"

func resourceICMPV4Objects() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceICMPV4ObjectsCreate,
		ReadContext:   resourceICMPV4ObjectsRead,
		UpdateContext: resourceICMPV4ObjectsUpdate,
		DeleteContext: resourceICMPV4ObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"icmp_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ICMP type for this resource",
			},
			"code": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ICMP code for this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func resourceICMPV4ObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var code *int
	if inputCode, ok := d.GetOk("code"); ok {
		intcode := inputCode.(int)
		code = &intcode
	}
	res, err := c.CreateICMPV4Object(ctx, &ICMPV4Object{
		Name:     d.Get("name").(string),
		Icmptype: d.Get("icmp_type").(string),
		Code:     code,
		Type:     icmpv4_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceICMPV4ObjectsRead(ctx, d, m)
}

func resourceICMPV4ObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetICMPV4Object(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("icmp_type", item.Icmptype); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("code", item.Code); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceICMPV4ObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	var code *int
	if inputCode, ok := d.GetOk("code"); ok {
		intcode := inputCode.(int)
		code = &intcode
	}
	if d.HasChanges("name", "description", "value") {
		_, err := c.UpdateICMPV4Object(ctx, id, &ICMPV4ObjectUpdateInput{
			Name:     d.Get("name").(string),
			Icmptype: d.Get("icmp_type").(string),
			Code:     code,
			Type:     icmpv4_type,
			ID:       id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update icmpv4 object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceICMPV4ObjectsRead(ctx, d, m)
}

func resourceICMPV4ObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteICMPV4Object(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete icmpv4 object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
