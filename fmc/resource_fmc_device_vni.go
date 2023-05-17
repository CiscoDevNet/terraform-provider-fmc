package fmc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVNI() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for VNI Interfaces in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_device_vni\" \"my_fmc_device_vni\" {\n" +
			"	 device_id = \"<ID of the ftd>\"\n" +
			"	 security_zone_id = \"<ID of the security zone>\"\n" +
			"	 if_name = \"Inside\"\n" +
			"	 description = \"<description>\"\n" +
			"	 priority = 3\n" +
			"	 vnid = 11\n" +
			"	 multicast_groupaddress = \"224.0.0.34\"\n" +
			"	 segment_id = 4011\n" +
			"	 enable_proxy= false\n" +
			"	 ipv4 {\n" +
			"	 	static {\n" +
			"		 address = \"3.3.3.3\"\n" +
			"	     netmask = 4\n" +
			"}\n" +
			"		dhcp {\n" +
			"	     enable_default_route_dhcp = false \n" +
			"		 dhcp_route_metric = 0\n" +
			"  		}\n" +
			"    }\n" +
			"}\n" +
			"```",

		CreateContext: resourceVNICreate,
		ReadContext:   resourceVNIRead,
		UpdateContext: resourceVNIUpdate,
		DeleteContext: resourceVNIDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Device Id of VNI",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of resource",
			},
			"vtep_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Device Id of VNI",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables VNI",
			},
			"if_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Logical Name of chosen interface",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VNI description",
			},
			"security_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "securityZone of the VNI",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Priority of the VNI",
			},
			"vnid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "VNID of the VNI",
			},
			"multicast_groupaddress": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "MulticastGroupAddress of the VNI",
			},
			"segment_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "SegmentId of the VNI",
			},
			"enable_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "EnableProxy of the VNI",
			},
			"ipv4": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP of the interface",
									},
									"netmask": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Net mask of the interface",
									},
								},
							},
						},
						"dhcp": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_default_route_dhcp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Dynamic IP of the interface",
									},
									"dhcp_route_metric": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
								},
							},
						},
					},
				},
				Description: "IPV4 information",
			},
		},
	}
}

func resourceVNIRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	log.Printf("VNI: Fetching VNI details")

	if d.Get("device_id") != nil && d.Get("id") != nil {

		deviceId := d.Get("device_id").(string)
		id := d.Get("id").(string)

		log.Printf("VNI: DeviceId=%s, Id=%s", deviceId, id)

		c := m.(*Client)

		vni, err := c.GetFmcVNIDetails(ctx, deviceId, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read vni",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("VNI: VNI Details Id=%s Name=%s IFName=%s", vni.ID, vni.Name, vni.Ifname)

		d.SetId(vni.ID)

	} else {
		log.Printf("VNI: No Proper Data Of Device Id=%s, Id=%s", d.Get("device_id"), d.Get("id"))
	}

	return diags
}

func resourceVNICreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VNI: Creating virtual interface")

	device_id := d.Get("device_id").(string)
	ifname := d.Get("if_name").(string)
	description := d.Get("description").(string)
	security_zone_id := d.Get("security_zone_id").(string)
	// priority := d.Get("priority").(int)
	vnid := d.Get("vnid").(int)
	vtepid := d.Get("vtep_id").(int)
	// multicastGroupaddress := d.Get("multicast_groupaddress").(string)
	// segmentId := d.Get("segment_id").(int)
	// enableProxy := d.Get("enable_proxy").(bool)

	var ipv4 = IPv4{}

	if d.Get("ipv4") != nil {
		if ipv4Entries, ok := d.GetOk("ipv4"); ok {
			log.Printf("IPv4s data=%v", ipv4Entries)
			entries := ipv4Entries.([]interface{})[0]
			ipV4Entry := entries.(map[string]interface{})

			isStatic := false

			if ipV4Entry["static"] != nil {
				statics := ipV4Entry["static"].([]interface{})
				if len(statics) > 0 {
					static := statics[0].(map[string]interface{})
					var IPv4Static = IPv4Static{
						Address: static["address"].(string),
						Netmask: static["netmask"].(int),
					}
					ipv4.Static = &IPv4Static
					isStatic = true
				}

			}

			if ipV4Entry["dhcp"] != nil && !isStatic {

				dhcps := ipV4Entry["dhcp"].([]interface{})
				dhcp := dhcps[0].(map[string]interface{})

				var IPv4DHCP = IPv4DHCP{
					Enable:      dhcp["enable_default_route_dhcp"].(bool),
					RouteMetric: dhcp["dhcp_route_metric"].(int),
				}
				ipv4.DHCP = &IPv4DHCP
			}

		}

	}

	log.Printf("IFname=%s Desc=%s, security_zone_id=%s, vnid=%v", ifname, description, security_zone_id, vnid)

	c := m.(*Client)

	var diags diag.Diagnostics

	var securityZone = VNISecurityZone{
		ID:   security_zone_id,
		Type: "SecurityZone",
	}

	vniResp, err := c.CreateFmcVNI(ctx, device_id, &VNIRequest{
		Ifname:                ifname,
		Description:           description,
		Priority:              d.Get("priority").(int),
		VniId:                 vnid,
		MulticastGroupAddress: d.Get("multicast_groupaddress").(string),
		SegmentId:             d.Get("segment_id").(int),
		EnableProxy:           d.Get("enable_proxy").(bool),
		SecurityZone:          securityZone,
		IPv4:                  ipv4,
		VtepID:                vtepid,
		Type:                  "VNIInterface",
		Enabled:               d.Get("enabled").(bool),
	})

	if err != nil {
		log.Printf("FPU: err=%s", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to crete VNI",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Printf("VNI Id=%s", vniResp.ID)

	d.SetId(vniResp.ID)
	return diags
}

func resourceVNIUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VNI: Updating virtual interface")

	var diags diag.Diagnostics

	id := d.Id()

	log.Printf("Updating Id=%v", id)

	if len(id) > 0 {
		device_id := d.Get("device_id").(string)

		ifname := d.Get("if_name").(string)
		description := d.Get("description").(string)
		security_zone_id := d.Get("security_zone_id").(string)
		// priority := d.Get("priority").(int)
		vnid := d.Get("vnid").(int)
		// multicastGroupaddress := d.Get("multicast_groupaddress").(string)
		// segmentId := d.Get("segment_id").(int)
		// enableProxy := d.Get("enable_proxy").(bool)
		vtepid := d.Get("vtep_id").(int)
		var ipv4 = IPv4{}

		if d.Get("ipv4") != nil {
			if ipv4Entries, ok := d.GetOk("ipv4"); ok {
				log.Printf("IPv4s data=%v", ipv4Entries)
				entries := ipv4Entries.([]interface{})[0]
				ipV4Entry := entries.(map[string]interface{})
				isStatic := false
				if ipV4Entry["static"] != nil {
					statics := ipV4Entry["static"].([]interface{})
					if len(statics) > 0 {
						static := statics[0].(map[string]interface{})
						var IPv4Static = IPv4Static{
							Address: static["address"].(string),
							Netmask: static["netmask"].(int),
						}
						log.Printf("IPv4Static=%v", IPv4Static)
						ipv4.Static = &IPv4Static
						isStatic = true
					}

				}

				if ipV4Entry["dhcp"] != nil && !isStatic {

					dhcps := ipV4Entry["dhcp"].([]interface{})
					if len(dhcps) > 0 {
						dhcp := dhcps[0].(map[string]interface{})

						var IPv4DHCP = IPv4DHCP{
							Enable:      dhcp["enable_default_route_dhcp"].(bool),
							RouteMetric: dhcp["dhcp_route_metric"].(int),
						}
						ipv4.DHCP = &IPv4DHCP
					}

				}

			}

		}

		log.Printf("Update vniUUID=%s IFname=%s Desc=%s, security_zone_id=%s, vnid=%v ", id, ifname, description, security_zone_id, vnid)

		c := m.(*Client)

		oldvni, err := c.GetFmcVNIDetails(ctx, device_id, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update VNI as not able to VNI details based on id",
				Detail:   err.Error(),
			})
			return diags
		}

		var securityZone = VNISecurityZone{
			ID:   security_zone_id,
			Type: "SecurityZone",
		}

		vniResp, err := c.UpdateFmcVNI(ctx, device_id, id, &VNIRequest{
			ID:                    id,
			Name:                  oldvni.Name,
			Ifname:                ifname,
			Description:           description,
			Priority:              d.Get("priority").(int),
			VniId:                 vnid,
			MulticastGroupAddress: d.Get("multicast_groupaddress").(string),
			SegmentId:             d.Get("segment_id").(int),
			EnableProxy:           d.Get("enable_proxy").(bool),
			SecurityZone:          securityZone,
			IPv4:                  ipv4,
			VtepID:                vtepid,
			Type:                  "VNIInterface",
			Enabled:               d.Get("enabled").(bool),
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update VNI",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Updated VNI Id=%s", vniResp.ID)

		d.SetId(vniResp.ID)

	} else {
		log.Printf("Update VNIUUID is null, so record cannot be updated")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update vni",
			Detail:   "VNIUUID is null, so record cannot be updated",
		})
	}

	return diags
}

func resourceVNIDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("VNI: Delete virtual interface")

	var diags diag.Diagnostics

	id := d.Id()

	log.Printf("Id=%v", id)

	if len(id) > 0 && len(d.Get("device_id").(string)) > 0 {
		device_id := d.Get("device_id").(string)

		log.Printf("device_id=%v ", device_id)

		c := m.(*Client)

		vniResp, err := c.DeleteFmcVNIDetails(ctx, device_id, id)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to delete VNI as not able to get VNI details based on id",
				Detail:   err.Error(),
			})
			return diags
		}

		log.Printf("Deleted VNI Id=%s", vniResp.ID)

		d.SetId(vniResp.ID)
	} else {
		log.Printf("Unable to delete the VNI as device_id or id is not present")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete the VNI as device_id or id is not present",
			Detail:   "ID is null, so record cannot be deleted",
		})
	}

	return diags
}
