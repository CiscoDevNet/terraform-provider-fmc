package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var range_type string = "Range"

func resourceRangeObjects() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRangeObjectsCreate,
		ReadContext:   resourceRangeObjectsRead,
		UpdateContext: resourceRangeObjectsUpdate,
		DeleteContext: resourceRangeObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of this resource",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of this resource",
			},
		},
	}
}

func resourceRangeObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateRangeObject(ctx, &RangeObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Value:       d.Get("value").(string),
		Type:        range_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create network object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceRangeObjectsRead(ctx, d, m)
}

func resourceRangeObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("value", item.Value); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceRangeObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value") {
		_, err := c.UpdateRangeObject(ctx, id, &RangeObjectUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Value:       d.Get("value").(string),
			Type:        range_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update network object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceRangeObjectsRead(ctx, d, m)
}

func resourceRangeObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete network object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
