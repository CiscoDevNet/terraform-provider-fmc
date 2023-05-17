package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var NWAtype string = "NetworkAnalysisPolicy"

func resourceFmcNetworkAnalysisPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Network Analysis Policy in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_network_analysis_policy\" \"example\" {\n" +
			"  name = \"NWA-test\"\n" +
			"  description = \"This is the description\"\n" +
			"  base_policy {\n" +
			"    name = \"Balanced Security and Connectivity | Connectivity Over Security | Maximum Detection | Security Over Connectivity\"\n" +
			"  }\n" +
			"  snort_engine = \"SNORT2 | SNORT3\"\n" +
			"}\n" +
			"```\n",
		CreateContext: resourceFmcNetworkAnalysisPolicyCreate,
		ReadContext:   resourceFmcNetworkAnalysisPolicyRead,
		UpdateContext: resourceFmcNetworkAnalysisPolicyUpdate,
		DeleteContext: resourceFmcNetworkAnalysisPolicyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of policy",
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
			"base_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of base policy",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of base policy",
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Description: "The base policy of this resource",
			},
			"snort_engine": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The snort engine of this resource",
			},
		},
	}
}

func resourceFmcNetworkAnalysisPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	base_policy := d.Get("base_policy").([]interface{})[0].(map[string]interface{})["name"].(string)

	r, _ := c.GetFmcNetworkAnalysisPolicyByName(ctx, base_policy)
	var base = NetworkAnalysisPolicyBasePolicyInput{
		ID:   r.ID,
		Type: r.Type,
	}

	// get snort_engine and if it's empty set it to SNORT3
	var se string
	if d.Get("snort_engine").(string) == "" {
		se = "SNORT3"
	} else {
		se = d.Get("snort_engine").(string)
	}

	item, err := c.CreateFmcNetworkAnalysisPolicy(ctx, &NetworkAnalysisPolicy{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        NWAtype,
		BasePolicy:  base,
		SnortEngine: se,
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create Network analysis policy",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(item.ID)
	return resourceFmcNetworkAnalysisPolicyRead(ctx, d, m)
}

func resourceFmcNetworkAnalysisPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	id := d.Id()

	item, err := c.GetFmcNetworkAnalysisPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read network analysis policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set name",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set description",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set type",
			Detail:   err.Error(),
		})
		return diags
	}

	basePolicy := []interface{}{
		map[string]interface{}{
			"name": item.BasePolicy.Name,
			"type": item.BasePolicy.Type,
			"id":   item.BasePolicy.ID,
		},
	}

	if err := d.Set("base_policy", basePolicy); err != nil {
		return returnWithDiag(diags, err)
	}

	// if err := d.Set("snort_engine", item.MetaData.MappedPolicy.SnortEngine); err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "unable to set snort engine",
	// 		Detail:   err.Error(),
	// 	})
	// 	return diags
	// }

	return diags
}

func resourceFmcNetworkAnalysisPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "base_policy") {
		var basep = NetworkAnalysisPolicyBasePolicyInput{}
		if inputEntries, ok := d.GetOk("base_policy"); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})

			basep = NetworkAnalysisPolicyBasePolicyInput{
				Name: entry["name"].(string),
				Type: entry["type"].(string),
				ID:   entry["id"].(string),
			}
		}

		res, err := c.UpdateFmcNetworkAnalysisPolicy(ctx, id, &NetworkAnalysisPolicy{
			ID:          id,
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			BasePolicy:  basep,
			SnortEngine: d.Get("snort_engine").(string),
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update network analysis policy",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcNetworkAnalysisPolicyRead(ctx, d, m)
}

func resourceFmcNetworkAnalysisPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcNetworkAnalysisPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete network analysis policy",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
