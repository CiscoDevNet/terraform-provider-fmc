package fmc

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var id_array_map []interface{}

func resourceFmcNetworkObjectsBulk() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Network Objects in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"locals {\n" +
			"object_names = [for i in range(3) : \"VLAN825-Private-DRsite-${i}\"]\n" +
			"}\n" +
			"resource \"fmc_network_objects_bulk\" \"test\" {\n" +
			"dynamic \"objects\" {\n" +
			"for_each = local.object_names\n" +

			"content {\n" +
			"  name        = objects.value\n" +
			"  value       = data.fmc_network_objects.PrivateVLAN.value\n" +
			"  description = \"testing terraform\"\n" +
			"}\n" +
			"}\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcNetworkObjectsCreateBulk,
		ReadContext:   resourceFmcNetworkObjectsReadBulk,
		UpdateContext: resourceFmcNetworkObjectsUpdateBulk,
		DeleteContext: resourceFmcNetworkObjectsDeleteBulk,
		Schema: map[string]*schema.Schema{
			"objects": {
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
							Required:    true,
							Description: "The name of this resource",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type this resource",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of this resource",
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

func resourceFmcNetworkObjectsCreateBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	objects := d.Get("objects").([]interface{})

	obj := NetworkObjects{
		Objects: make([]NetworkObject, len(objects)),
	}
	for i, o := range objects {
		obj.Objects[i] = NetworkObject{
			Name:        o.(map[string]interface{})["name"].(string),
			Description: o.(map[string]interface{})["description"].(string),
			Value:       o.(map[string]interface{})["value"].(string),
			Type:        network_type,
		}
	}

	res, err := c.CreateFmcNetworkObjectBulk(ctx, &obj)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create network objects",
			Detail:   err.Error(),
		})
		return diags
	}

	for _, o := range res.Items {
		id_array_map = append(id_array_map, map[string]interface{}{
			"id":   o.ID,
			"name": o.Name,
		})
	}
	d.Set("id_mappings", id_array_map)

	return resourceFmcNetworkObjectsReadBulk(ctx, d, m)
}

func resourceFmcNetworkObjectsReadBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ids := d.Get("id_mappings").([]interface{})

	var all_obj []interface{}
	for _, elem := range ids {
		item, err := c.GetFmcNetworkObject(ctx, elem.(map[string]interface{})["id"].(string))
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				Log.info("Network object not found")
				return diags
			} else {
				return returnWithDiag(diags, err)
			}
		}

		obj := []interface{}{
			map[string]interface{}{
				"name":        item.Name,
				"type":        item.Type,
				"id":          item.ID,
				"description": item.Description,
				"value":       item.Value,
			},
		}
		all_obj = append(all_obj, obj[0])
	}
	if err := d.Set("objects", all_obj); err != nil {
		return returnWithDiag(diags, err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func resourceFmcNetworkObjectsUpdateBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics

	if d.HasChanges("objects") {
		// newObjects := d.Get("objects").([]interface{})
		_, new := d.GetChange("objects")
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
			obj := NetworkObjects{
				Objects: make([]NetworkObject, len(newObjectsArray)),
			}
			for i, o := range newObjectsArray {
				obj.Objects[i] = NetworkObject{
					Name:        o.(map[string]interface{})["name"].(string),
					Description: o.(map[string]interface{})["description"].(string),
					Value:       o.(map[string]interface{})["value"].(string),
					Type:        network_type,
				}
			}
			res, err := c.CreateFmcNetworkObjectBulk(ctx, &obj)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to create network object",
					Detail:   err.Error(),
				})
				return diags
			}

			id_array_map = d.Get("id_mappings").([]interface{})
			for _, o := range res.Items {
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
						err := c.DeleteFmcNetworkObject(ctx, idMapping["id"].(string))
						if err != nil {
							diags = append(diags, diag.Diagnostic{
								Severity: diag.Error,
								Summary:  "unable to delete network object",
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
			item, err := c.GetFmcNetworkObject(ctx, id)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to read network object",
					Detail:   err.Error(),
				})
				return diags
			}
			obj := map[string]interface{}{
				"name":        item.Name,
				"id":          item.ID,
				"description": item.Description,
				"value":       item.Value,
				"type":        item.Type,
			}

			// Rest of the code using obj
			keysToCompare := []string{"name", "description", "value"}
			if !isEqual(obj, newObject, keysToCompare...) {
				Log.info("changed object: ", newObject)
				_, err = c.UpdateFmcNetworkObject(ctx, id, &NetworkObjectUpdateInput{
					Name:        newObject["name"].(string),
					Description: newObject["description"].(string),
					Value:       newObject["value"].(string),
					Type:        network_type,
					ID:          id,
				})

				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "unable to update network object",
						Detail:   err.Error(),
					})
					return diags
				}

				// Update the id_mappings
				id_mappings[i] = map[string]interface{}{
					"name": newObject["name"].(string),
					"id":   id,
				}

				// Update the data
				d.Set("id_mappings", id_mappings)
			}
		}
	}
	return resourceFmcNetworkObjectsReadBulk(ctx, d, m)
}

func resourceFmcNetworkObjectsDeleteBulk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ids := d.Get("id_mappings").([]interface{})
	for _, elem := range ids {
		id := elem.(map[string]interface{})["id"].(string)
		err := c.DeleteFmcNetworkObject(ctx, id)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to delete network object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return diags
}
