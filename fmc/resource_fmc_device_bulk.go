package fmc

import (
	"context"
	"strconv"
	"strings"
	"time"

	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcDevicesBulk() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for adding device in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_devices\" \"device1\" {\n" +
			"    name = \"ftd\"\n" +
			"    hostname = \"<IP ADDR OF HOST>\"\n" +
			"    regkey = \"<Reg key used in FTD>\"\n" +
			"    metric_value = 22\n" +
			"    license_caps = [\n" +
			"		\"MALWARE\"\n" +
			"    ]\n" +
			"    access_policy {\n" +
			"        id = data.fmc_access_policies.access_policy.id\n" +
			"	}\n" +
			"}\n" +
			"```\n" +
			"**Note:** If creating multiple rules during a single `terraform apply`, remember to use `depends_on` to chain the rules so that terraform creates it in the same order that you intended.\n" +
			"**Note:** Please use a depends_on block to create multiple devices from the same plan such that second device only starts registering after device one is finished.",
		CreateContext: resourceFmcDeviceCreateBulk,
		ReadContext:   resourceFmcDeviceReadBulk,
		UpdateContext: resourceFmcDeviceUpdateBulk,
		DeleteContext: resourceFmcDeviceDeleteBulk,
		Schema: map[string]*schema.Schema{
			"devices": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    100,
				MinItems:    1,
				Description: "The list of network objects to create",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of this resource",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of FTD",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The hostname of FTD",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of this resource",
						},
						"regkey": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Same regkey as entered in FTD",
						},
						"nat_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT_ID is required, if configured in FTD ",
						},
						"performance_tier": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Select the desired performace tier",
						},
						"license_caps": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "License caps for this resource",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"access_policy": {
							Type:     schema.TypeList,
							Required: true,
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
										Optional:    true,
										Description: "The type of this resource",
									},
								},
							},
							Description: "access policy for this resource",
						},
					},
				},
			},
			"id_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of this resource",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of this resource",
						},
					},
				},
				Description: "id array",
			},
		},
	}
}

func resourceFmcDeviceCreateBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	devices := d.Get("devices").([]interface{})
	Log.info(devices)

	obj := DeviceBulk{
		Devices: make([]Device, len(devices)),
	}

	for ind, device := range devices {
		device := device.(map[string]interface{})
		Log.info(device)
		var accpolicy *AccessPolicyItem
		dynamicObjects2 := []**AccessPolicyItem{
			&accpolicy,
		}

		if inputEntries, ok := device["access_policy"]; ok {
			entry := inputEntries.([]interface{})[0].(map[string]interface{})

			*dynamicObjects2[0] = &AccessPolicyItem{
				ID:   entry["id"].(string),
				Type: entry["type"].(string),
			}
		}

		lcap := []string{}
		for _, lic := range device["license_caps"].([]interface{}) {
			lcap = append(lcap, lic.(string))
		}

		obj.Devices[ind] = Device{
			Name:            device["name"].(string),
			HostName:        device["hostname"].(string),
			Type:            device_type,
			RegKey:          device["regkey"].(string),
			NatID:           device["nat_id"].(string),
			PerformanceTier: device["performance_tier"].(string),
			LicenseCaps:     lcap,
			AccessPolicy:    accpolicy,
		}
	}

	_, err := c.CreateFmcDeviceBulk(ctx, &obj)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to add device",
			Detail:   err.Error(),
		})
		return diags
	}

	time.Sleep(300 * time.Second)
	//ADD CODE TO GET ID

	for _, dev := range devices {
		device, err := c.GetFmcDeviceByName(ctx, dev.(map[string]interface{})["name"].(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get device",
				Detail:   err.Error(),
			})
			return diags
		}

		id_array_map = append(id_array_map, map[string]interface{}{
			"id":   device.ID,
			"name": device.Name,
		})
		d.Set("id_mappings", id_array_map)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	//Code ENDED
	return resourceFmcDeviceReadBulk(ctx, d, m)
}

func resourceFmcDeviceReadBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var all_obj []interface{}
	for _, dev := range d.Get("id_mappings").([]interface{}) {
		item, err := c.GetFmcDevice(ctx, dev.(map[string]interface{})["id"].(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read device",
				Detail:   err.Error(),
			})
			return diags
		}

		obj := []interface{}{
			map[string]interface{}{
				"name":     item.Name,
				"type":     item.Type,
				"id":       item.ID,
				"hostname": item.HostName,
				"nat_id":   item.NatID,
			},
		}
		all_obj = append(all_obj, obj[0])
	}

	if err := d.Set("devices", all_obj); err != nil {
		return returnWithDiag(diags, err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func resourceFmcDeviceUpdateBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("devices") {
		_, new := d.GetChange("devices")
		newObjects := new.([]interface{})

		id_mappings := d.Get("id_mappings").([]interface{})

		// Check if objects are added in the update cycle
		if len(id_mappings) < len(newObjects) {
			Log.info("New objects found")
			idMappingsMap := make(map[string]interface{})
			for _, idMapping := range id_mappings {
				idMapping := idMapping.(map[string]interface{})
				idMappingsMap[idMapping["name"].(string)] = struct{}{}
			}
			var newObjectsArray []interface{}
			for _, newObject := range newObjects {
				newObject := newObject.(map[string]interface{})
				if _, found := idMappingsMap[newObject["name"].(string)]; !found {
					newObjectsArray = append(newObjectsArray, newObject)
				}
			}

			// Create the new objects
			obj := DeviceBulk{
				Devices: make([]Device, len(newObjectsArray)),
			}

			for ind, device := range newObjectsArray {
				device := device.(map[string]interface{})
				Log.info(device)
				var accpolicy *AccessPolicyItem
				dynamicObjects2 := []**AccessPolicyItem{
					&accpolicy,
				}

				if inputEntries, ok := device["access_policy"]; ok {
					entry := inputEntries.([]interface{})[0].(map[string]interface{})

					*dynamicObjects2[0] = &AccessPolicyItem{
						ID:   entry["id"].(string),
						Type: entry["type"].(string),
					}
				}

				lcap := []string{}
				for _, lic := range device["license_caps"].([]interface{}) {
					lcap = append(lcap, lic.(string))
				}

				obj.Devices[ind] = Device{
					Name:            device["name"].(string),
					HostName:        device["hostname"].(string),
					Type:            device_type,
					RegKey:          device["regkey"].(string),
					NatID:           device["nat_id"].(string),
					PerformanceTier: device["performance_tier"].(string),
					LicenseCaps:     lcap,
					AccessPolicy:    accpolicy,
				}
			}

			res, err := c.CreateFmcDeviceBulk(ctx, &obj)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to create network object",
					Detail:   err.Error(),
				})
				return diags
			}

			id_array_map = d.Get("id_mappings").([]interface{})
			for _, o := range res.DeviceResponses {
				id_array_map = append(id_array_map, map[string]interface{}{
					"id":   o.ID,
					"name": o.Name,
				})
			}

			d.Set("id_mappings", id_array_map)
		}

		// Check if objects are deleted in the update cycle
		if len(id_mappings) > len(newObjects) {
			// Delete stuff now.
			idMappingsMap := make(map[string]interface{})
			for _, idMapping := range id_mappings {
				idMapping := idMapping.(map[string]interface{})
				idMappingsMap[idMapping["name"].(string)] = struct{}{}
			}

			var missingKeys []string
			// Check for missing keys in the new objects
			for key := range idMappingsMap {
				found := false

				for _, newObject := range newObjects {
					newObject := newObject.(map[string]interface{})
					if newObject["name"].(string) == key {
						found = true
						break
					}
				}

				if !found {
					missingKeys = append(missingKeys, key)
				}
			}

			// Delete the missing keys
			for _, key := range missingKeys {
				for _, idMapping := range id_mappings {
					idMapping := idMapping.(map[string]interface{})
					if idMapping["name"].(string) == key {
						err := c.DeleteFmcDevice(ctx, m, idMapping["id"].(string), idMapping["name"].(string), d.Get("cdo_host").(string), d.Get("cdo_region").(string))
						if err != nil {
							diags = append(diags, diag.Diagnostic{
								Severity: diag.Error,
								Summary:  "unable to delete Device",
								Detail:   err.Error(),
							})
							return diags
						}
					}
				}
			}

			// Remove the keys from the id_mappings
			var newIdMappings []interface{}
			for _, idMapping := range id_mappings {
				idMapping := idMapping.(map[string]interface{})
				found := false
				for _, key := range missingKeys {
					if idMapping["name"].(string) == key {
						found = true
						break
					}
				}
				if !found {
					newIdMappings = append(newIdMappings, idMapping)
				}
			}

			Log.info("New id_mappings: ", newIdMappings)
			// Update the id_mappings
			d.Set("id_mappings", newIdMappings)
		}

		// Check if objects are updated in the update cycle
		for i := 0; i < len(id_mappings) && i < len(newObjects); i++ {
			elem := id_mappings[i]
			newObject := newObjects[i].(map[string]interface{})

			// Rest of the code for processing elem and newObject
			id := elem.(map[string]interface{})["id"].(string)
			item, err := c.GetFmcDevice(ctx, id)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to read device",
					Detail:   err.Error(),
				})
				return diags
			}
			obj := map[string]interface{}{
				"name":     item.Name,
				"id":       item.ID,
				"hostname": item.HostName,
				"type":     item.Type,
			}

			// obj["access_policy"] = ConvertStructToMap(obj["access_policy"])
			// interfaceSlice := newObject["access_policy"].([]interface{})
			// newSlice := make([]map[string]interface{}, len(interfaceSlice))

			// for i, item := range interfaceSlice {
			// 	newSlice[i] = item.(map[string]interface{})
			// }

			// newObject["access_policy"] = newSlice

			keysToCheck := []string{"name", "hostname", "type", "performance_tier"}
			if !isEqual(obj, newObject, keysToCheck...) {
				// Update the object
				Log.info("changed object: ", newObject)
				lcap := []string{}
				for _, lic := range newObject["license_caps"].([]interface{}) {
					lcap = append(lcap, lic.(string))
				}

				res, err := c.UpdateFmcDevice(ctx, id, &Device{
					ID:              id,
					Name:            newObject["name"].(string),
					HostName:        newObject["hostname"].(string),
					PerformanceTier: newObject["performance_tier"].(string),
					Type:            device_type,
					LicenseCaps:     lcap,
				})
				Log.info(res)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "unable to add device",
						Detail:   err.Error(),
					})
					return diags
				}

			}
		}
	}
	return resourceFmcDeviceReadBulk(ctx, d, m)
}

func resourceFmcDeviceDeleteBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var ids string
	id_array := d.Get("id_mappings").([]interface{})
	for _, id := range id_array {
		ids = ids + id.(map[string]interface{})["id"].(string) + ","
	}
	ids = strings.TrimSuffix(ids, ",")

	err := c.DeleteFmcDeviceBulk(ctx, ids)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete device",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.Set("id_mappings", []interface{}{})
	d.Set("devices", []interface{}{})
	d.SetId("")

	return diags
}
