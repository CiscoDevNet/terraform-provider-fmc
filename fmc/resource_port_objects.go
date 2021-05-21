package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var port_type string = "ProtocolPortObject"

func resourcePortObjects() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortObjectsCreate,
		ReadContext:   resourcePortObjectsRead,
		UpdateContext: resourcePortObjectsUpdate,
		DeleteContext: resourcePortObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"overridable": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePortObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreatePortObject(ctx, &PortObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Port:        d.Get("port").(string),
		Protocol:    d.Get("protocol").(string),
		Overridable: d.Get("overridable").(bool),
		Type:        d.Get("type").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create port object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourcePortObjectsRead(ctx, d, m)
}

func resourcePortObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetPortObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("port", item.Port); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("protocol", item.Protocol); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourcePortObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value") {
		_, err := c.UpdatePortObject(ctx, id, &PortObjectUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Port:        d.Get("port").(string),
			Protocol:    d.Get("protocol").(string),
			Overridable: d.Get("overridable").(bool),
			Type:        port_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update port object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourcePortObjectsRead(ctx, d, m)
}

func resourcePortObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeletePortObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete port object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
