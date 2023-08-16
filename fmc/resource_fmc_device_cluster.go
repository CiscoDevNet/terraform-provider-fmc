package fmc

import (
	"context"
	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var cluster_type string = "DeviceCluster"

func resourceFmcDeviceCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for adding device in a cluster\n" +
			"\n" +
			"**Note: this will only work on VMware, not on public cloud.**\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			" resource \"fmc_device_cluster\" \"cluster\" { \n"+
			"	name = \"ftd_cluster\" \n"+
			"	control_device {\n"+
			"	  cluster_node_bootstrap {\n"+
			"		priority = 1\n"+
			"		cclip = \"10.10.11.1\"\n"+
			"	  }\n"+
			"	  device_details  {\n"+
			"		id = data.fmc_devices.device1.id\n"+
			"		name = data.fmc_devices.device1.name\n"+
			"	  }\n"+
			"	}\n"+
			"	common_bootstrap {\n"+
			"	  ccl_interface {\n" +
			"		id = data.fmc_device_physical_interfaces.ccl_physical_interface.id\n" +
			"		name = data.fmc_device_physical_interfaces.ccl_physical_interface.name\n" +
			"	  }\n" +
			"	  ccl_network = \"10.10.11.0/27\"\n" +
			"	  vni_network = \"10.10.10.0/27\"\n" +
			"	}\n" +  
			"	data_devices{\n" +
			"		cluster_node_bootstrap{\n" +
			"		  priority = 2\n" +
			"		  cclip = \"10.10.11.2\"\n" +
			"		}\n" +
			"		device_details{\n" +
			"		  id = data.fmc_devices.device2.id\n" +
			"		  name = data.fmc_devices.device2.name\n" +
			"		}\n" +
			"	}\n" +
			"	data_devices{\n" +
			"		cluster_node_bootstrap{\n" +
			"		  priority = 3\n" +
			"		  cclip = \"10.10.11.3\"\n" +
			"		}\n" +
			"		device_details{\n" +
			"		  id = data.fmc_devices.device3.id\n" +
			"		  name = data.fmc_devices.device3.name\n" +
			"		}\n" +
			"	}\n" +
			"  }\n" +
			"**Note:** This feature is only supported for VMWare cloud platform.\n"+
			"**Note:** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.\n" +
			"**Note:** Deleting a cluster will delete all it's data nodes with control node as well.\n" +
			"```",
		CreateContext: resourceFmcDeviceClusterCreate,
		ReadContext:   resourceFmcDeviceClusterRead,
		UpdateContext: resourceFmcDeviceClusterUpdate,
		DeleteContext: resourceFmcDeviceClusterDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of cluster",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"control_device": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_node_bootstrap": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Set the priority of cluster node",
									},
									"cclip": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Cluster control IP",
									},
								},
							},
						},
						"device_details": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Device ID",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Description: "Device Name",
									},
									"type":{
										Type:     schema.TypeString,
										Computed: true,
										Description: "Device Type", 
									},
								},
							},
						},
					},
				},
				Description: "Control device information",
			},
			"data_devices": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 500,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_node_bootstrap": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Set the priority of cluster node",
									},
									"cclip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cluster control IP",
									},
								},
							},
						},
						"device_details": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Device ID",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Description: "Device Name",
									},
									"type":{
										Type:     schema.TypeString,
										Computed: true,
										Description: "Device Type", 
									},
								},
							},
						},
					},
				},
				Description: "Data device information",
			},
			"common_bootstrap": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ccl_interface": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface ID",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Description: "Interface Name",
									},
									"type":{
										Type:     schema.TypeString,
										Computed: true,
										Description: "Interface Type", 
									},
								},
							},
						},
						"ccl_network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster control network",
						},
						"vni_network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VNI network",
						},
					},
				},
				Description: "Control device information",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFmcDeviceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var control_device = ControlDevice{}

	if d.Get("control_device") != nil {
		if ControlEntries, ok := d.GetOk("control_device"); ok {
			entries := ControlEntries.([]interface{})[0]
			controlEntry := entries.(map[string]interface{})

			if controlEntry["cluster_node_bootstrap"] != nil {
				node_bootstraps := controlEntry["cluster_node_bootstrap"].([]interface{})
				if len(node_bootstraps) > 0 {
					node_bootstrap := node_bootstraps[0].(map[string]interface{})
					var ClusterNodeBootstrap = ClusterNodeBootstrap{
						CclIP: node_bootstrap["cclip"].(string),
						Priority: node_bootstrap["priority"].(int),
					}
					control_device.ClusterNode = &ClusterNodeBootstrap
				}

			}

			if controlEntry["device_details"] != nil {

				details := controlEntry["device_details"].([]interface{})
				detail := details[0].(map[string]interface{})

				var DeviceDetails = DeviceDetails{
					ID:      detail["id"].(string),
					Name: detail["name"].(string),
					Type: "Device",
				}
				control_device.DeviceDetail = &DeviceDetails
			}

		}

	}

	// Funky stuff starts here

	dataDevices := d.Get("data_devices").([]interface{})
	Log.info(dataDevices)

	obj := Cluster{
		DataDevices: make([]ControlDevice, len(dataDevices)),
	}
	for ind, node := range dataDevices {
		node := node.(map[string]interface{})
		var devDet *DeviceDetails
		dynamicObjects2 := []**DeviceDetails{
			&devDet,
		}

		if inputEntries, ok := node["device_details"]; ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects2[0] = &DeviceDetails{
				ID:   entry["id"].(string),
				Type: "Device",
				Name: entry["name"].(string),
			}
		}

		var clusno *ClusterNodeBootstrap
		dynamicObjects3 := []**ClusterNodeBootstrap{
			&clusno,
		}

		if inputEntries, ok := node["cluster_node_bootstrap"]; ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})
			*dynamicObjects3[0] = &ClusterNodeBootstrap{
				CclIP:   entry["cclip"].(string),
				Priority: entry["priority"].(int),
			}
		}

		obj.DataDevices[ind] = ControlDevice{
			ClusterNode:     clusno,
			DeviceDetail:    devDet,
		}
	}
	// Funky stuff ends here
	var common_bootstrap = CommonBootstrap{}
	if d.Get("common_bootstrap") != nil {
		if interfaceEntries, ok := d.GetOk("common_bootstrap"); ok {
			entries := interfaceEntries.([]interface{})[0]
			interfaceEntry := entries.(map[string]interface{})

			if interfaceEntry["ccl_interface"] != nil {
				interafaces := interfaceEntry["ccl_interface"].([]interface{})
				if len(interafaces) > 0 {
					interfac := interafaces[0].(map[string]interface{})
					var CclInterface = DeviceDetails{
						ID:      interfac["id"].(string),
						Name:    interfac["name"].(string),
						Type:    "PhysicalInterface",
					}
					common_bootstrap.CclInterface = &CclInterface
				}

			}
			common_bootstrap.CclNetwork = interfaceEntry["ccl_network"].(string)
			common_bootstrap.VniNetwork = interfaceEntry["vni_network"].(string)
		}

	}
	_, err := c.CreateFmcDeviceCluster(ctx, &Cluster{
		Name:            d.Get("name").(string),
		Type:            cluster_type,
		ControlDevice:   control_device,
		CommonBootstrap: common_bootstrap,
		DataDevices : obj.DataDevices,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to add device",
			Detail:   err.Error(),
		})
		return diags
	}
	//ADD CODE TO GET ID
	cluster, err := c.GetFmcDeviceClusterByName(ctx, d.Get("name").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get cluster",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(cluster.ID)

	//Code ENDED
	return resourceFmcDeviceClusterRead(ctx, d, m)
}

func resourceFmcDeviceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcDeviceCluster(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read device",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcDeviceClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	if d.HasChanges("name") {

			res, err := c.UpdateFmcDeviceCluster(ctx, d.Id(), &Cluster{
			ID:              d.Id(),
			Name:            d.Get("name").(string),
			Action: 		"UPDATE_CLUSTER_NAME",

		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to create cluster",
				Detail:   err.Error(),
			})
			return diags
		}
		d.SetId(res.ID)
	}
	return resourceFmcDeviceClusterRead(ctx, d, m)
}

func resourceFmcDeviceClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcDeviceCluster(ctx, m, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete cluster",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
