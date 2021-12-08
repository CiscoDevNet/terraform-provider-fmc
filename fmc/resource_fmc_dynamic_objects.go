package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dynamicObjectType string = "DynamicObject"

func resourceFmcDynamicObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Dynamic Object in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_dynamic_object\" \"test\" {\n" +
			"  name        = \"test\"\n" +
			"  object_type = \"IP\"\n" +
			"  description = \"test dynamic object\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcDynamicObjectsCreate,
		ReadContext:   resourceFmcDynamicObjectsRead,
		UpdateContext: resourceFmcDynamicObjectsUpdate,
		DeleteContext: resourceFmcDynamicObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of this resource",
			},
			"object_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of Dynamic Object. Allowed values: IP",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"IP"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this resource",
				Default:     " ",
				StateFunc: func(val interface{}) string {
					state := val.(string)
					if val == nil || state == "" || state == " " {
						return " "
					}
					return state
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Fix for bug in the FMC API which returns " " for empty description
					if (new == " " && old == "") || (old == " " && new == "") {
						return true
					}
					return old == new
				},
			},
		},
	}
}

func resourceFmcDynamicObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcDynamicObject(ctx, &DynamicObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        dynamicObjectType,
		ObjectType:  d.Get("object_type").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcDynamicObjectsRead(ctx, d, m)
}

func resourceFmcDynamicObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcDynamicObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("object_type", item.ObjectType); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcDynamicObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "object_type") {
		_, err := c.UpdateFmcDynamicObject(ctx, id, &DynamicObjectUpdated{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			ObjectType:  d.Get("object_type").(string),
			Type:        network_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update dynamic object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcDynamicObjectsRead(ctx, d, m)
}

func resourceFmcDynamicObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcDynamicObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
