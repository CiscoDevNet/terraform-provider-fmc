package fmc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var network_type string = "Network"

func resourceFmcNetworkObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Network Objects in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_network_objects\" \"PrivateVLANDR\" {\n" +
			"  name        = \"VLAN-Private-DRsite\"\n" +
			"  value       = \"10.10.10.0/24\"\n" +
			"  description = \"Terraform DR network object\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcNetworkObjectsCreate,
		ReadContext:   resourceFmcNetworkObjectsRead,
		UpdateContext: resourceFmcNetworkObjectsUpdate,
		DeleteContext: resourceFmcNetworkObjectsDelete,
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
				Description: "The type this resource",
			},
		},
	}
}

func resourceFmcNetworkObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcNetworkObject(ctx, &NetworkObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Value:       d.Get("value").(string),
		Type:        network_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create network object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcNetworkObjectsRead(ctx, d, m)
}

func resourceFmcNetworkObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcNetworkObject(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read network object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("value", item.Value); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network object",
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

func resourceFmcNetworkObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value") {
		_, err := c.UpdateFmcNetworkObject(ctx, id, &NetworkObjectUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Value:       d.Get("value").(string),
			Type:        network_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update network object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcNetworkObjectsRead(ctx, d, m)
}

func resourceFmcNetworkObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcNetworkObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete network object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
