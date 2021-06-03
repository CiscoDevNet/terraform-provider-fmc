package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHostObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Host Objects in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_host_objects\" \"existing_host\" {\n" +
			"	name = \"CUCM-Pub\"\n" +
			"}\n" +
			"```\n" +
			"Any one of the id, name or value can be specified. The first filter in the order of id, name and value will be used, and the rest will be ignored if multiple are specified.",
		ReadContext: dataSourceHostObjectsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of this resource",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The value of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
	}
}

func dataSourceHostObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	idInput, okId := d.GetOk("id")
	nameInput, okName := d.GetOk("name")
	valueInput, okValue := d.GetOk("value")
	var (
		item *HostObjectResponse
		err  error
	)
	if (okId && (okName || okValue)) || (okName && okValue) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "More than one filter provided",
			Detail:   "The first filter in the order of id, name and value will be used, and the rest will be ignored",
		})
	}
	switch {
	case okId:
		item, err = c.GetHostObject(ctx, idInput.(string))
	case okName:
		item, err = c.GetHostObjectByNameOrValue(ctx, nameInput.(string))
	case okValue:
		item, err = c.GetHostObjectByNameOrValue(ctx, valueInput.(string))
	default:
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No id, name, value not provided, please provide any one",
			Detail:   "Please set one of the values to filter the datasource by",
		})
		return diags
	}

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get host object",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(item.ID)

	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read host object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("value", item.Value); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read host object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read host object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
