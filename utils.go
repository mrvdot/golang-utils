// Package utils is a collection of helpful utilities for common actions within GoLang development
package utils

import (
	"reflect"
	"strings"
	"unicode"
)

// type ApiResponse is a generic API response struct
type ApiResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Result  interface{}            `json:"result"`
	Data    map[string]interface{} `json:"data"` // Generic extra data to be sent along in response
}

// GenerateSlug converts a string into a lowercase dasherized slug
// For example: GenerateSlug("My cool object") returns "my-cool-object"
func GenerateSlug(str string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
		return -1
	}, strings.ToLower(strings.TrimSpace(str)))
}

// InChain returns a boolean if a string is already in a slice of strings
// Todo - Extend this to work for all standard types
func InChain(needle string, haystack []string) bool {
	if haystack == nil {
		return false
	}
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}

// Similar to "extend" in JS, only updates fields that are specified and not empty in newData
// Both newData and mainObj must be pointers to struct objects
func Update(mainObj interface{}, newData interface{}) bool {
	newDataVal, mainObjVal := reflect.ValueOf(newData).Elem(), reflect.ValueOf(mainObj).Elem()
	fieldCount := newDataVal.NumField()
	changed := false
	for i := 0; i < fieldCount; i++ {
		newField := newDataVal.Field(i)
		// They passed in a value for this field, update our DB user
		if newField.IsValid() && !IsEmpty(newField) {
			dbField := mainObjVal.Field(i)
			dbField.Set(newField)
			changed = true
		}
	}
	return changed
}

// IsEmpty checks to see if a field has a set value
// Goes beyond usual reflect.IsZero check to handle numbers, strings, and slices
// For structs, iterates over all accessible properties and returns true only if all nested fields
// are also empty.
func IsEmpty(val reflect.Value) bool {
	typeStr := val.Kind().String()
	switch typeStr {
	case "int", "int8", "int16", "int32", "int64":
		return val.Int() == 0
	case "float", "float8", "float16", "float32", "float64":
		return val.Float() == 0
	case "string":
		return val.String() == ""
	case "slice", "ptr", "map", "chan", "func":
		// Check for empty slices and props
		return val.IsNil()
	case "struct":
		fieldCount := val.NumField()
		for i := 0; i < fieldCount; i++ {
			field := val.Field(i)
			if field.IsValid() && !IsEmpty(field) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
