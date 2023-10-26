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
			"resource \"fmc_extended_acl\" \"new_acl\" {\n" +
			"\tname         = \"new_acl_test\"\n" +
			"\taction       = \"BLOCK\" //\"PERMIT\"\n" +
			"\tlog_level    = \"ERROR\" //\"INFORMATIONAL\"\n" +
			"\tlogging      = \"PER_ACCESS_LIST_ENTRY\"\n" +
			"\tlog_interval = 545\n" +
			"\n\n" +
			"\tsource_port_object_ids       = [data.fmc_port_objects.port1.id, data.fmc_port_objects.port2.id]\n" +
			"\tsource_port_literal_port     = \"12311\"\n" +
			"\tsource_port_literal_protocol = \"6\"\n" +
			"\n\n" +
			"\tdestination_port_object_ids       = [data.fmc_port_objects.port2.id, data.fmc_port_objects.port3.id]\n" +
			"\tdestination_port_literal_port     = \"12311\"\n" +
			"\tdestination_port_literal_protocol = \"6\"\n" +
			"\n\n" +
			"\tsource_network_object_ids    = [data.fmc_network_objects.nw2.id]\n" +
			"\tsource_network_literal_type  = \"Host\"\n" +
			"\tsource_network_literal_value = \"172.16.1.2\"\n" +
			"\n\n" +
			"\tdestination_network_object_ids    = [data.fmc_network_objects.nw1.id, data.fmc_network_objects.nw2.id]\n" +
			"\tdestination_network_literal_type  = \"Host\"\n" +
			"\tdestination_network_literal_value = \"172.16.1.2\"\n" +
			"}\n\n" +
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
				Required:    true,
				Description: "The action of this resource",
			},
			"log_level": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The loglevel of this resource",
			},
			"logging": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The logging of this resource",
			},
			"log_interval": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The log interval of this resource",
			},
			"source_port_object_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Source Port Object IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_port_object_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Destination Port Object IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_network_object_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Source Network Object IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_network_object_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Destination Network Object IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"destination_port_literal_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination Port Literal Port",
			},
			"destination_port_literal_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination Port Literal Protocol",
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFmcExtendedAclCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	// Source port objects
	var source_port_objects []Object_data

	for _, id := range d.Get("source_port_object_ids").([]interface{}) {
		source_port_objects = append(source_port_objects, Object_data{
			ID: id.(string),
		})
	}

	// Destination port objects
	var dest_port_objects []Object_data

	for _, id := range d.Get("destination_port_object_ids").([]interface{}) {
		dest_port_objects = append(dest_port_objects, Object_data{
			ID: id.(string),
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

	//Destination port Literal
	DestinationPort_lit_Port := d.Get("destination_port_literal_port").(string)
	DestinationPort_lit_Protocol := d.Get("destination_port_literal_protocol").(string)

	var DestinationPort_Lit_input []Literals_Port_data
	if len(DestinationPort_lit_Port) > 0 && len(DestinationPort_lit_Protocol) > 0 {
		DestinationPort_Lit_input = append(DestinationPort_Lit_input, Literals_Port_data{
			Type:     "PortLiteral",
			Port:     DestinationPort_lit_Port,
			Protocol: DestinationPort_lit_Protocol,
		})
	}

	// Source network objects
	var source_nw_objects []Object_data

	for _, id := range d.Get("source_network_object_ids").([]interface{}) {
		source_nw_objects = append(source_nw_objects, Object_data{
			ID: id.(string),
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

	// Destination network objects
	var dest_nw_objects []Object_data

	for _, id := range d.Get("destination_network_object_ids").([]interface{}) {
		dest_nw_objects = append(dest_nw_objects, Object_data{
			ID: id.(string),
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
					Objects:  source_port_objects,
					Literals: SourcePort_Lit_input,
				},
				DestinationPorts: Data_Ports{
					Objects:  dest_port_objects,
					Literals: DestinationPort_Lit_input,
				},
				SourceNetworks: Data_Nw{
					Objects:  source_nw_objects,
					Literals: SourceNw_Lit_input,
				},
				DestinationNetworks: Data_Nw{
					Objects:  dest_nw_objects,
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
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("action", item.Entries[0].Action); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("log_level", item.Entries[0].LogLevel); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("logging", item.Entries[0].Logging); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("log_interval", item.Entries[0].LogInterval); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	var source_port_object_ids []string
	for _, obj := range item.Entries[0].SourcePorts.Objects {
		source_port_object_ids = append(source_port_object_ids, obj.ID)
	}
	if err := d.Set("source_port_object_ids", source_port_object_ids); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	var dest_port_object_ids []string
	for _, obj := range item.Entries[0].DestinationPorts.Objects {
		dest_port_object_ids = append(dest_port_object_ids, obj.ID)
	}
	if err := d.Set("destination_port_object_ids", dest_port_object_ids); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	var source_nw_object_ids []string
	for _, obj := range item.Entries[0].SourceNetworks.Objects {
		source_nw_object_ids = append(source_nw_object_ids, obj.ID)
	}
	if err := d.Set("source_network_object_ids", source_nw_object_ids); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
			Detail:   err.Error(),
		})
		return diags
	}

	var dest_nw_object_ids []string
	for _, obj := range item.Entries[0].DestinationNetworks.Objects {
		dest_nw_object_ids = append(dest_nw_object_ids, obj.ID)
	}
	if err := d.Set("destination_network_object_ids", dest_nw_object_ids); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read extended acl",
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

	// Source port objects
	var source_port_objects []Object_data

	for _, id := range d.Get("source_port_object_ids").([]interface{}) {
		source_port_objects = append(source_port_objects, Object_data{
			ID: id.(string),
		})
	}

	// Destination port objects
	var dest_port_objects []Object_data

	for _, id := range d.Get("destination_port_object_ids").([]interface{}) {
		dest_port_objects = append(dest_port_objects, Object_data{
			ID: id.(string),
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

	//Destination port Literal
	DestinationPort_lit_Port := d.Get("destination_port_literal_port").(string)
	DestinationPort_lit_Protocol := d.Get("destination_port_literal_protocol").(string)

	var DestinationPort_Lit_input []Literals_Port_data
	if len(DestinationPort_lit_Port) > 0 && len(DestinationPort_lit_Protocol) > 0 {
		DestinationPort_Lit_input = append(DestinationPort_Lit_input, Literals_Port_data{
			Type:     "PortLiteral",
			Port:     DestinationPort_lit_Port,
			Protocol: DestinationPort_lit_Protocol,
		})
	}

	// Source network objects
	var source_nw_objects []Object_data

	for _, id := range d.Get("source_network_object_ids").([]interface{}) {
		source_nw_objects = append(source_nw_objects, Object_data{
			ID: id.(string),
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

	// Destination network objects
	var dest_nw_objects []Object_data

	for _, id := range d.Get("destination_network_object_ids").([]interface{}) {
		dest_nw_objects = append(dest_nw_objects, Object_data{
			ID: id.(string),
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
	if d.HasChanges("name", "action", "log_level", "logging", "log_interval", "source_port_object_ids", "destination_port_object_ids", "source_network_object_ids", "destination_network_object_ids", "source_port_literal_port", "source_port_literal_protocol", "destination_port_literal_port", "destination_port_literal_protocol", "source_network_literal_type", "source_network_literal_value", "destination_network_literal_type", "destination_network_literal_value") {
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
						Objects:  source_port_objects,
						Literals: SourcePort_Lit_input,
					},
					DestinationPorts: Data_Ports{
						Objects:  dest_port_objects,
						Literals: DestinationPort_Lit_input,
					},
					SourceNetworks: Data_Nw{
						Objects:  source_nw_objects,
						Literals: SourceNw_Lit_input,
					},
					DestinationNetworks: Data_Nw{
						Objects:  dest_nw_objects,
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
