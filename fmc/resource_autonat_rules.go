package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var autonat_rules_type string = "AutoNatRule"

func resourceAutoNatRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutoNatRulesCreate,
		ReadContext:   resourceAutoNatRulesRead,
		UpdateContext: resourceAutoNatRulesUpdate,
		DeleteContext: resourceAutoNatRulesDelete,
		Schema: map[string]*schema.Schema{
			"nat_policy": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
			"original_network": {
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
			"translated_network": {
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
			"translated_network_is_destination_interface": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"original_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(int)
								if v > 0 && v < 65536 {
									return
								}
								errs = append(errs, fmt.Errorf("%q must be in 1-65535, got: %q", key, v))
								return
							},
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := strings.ToUpper(val.(string))
								allowedValues := []string{"TCP", "UDP"}
								for _, allowed := range allowedValues {
									if v == allowed {
										return
									}
								}
								errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
								return
							},
						},
					},
				},
			},
			"translated_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v > 0 && v < 65536 {
						return
					}
					errs = append(errs, fmt.Errorf("%q must be in 1-65535, got: %q", key, v))
					return
				},
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

func resourceAutoNatRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var sourceInterface, destinationInterface, originalNetwork, translatedNetwork *AutoNatRuleSubConfig
	dynamicObjects := []**AutoNatRuleSubConfig{
		&sourceInterface, &destinationInterface, &originalNetwork, &translatedNetwork,
	}
	for i, objType := range []string{"source_interface", "destination_interface", "original_network", "translated_network"} {
		if inputEntries, ok := d.GetOk(objType); ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects[i] = &AutoNatRuleSubConfig{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}
	}
	var original_port map[string]interface{}
	var patOptions *AutoNatRulePatOptions
	if or_port := d.Get("original_port").([]interface{}); len(or_port) > 0 {
		original_port = or_port[0].(map[string]interface{})
	} else {
		original_port = make(map[string]interface{})
		original_port["port"] = 0
		original_port["protocol"] = ""
	}
	if pat_o := d.Get("pat_options").([]interface{}); len(pat_o) > 0 {
		pat_options := pat_o[0].(map[string]interface{})
		pat_pool_options := pat_options["pat_pool_address"].([]interface{})[0].(map[string]interface{})
		pat_pool := AutoNatRuleSubConfig{
			ID:   pat_pool_options["id"].(string),
			Type: pat_pool_options["type"].(string),
		}
		patOptions = &AutoNatRulePatOptions{
			Patpooladdress: pat_pool,
			Includereserve: pat_options["include_reserve_ports"].(bool),
			Interfacepat:   pat_options["interface_pat"].(bool),
			Extendedpat:    pat_options["extended_pat_table"].(bool),
			Roundrobin:     pat_options["round_robin"].(bool),
		}
	}

	res, err := c.CreateAutoNatRule(ctx, d.Get("nat_policy").(string), &AutoNatRule{
		Type:                         autonat_rules_type,
		Description:                  d.Get("description").(string),
		Nattype:                      strings.ToUpper(d.Get("nat_type").(string)),
		Sourceinterface:              sourceInterface,
		Destinationinterface:         destinationInterface,
		Originalnetwork:              originalNetwork,
		Translatednetwork:            translatedNetwork,
		Interfaceintranslatednetwork: d.Get("translated_network_is_destination_interface").(bool),
		Originalport:                 original_port["port"].(int),
		Serviceprotocol:              strings.ToUpper(original_port["protocol"].(string)),
		Translatedport:               d.Get("translated_port").(int),
		Fallthrough:                  d.Get("fallthrough").(bool),
		DNS:                          d.Get("translate_dns").(bool),
		Noproxyarp:                   d.Get("no_proxy_arp").(bool),
		Routelookup:                  d.Get("perform_route_lookup").(bool),
		Nettonet:                     d.Get("net_to_net").(bool),
		Interfaceipv6:                d.Get("ipv6").(bool),
		Patoptions:                   patOptions,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceAutoNatRulesRead(ctx, d, m)
}

func resourceAutoNatRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id())
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

	if err := d.Set("source_interface", convertTo1ListMapStringGeneric(item.Sourceinterface)); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("destination_interface", convertTo1ListMapStringGeneric(item.Destinationinterface)); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("original_network", convertTo1ListMapStringGeneric(item.Originalnetwork)); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("translated_network", convertTo1ListMapStringGeneric(item.Translatednetwork)); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("translated_network_is_destination_interface", item.Interfaceintranslatednetwork); err != nil {
		return returnWithDiag(diags, err)
	}

	original_port := make(map[string]interface{})
	original_port["port"] = item.Originalport
	original_port["protocol"] = item.Serviceprotocol
	if err := d.Set("original_port", convertTo1ListGeneric(original_port)); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("translated_port", item.Translatedport); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("fallthrough", item.Fallthrough); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("translate_dns", item.DNS); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("no_proxy_arp", item.Noproxyarp); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("perform_route_lookup", item.Nettonet); err != nil {
		return returnWithDiag(diags, err)
	}
	if err := d.Set("ipv6", item.Interfaceipv6); err != nil {
		return returnWithDiag(diags, err)
	}

	pat_options := make(map[string]interface{})
	pat_options["pat_pool_address"] = convertTo1ListMapStringGeneric(item.Patoptions.Patpooladdress)
	pat_options["interface_pat"] = item.Patoptions.Interfacepat
	pat_options["include_reserve_ports"] = item.Patoptions.Interfacepat
	pat_options["extended_pat_table"] = item.Patoptions.Interfacepat
	pat_options["round_robin"] = item.Patoptions.Interfacepat
	if err := d.Set("pat_options", convertTo1ListGeneric(pat_options)); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceAutoNatRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("type", "description", "nat_type", "source_interface", "destination_interface", "original_network", "translated_network", "translated_network_is_destination_interface", "original_port", "translated_port", "fallthrough", "translate_dns", "no_proxy_arp", "perform_route_lookup", "net_to_net", "ipv6", "pat_options") {
		var sourceInterface, destinationInterface, originalNetwork, translatedNetwork *AutoNatRuleSubConfig
		dynamicObjects := []**AutoNatRuleSubConfig{
			&sourceInterface, &destinationInterface, &originalNetwork, &translatedNetwork,
		}
		for i, objType := range []string{"source_interface", "destination_interface", "original_network", "translated_network"} {
			if inputEntries, ok := d.GetOk(objType); ok {
				entry := inputEntries.([]interface{})[0].(map[string]interface{})
				*dynamicObjects[i] = &AutoNatRuleSubConfig{
					ID:   entry["id"].(string),
					Type: entry["type"].(string),
				}
			}
		}

		var original_port map[string]interface{}
		var patOptions *AutoNatRulePatOptions
		if or_port := d.Get("original_port").([]interface{}); len(or_port) > 0 {
			original_port = or_port[0].(map[string]interface{})
		} else {
			original_port = make(map[string]interface{})
			original_port["port"] = 0
			original_port["protocol"] = ""
		}
		if pat_o := d.Get("pat_options").([]interface{}); len(pat_o) > 0 {
			pat_options := pat_o[0].(map[string]interface{})
			pat_pool_options := pat_options["pat_pool_address"].([]interface{})[0].(map[string]interface{})
			pat_pool := AutoNatRuleSubConfig{
				ID:   pat_pool_options["id"].(string),
				Type: pat_pool_options["type"].(string),
			}
			patOptions = &AutoNatRulePatOptions{
				Patpooladdress: pat_pool,
				Includereserve: pat_options["include_reserve_ports"].(bool),
				Interfacepat:   pat_options["interface_pat"].(bool),
				Extendedpat:    pat_options["extended_pat_table"].(bool),
				Roundrobin:     pat_options["round_robin"].(bool),
			}
		}
		res, err := c.UpdateAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id(), &AutoNatRule{
			ID:                           d.Id(),
			Type:                         autonat_rules_type,
			Description:                  d.Get("description").(string),
			Nattype:                      strings.ToUpper(d.Get("nat_type").(string)),
			Sourceinterface:              sourceInterface,
			Destinationinterface:         destinationInterface,
			Originalnetwork:              originalNetwork,
			Translatednetwork:            translatedNetwork,
			Interfaceintranslatednetwork: d.Get("translated_network_is_destination_interface").(bool),
			Originalport:                 original_port["port"].(int),
			Serviceprotocol:              strings.ToUpper(original_port["protocol"].(string)),
			Translatedport:               d.Get("translated_port").(int),
			Fallthrough:                  d.Get("fallthrough").(bool),
			DNS:                          d.Get("translate_dns").(bool),
			Noproxyarp:                   d.Get("no_proxy_arp").(bool),
			Routelookup:                  d.Get("perform_route_lookup").(bool),
			Nettonet:                     d.Get("net_to_net").(bool),
			Interfaceipv6:                d.Get("ipv6").(bool),
			Patoptions:                   patOptions,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)
	}
	return resourceAutoNatRulesRead(ctx, d, m)
}

func resourceAutoNatRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
