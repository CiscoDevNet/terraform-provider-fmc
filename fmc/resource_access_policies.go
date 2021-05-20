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

func resourceAccessPolicies() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccessPoliciesCreate,
		ReadContext:   resourceAccessPoliciesRead,
		DeleteContext: resourceAccessPoliciesDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_action": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			},
			"default_action_base_intrusion_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_action_send_events_to_fmc": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_action_log_begin": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_action_log_end": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_action_syslog_config_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_action_type": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAccessPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateAccessPolicy(ctx, &AccessPolicy{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Defaultaction: AccessPolicyDefaultAction{
			Type: access_policy_default_action_type,
			Intrusionpolicy: AccessPolicyDefaultActionIntrusionPolicy{
				ID:   d.Get("default_action_base_intrusion_policy_id").(string),
				Type: access_policy_default_action_type,
			},
			Syslogconfig: AccessPolicyDefaultActionSyslogConfig{
				ID:   d.Get("default_action_syslog_config_id").(string),
				Type: access_policy_default_syslog_alert_type,
			},
			Logbegin:        d.Get("default_action_log_begin").(string),
			Logend:          d.Get("default_action_log_end").(string),
			Sendeventstofmc: d.Get("default_action_send_events_to_fmc").(string),
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
	return resourceAccessPoliciesRead(ctx, d, m)
}

func resourceAccessPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetAccessPolicy(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read access policy",
			Detail:   err.Error(),
		})
		return diags
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

	return diags
}

func resourceAccessPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteAccessPolicy(ctx, id)
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
