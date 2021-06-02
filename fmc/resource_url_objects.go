package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var url_type string = "Url"

func resourceURLObjects() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceURLObjectsCreate,
		ReadContext:   resourceURLObjectsRead,
		UpdateContext: resourceURLObjectsUpdate,
		DeleteContext: resourceURLObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL for this resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func resourceURLObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateURLObject(ctx, &URLObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Url:         d.Get("url").(string),
		Type:        url_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create url object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceURLObjectsRead(ctx, d, m)
}

func resourceURLObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetURLObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("url", item.URL); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object",
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

func resourceURLObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "url") {
		_, err := c.UpdateURLObject(ctx, id, &URLObjectUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Url:         d.Get("url").(string),
			Type:        url_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update url object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceURLObjectsRead(ctx, d, m)
}

func resourceURLObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteURLObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete url object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
