package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcTimeRangeObject() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Time Range Object in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_time_range_object\" \"test\" {\n" +
			"    name        		  = \"test-time-range\"\n" +
			"    description 		  = \"Test time range\"\n" +
			"    effective_start_date = \"2019-09-19T15:53:00\"\n" +
			"    effective_end_date   = \"2019-09-21T17:53:00\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcTimeRangeObjectCreate,
		ReadContext:   resourceFmcTimeRangeObjectRead,
		UpdateContext: resourceFmcTimeRangeObjectUpdate,
		DeleteContext: resourceFmcTimeRangeObjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"effective_start_date": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Effective start date for this time range object (time in RFC3339 format)",
			},
			"effective_end_date": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Effective end date for this time range object (time in RFC3339 format)",
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

func resourceFmcTimeRangeObjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcTimeRangeObject(ctx, &TimeRangeObjectInput{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		EffectiveStartDate: d.Get("effective_start_date").(string),
		EffectiveEndDate:   d.Get("effective_end_date").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create time range object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcTimeRangeObjectRead(ctx, d, m)
}

func resourceFmcTimeRangeObjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcTimeRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("effective_start_date", item.EffectiveStartDate); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("effective_end_date", item.EffectiveEndDate); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcTimeRangeObjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "effective_start_date", "effective_end_date") {
		_, err := c.UpdateFmcTimeRangeObject(ctx, id, &TimeRangeObject{
			Name:               d.Get("name").(string),
			Description:        d.Get("description").(string),
			EffectiveStartDate: d.Get("effective_start_date").(string),
			EffectiveEndDate:   d.Get("effective_end_date").(string),
			ID:                 id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update time range object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcTimeRangeObjectRead(ctx, d, m)
}

func resourceFmcTimeRangeObjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcTimeRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
