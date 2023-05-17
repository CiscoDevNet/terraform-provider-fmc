package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcExtendedAcl() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Extended ACL in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_extended_acl\" \"acl1\" {\n" +
			"    name = \"ACL-1\"\n" +
			"    action = \"DENY\"\n" +
			"    log_level = \"ERROR\"\n" +
			"     logging = \"PER_ACCESS_LIST_ENTRY\"\n" +
			"     log_interval= 545\n" +
			"}\n" +
			"```\n",
		CreateContext: resourceFmcExtendedAclCreate,
		ReadContext:   resourceFmcExtendedAclRead,
		UpdateContext: resourceFmcExtendedAclUpdate,
		DeleteContext: resourceFmcExtendedAclDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The action of this resource",
			},
			"log_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The loglevel of this resource",
			},
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The logging of this resource",
			},
			"log_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The log interval of this resource",
			},
			"source_port_object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Port Object ID",
			},
			"source_network_object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Network Object ID",
			},
			"destination_network_object_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination Network Object ID",
			},

			"source_port_literal_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Port Literal Port",
			},
			"source_port_literal_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Port Literal Protocol",
			},
			"source_network_literal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Network Literal Type",
			},
			"source_network_literal_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Network Literal Value",
			},

			"destination_network_literal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination Network Literal Type",
			},
			"destination_network_literal_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination Network Literal Value",
			},
		},
	}
}

func resourceFmcExtendedAclCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	//Source port object
	SourcePort_objectID := d.Get("source_port_object_id").(string)
	var SourcePort_Object_input []Object_data
	if len(SourcePort_objectID) > 0 {
		SourcePort_Object_input = append(SourcePort_Object_input, Object_data{
			ID: SourcePort_objectID,
		})
	}

	//Source port Literal
	SourcePort_lit_Port := d.Get("source_port_literal_port").(string)
	SourcePort_lit_Protocol := d.Get("source_port_literal_protocol").(string)

	var SourcePort_Lit_input []Literals_Port_data
	if len(SourcePort_lit_Port) > 0 && len(SourcePort_lit_Protocol) > 0 {
		SourcePort_Lit_input = append(SourcePort_Lit_input, Literals_Port_data{
			Type:     "PortLiteral",
			Port:     SourcePort_lit_Port,
			Protocol: SourcePort_lit_Protocol,
		})
	}

	//Source Network Object
	SourceNw_objectID := d.Get("source_network_object_id").(string)
	var SourceNw_Object_input []Object_data
	if len(SourceNw_objectID) > 0 {
		SourceNw_Object_input = append(SourceNw_Object_input, Object_data{
			ID: SourceNw_objectID,
		})
	}

	////Source Network Literal
	SourceNw_litType := d.Get("source_network_literal_type").(string)
	SourceNw_litValue := d.Get("source_network_literal_value").(string)
	var SourceNw_Lit_input []Literals_Nw_data
	if len(SourceNw_litType) > 0 && len(SourceNw_litValue) > 0 {
		SourceNw_Lit_input = append(SourceNw_Lit_input, Literals_Nw_data{
			Type:  SourceNw_litType,
			Value: SourceNw_litValue,
		})
	}

	// Destination Network Object
	DestinationNw_objectID := d.Get("destination_network_object_id").(string)
	var DestinationNw_Object_input []Object_data
	if len(DestinationNw_objectID) > 0 {
		DestinationNw_Object_input = append(DestinationNw_Object_input, Object_data{
			ID: DestinationNw_objectID,
		})
	}

	////Destination Network Literal
	DestinationNw_litType := d.Get("destination_network_literal_type").(string)
	DestinationNw_litValue := d.Get("destination_network_literal_value").(string)
	var DestinationNw_Lit_input []Literals_Nw_data
	if len(DestinationNw_litType) > 0 && len(DestinationNw_litValue) > 0 {
		DestinationNw_Lit_input = append(DestinationNw_Lit_input, Literals_Nw_data{
			Type:  DestinationNw_litType,
			Value: DestinationNw_litValue,
		})
	}

	res, err := c.CreateFmcExtendedAcl(ctx, &ExtendedAcl{
		Name: d.Get("name").(string),
		Type: "ExtendeddAccessList",
		Entries: []Entries_data{
			{
				Action:      d.Get("action").(string),
				LogLevel:    d.Get("log_level").(string),
				Logging:     d.Get("logging").(string),
				LogInterval: d.Get("log_interval").(int),
				SourcePorts: Data_Ports{
					Objects:  SourcePort_Object_input,
					Literals: SourcePort_Lit_input,
				},
				SourceNetworks: Data_Nw{
					Objects:  SourceNw_Object_input,
					Literals: SourceNw_Lit_input,
				},
				DestinationNetworks: Data_Nw{
					Objects:  DestinationNw_Object_input,
					Literals: DestinationNw_Lit_input,
				},
			},
		},
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create extended access list",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcExtendedAclRead(ctx, d, m)
}

func resourceFmcExtendedAclRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcExtendedAcl(ctx, d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read std acl",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read std acl",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcExtendedAclUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	//Source port object
	SourcePort_objectID := d.Get("source_port_object_id").(string)
	var SourcePort_Object_input []Object_data
	if len(SourcePort_objectID) > 0 {
		SourcePort_Object_input = append(SourcePort_Object_input, Object_data{
			ID: SourcePort_objectID,
		})
	}

	//Source port Literal
	SourcePort_lit_Port := d.Get("source_port_literal_port").(string)
	SourcePort_lit_Protocol := d.Get("source_port_literal_protocol").(string)

	var SourcePort_Lit_input []Literals_Port_data
	if len(SourcePort_lit_Port) > 0 && len(SourcePort_lit_Protocol) > 0 {
		SourcePort_Lit_input = append(SourcePort_Lit_input, Literals_Port_data{
			Type:     "PortLiteral",
			Port:     SourcePort_lit_Port,
			Protocol: SourcePort_lit_Protocol,
		})
	}

	//Source Network Object
	SourceNw_objectID := d.Get("source_network_object_id").(string)
	var SourceNw_Object_input []Object_data
	if len(SourceNw_objectID) > 0 {
		SourceNw_Object_input = append(SourceNw_Object_input, Object_data{
			ID: SourceNw_objectID,
		})
	}

	////Source Network Literal
	SourceNw_litType := d.Get("source_network_literal_type").(string)
	SourceNw_litValue := d.Get("source_network_literal_value").(string)
	var SourceNw_Lit_input []Literals_Nw_data
	if len(SourceNw_litType) > 0 && len(SourceNw_litValue) > 0 {
		SourceNw_Lit_input = append(SourceNw_Lit_input, Literals_Nw_data{
			Type:  SourceNw_litType,
			Value: SourceNw_litValue,
		})
	}

	// Destination Network Object
	DestinationNw_objectID := d.Get("destination_network_object_id").(string)
	var DestinationNw_Object_input []Object_data
	if len(DestinationNw_objectID) > 0 {
		DestinationNw_Object_input = append(DestinationNw_Object_input, Object_data{
			ID: DestinationNw_objectID,
		})
	}

	////Destination Network Literal
	DestinationNw_litType := d.Get("destination_network_literal_type").(string)
	DestinationNw_litValue := d.Get("destination_network_literal_value").(string)
	var DestinationNw_Lit_input []Literals_Nw_data
	if len(DestinationNw_litType) > 0 && len(DestinationNw_litValue) > 0 {
		DestinationNw_Lit_input = append(DestinationNw_Lit_input, Literals_Nw_data{
			Type:  DestinationNw_litType,
			Value: DestinationNw_litValue,
		})
	}
	if d.HasChanges("name", "action", "log_level", "logging", "log_interval", "source_port_object_id", "source_network_object_id", "destination_network_object_id", "source_port_literal_port", "source_port_literal_protocol", "source_network_literal_type", "source_network_literal_value", "destination_network_literal_type", "destination_network_literal_value") {
		res, err := c.UpdateFmcExtendedAcl(ctx, d.Id(), &ExtendedAcl{
			ID:   d.Id(),
			Name: d.Get("name").(string),
			Type: "ExtendeddAccessList",
			Entries: []Entries_data{
				{
					Action:      d.Get("action").(string),
					LogLevel:    d.Get("log_level").(string),
					Logging:     d.Get("logging").(string),
					LogInterval: d.Get("log_interval").(int),
					SourcePorts: Data_Ports{
						Objects:  SourcePort_Object_input,
						Literals: SourcePort_Lit_input,
					},
					SourceNetworks: Data_Nw{
						Objects:  SourceNw_Object_input,
						Literals: SourceNw_Lit_input,
					},
					DestinationNetworks: Data_Nw{
						Objects:  DestinationNw_Object_input,
						Literals: DestinationNw_Lit_input,
					},
				},
			},
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update extended access list",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcExtendedAclRead(ctx, d, m)

}

func resourceFmcExtendedAclDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteFmcExtendedAcl(ctx, d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete extended ACL",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
