package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcStandardAcl() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Standard ACL in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_standard_acl\" \"acl1\" {\n" +
			"    name = \"ACL-1\"\n" +
			"    action = \"DENY\"\n" +
			"    object_id = data.fmc_network_objects.any.id\n" +
			"    literal_type = \"Host\"\n" +
			"    literal_value = \"1.1.1.1\"\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcStandardAclCreate,
		ReadContext:   resourceFmcStandardAclRead,
		UpdateContext: resourceFmcStandardAclUpdate,
		DeleteContext: resourceFmcStandardAclDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type this resource",
			},
			"object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of this resource",
			},
			"literal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of this resource",
			},
			"literal_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of this resource",
			},
		},
	}
}

func resourceFmcStandardAclCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*Client)
    // Warning or errors can be collected in a slice type
    // var diags diag.Diagnostics
    var diags diag.Diagnostics
    
    objectID := d.Get("object_id").(string)
    var Object_input []Object_nw
    if len(objectID) > 0 {
        Object_input = append(Object_input, Object_nw{
            ID:      objectID,
        })
    }

    litType := d.Get("literal_type").(string)
    litValue := d.Get("literal_value").(string)
    var Lit_input []Literal_nw
    if len(litType) > 0 && len(litValue) > 0 {
        Lit_input = append(Lit_input, Literal_nw{
            Type:      litType,
            Value:     litValue,
        })
    }

    // var Network_in = Network{Objects: Object_input, Literals: Lit_input }

    res, err := c.CreateFmcStandardAcl(ctx, &StandardAcl{
        Name: d.Get("name").(string),
        Type: "StandardAccessList",
        Entries: []ObjectEntries{
            {
            Action: d.Get("action").(string),
            Networks: Network{
                Objects: Object_input,
                Literals: Lit_input,
            },
            },
        },
    })
    if err != nil {
        diags = append(diags, diag.Diagnostic{
            Severity: diag.Error,
            Summary:  "unable to create standard access list",
            Detail:   err.Error(),
        })
        return diags
    }
    d.SetId(res.ID)
    return resourceFmcStandardAclRead(ctx, d, m)
}


func resourceFmcStandardAclRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcStandardAcl(ctx, d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read std acl",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read std acl",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read std acl",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}
func resourceFmcStandardAclUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	objectID := d.Get("object_id").(string)
    var Object_input []Object_nw
    if len(objectID) > 0 {
        Object_input = append(Object_input, Object_nw{
            ID:      objectID,
        })
    }

    litType := d.Get("literal_type").(string)
    litValue := d.Get("literal_value").(string)
    var Lit_input []Literal_nw
    if len(litType) > 0 && len(litValue) > 0 {
        Lit_input = append(Lit_input, Literal_nw{
            Type:      litType,
            Value:     litValue,
        })
    }

	if d.HasChanges("name", "object_id", "action","literal_value","literal_type" ) {
		res, err := c.UpdateFmcStandardAcl(ctx, d.Id(), &StandardAcl{
			ID:   d.Id(),
			Name: d.Get("name").(string),
			Type: "StandardAccessList",
			Entries: []ObjectEntries{
				{
				Action: d.Get("action").(string),
				Networks: Network{
                Objects: Object_input,
                Literals: Lit_input,
           		 },
				},
			},
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update standard acl",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcStandardAclRead(ctx, d, m)
}

func resourceFmcStandardAclDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcStandardAcl(ctx, d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete route",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}