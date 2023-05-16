package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcNetworkAnalysisPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Network Analysis Policies in FMC\n\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"data \"fmc_network_analysis_policies\" \"network_analysis_policy\" {\n" +
			"	name = \"AMP Policy\"\n" +
			"}\n" +
			"```",
		ReadContext: dataSourceFmcNetworkAnalysisPoliciesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the network analysis policy",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"base_policy": {
				Type:     schema.TypeList,
				Optional: true,
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
							Optional:    true,
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
				Description: "Snort engine",
			},
		},
	}
}

func dataSourceFmcNetworkAnalysisPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	networkAnalysisPolicy, err := c.GetFmcNetworkAnalysisPolicyByName(ctx, d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get network analysis policy",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(networkAnalysisPolicy.ID)

	if err := d.Set("id", networkAnalysisPolicy.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set id",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", networkAnalysisPolicy.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set type",
			Detail:   err.Error(),
		})
		return diags
	}

	// Log.Debug(networkAnalysisPolicy.BasePolicy)

	basePolicy := []interface{}{
		map[string]interface{}{
			"name": networkAnalysisPolicy.BasePolicy.Name,
			"type": networkAnalysisPolicy.BasePolicy.Type,
			"id":   networkAnalysisPolicy.BasePolicy.ID,
		},
	}

	if err := d.Set("base_policy", basePolicy); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("name", networkAnalysisPolicy.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set name",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("snort_engine", networkAnalysisPolicy.MetaData.MappedPolicy.SnortEngine); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set snort engine version",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}
