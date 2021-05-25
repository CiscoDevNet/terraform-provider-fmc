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

func resourceManualNatRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceManualNatRulesCreate,
		ReadContext:   resourceManualNatRulesRead,
		UpdateContext: resourceManualNatRulesUpdate,
		DeleteContext: resourceManualNatRulesDelete,
		Schema: map[string]*schema.Schema{
			"nat_policy": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			},
			"source_interface": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"destination_interface": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"interface_in_original_destination": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"original_destination": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"original_destination_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"original_source": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"original_source_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"translated_destination": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"translated_destination_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"interface_in_translated_source": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"translated_source": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"translated_source_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"unidirectional": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fallthrough": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"translate_dns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"no_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"perform_route_lookup": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"net_to_net": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
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
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
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
			},
		},
	}
}

func resourceManualNatRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	res, err := c.CreateManualNatRule(ctx, d.Get("nat_policy").(string), strings.ToLower(d.Get("section").(string)), d.Get("target_index").(string), &ManualNatRule{
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
	return resourceManualNatRulesRead(ctx, d, m)
}

func resourceManualNatRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetManualNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
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
		if err := d.Set(dynamicObjectNames[i], convertTo1ListMapStringGeneric(objs)); err != nil {
			return returnWithDiag(diags, err)
		}
	}

	return diags
}

func resourceManualNatRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		res, err := c.UpdateManualNatRule(ctx, d.Get("nat_policy").(string), d.Id(), &ManualNatRule{
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
	return resourceManualNatRulesRead(ctx, d, m)
}

func resourceManualNatRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteManualNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
