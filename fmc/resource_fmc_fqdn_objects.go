package fmc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var fqdn_type string = "FQDN"

func resourceFmcFQDNObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for FQDN Objects in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_fqdn_objects\" \"new\" {\n" +
			"  name        = \"Cisco\"\n" +
			"  value       = \"cisco.com\"\n" +
			"  description = \"Cisco domain\"\n" +
			"  dns_resolution = \"IPV4_ONLY\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcFQDNObjectsCreate,
		ReadContext:   resourceFmcFQDNObjectsRead,
		UpdateContext: resourceFmcFQDNObjectsUpdate,
		DeleteContext: resourceFmcFQDNObjectsDelete,
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
			"dns_resolution": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `DNS resolution, "IPV4_ONLY", "IPV6_ONLY" or "IPV4_AND_IPV6"`,
			},
		},
	}
}

func resourceFmcFQDNObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcFQDNObject(ctx, &FQDNObject{
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
	return resourceFmcFQDNObjectsRead(ctx, d, m)
}

func resourceFmcFQDNObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcFQDNObject(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read fqdn object",
				Detail:   err.Error(),
			})
			return diags
		}
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
	if err := d.Set("dns_resolution", item.DNSResolution); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read fqdn object",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceFmcFQDNObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value", "dns_resolution") {
		_, err := c.UpdateFmcFQDNObject(ctx, id, &FQDNObjectUpdateInput{
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
	return resourceFmcFQDNObjectsRead(ctx, d, m)
}

func resourceFmcFQDNObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcFQDNObject(ctx, id)
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
