package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcAccessPoliciesCategory() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for access policy category in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_access_policies_category\" \"category\" {\n" +
			"    name        		  = \"test-time-range\"\n" +
			"    access_policy_id     = \"BB62F664-7168-4C8E-B4CE-F70D522889D2\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcAccessPoliciesCategoryCreate,
		ReadContext:   resourceFmcAccessPoliciesCategoryRead,
		DeleteContext: resourceFmcAccessPoliciesCategoryDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of this category",
			},
			"access_policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of access policy this category belongs to",
			},
		},
	}
}

func resourceFmcAccessPoliciesCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcAccessPoliciesCategory(ctx, d.Get("access_policy_id").(string), &AccessPolicyCategory{
		Name: d.Get("name").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create access policy category",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcAccessPoliciesCategoryRead(ctx, d, m)
}

func resourceFmcAccessPoliciesCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	accessPolicyID := d.Get("access_policy_id").(string)

	item, err := c.GetFmcAccessPoliciesCategory(ctx, id, accessPolicyID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy category",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy category",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("access_policy_id", accessPolicyID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcAccessPoliciesCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	accessPolicyID := d.Get("access_policy_id").(string)

	err := c.DeleteFmcAccessPoliciesCategory(ctx, id, accessPolicyID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete access policy category",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
