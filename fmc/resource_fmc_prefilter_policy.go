package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcPrefilterPolicy() *schema.Resource {
	return &schema.Resource{
		Description: `
			resource "fmc_prefilter_policy" "prefilter_policy" {
				name        = "Prefilter Policy"
				description = "Terraform Prefilter Policy description"
				default_action { 
					log_end = true
					log_begin = true
					send_events_to_fmc = true
					action = "BLOCK_TUNNELS"
				}
			}
			`,
		CreateContext: resourceFmcPrefilterPolicyCreate,
		ReadContext:   resourceFmcPrefilterPolicyRead,
		UpdateContext: resourceFmcPrefilterPolicyUpdate,
		DeleteContext: resourceFmcPrefilterPolicyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
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
			"default_action": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//"log_end": {
						//	Type:        schema.TypeBool,
						//	Optional: true,
						//	Description: "Log end",
						//},
						"log_begin": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Log begin",
						},
						"send_events_to_fmc": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Send events to FMC",
						},
						"action": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Action. Should be BLOCK_TUNNELS or ANALYZE_TUNNELS",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := strings.ToUpper(val.(string))
								allowedValues := []string{"BLOCK_TUNNELS", "ANALYZE_TUNNELS"}
								for _, allowed := range allowedValues {
									if v == allowed {
										return
									}
								}
								errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
								return
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Description: "Source interface for this resource",
			},
		},
	}
}

func resourceFmcPrefilterPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var defaultAction = PrefilterPolicyDefaultActionInput{}
	if inputEntries, ok := d.GetOk("default_action"); ok {
		entry := inputEntries.([]interface{})[0].(map[string]interface{})

		defaultAction = PrefilterPolicyDefaultActionInput{
			LogBegin: entry["log_begin"].(bool),
			//LogEnd:          entry["log_end"].(bool),
			SendEventsToFMC: entry["send_events_to_fmc"].(bool),
			Action:          entry["action"].(string),
		}
	}

	res, err := c.CreateFmcPrefilterPolicy(ctx, &PrefilterPolicyInput{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		DefaultAction: defaultAction,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create prefilter policy",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcPrefilterPolicyRead(ctx, d, m)
}

func resourceFmcPrefilterPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcPrefilterPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read prefilter policy",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read prefilter policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read prefilter policy",
			Detail:   err.Error(),
		})
		return diags
	}

	defaultActionsList := []interface{}{
		map[string]interface{}{
			//"log_end": item.DefaultAction.LogEnd,
			"log_begin":          item.DefaultAction.LogBegin,
			"send_events_to_fmc": item.DefaultAction.SendEventsToFMC,
			"action":             item.DefaultAction.Action,
			"id":                 item.DefaultAction.ID,
		},
	}

	if err := d.Set("default_action", defaultActionsList); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceFmcPrefilterPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name", "description", "default_action") {
		var defaultAction = PrefilterPolicyDefaultAction{}
		if inputEntries, ok := d.GetOk("default_action"); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})

			defaultAction = PrefilterPolicyDefaultAction{
				LogBegin: entry["log_begin"].(bool),
				//LogEnd:          entry["log_end"].(bool),
				SendEventsToFMC: entry["send_events_to_fmc"].(bool),
				Action:          entry["action"].(string),
				ID:              entry["id"].(string),
			}
		}

		res, err := c.UpdateFmcPrefilterPolicy(ctx, &PrefilterPolicy{
			ID:            d.Id(),
			Name:          d.Get("name").(string),
			Description:   d.Get("description").(string),
			DefaultAction: defaultAction,
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update prefilter policy",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcPrefilterPolicyRead(ctx, d, m)
}

func resourceFmcPrefilterPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcPrefilterPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete prefilter policy",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
