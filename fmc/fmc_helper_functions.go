package fmc

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

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

func ConvertStructToMap(obj interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)
	result[0] = make(map[string]interface{})

	// Iterate through the fields of the object
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		fieldName := strings.ToLower(field.Name) // Convert field name to lowercase

		// Check if the field value is a struct
		if fieldValue.Kind() == reflect.Struct {
			// Recursively convert the struct to a map
			result[0][fieldName] = ConvertStructToMap(fieldValue.Interface())
		} else {
			// Convert field value to lowercase if it is a string
			if fieldValue.Kind() == reflect.String {
				strValue := fieldValue.Interface().(string)
				hexBytes, err := hex.DecodeString(strValue)
				if err == nil {
					result[0][fieldName] = hex.EncodeToString(hexBytes)
				} else {
					result[0][fieldName] = strings.ToLower(strValue)
				}
			} else {
				// Otherwise, use the field value as is
				result[0][fieldName] = fieldValue.Interface()
			}
		}
	}

	return result
}

func isEqual(obj1, obj2 interface{}, keysToCompare ...string) bool {
	if reflect.TypeOf(obj1).Kind() == reflect.Map {
		map1 := reflect.ValueOf(obj1)
		map2 := reflect.ValueOf(obj2)

		// Check if the specified keys are present and equal
		for _, key := range keysToCompare {
			val1 := map1.MapIndex(reflect.ValueOf(key))
			val2 := map2.MapIndex(reflect.ValueOf(key))

			if !val1.IsValid() || !val2.IsValid() {
				return false
			}

			// If the values have interfaces, access the id fields for comparison
			if val1.Elem().Index(0).Kind() == reflect.Map && val2.Elem().Index(0).Kind() == reflect.Map {
				id1 := val1.Elem().Index(0).MapIndex(reflect.ValueOf("id")).Interface()
				id2 := val2.Elem().Index(0).MapIndex(reflect.ValueOf("id")).Interface()
				id1 = strings.ToLower(id1.(string))
				id2 = strings.ToLower(id2.(string))

				if id1 != id2 {
					return false
				}
			} else {
				// Use reflect.DeepEqual for non-interface values
				if !reflect.DeepEqual(val1.Interface(), val2.Interface()) {
					return false
				}
			}
		}

		return true
	}

	return false
}
