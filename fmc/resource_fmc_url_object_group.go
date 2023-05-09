package fmc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var url_object_group_type string = "Url"

func resourceFmcURLObjectGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for URL Object Groups in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_url_object_group\" \"URLGroup\" {\n" +
			"    name = \"URLGroup\"\n" +
			"    description = \"Websites\"\n" +
			"    objects {\n" +
			"        id = data.fmc_url_objects.cisco-home.id\n" +
			"        type = data.fmc_url_objects.cisco-home.type\n" +
			"    }\n" +
			"    objects {\n" +
			"        id = fmc_url_objects.cisco-tac.id\n" +
			"        type = fmc_url_objects.cisco-tac.type\n" +
			"    }\n" +
			"    literals {\n" +
			"        url = \"https://www.terraform.io/\"\n" +
			"        type = \"Url\"\n" +
			"    }\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcURLObjectGroupCreate,
		ReadContext:   resourceFmcURLObjectGroupRead,
		UpdateContext: resourceFmcURLObjectGroupUpdate,
		DeleteContext: resourceFmcURLObjectGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
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
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"objects": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
					},
				},
				Description: "List of URL objects to add",
			},
			"literals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value URL for this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of this resource",
						},
					},
				},
				Description: "List of URL literals to add",
			},
		},
	}
}

func resourceFmcURLObjectGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var objs []URLObjectGroupObjects
	var lits []URLObjectGroupLiterals

	if inputObjs, ok := d.GetOk("objects"); ok {
		for _, obj := range inputObjs.([]interface{}) {
			obji := obj.(map[string]interface{})
			objs = append(objs, URLObjectGroupObjects{
				ID:   obji["id"].(string),
				Type: obji["type"].(string),
			})
		}
	}

	if inputLits, ok := d.GetOk("literals"); ok {
		for _, lit := range inputLits.([]interface{}) {
			liti := lit.(map[string]interface{})
			lits = append(lits, URLObjectGroupLiterals{
				URL:  liti["url"].(string),
				Type: url_type,
			})
		}
	}
	res, err := c.CreateFmcURLObjectGroup(ctx, &URLObjectGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        url_object_group_type,
		Objects:     objs,
		Literals:    lits,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create network group object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcURLObjectGroupRead(ctx, d, m)
}

func resourceFmcURLObjectGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcURLObjectGroup(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read url object group",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read url object group",
			Detail:   err.Error(),
		})
		return diags
	}

	objects := make([]interface{}, 0, len(item.Objects))
	for _, obj := range item.Objects {
		obji := make(map[string]interface{})
		obji["id"] = obj.ID
		obji["type"] = obj.Type
		objects = append(objects, obji)
	}

	if err := d.Set("objects", objects); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network group object",
			Detail:   err.Error(),
		})
		return diags
	}

	literals := make([]interface{}, 0, len(item.Literals))
	for _, lit := range item.Literals {
		liti := make(map[string]interface{})
		liti["url"] = lit.URL
		liti["type"] = lit.Type
		literals = append(literals, liti)
	}

	if err := d.Set("literals", literals); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network group object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcURLObjectGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "objects", "literals") {
		var objs []URLObjectGroupObjects
		var lits []URLObjectGroupLiterals

		if inputObjs, ok := d.GetOk("objects"); ok {
			for _, obj := range inputObjs.([]interface{}) {
				obji := obj.(map[string]interface{})
				objs = append(objs, URLObjectGroupObjects{
					ID:   obji["id"].(string),
					Type: obji["type"].(string),
				})
			}
		}

		if inputLits, ok := d.GetOk("literals"); ok {
			for _, lit := range inputLits.([]interface{}) {
				liti := lit.(map[string]interface{})
				lits = append(lits, URLObjectGroupLiterals{
					URL:  liti["url"].(string),
					Type: url_type,
				})
			}
		}
		_, err := c.UpdateFmcURLObjectGroup(ctx, id, &URLObjectGroupUpdateInput{
			ID:          d.Id(),
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        url_object_group_type,
			Objects:     objs,
			Literals:    lits,
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update url object group",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcURLObjectGroupRead(ctx, d, m)
}

func resourceFmcURLObjectGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcURLObjectGroup(ctx, id)
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
