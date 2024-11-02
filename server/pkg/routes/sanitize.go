package routes

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"html"
	"reflect"
)

// SanitizeFields recursively sanitizes string fields of a given interface{} object
func SanitizeFields(input interface{}) {
	p := bluemonday.UGCPolicy()

	sanitizeFn := func(userInput string) string {
		return html.UnescapeString(p.Sanitize(userInput))
	}

	// Use reflection to inspect the type of the input
	val := reflect.ValueOf(input)

	// If the input is not a pointer or interface, return
	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Interface {
		fmt.Println("Input must be a pointer or interface")
		return
	}

	// Dereference the pointer or interface to get the underlying value
	val = val.Elem()

	// If the value is not a struct, return
	if val.Kind() != reflect.Struct {
		fmt.Println("Input must be a struct")
		return
	}

	// Iterate through the fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// If the field is a string, sanitize it
		if field.Kind() == reflect.String {
			oldValue := field.String()
			newValue := sanitizeFn(oldValue)

			// Set the sanitized value back to the field
			field.SetString(newValue)
		}

		// Handle *string type
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.String {
			if !field.IsNil() {
				// If the pointer is not nil, sanitize the string value it points to
				oldValue := field.Elem().String()
				newValue := sanitizeFn(oldValue)
				field.Elem().SetString(newValue)
			}
		}

		// If the field is a nested struct, recursively call the function
		if field.Kind() == reflect.Struct {
			SanitizeFields(field.Addr().Interface())
		}
	}
}
