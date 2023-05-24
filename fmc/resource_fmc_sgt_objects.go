package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var tag_type string = "SecurityGroupTag"

func resourceFmcSGTObjects() *schema.Resource {
	return &schema.Resource{

		Description: "Resource for Security Group Tag in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_sgt_objects\" \"sgt\" {\n" +
			"    name = \"test\"\n" +
			"    description = <description>\n" +
			"    tag = \"27\"\n" +
			"}\n" +
			"```",

		CreateContext: resourceFmcSGTObjectsCreate,
		ReadContext:   resourceFmcSGTObjectsRead,
		UpdateContext: resourceFmcSGTObjectsUpdate,
		DeleteContext: resourceFmcSGTObjectsDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"tag": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The tag of this resource",
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
				Optional:    true,
				Description: "The type of this resource",
				Default:     tag_type,
			},
		},
	}
}

func resourceFmcSGTObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcSGTObjects(ctx, &SecurityGroupTag{
		Type:        tag_type,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tag:         d.Get("tag").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create SGT object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcSGTObjectsRead(ctx, d, m)

}

func resourceFmcSGTObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	id := d.Id()
	securityGroupTag, err := c.GetFmcSGTObjects(ctx, id)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No Security Group Tag found with that name",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(securityGroupTag.ID)

	if err := d.Set("name", securityGroupTag.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set name of security group object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", securityGroupTag.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set the type of the security group tag",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcSGTObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "tag") {
		_, err := c.UpdateFmcSGTObjects(ctx, id, &SecurityGroupTagUpdateInput{
			ID:          id,
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Tag:         d.Get("tag").(string),
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update security group tag",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcSGTObjectsRead(ctx, d, m)
}

func resourceFmcSGTObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	err := c.DeleteFmcSGTObjects(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete security group tag",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
