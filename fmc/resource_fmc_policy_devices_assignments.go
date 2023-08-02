package fmc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var policy_devices_assignments_type string = "PolicyAssignment"

func resourceFmcPolicyDevicesAssignments() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Policy Device Assignments in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_policy_devices_assignments\" \"policy_assignment\" {\n" +
			"    policy {\n" +
			"        id = fmc_access_policies.access_policy.id\n" +
			"        type = fmc_access_policies.access_policy.type\n" +
			"    }\n" +
			"    target_devices {\n" +
			"        id = data.fmc_devices.device.id\n" +
			"        type = data.fmc_devices.device.type\n" +
			"    }\n" +
			"}\n" +
			"```\n" +
			"**Note** You cannot delete a policy assignment, only reassign the devices to another policy. So, the delete operation on terraform does nothing, but the assignment is not deleted until you have manually moved the devices to another policy.",
		CreateContext: resourceFmcPolicyDevicesAssignmentsCreate,
		ReadContext:   resourceFmcPolicyDevicesAssignmentsRead,
		UpdateContext: resourceFmcPolicyDevicesAssignmentsUpdate,
		DeleteContext: resourceFmcPolicyDevicesAssignmentsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of this resource",
			},
			"policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
					},
				},
				Description: "Policy (ACP/NAT) for this resource",
			},
			"target_devices": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
					},
				},
				Description: "Target devices for this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFmcPolicyDevicesAssignmentsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var policy PolicyDevicesAssignmentSubConfig

	if inputObjs, ok := d.GetOk("policy"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		policy = PolicyDevicesAssignmentSubConfig{
			ID:   obj["id"].(string),
			Type: obj["type"].(string),
		}
	}

	var devices []PolicyDevicesAssignmentSubConfig

	if inputObjs, ok := d.GetOk("target_devices"); ok {
		for _, obj := range inputObjs.([]interface{}) {
			obji := obj.(map[string]interface{})
			devices = append(devices, PolicyDevicesAssignmentSubConfig{
				ID:   obji["id"].(string),
				Type: obji["type"].(string),
			})
		}
	}

	res, err := c.CreateFmcPolicyDevicesAssignment(ctx, &PolicyDevicesAssignment{
		Name:        d.Get("name").(string),
		Policy:      policy,
		Targets:     devices,
		Type:        policy_devices_assignments_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create policy devices assignment",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcPolicyDevicesAssignmentsRead(ctx, d, m)
}

func resourceFmcPolicyDevicesAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcPolicyDevicesAssignment(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read policy devices assignment",
				Detail:   err.Error(),
			})
			return diags
		}

	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read policy devices assignment",
			Detail:   err.Error(),
		})
		return diags
	}

	policy := make([]interface{}, 0, 1)
	policyObj := make(map[string]interface{})
	policyObj["id"] = item.Policy.ID
	policyObj["type"] = item.Policy.Type
	policy = append(policy, policyObj)

	if err := d.Set("policy", policy); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read policy devices assignment",
			Detail:   err.Error(),
		})
		return diags
	}

	devices := make([]interface{}, 0, len(item.Targets))
	for _, obj := range item.Targets {
		device := make(map[string]interface{})
		device["id"] = obj.ID
		device["type"] = obj.Type
		devices = append(devices, device)
	}

	if err := d.Set("target_devices", devices); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read policy devices assignment",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourceFmcPolicyDevicesAssignmentsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "policy", "target_devices") {
		var policy PolicyDevicesAssignmentSubConfig

		if inputObjs, ok := d.GetOk("policy"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			policy = PolicyDevicesAssignmentSubConfig{
				ID:   obj["id"].(string),
				Type: obj["type"].(string),
			}
		}

		var devices []PolicyDevicesAssignmentSubConfig

		if inputObjs, ok := d.GetOk("target_devices"); ok {
			for _, obj := range inputObjs.([]interface{}) {
				obji := obj.(map[string]interface{})
				devices = append(devices, PolicyDevicesAssignmentSubConfig{
					ID:   obji["id"].(string),
					Type: obji["type"].(string),
				})
			}
		}

		_, err := c.UpdateFmcPolicyDevicesAssignment(ctx, id, &PolicyDevicesAssignment{
			Name:        d.Get("name").(string),
			Policy:      policy,
			Type:        policy_devices_assignments_type,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to create policy devices assignment",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcPolicyDevicesAssignmentsRead(ctx, d, m)
}

func resourceFmcPolicyDevicesAssignmentsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*Client)
	var policy PolicyDevicesAssignmentSubConfig
	id := d.Id()
	if inputObjs, ok := d.GetOk("policy"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		policy = PolicyDevicesAssignmentSubConfig{
			ID:   obj["id"].(string),
			Type: obj["type"].(string),
		}
	}

	var devices []PolicyDevicesAssignmentSubConfig

	if inputObjs, ok := d.GetOk("target_devices"); ok {
		for _, obj := range inputObjs.([]interface{}) {
			obji := obj.(map[string]interface{})
			devices = append(devices, PolicyDevicesAssignmentSubConfig{
				ID:   obji["id"].(string),
				Type: obji["type"].(string),
			})
		}
	}

	_, err := c.UpdateFmcPolicyDevicesAssignment(ctx, id, &PolicyDevicesAssignment{
		Name:        d.Get("name").(string),
		Policy:      policy,
		Type:        policy_devices_assignments_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete policy devices assignment",
			Detail:   err.Error(),
		})
		return diags
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
