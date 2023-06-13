package fmc

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func returnWithDiag(diags diag.Diagnostics, err error) diag.Diagnostics {
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Error in access rule",
		Detail:   err.Error(),
	})
	return diags
}

func convertTo1ListGeneric(items interface{}) []interface{} {
	list := append(make([]interface{}, 0), items)
	return list
}

func convertMapStringToGeneric(items interface{}) map[string]interface{} {
	if item, err := ToMap(items, "json"); err == nil {
		return item
	}
	if item, err := items.(map[string]interface{}); !err {
		return item
	}
	panic(fmt.Errorf("cannot convert %T to map string generic", items))
}

func convertTo1ListMapStringGeneric(item interface{}) []interface{} {
	return convertTo1ListGeneric(convertMapStringToGeneric(item))
}

// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func isEqual(obj1 map[string]interface{}, obj2 map[string]interface{}) bool {
	// Compare each attribute of the objects
	// Return true if all attributes are equal, false otherwise
	// You can modify this function based on the structure of your objects

	// Example implementation: Check if the "name" attribute is equal

	// if obj1["id"] != obj2["id"] {
	// 	return false
	// }

	if obj1["name"] != obj2["name"] {
		return false
	}

	if obj1["description"] != obj2["description"] {
		return false
	}

	if obj1["value"] != obj2["value"] {
		return false
	}

	return true
}
