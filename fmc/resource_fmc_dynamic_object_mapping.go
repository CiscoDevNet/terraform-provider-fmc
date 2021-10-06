package fmc

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func resourceFmcDynamicObjectMapping() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Dynamic Object Mapping in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_dynamic_object\" \"test\" {\n" +
			"  name        = \"test\"\n" +
			"  object_type = \"IP\"\n" +
			"  description = \"test dynamic object\"\n" +
			"}\n" +
			"resource \"fmc_dynamic_object_mapping\" \"test\" {\n" +
			"  dynamic_object_id = fmc_dynamic_object.test.id\n" +
			"  mappings = \"8.8.8.8\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcDynamicObjectMappingCreate,
		DeleteContext: resourceFmcDynamicObjectMappingDelete,
		ReadContext: resourceFmcDynamicObjectMappingRead,
		Schema: map[string]*schema.Schema{
			"dynamic_object_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of dynamic object to be used for mapping",
			},
			"mappings": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IP or list of IPs to be mapped to dynamic object",
			},
		},
	}
}

func resourceFmcDynamicObjectMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcDynamicObjectMapping(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("dynamic_object_id", item.DynamicObject.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object mapping",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("mappings", item.Mappings); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object mapping",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcDynamicObjectMappingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	id, err := c.CreateFmcDynamicObjectMapping(ctx,
		d.Get("dynamic_object_id").(string),
		d.Get("mappings").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create dynamic object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(id)
	return diags
}

func resourceFmcDynamicObjectMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcDynamicObjectMapping(ctx, d.Get("dynamic_object_id").(string),
		d.Get("mappings").(string))
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
