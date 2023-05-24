package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var access_policy_type string = "AccessPolicy"
var access_policy_default_action_type string = "AccessPolicyDefaultAction"
var access_policy_default_syslog_alert_type string = "SyslogAlert"

func resourceFmcAccessPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Access Control Policies in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_access_policies\" \"access_policy\" {\n" +
			"    name = \"Terraform Access Policy\"\n" +
			"    # default_action = \"block\" # Cannot have block with base IPS policy\n" +
			"    default_action = \"permit\"\n" +
			"    default_action_base_intrusion_policy_id = data.fmc_ips_policies.ips_policy.id\n" +
			"    default_action_send_events_to_fmc = \"true\"\n" +
			"    default_action_log_end = \"true\"\n" +
			"    default_action_syslog_config_id = data.fmc_syslog_alerts.syslog_alert.id\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcAccessPoliciesCreate,
		ReadContext:   resourceFmcAccessPoliciesRead,
		UpdateContext: resourceFmcAccessPoliciesUpdate,
		DeleteContext: resourceFmcAccessPoliciesDelete,
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
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"default_action_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of default action of this resource",
			},
			"default_action": {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"BLOCK", "TRUST", "PERMIT", "NETWORK_DISCOVERY", "INHERIT_FROM_PARENT"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
				Description: `Default action for this resource, "BLOCK", "TRUST", "PERMIT", "NETWORK_DISCOVERY" or "INHERIT_FROM_PARENT".`,
			},
			"default_action_base_intrusion_policy_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default action base policy ID to inherit from for this resource",
			},
			"default_action_send_events_to_fmc": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enable sending events to FMC for this resource, "true" or "false"`,
			},
			"default_action_log_begin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enable logging at the beginning of the connection for this resource, "true" or "false`,
			},
			"default_action_log_end": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enable logging at the end of the connection for this resource, "true" or "false"`,
			},
			"default_action_syslog_config_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Syslog configuration ID for this resource",
			},
			"default_action_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of default action of this resource",
			},
		},
	}
}

func resourceFmcAccessPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var intrusionPolicy, syslogConfig *AccessPolicySubConfig
	if val, ok := d.GetOk("default_action_base_intrusion_policy_id"); ok {
		intrusionPolicy = &AccessPolicySubConfig{
			ID:   val.(string),
			Type: access_policy_default_action_type,
		}
	}

	if val, ok := d.GetOk("default_action_syslog_config_id"); ok {
		syslogConfig = &AccessPolicySubConfig{
			ID:   val.(string),
			Type: access_policy_default_syslog_alert_type,
		}
	}

	res, err := c.CreateFmcAccessPolicy(ctx, &AccessPolicy{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Defaultaction: AccessPolicyDefaultAction{
			Type:            access_policy_default_action_type,
			Intrusionpolicy: intrusionPolicy,
			Syslogconfig:    syslogConfig,
			Logbegin:        d.Get("default_action_log_begin").(bool),
			Logend:          d.Get("default_action_log_end").(bool),
			Sendeventstofmc: d.Get("default_action_send_events_to_fmc").(bool),
			Action:          strings.ToUpper(d.Get("default_action").(string)),
		},
		Type: access_policy_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create access policy",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcAccessPoliciesRead(ctx, d, m)
}

func resourceFmcAccessPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcAccessPolicy(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read access policy",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("default_action", item.Defaultaction.Action); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("default_action_id", item.Defaultaction.ID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcAccessPoliciesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name", "description", "type", "default_action", "default_action_base_intrusion_policy_id", "default_action_send_events_to_fmc", "default_action_log_begin", "default_action_log_end", "default_action_syslog_config_id", "default_action_type") {
		var intrusionPolicy, syslogConfig *AccessPolicySubConfig
		if val, ok := d.GetOk("default_action_base_intrusion_policy_id"); ok {
			intrusionPolicy = &AccessPolicySubConfig{
				ID:   val.(string),
				Type: access_policy_default_action_type,
			}
		}

		if val, ok := d.GetOk("default_action_syslog_config_id"); ok {
			syslogConfig = &AccessPolicySubConfig{
				ID:   val.(string),
				Type: access_policy_default_syslog_alert_type,
			}
		}
		res, err := c.UpdateFmcAccessPolicy(ctx, d.Id(), &AccessPolicy{
			ID:          d.Id(),
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Defaultaction: AccessPolicyDefaultAction{
				ID:              d.Get("default_action_id").(string),
				Type:            access_policy_default_action_type,
				Intrusionpolicy: intrusionPolicy,
				Syslogconfig:    syslogConfig,
				Logbegin:        d.Get("default_action_log_begin").(bool),
				Logend:          d.Get("default_action_log_end").(bool),
				Sendeventstofmc: d.Get("default_action_send_events_to_fmc").(bool),
				Action:          strings.ToUpper(d.Get("default_action").(string)),
			},
			Type: access_policy_type,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to create access policy",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcAccessPoliciesRead(ctx, d, m)
}

func resourceFmcAccessPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcAccessPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete access policy",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
