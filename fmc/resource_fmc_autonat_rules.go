package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var autonat_rules_type string = "AutoNatRule"

func resourceFmcAutoNatRules() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Auto NAT Rules in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_ftd_autonat_rules\" \"new_rule\" {\n" +
			"    nat_policy = fmc_ftd_nat_policies.nat_policy.id\n" +
			"    description = \"Testing Auto NAT priv-pub\"\n" +
			"    nat_type = \"static\"\n" +
			"    source_interface {\n" +
			"        id = data.fmc_security_zones.inside.id\n" +
			"        type = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    destination_interface {\n" +
			"        id = data.fmc_security_zones.outside.id\n" +
			"        type = data.fmc_security_zones.outside.type\n" +
			"    }\n" +
			"    original_network {\n" +
			"        id = data.fmc_network_objects.private.id\n" +
			"        type = data.fmc_network_objects.private.type\n" +
			"    }\n" +
			"    translated_network {\n" +
			"        id = data.fmc_network_objects.public.id\n" +
			"        type = data.fmc_network_objects.public.type\n" +
			"    }\n" +
			"    translated_network_is_destination_interface = false\n" +
			"    original_port {\n" +
			"        port = 53\n" +
			"        protocol = \"udp\"\n" +
			"    }\n" +
			"    translated_port = 5353\n" +
			"    ipv6 = true\n" +
			"}\n" +
			"\n" +
			"resource \"fmc_ftd_autonat_rules\" \"new_rule_2\" {\n" +
			"    nat_policy = fmc_ftd_nat_policies.nat_policy.id\n" +
			"    description = \"Testing Auto NAT pub-priv\"\n" +
			"    nat_type = \"dynamic\"\n" +
			"    source_interface {\n" +
			"        id = data.fmc_security_zones.inside.id\n" +
			"        type = data.fmc_security_zones.inside.type\n" +
			"    }\n" +
			"    destination_interface {\n" +
			"        id = data.fmc_security_zones.outside.id\n" +
			"        type = data.fmc_security_zones.outside.type\n" +
			"    }\n" +
			"    original_network {\n" +
			"        id = data.fmc_host_objects.CUCMPub.id\n" +
			"        type = data.fmc_host_objects.CUCMPub.type\n" +
			"    }\n" +
			"    translated_network_is_destination_interface = false\n" +
			"    pat_options {\n" +
			"        pat_pool_address {\n" +
			"            id = data.fmc_network_objects.private.id\n" +
			"            type = data.fmc_network_objects.private.type\n" +
			"        }\n" +
			"        extended_pat_table = true\n" +
			"        round_robin = true\n" +
			"    }\n" +
			"    ipv6 = true\n" +
			"}\n" +
			"```\n" +
			"**Note** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.",
		CreateContext: resourceFmcAutoNatRulesCreate,
		ReadContext:   resourceFmcAutoNatRulesRead,
		UpdateContext: resourceFmcAutoNatRulesUpdate,
		DeleteContext: resourceFmcAutoNatRulesDelete,
		Schema: map[string]*schema.Schema{
			"nat_policy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the NAT policy this resource belongs to",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
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
				Description: "Destination interface of this resource",
			},
			"original_network": {
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
				Description: "Original network for this resource",
			},
			"translated_network": {
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
				Description: "Translated interface for this resource",
			},
			"translated_network_is_destination_interface": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Interface is the destination translated network",
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
							StateFunc: func(val interface{}) string {
								return strings.ToUpper(val.(string))
							},
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return strings.EqualFold(old, new)
							},
						},
					},
				},
				Description: "Original port for this resource",
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
				Description: "Translated port for this resource",
			},
			"fallthrough": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable Fallthrough",
			},
			"translate_dns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable Translate DNS",
			},
			"no_proxy_arp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable Proxy ARP",
			},
			"perform_route_lookup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable perform route lookups for this resource",
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
							Description: "Network Pool for PAT",
						},
						"interface_pat": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable interface PAT",
						},
						"include_reserve_ports": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Include Reserve ports",
						},
						"extended_pat_table": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Extended PAT table",
						},
						"round_robin": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Round Robin",
						},
					},
				},
				Description: "PAT options for this resource",
			},
		},
	}
}

func resourceFmcAutoNatRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		pat_pool_address_list := pat_options["pat_pool_address"].([]interface{})
		var pat_pool *AutoNatRuleSubConfig
		if len(pat_pool_address_list) != 0 {
			pat_pool_address := pat_pool_address_list[0].(map[string]interface{})
			pat_pool = &AutoNatRuleSubConfig{
				ID:   pat_pool_address["id"].(string),
				Type: pat_pool_address["type"].(string),
			}
		}
		patOptions = &AutoNatRulePatOptions{
			Patpooladdress: pat_pool,
			Includereserve: pat_options["include_reserve_ports"].(bool),
			Interfacepat:   pat_options["interface_pat"].(bool),
			Extendedpat:    pat_options["extended_pat_table"].(bool),
			Roundrobin:     pat_options["round_robin"].(bool),
		}
	}

	res, err := c.CreateFmcAutoNatRule(ctx, d.Get("nat_policy").(string), &AutoNatRule{
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
	return resourceFmcAutoNatRulesRead(ctx, d, m)
}

func resourceFmcAutoNatRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id())
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

	// See https://gitlab-sjc.cisco.com/tfprovider/fmc-terraform/issues/33 , API does not return this field.
	if err := d.Set("description", d.Get("description")); err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("nat_type", item.Nattype); err != nil {
		return returnWithDiag(diags, err)
	}

	if item.Sourceinterface != (AutoNatRuleSubConfig{}) {
		if err := d.Set("source_interface", convertTo1ListMapStringGeneric(item.Sourceinterface)); err != nil {
			return returnWithDiag(diags, err)
		}
	}
	if item.Destinationinterface != (AutoNatRuleSubConfig{}) {
		if err := d.Set("destination_interface", convertTo1ListMapStringGeneric(item.Destinationinterface)); err != nil {
			return returnWithDiag(diags, err)
		}
	}
	if item.Originalnetwork != (AutoNatRuleSubConfig{}) {
		if err := d.Set("original_network", convertTo1ListMapStringGeneric(item.Originalnetwork)); err != nil {
			return returnWithDiag(diags, err)
		}
	}
	if item.Translatednetwork != (AutoNatRuleSubConfig{}) {
		if err := d.Set("translated_network", convertTo1ListMapStringGeneric(item.Translatednetwork)); err != nil {
			return returnWithDiag(diags, err)
		}
	}
	if err := d.Set("translated_network_is_destination_interface", item.Interfaceintranslatednetwork); err != nil {
		return returnWithDiag(diags, err)
	}

	if item.Originalport != 0 && item.Serviceprotocol != "" {
		original_port := make(map[string]interface{})
		original_port["port"] = item.Originalport
		original_port["protocol"] = item.Serviceprotocol
		if err := d.Set("original_port", convertTo1ListGeneric(original_port)); err != nil {
			return returnWithDiag(diags, err)
		}
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

	if item.Patoptions != (AutoNatRulePatOptions{}) {
		pat_options := make(map[string]interface{})
		if item.Patoptions.Patpooladdress != nil && *item.Patoptions.Patpooladdress != (AutoNatRuleSubConfig{}) {
			pat_options["pat_pool_address"] = convertTo1ListMapStringGeneric(item.Patoptions.Patpooladdress)
		}
		pat_options["interface_pat"] = item.Patoptions.Interfacepat
		pat_options["include_reserve_ports"] = item.Patoptions.Includereserve
		pat_options["extended_pat_table"] = item.Patoptions.Extendedpat
		pat_options["round_robin"] = item.Patoptions.Roundrobin
		if err := d.Set("pat_options", convertTo1ListGeneric(pat_options)); err != nil {
			return returnWithDiag(diags, err)
		}
	}

	return diags
}

func resourceFmcAutoNatRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
			pat_pool_address_list := pat_options["pat_pool_address"].([]interface{})
			var pat_pool *AutoNatRuleSubConfig
			if len(pat_pool_address_list) != 0 {
				pat_pool_address := pat_pool_address_list[0].(map[string]interface{})
				pat_pool = &AutoNatRuleSubConfig{
					ID:   pat_pool_address["id"].(string),
					Type: pat_pool_address["type"].(string),
				}
			}
			patOptions = &AutoNatRulePatOptions{
				Patpooladdress: pat_pool,
				Includereserve: pat_options["include_reserve_ports"].(bool),
				Interfacepat:   pat_options["interface_pat"].(bool),
				Extendedpat:    pat_options["extended_pat_table"].(bool),
				Roundrobin:     pat_options["round_robin"].(bool),
			}
		}
		res, err := c.UpdateFmcAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id(), &AutoNatRule{
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
	return resourceFmcAutoNatRulesRead(ctx, d, m)
}

func resourceFmcAutoNatRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcAutoNatRule(ctx, d.Get("nat_policy").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
