package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var VNI_interface_type string = "VNIInterface"

func resourceFmcDeviceVNIInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for VNI Interfaces in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_VNI_interfaces\" \"VNI_interfaces_1\" {\n" +
			"    id = fpVNIinterfaceUUID1\n" +
			"    name = \"s1p1\"\n" +
			"    type = \"FPVNIInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
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
			"\n" +
			"resource \"fmc_VNI_interfaces\" \"VNI_interfaces_2\" {\n" +
			"    id = fpVNIinterfaceUUID2\n" +
			"    name = \"s1p2\"\n" +
			"    type = \"FPVNIInterface\"\n" +
			"    enabled = false\n" +
			"    MTU = 1500\n" +
			"    mode = true\n" +
			"    ifname = false\n" +
			"    nveOnly = true\n" +
			"    managementOnly = true\n" +
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
			"\n" +
			"}\n" +
			"```\n",
		UpdateContext: resourceFmcDeviceVNIInterfacesUpdate,
		ReadContext:   resourceFmcDeviceVNIInterfacesRead,
		DeleteContext: resourceFmcDeviceVNIInterfacesDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this VNI interfaces",
			},
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of VNI interfaces",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of this VNI interfaces",
			},
			"vni_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The description of this VNI interfaces",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of this VNI interfaces",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enabled of this VNI interfaces",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "MTU of this VNI interfaces",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "mode of this VNI interfaces",
			},
			"ifname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ifname of this VNI interfaces",
			},
			"enable_sgt_propagate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableSGTPropagate of this VNI interfaces",
			},
			"management_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "managementOnly of this VNI interfaces",
			},
			"enable_anti_spoofing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableAntiSpoofing of this VNI interfaces",
			},
			"active_mac_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "activeMACAddress of this VNI interfaces",
			},
			"fragment_reassembly": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "activeMACAddress of this VNI interfaces",
			},
			"standby_mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "standbyMACAddress of this VNI interfaces",
			},
			"enable_dns_lookup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableDNSLookup of this VNI interfaces",
			},
			"enable_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enableProxy of this VNI interfaces",
			},
			"vtep_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "vtepID of this VNI interfaces",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "priority of this VNI interfaces",
			},
			"multicast_group_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "multicast_group_address of this VNI interfaces",
			},
			"override_default_fragment_setting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The ID of this resource",
						},
						"chain": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The type of this resource",
						},
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The name of this resource",
						},
					},
				},
			},
			"fmc_access_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_networks": {
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
									"description": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The description of this resource",
									},
								},
							},
						},
						"enable_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "enableAccess of this resource",
						},
					},
				},
				Description: "fmcAccessConfig for this resource",
			},
			"security_zone": {
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
				Description: "Destination zones for this resource",
			},
		},
	}
}

func resourceFmcDeviceVNIInterfacesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	//var seczone SecurityZone_VNI
	// var overrideDefaultFragmentSetting OverrideDefaultFragmentSetting
	// var arpConfig []ArpConfig

	// if inputSeczone, ok := d.GetOk("security_zone"); ok {
	// 	szs := inputSeczone.([]interface{})[0].(map[string]interface{})
	// 	seczone = SecurityZone_VNI{
	// 		ID:    szs["id"].(string),
	// 		Type:  szs["type"].(string),
	// 	}
	// }

   var seczone []SecurityZone_VNI
   if inputSeczone, ok := d.GetOk("security_zone"); ok {
	   for _, sz := range inputSeczone.([]interface{}) {
		   szi := sz.(map[string]interface{})
		   seczone = append(seczone, SecurityZone_VNI{
			   ID:   szi["id"].(string),
			   Type: szi["type"].(string),
		   })
	   }
	}
	// if inputObjs, ok := d.GetOk("override_default_fragment_setting"); ok {
	// 	obj := inputObjs.([]interface{})[0].(map[string]interface{})
	// 	overrideDefaultFragmentSetting = OverrideDefaultFragmentSetting{
	// 		Size:    obj["size"].(int),
	// 		Timeout: obj["timeout"].(int),
	// 		Chain:   obj["chain"].(int),
	// 	}
	// }

	// if inputObjs, ok := d.GetOk("arpConfig"); ok {
	// 	obj := inputObjs.([]interface{})[0].(map[string]interface{})
	// 	arpConfig = append(arpConfig, ArpConfig{
	// 		IpAddress:   obj["ipAddress"].(string),
	// 		MacAddress:  obj["macAddress"].(string),
	// 		EnableAlias: obj["enableAlias"].(bool),
	// 	})
	// }
	res, err := c.CreateVNIInterfaces(ctx, d.Get("device_id").(string), &FTDVNIInterfaces{
		Type:                           VNI_interface_type,
		// ArpConfig:                      arpConfig,
		// OverrideDefaultFragmentSetting: overrideDefaultFragmentSetting,
		Name:                           d.Get("name").(string),
		VNIID:                          d.Get("vni_id").(int),
		MTU: 							d.Get("mtu").(int),
		EnableProxy: 	 				d.Get("enable_proxy").(bool),
		MulticastGroupAddress:          d.Get("multicast_group_address").(string),
		VTEPID:                         d.Get("vtep_id").(int),
		Mode: 							d.Get("mode").(string),
		Enabled:                        d.Get("enabled").(bool),
		SecurityZone: seczone,
	})
	if err != nil {
		return returnWithDiag(diags, err)
	}
	d.SetId(res.ID)
	return resourceFmcDeviceVNIInterfacesRead(ctx, d, m)
}

func resourceFmcDeviceVNIInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetVNIInterfaces(ctx, d.Get("name").(string), d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	if err := d.Set("type", item.Type); err != nil {
		return returnWithDiag(diags, err)
	}

	return diags
}

func resourceFmcDeviceVNIInterfacesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("override_default_fragment_setting", "arpConfig", "name", "vni_id", "multicast_group_address", "vtep_id", "enabled", "mode") {
		// var overrideDefaultFragmentSetting OverrideDefaultFragmentSetting
		// var arpConfig []ArpConfig
		// var seczone SecurityZone_VNI

		// if inputObjs, ok := d.GetOk("override_default_fragment_setting"); ok {
		// 	obj := inputObjs.([]interface{})[0].(map[string]interface{})
		// 	overrideDefaultFragmentSetting = OverrideDefaultFragmentSetting{
		// 		Size:    obj["size"].(int),
		// 		Timeout: obj["timeout"].(int),
		// 		Chain:   obj["chain"].(int),
		// 	}
		// }

		// if inputSeczone, ok := d.GetOk("security_zone"); ok {
		// 	szs := inputSeczone.([]interface{})[0].(map[string]interface{})
		// 	seczone = SecurityZone_VNI{
		// 		ID:    szs["id"].(string),
		// 		Type: szs["type"].(string),
		// 	}
		// }
		var seczone []SecurityZone_VNI
   		if inputSeczone, ok := d.GetOk("security_zone"); ok {
	   	for _, sz := range inputSeczone.([]interface{}) {
		   szi := sz.(map[string]interface{})
		   seczone = append(seczone, SecurityZone_VNI{
			   ID:   szi["id"].(string),
			   Type: szi["type"].(string),
		   })
	   }
	}
		// if inputObjs, ok := d.GetOk("arpConfig"); ok {
		// 	obj := inputObjs.([]interface{})[0].(map[string]interface{})
		// 	arpConfig = append(arpConfig, ArpConfig{
		// 		IpAddress:   obj["ipAddress"].(string),
		// 		MacAddress:  obj["macAddress"].(string),
		// 		EnableAlias: obj["enableAlias"].(bool),
		// 	})
		// }
		res, err := c.UpdateVNIInterfaces(ctx, d.Get("device_id").(string), &FTDVNIInterfaces{
			Type:                           VNI_interface_type,
			// ArpConfig:                      arpConfig,
			// OverrideDefaultFragmentSetting: overrideDefaultFragmentSetting,
			Name:                           d.Get("name").(string),
			VNIID:                          d.Get("vni_id").(int),
			MTU: 							d.Get("mtu").(int),
			EnableProxy: 	 				d.Get("enable_proxy").(bool),
			MulticastGroupAddress:          d.Get("multicast_group_address").(string),
			VTEPID:                         d.Get("vtep_id").(int),
			Mode: 							d.Get("mode").(string),
			Enabled:                        d.Get("enabled").(bool),
			SecurityZone: seczone,
		})
		if err != nil {
			return returnWithDiag(diags, err)
		}
		d.SetId(res.ID)

	}
	return resourceFmcDeviceVNIInterfacesRead(ctx, d, m)
}

func resourceFmcDeviceVNIInterfacesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteVNIInterfaces(ctx, d.Id())
	if err != nil {
		return returnWithDiag(diags, err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}