package fmc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var access_policies_type string = "AccessRule"

func resourceFmcAccessRules() *schema.Resource {
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
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"    destination_zones {\n" +
			"        destination_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"    source_networks {\n" +
			"        source_network {\n" +
			"            id = data.fmc_network_objects.source.id\n" +
			"            type =  data.fmc_network_objects.source.type\n" +
			"        }\n" +
			"    }\n" +
			"    destination_networks {\n" +
			"        destination_network {\n" +
			"            id = data.fmc_network_objects.dest.id\n" +
			"            type =  data.fmc_network_objects.dest.type\n" +
			"        }\n" +
			"    }\n" +
			"    destination_ports {\n" +
			"        destination_port {\n" +
			"            id = data.fmc_port_objects.http.id\n" +
			"            type =  data.fmc_port_objects.http.type\n" +
			"        }\n" +
			"    }\n" +
			"    urls {\n" +
			"        url {\n" +
			"            id = fmc_url_objects.dest_url.id\n" +
			"            type = \"Url\"\n" +
			"        }\n" +
			"    }\n" +
			"    ips_policy = data.fmc_ips_policies.ips_policy.id\n" +
			"    syslog_config = data.fmc_syslog_alerts.syslog_alert.id\n" +
			"    new_comments = [ \"New\", \"comment\" ]\n" +
			"}\n" +
			"\n" +
			"resource \"fmc_access_rules\" \"access_rule_2\" {\n" +
			"    acp = fmc_access_policies.access_policy.id\n" +
			"    section = \"mandatory\"\n" +
			"    insert_before = 1 # Wont work as assumed since terraform does not \n" +
			"    name = \"Test rule 2\"\n" +
			"    action = \"allow\"\n" +
			"    enabled = true\n" +
			"    enable_syslog = true\n" +
			"    syslog_severity = \"alert\"\n" +
			"    send_events_to_fmc = true\n" +
			"    log_files = false\n" +
			"    log_end = true\n" +
			"    source_zones {\n" +
			"        source_zone {\n" +
			"            id = data.fmc_security_zones.inside.id\n" +
			"            type =  data.fmc_security_zones.inside.type\n" +
			"        }\n" +
			"        # source_zone {\n" +
			"        #     id = data.fmc_security_zones.outside.id\n" +
			"        #     type =  data.fmc_security_zones.outside.type\n" +
			"        # }\n" +
			"    }\n" +
			"    destination_zones {\n" +
			"        destination_zone {\n" +
			"            id = data.fmc_security_zones.outside.id\n" +
			"            type =  data.fmc_security_zones.outside.type\n" +
			"        }\n" +
			"    }\n" +
			"    source_networks {\n" +
			"        source_network {\n" +
			"            id = data.fmc_network_objects.source.id\n" +
			"            type =  data.fmc_network_objects.source.type\n" +
			"        }\n" +
			"    }\n" +
			"    destination_networks {\n" +
			"        destination_network {\n" +
			"            id = data.fmc_network_objects.dest.id\n" +
			"            type =  data.fmc_network_objects.dest.type\n" +
			"        }\n" +
			"    }\n" +
			"    destination_ports {\n" +
			"        destination_port {\n" +
			"            id = data.fmc_port_objects.http.id\n" +
			"            type =  data.fmc_port_objects.http.type\n" +
			"        }\n" +
			"    }\n" +
			"    urls {\n" +
			"        url {\n" +
			"            id = fmc_url_objects.dest_url.id\n" +
			"            type = \"Url\"\n" +
			"        }\n" +
			"    }\n" +
			"    ips_policy = data.fmc_ips_policies.ips_policy.id\n" +
			"    syslog_config = data.fmc_syslog_alerts.syslog_alert.id\n" +
			"    new_comments = [ \"New comment\" ]\n" +
			"    depends_on = [\n" +
			"        fmc_access_rules.access_rule_1\n" +
			"    ]\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcAccessRulesCreate,
		ReadContext:   resourceFmcAccessRulesRead,
		UpdateContext: resourceFmcAccessRulesUpdate,
		DeleteContext: resourceFmcAccessRulesDelete,
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
				Required: true,
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
			"destination_dynamic_objects": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_dynamic_object": {
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
				Description: "Destination dynamic objects ports for this resource",
			},
			"source_dynamic_objects": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_dynamic_object": {
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
				Description: "Source dynamic objects for this resource",
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
			"source_security_group_tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_security_group_tag": {
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
				Description: "Source SGTs",
			},
			"destination_security_group_tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_security_group_tag": {
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
				Description: "Destination SGTs",
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFmcAccessRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	var sourceZones, destinationZones, sourceNetworks, destinationNetworks, sourcePorts, destinationPorts, destinationDynamicObjects, sourceDynamicObjects, urls, sourceSecurityGroupTag, destinationSecurityGroupTag []AccessRuleSubConfig
	dynamicObjects := []*[]AccessRuleSubConfig{
		&sourceZones, &destinationZones, &sourceNetworks, &destinationNetworks, &sourcePorts, &destinationPorts, &destinationDynamicObjects, &sourceDynamicObjects, &urls, &sourceSecurityGroupTag, &destinationSecurityGroupTag,
	}
	for i, objType := range []string{"source_zones", "destination_zones", "source_networks", "destination_networks", "source_ports", "destination_ports", "destination_dynamic_objects", "source_dynamic_objects", "urls", "source_security_group_tags", "destination_security_group_tags"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entries := inputEntries.([]interface{})[0].(map[string]interface{})[objType[:len(objType)-1]]
			for _, ent := range entries.([]interface{}) {
				entry := ent.(map[string]interface{})
				*dynamicObjects[i] = append(*dynamicObjects[i], AccessRuleSubConfig{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
				})
			}
		}
	}

	var ipsPolicy, filePolicy, syslogConfig *AccessRuleSubConfig
	dynamicSimpleObjects := []**AccessRuleSubConfig{
		&ipsPolicy, &filePolicy, &syslogConfig,
	}
	for i, objType := range []string{"ips_policy", "file_policy", "syslog_config"} {
		if inputEntry, ok := d.GetOk(objType); ok {
			*dynamicSimpleObjects[i] = &AccessRuleSubConfig{
				ID: inputEntry.(string),
			}
		}
	}

	comments := []string{}
	for _, comment := range d.Get("new_comments").([]interface{}) {
		comments = append(comments, comment.(string))
	}
	var insertBefore, insertAfter string
	if entry, ok := d.GetOk("insert_before"); ok {
		insertBefore = strconv.Itoa(entry.(int))
	}
	if entry, ok := d.GetOk("insert_after"); ok {
		insertAfter = strconv.Itoa(entry.(int))
	}

	res, err := c.CreateFmcAccessRule(ctx, d.Get("acp").(string), strings.ToLower(d.Get("section").(string)), insertBefore, insertAfter, d.Get("category").(string), &AccessRule{
		Name:            d.Get("name").(string),
		Type:            access_policies_type,
		Action:          strings.ToUpper(d.Get("action").(string)),
		Syslogseverity:  strings.ToUpper(d.Get("syslog_severity").(string)),
		Enablesyslog:    d.Get("enable_syslog").(bool),
		Enabled:         d.Get("enabled").(bool),
		Sendeventstofmc: d.Get("send_events_to_fmc").(bool),
		Logfiles:        d.Get("log_files").(bool),
		Logbegin:        d.Get("log_begin").(bool),
		Logend:          d.Get("log_end").(bool),
		Sourcezones: AccessRuleSubConfigs{
			Objects: sourceZones,
		},
		Destinationzones: AccessRuleSubConfigs{
			Objects: destinationZones,
		},
		Sourcenetworks: AccessRuleSubConfigs{
			Objects: sourceNetworks,
		},
		Destinationnetworks: AccessRuleSubConfigs{
			Objects: destinationNetworks,
		},
		Sourceports: AccessRuleSubConfigs{
			Objects: sourcePorts,
		},
		Destinationports: AccessRuleSubConfigs{
			Objects: destinationPorts,
		},
		SourceDynamicObjects: AccessRuleSubConfigs{
			Objects: sourceDynamicObjects,
		},
		DestinationDynamicObjects: AccessRuleSubConfigs{
			Objects: destinationDynamicObjects,
		},
		Urls: AccessRuleSubConfigs{
			Objects: urls,
		},
		SourceSecurityGroupTags: AccessRuleSubConfigs{
			Objects: sourceSecurityGroupTag,
		},
		DestinationSecurityGroupTags: AccessRuleSubConfigs{
			Objects: destinationSecurityGroupTag,
		},
		Ipspolicy:    ipsPolicy,
		Filepolicy:   filePolicy,
		Syslogconfig: syslogConfig,
		Newcomments:  comments,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcAccessRulesRead(ctx, d, m)
}

func resourceFmcAccessRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcAccessRule(ctx, d.Get("acp").(string), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			return returnWithDiag(diags, err)
		}
	}
	if err := d.Set("name", item.Name); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("action", item.Action); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("syslog_severity", item.Syslogseverity); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("enable_syslog", item.Enablesyslog); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("enabled", item.Enabled); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("send_events_to_fmc", item.Sendeventstofmc); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("log_files", item.Logfiles); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("log_begin", item.Logbegin); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("log_end", item.Logend); err != nil {
		return returnWithDiag(diags, err)
	}
	// seems that category is not returned within API response, so that's the only way
	if err := d.Set("category", d.Get("category").(string)); err != nil {
		return returnWithDiag(diags, err)
	}

	dynamicObjects := []*[]AccessRuleResponseObject{
		&item.Sourcezones.Objects,
		&item.Destinationzones.Objects,
		&item.Sourcenetworks.Objects,
		&item.Destinationnetworks.Objects,
		&item.Sourceports.Objects,
		&item.Destinationports.Objects,
		&item.SourceDynamicObjects.Objects,
		&item.DestinationDynamicObjects.Objects,
		&item.SourceSecurityGroupTags.Objects,
		&item.DestinationSecurityGroupTags.Objects,
		&item.Urls.Objects,
	}

	dynamicObjectNames := []string{"source_zones", "destination_zones", "source_networks", "destination_networks", "source_ports", "destination_ports", "source_dynamic_objects", "destination_dynamic_objects","source_security_group_tags","destination_security_group_tags", "urls"}

	for i, objs := range dynamicObjects {
		mainResponse := make([]map[string]interface{}, 0)
		subResponse := make(map[string]interface{}, len(*objs))
		response := make([]map[string]interface{}, 0, len(*objs))
		for _, obj := range *objs {
			responseObj := make(map[string]interface{})
			responseObj["id"] = obj.ID
			responseObj["type"] = obj.Type
			response = append(response, responseObj)
		}
		if len(response) != 0 {
			subResponse[dynamicObjectNames[i][:len(dynamicObjectNames[i])-1]] = response
			mainResponse = append(mainResponse, subResponse)
		}
		if err := d.Set(dynamicObjectNames[i], mainResponse); err != nil {
			return returnWithDiag(diags, err)
		}
	}

	dynamicSimpleObjects := []*AccessRuleResponseObject{
		&item.Ipspolicy, &item.Filepolicy, &item.Syslogconfig,
	}
	for i, objType := range []string{"ips_policy", "file_policy", "syslog_config"} {
		id := &dynamicSimpleObjects[i].ID
		if *id == "" {
			id = nil
		}
		if err := d.Set(objType, id); err != nil {
			return returnWithDiag(diags, err)
		}
	}

	return diags
}

func resourceFmcAccessRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name", "type", "action", "syslog_severity", "enable_syslog", "enabled", "send_events_to_fmc", "log_files", "log_begin", "log_end", "source_zones", "destination_zones", "source_networks", "destination_networks", "source_ports", "destination_ports", "destination_dynamic_objects", "source_dynamic_objects", "urls", "source_security_group_tags", "destination_security_group_tags", "ips_policy", "file_policy", "syslog_config", "new_comments") {
		var sourceZones, destinationZones, sourceNetworks, destinationNetworks, sourcePorts, destinationPorts, destinationDynamicObjects, sourceDynamicObjects, urls, sourceSecurityGroupTag, destinationSecurityGroupTag []AccessRuleSubConfig
		dynamicObjects := []*[]AccessRuleSubConfig{
			&sourceZones, &destinationZones, &sourceNetworks, &destinationNetworks, &sourcePorts, &destinationPorts, &destinationDynamicObjects, &sourceDynamicObjects, &urls, &sourceSecurityGroupTag, &destinationSecurityGroupTag,
		}
		for i, objType := range []string{"source_zones", "destination_zones", "source_networks", "destination_networks", "source_ports", "destination_ports", "destination_dynamic_objects", "source_dynamic_objects", "urls", "source_security_group_tags", "destination_security_group_tags"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entries := inputEntries.([]interface{})[0].(map[string]interface{})[objType[:len(objType)-1]]
				for _, ent := range entries.([]interface{}) {
					entry := ent.(map[string]interface{})
					*dynamicObjects[i] = append(*dynamicObjects[i], AccessRuleSubConfig{
						ID:   entry["id"].(string),
						Type: entry["type"].(string),
					})
				}
			}
		}

		var ipsPolicy, filePolicy, syslogConfig *AccessRuleSubConfig
		dynamicSimpleObjects := []**AccessRuleSubConfig{
			&ipsPolicy, &filePolicy, &syslogConfig,
		}
		for i, objType := range []string{"ips_policy", "file_policy", "syslog_config"} {
			if inputEntry, ok := d.GetOk(objType); ok {
				*dynamicSimpleObjects[i] = &AccessRuleSubConfig{
					ID: inputEntry.(string),
				}
			}
		}

		comments := []string{}
		for _, comment := range d.Get("new_comments").([]interface{}) {
			comments = append(comments, comment.(string))
		}
		res, err := c.UpdateFmcAccessRule(ctx, d.Get("acp").(string), d.Id(), &AccessRuleUpdate{
			ID:              d.Id(),
			Name:            d.Get("name").(string),
			Type:            access_policies_type,
			Action:          strings.ToUpper(d.Get("action").(string)),
			Syslogseverity:  strings.ToUpper(d.Get("syslog_severity").(string)),
			Enablesyslog:    d.Get("enable_syslog").(bool),
			Enabled:         d.Get("enabled").(bool),
			Sendeventstofmc: d.Get("send_events_to_fmc").(bool),
			Logfiles:        d.Get("log_files").(bool),
			Logbegin:        d.Get("log_begin").(bool),
			Logend:          d.Get("log_end").(bool),
			Sourcezones: AccessRuleSubConfigs{
				Objects: sourceZones,
			},
			Destinationzones: AccessRuleSubConfigs{
				Objects: destinationZones,
			},
			Sourcenetworks: AccessRuleSubConfigs{
				Objects: sourceNetworks,
			},
			Destinationnetworks: AccessRuleSubConfigs{
				Objects: destinationNetworks,
			},
			Sourceports: AccessRuleSubConfigs{
				Objects: sourcePorts,
			},
			Destinationports: AccessRuleSubConfigs{
				Objects: destinationPorts,
			},
			SourceDynamicObjects: AccessRuleSubConfigs{
				Objects: sourceDynamicObjects,
			},
			DestinationDynamicObjects: AccessRuleSubConfigs{
				Objects: destinationDynamicObjects,
			},
			Urls: AccessRuleSubConfigs{
				Objects: urls,
			},
			SourceSecurityGroupTags: AccessRuleSubConfigs{
				Objects: sourceSecurityGroupTag,
			},
			DestinationSecurityGroupTags: AccessRuleSubConfigs{
				Objects: destinationSecurityGroupTag,
			},
			Ipspolicy:    ipsPolicy,
			Filepolicy:   filePolicy,
			Syslogconfig: syslogConfig,
			Newcomments:  comments,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceFmcAccessRulesRead(ctx, d, m)
}

func resourceFmcAccessRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcAccessRule(ctx, d.Get("acp").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
