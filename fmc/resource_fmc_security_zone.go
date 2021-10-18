package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var securityZoneType string = "SecurityZone"

func resourceFmcSecurityZone() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Security Zone in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_security_zone\" \"test\" {\n" +
			"  name          = \"test\"\n" +
			"  interface_mode = \"INLINE\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcSecurityZoneCreate,
		ReadContext:   resourceFmcSecurityZoneRead,
		UpdateContext: resourceFmcSecurityZoneUpdate,
		DeleteContext: resourceFmcSecurityZoneDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of security zone",
			},
			"interface_mode": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Interface mode for this security zone",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"PASSIVE", "INLINE", "SWITCHED", "ROUTED", "ASA"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
			},
		},
	}
}

func resourceFmcSecurityZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcSecurityZone(ctx, &SecurityZoneRequest{
		Name:          d.Get("name").(string),
		Type:          securityZoneType,
		InterfaceMode: d.Get("interface_mode").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create security zone",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcSecurityZoneRead(ctx, d, m)
}

func resourceFmcSecurityZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcSecurityZone(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read security zone",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read security zone",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("interface_mode", item.InterfaceMode); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read security zone",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcSecurityZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description") {
		_, err := c.UpdateFmcSecurityZone(ctx, id, &SecurityZoneRequest{
			Name:          d.Get("name").(string),
			InterfaceMode: d.Get("interface_mode").(string),
			Type:          securityZoneType,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update security zone",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcSecurityZoneRead(ctx, d, m)
}

func resourceFmcSecurityZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcSecurityZone(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete security zone",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
