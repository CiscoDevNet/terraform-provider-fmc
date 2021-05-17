package fmc

import "C"
import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceURLObjectGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceURLObjectGroupCreate,
		ReadContext:   resourceURLObjectGroupRead,
		UpdateContext: resourceURLObjectGroupUpdate,
		DeleteContext: resourceURLObjectGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"objects": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"literals": {
				Type:     schema.TypeSet,
				Required: false,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}


func resourceURLObjectGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	objects := d.Get("objects").(*schema.Set)
	var allobjects map[interface{}]struct{}
	for i, vRaw := range objects.List() {
		val := vRaw.(map[string]interface{})

		allobjects = append(allobjects, i{
			ID:		val["id"].(string),
			Name:	val["name"].(string),
			Type:	val["type"].(string),
		})
	}
	d.Objects = allobjects


	res, err := c.CreateURLObjectGroup(ctx, &URLObjectGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Objects:     d.Get("Objects").(interface{}),
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create url group",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceURLObjectsRead(ctx, d, m)
}

func resourceURLObjectGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetURLObjectGroup(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group name",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group description",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("objects", item.URL); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group: objects",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group type",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceURLObjectGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "objects", "type") {
		_, err := c.UpdateURLObjectGroup(ctx, id, &URLObjectGroupUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Objects:       d.Get("Objects").(interface{}),
			Type:        d.Get("type").(string),
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
	return resourceURLObjectGroupRead(ctx, d, m)
}

func resourceURLObjectGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteURLObjectGroup(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete url object group",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
