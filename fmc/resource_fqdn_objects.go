package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var fqdn_type string = "FQDN"

func resourceFQDNObjects() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFQDNObjectsCreate,
		ReadContext:   resourceFQDNObjectsRead,
		UpdateContext: resourceFQDNObjectsUpdate,
		DeleteContext: resourceFQDNObjectsDelete,
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
				Description: "The description for this resource",
			},
			"dns_resolution": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `DNS resolution, "IPV4_ONLY", "IPV6_ONLY" or "IPV4_AND_IPV6"`,
			},
		},
	}
}

func resourceFQDNObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFQDNObject(ctx, &FQDNObject{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Value:         d.Get("value").(string),
		DNSResolution: d.Get("dns_resolution").(string),
		Type:          fqdn_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFQDNObjectsRead(ctx, d, m)
}

func resourceFQDNObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFQDNObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("value", item.Value); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("dns_resolution", item.Value); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceFQDNObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value", "dns_resolution") {
		_, err := c.UpdateFQDNObject(ctx, id, &FQDNObjectUpdateInput{
			Name:          d.Get("name").(string),
			Description:   d.Get("description").(string),
			Value:         d.Get("value").(string),
			DNSResolution: d.Get("dns_resolution").(string),
			Type:          fqdn_type,
			ID:            id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update fqdn object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFQDNObjectsRead(ctx, d, m)
}

func resourceFQDNObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFQDNObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
