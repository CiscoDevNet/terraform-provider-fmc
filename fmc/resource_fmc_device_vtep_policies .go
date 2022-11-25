package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var vtep_policies_type string = "VTEPPolicy"

func resourceFmcDeviceVTEPPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Access Rules in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_access_rules\" \"access_rule_1\" {\n" +
			"    acp = fmc_access_policies.access_policy.id\n" +
			"    section = \"mandatory\"\n" +
			"    name = \"Test rule 1\"\n" +
			"    action = \"allow\"\n" +
			"    enabled = true\n" +
			"    enable_syslog = true\n" +
			"    syslog_severity = \"alert\"\n" +
			"    send_events_to_fmc = true\n" +
			"    log_files = false\n" +
			"    log_end = true\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcDeviceVTEPPoliciesCreate,
		ReadContext:   resourceFmcDeviceVTEPPoliciessRead,
		UpdateContext: resourceFmcDeviceVTEPPoliciessUpdate,
		DeleteContext: resourceFmcDeviceVTEPPoliciesDelete,
		Schema: map[string]*schema.Schema{
			"acp": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the ACP this resource belongs to",
			},
			"category": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Category of the ACP this resource belongs to. Should be created upfront with fmc_access_policies_category resource",
			},
			"section": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToLower(val.(string))
					allowedValues := []string{"mandatory", "default"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
				Description: `Section for this resource, "mandatory" or "default"`,
			},
			"insert_before": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v > 0 {
						return
					}
					errs = append(errs, fmt.Errorf("%q must be greater than 0, got: %q", key, v))
					return
				},
				Description: "The rule number before which to insert this resource",
			},
			"insert_after": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v > 0 {
						return
					}
					errs = append(errs, fmt.Errorf("%q must be greater than 0, got: %q", key, v))
					return
				},
				Description: "The rule number after which to insert this resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resourceFmc",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE", "BLOCK_RESET_INTERACTIVE"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
				Description: `Action for this resource, "ALLOW", "TRUST", "BLOCK", "MONITOR", "BLOCK_RESET", "BLOCK_INTERACTIVE" or "BLOCK_RESET_INTERACTIVE"`,
			},
			"syslog_severity": {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE", "WARNING"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
				Description: `Syslog severity for this resource, "ALERT", "CRIT", "DEBUG", "EMERG", "ERR", "INFO", "NOTICE" or "WARNING"`,
			},
			"enable_syslog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable syslog for this resource",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Enable the resourceFmc",
			},
			"send_events_to_fmc": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable sending events to FMC for this resource",
			},
			"log_files": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable logging files for this resource",
			},
			"log_begin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable logging at the beginning of connection for this resource",
			},
			"log_end": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable logging at the end of connection for this resource",
			},
			"source_zones": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_zone": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Source zones for this resource",
			},
			"destination_zones": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_zone": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Destination zones for this resource",
			},
			"source_networks": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_network": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Source networks for this resource",
			},
			"destination_networks": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_network": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Destination networks for this resource",
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_port": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Source ports for this resource",
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_port": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "Destination ports for this resource",
			},
			"urls": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeList,
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
						},
					},
				},
				Description: "URLs for this resource",
			},
			"ips_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPS policy for this resource",
			},
			"file_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File policy for this resource",
			},
			"syslog_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Syslog configuration ID for this resource",
			},
			"new_comments": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "New comments to be added for this resource",
			},
		},
	}
}

func resourceFmcDeviceVTEPPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	var vtepEntry []VTEPEntry

	var sourceInterface SourceInterface

	if inputObjs, ok := d.GetOk("sourceInterface"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		sourceInterface = SourceInterface{
			ID:   obj["id"].(string),
			Name: obj["name"].(string),
			Type: obj["type"].(string),
		}
	}

	if inputObjs, ok := d.GetOk("vtepEntries"); ok {
		obj := inputObjs.([]interface{})[0].(map[string]interface{})
		vtepEntry = append(vtepEntry, VTEPEntry{
			SourceInterface:          sourceInterface,
			NveVtepId:                obj["nveVtepId"].(int),
			NveDestinationPort:       obj["nveDestinationPort"].(int),
			NveEncapsulationType:     obj["NveEncapsulationType"].(string),
			NveNeighborDiscoveryType: obj["NveNeighborDiscoveryType"].(string),
		})
	}

	res, err := c.CreateVTEPPolicies(ctx, d.Get("acp").(string), &VTEPPolicy{
		Type:        vtep_policies_type,
		NveEnable:   d.Get("nveEnable").(bool),
		VTEPEntries: vtepEntry,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcDeviceVTEPPoliciessRead(ctx, d, m)
}

func resourceFmcDeviceVTEPPoliciessRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetVTEPPoliciesByName(ctx, d.Get("acp").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceFmcDeviceVTEPPoliciessUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("vtepEntries") {
		var vtepEntry []VTEPEntry

		var sourceInterface SourceInterface

		if inputObjs, ok := d.GetOk("sourceInterface"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			sourceInterface = SourceInterface{
				ID:   obj["id"].(string),
				Name: obj["name"].(string),
				Type: obj["type"].(string),
			}
		}

		if inputObjs, ok := d.GetOk("vtepEntries"); ok {
			obj := inputObjs.([]interface{})[0].(map[string]interface{})
			vtepEntry = append(vtepEntry, VTEPEntry{
				SourceInterface:          sourceInterface,
				NveVtepId:                obj["nveVtepId"].(int),
				NveDestinationPort:       obj["nveDestinationPort"].(int),
				NveEncapsulationType:     obj["NveEncapsulationType"].(string),
				NveNeighborDiscoveryType: obj["NveNeighborDiscoveryType"].(string),
			})
		}

		res, err := c.UpdateVTEPPolicies(ctx, d.Id(), &VTEPPolicy{
			Type:        vtep_policies_type,
			NveEnable:   d.Get("nveEnable").(bool),
			VTEPEntries: vtepEntry,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceFmcDeviceVTEPPoliciessRead(ctx, d, m)
}

func resourceFmcDeviceVTEPPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteVTEPPolicies(ctx, d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
