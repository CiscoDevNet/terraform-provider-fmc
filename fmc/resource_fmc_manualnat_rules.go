package fmc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var manualnat_rules_type string = "FTDManualNatRule"

func resourceFmcManualNatRules() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Manual NAT Rules in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_ftd_manualnat_rules\" \"new_rule\" {\n" +
			"    nat_policy = fmc_ftd_nat_policies.nat_policy.id\n" +
			"    description = \"Testing Manual NAT priv-pub\"\n" +
			"    nat_type = \"static\"\n" +
			"    source_interface {\n" +
			"        id = data.fmc_security_zones.inside.id\n" +
			"        type = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    destination_interface {\n" +
			"        id = data.fmc_security_zones.outside.id\n" +
			"        type = data.fmc_security_zones.outside.type\n" +
			"    }\n" +
			"    original_source {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    translated_destination {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    interface_in_original_destination = true\n" +
			"    interface_in_translated_source = true\n" +
			"    ipv6 = true\n" +
			"}\n" +
			"\n" +
			"resource \"fmc_ftd_manualnat_rules\" \"new_rule_after\" {\n" +
			"    nat_policy = fmc_ftd_nat_policies.nat_policy.id\n" +
			"    description = \"Testing Manual NAT priv-pub\"\n" +
			"    nat_type = \"static\"\n" +
			"    section = \"after_auto\"\n" +
			"    source_interface {\n" +
			"        id = data.fmc_security_zones.inside.id\n" +
			"        type = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    destination_interface {\n" +
			"        id = data.fmc_security_zones.outside.id\n" +
			"        type = data.fmc_security_zones.outside.type\n" +
			"    }\n" +
			"    original_source {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    translated_destination {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    interface_in_original_destination = true\n" +
			"    interface_in_translated_source = true\n" +
			"    ipv6 = true\n" +
			"}\n" +
			"\n" +
			"resource \"fmc_ftd_manualnat_rules\" \"new_rule_before_1\" {\n" +
			"    nat_policy = fmc_ftd_nat_policies.nat_policy.id\n" +
			"    description = \"Testing Manual NAT before priv-pub\"\n" +
			"    nat_type = \"static\"\n" +
			"    section = \"before_auto\"\n" +
			"    target_index = 1\n" +
			"    source_interface {\n" +
			"        id = data.fmc_security_zones.inside.id\n" +
			"        type = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    destination_interface {\n" +
			"        id = data.fmc_security_zones.outside.id\n" +
			"        type = data.fmc_security_zones.outside.type\n" +
			"    }\n" +
			"    original_source {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    translated_destination {\n" +
			"        id = data.fmc_host_objects.CUCMPub.id\n" +
			"        type = data.fmc_host_objects.CUCMPub.type\n" +
			"    }\n" +
			"    interface_in_original_destination = true\n" +
			"    interface_in_translated_source = true\n" +
			"    ipv6 = true\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcManualNatRulesCreate,
		ReadContext:   resourceFmcManualNatRulesRead,
		UpdateContext: resourceFmcManualNatRulesUpdate,
		DeleteContext: resourceFmcManualNatRulesDelete,
		Schema: map[string]*schema.Schema{
			"nat_policy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the NAT policy this resource belongs to",
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
					allowedValues := []string{"after_auto", "before_auto"}
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
				Description: `Section, "after_auto" or "before_auto"`,
			},
			"target_index": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v, err := strconv.Atoi(val.(string))
					if err == nil && v > 0 {
						return
					}
					errs = append(errs, fmt.Errorf("%q must be greater than 0, got: %q and error: %s", key, v, err))
					return
				},
				Description: "Target index to place this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
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
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable this resource",
			},
			"nat_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					allowedValues := []string{"STATIC", "DYNAMIC"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
				Description: `The type of this resource, "static" or "dynamic"`,
			},
			"source_interface": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Source interface for this resource",
			},
			"destination_interface": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Destination interface for this resource",
			},
			"interface_in_original_destination": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Interface is the original destination",
			},
			"original_destination": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Original destination for this resource",
			},
			"original_destination_port": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Original destination port for this resource",
			},
			"original_source": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Original source for this resource",
			},
			"original_source_port": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Original source port for this resource",
			},
			"translated_destination": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Translated destination for this resource",
			},
			"translated_destination_port": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Translated destination port for this resource",
			},
			"interface_in_translated_source": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Interface is the translated source",
			},
			"translated_source": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Translated source for this resource",
			},
			"translated_source_port": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "Translated source port for this resource",
			},
			"unidirectional": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Make this resource unidirectional",
			},
			"fallthrough": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable fallthrough",
			},
			"translate_dns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable translate DNS",
			},
			"no_proxy_arp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable proxy ARP",
			},
			"perform_route_lookup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable perform route lookup",
			},
			"net_to_net": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable Net to Net",
			},
			"ipv6": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable IPv6",
			},
			"pat_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pat_pool_address": {
							Type:     schema.TypeList,
							Optional: true,
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
						},
						"interface_pat": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"include_reserve_ports": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"extended_pat_table": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"round_robin": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				Description: "PAT Options for this resource",
			},
		},
	}
}

func resourceFmcManualNatRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var sourceInterface, destinationInterface, originalDestination, originalDestinationPort, originalSource, originalSourcePort, translatedDestination, translatedDestinationPort, translatedSource, translatedSourcePort *ManualNatRuleSubConfig
	dynamicObjects := []**ManualNatRuleSubConfig{
		&sourceInterface, &destinationInterface, &originalDestination, &originalDestinationPort, &originalSource, &originalSourcePort, &translatedDestination, &translatedDestinationPort, &translatedSource, &translatedSourcePort,
	}
	for i, objType := range []string{"source_interface", "destination_interface", "original_destination", "original_destination_port", "original_source", "original_source_port", "translated_destination", "translated_destination_port", "translated_source", "translated_source_port"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects[i] = &ManualNatRuleSubConfig{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}
	}
	var patOptions *ManualNatRulePatOptions
	if pat_o := d.Get("pat_options").([]interface{}); len(pat_o) > 0 {
		pat_options := pat_o[0].(map[string]interface{})
		pat_pool_options := pat_options["pat_pool_address"].([]interface{})[0].(map[string]interface{})
		pat_pool := ManualNatRuleSubConfig{
			ID:   pat_pool_options["id"].(string),
			Type: pat_pool_options["type"].(string),
		}
		patOptions = &ManualNatRulePatOptions{
			Patpooladdress: pat_pool,
			Includereserve: pat_options["include_reserve_ports"].(bool),
			Interfacepat:   pat_options["interface_pat"].(bool),
			Extendedpat:    pat_options["extended_pat_table"].(bool),
			Roundrobin:     pat_options["round_robin"].(bool),
		}
	}
	res, err := c.CreateFmcManualNatRule(ctx, d.Get("nat_policy").(string), strings.ToLower(d.Get("section").(string)), d.Get("target_index").(string), &ManualNatRule{
		Type:                           manualnat_rules_type,
		Description:                    d.Get("description").(string),
		Enabled:                        d.Get("enabled").(bool),
		Nattype:                        strings.ToUpper(d.Get("nat_type").(string)),
		Sourceinterface:                sourceInterface,
		Destinationinterface:           destinationInterface,
		Interfaceinoriginaldestination: d.Get("interface_in_original_destination").(bool),
		Originaldestination:            originalDestination,
		Originaldestinationport:        originalDestinationPort,
		Originalsource:                 originalSource,
		Originalsourceport:             originalSourcePort,
		Translateddestination:          translatedDestination,
		Translateddestinationport:      translatedDestinationPort,
		Interfaceintranslatedsource:    d.Get("interface_in_translated_source").(bool),
		Translatedsource:               translatedSource,
		Translatedsourceport:           translatedSourcePort,
		Unidirectional:                 d.Get("unidirectional").(bool),
		Fallthrough:                    d.Get("fallthrough").(bool),
		DNS:                            d.Get("translate_dns").(bool),
		Noproxyarp:                     d.Get("no_proxy_arp").(bool),
		Routelookup:                    d.Get("perform_route_lookup").(bool),
		Nettonet:                       d.Get("net_to_net").(bool),
		Interfaceipv6:                  d.Get("ipv6").(bool),
		Patoptions:                     patOptions,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcManualNatRulesRead(ctx, d, m)
}

func resourceFmcManualNatRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcManualNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
		} else {
			return returnWithDiag(diags, err)
		}
	}
	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("description", item.Description); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("nat_type", item.Nattype); err != nil {
		return returnWithDiag(diags, err)
	}

	dynamicObjects := []*ManualNatRuleSubConfig{
		&item.Sourceinterface,
		&item.Destinationinterface,
		&item.Originaldestination,
		&item.Originaldestinationport,
		&item.Originalsource,
		&item.Originalsourceport,
		&item.Translateddestination,
		&item.Translateddestinationport,
		&item.Translatedsource,
		&item.Translatedsourceport,
	}

	dynamicObjectNames := []string{"source_interface", "destination_interface", "original_destination", "original_destination_port", "original_source", "original_source_port", "translated_destination", "translated_destination_port", "translated_source", "translated_source_port"}

	for i, objs := range dynamicObjects {
		if *objs != (ManualNatRuleSubConfig{}) {
			if err := d.Set(dynamicObjectNames[i], convertTo1ListMapStringGeneric(objs)); err != nil {
				return returnWithDiag(diags, err)
			}
		}
	}

	return diags
}

func resourceFmcManualNatRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("type", "description", "nat_type", "source_interface", "destination_interface", "original_destination", "original_destination_port", "original_source", "original_source_port", "translated_destination", "translated_destination_port", "translated_source", "translated_source_port", "unidirectional", "fallthrough", "translate_dns", "no_proxy_arp", "perform_route_lookup", "net_to_net", "ipv6", "pat_options") {
		var sourceInterface, destinationInterface, originalDestination, originalDestinationPort, originalSource, originalSourcePort, translatedDestination, translatedDestinationPort, translatedSource, translatedSourcePort *ManualNatRuleSubConfig
		dynamicObjects := []**ManualNatRuleSubConfig{
			&sourceInterface, &destinationInterface, &originalDestination, &originalDestinationPort, &originalSource, &originalSourcePort, &translatedDestination, &translatedDestinationPort, &translatedSource, &translatedSourcePort,
		}
		for i, objType := range []string{"source_interface", "destination_interface", "original_destination", "original_destination_port", "original_source", "original_source_port", "translated_destination", "translated_destination_port", "translated_source", "translated_source_port"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entry := inputEntries.([]interface{})[0].(map[string]interface{})
				*dynamicObjects[i] = &ManualNatRuleSubConfig{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
				}
			}
		}
		var patOptions *ManualNatRulePatOptions
		if pat_o := d.Get("pat_options").([]interface{}); len(pat_o) > 0 {
			pat_options := pat_o[0].(map[string]interface{})
			pat_pool_options := pat_options["pat_pool_address"].([]interface{})[0].(map[string]interface{})
			pat_pool := ManualNatRuleSubConfig{
				ID:   pat_pool_options["id"].(string),
				Type: pat_pool_options["type"].(string),
			}
			patOptions = &ManualNatRulePatOptions{
				Patpooladdress: pat_pool,
				Includereserve: pat_options["include_reserve_ports"].(bool),
				Interfacepat:   pat_options["interface_pat"].(bool),
				Extendedpat:    pat_options["extended_pat_table"].(bool),
				Roundrobin:     pat_options["round_robin"].(bool),
			}
		}
		res, err := c.UpdateFmcManualNatRule(ctx, d.Get("nat_policy").(string), d.Id(), &ManualNatRule{
			ID:                             d.Id(),
			Type:                           manualnat_rules_type,
			Description:                    d.Get("description").(string),
			Enabled:                        d.Get("enabled").(bool),
			Nattype:                        strings.ToUpper(d.Get("nat_type").(string)),
			Sourceinterface:                sourceInterface,
			Destinationinterface:           destinationInterface,
			Interfaceinoriginaldestination: d.Get("interface_in_original_destination").(bool),
			Originaldestination:            originalDestination,
			Originaldestinationport:        originalDestinationPort,
			Originalsource:                 originalSource,
			Originalsourceport:             originalSourcePort,
			Translateddestination:          translatedDestination,
			Translateddestinationport:      translatedDestinationPort,
			Interfaceintranslatedsource:    d.Get("interface_in_translated_source").(bool),
			Translatedsource:               translatedSource,
			Translatedsourceport:           translatedSourcePort,
			Unidirectional:                 d.Get("unidirectional").(bool),
			Fallthrough:                    d.Get("fallthrough").(bool),
			DNS:                            d.Get("translate_dns").(bool),
			Noproxyarp:                     d.Get("no_proxy_arp").(bool),
			Routelookup:                    d.Get("perform_route_lookup").(bool),
			Nettonet:                       d.Get("net_to_net").(bool),
			Interfaceipv6:                  d.Get("ipv6").(bool),
			Patoptions:                     patOptions,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceFmcManualNatRulesRead(ctx, d, m)
}

func resourceFmcManualNatRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcManualNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
