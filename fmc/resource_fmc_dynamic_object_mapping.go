package fmc

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
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
		ReadContext:   resourceFmcDynamicObjectMappingRead,
		Schema: map[string]*schema.Schema{
			"dynamic_object_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of dynamic object to be used for mapping",
			},
			"mappings": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "List of IPs to be mapped to dynamic object",
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	dynamicObjectMapping, err := parseDynamicObjectMappingId(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object mapping",
			Detail:   err.Error(),
		})
		return diags
	}

	item, err := c.GetFmcDynamicObjectMapping(ctx, dynamicObjectMapping)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object mapping",
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

	mappings := []string{}
	for _, mapping := range d.Get("mappings").([]interface{}) {
		mappings = append(mappings, mapping.(string))
	}

	err := c.CreateFmcDynamicObjectMapping(ctx,
		&DynamicObjectMapping{
			DynamicObject: DynamicObjectMappingObject{
				ID: d.Get("dynamic_object_id").(string),
			},
			Mappings: mappings,
		})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create dynamic object mapping",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(generateDynamicObjectMappingId(d.Get("dynamic_object_id").(string), mappings))
	return diags
}

func resourceFmcDynamicObjectMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	dynamicObjectMapping, err := parseDynamicObjectMappingId(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read dynamic object mapping",
			Detail:   err.Error(),
		})
		return diags
	}

	err = c.DeleteFmcDynamicObjectMapping(ctx, dynamicObjectMapping)
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

func generateDynamicObjectMappingId(dynamicObjectId string, mappings []string) string {
	return fmt.Sprintf("%s+%s", dynamicObjectId, strings.Join(mappings, "+"))
}

func parseDynamicObjectMappingId(id string) (*DynamicObjectMapping, error) {
	substrings := strings.Split(id, "+")
	if len(substrings) < 2 {
		return nil, fmt.Errorf("unable to parse dynamic object mapping id")
	}

	return &DynamicObjectMapping{
		DynamicObject: DynamicObjectMappingObject{
			ID: substrings[0],
		},
		Mappings: substrings[1:],
	}, nil
}
