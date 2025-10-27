package gasm

import (
	"fmt"
	"reflect"
	"strings"
	"syscall/js"

	"github.com/SirMoM/go-wasm/shared"
)

func encodeJsObject[T any](strct *T) js.Value {
	val := reflect.ValueOf(strct)

	// Ensure the input is a non-nil pointer to a struct
	if val.Kind() != reflect.Ptr || val.IsNil() {
		shared.ERR("encodeJsObject requires a non-nil pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		shared.ERR("encodeJsObject expects a pointer to a struct")
	}

	// Create a real JS object so syscall/gasm doesn't need to convert Go maps/slices
	obj := js.Global().Get("Object").New()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)

		jsTag, ok := typeField.Tag.Lookup("gasm")
		jsName := ""
		isClamped := false
		if ok {
			name, opts := parseJsTag(jsTag)
			shared.Warn(name)
			shared.Warn(opts)
			if name != "" {
				jsName = name
			} else {
				jsName = typeField.Name
			}
			if opts["clamped"] {
				isClamped = true
			}
		} else {
			jsName = typeField.Name
		}

		shared.Info(fmt.Sprintf("Encoding field: %s", jsName))

		switch field.Kind() {
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.Uint8 || field.Type().Elem().Kind() == reflect.Int8 {
				b := field.Bytes()
				u8 := js.Global().Get("Uint8Array").New(len(b))
				js.CopyBytesToJS(u8, b)
				if isClamped {
					clamped := js.Global().Get("Uint8ClampedArray").New(u8.Get("buffer"))
					obj.Set(jsName, clamped)
				} else {
					obj.Set(jsName, u8)
				}
				break
			}
			// Other slice kinds are not supported in this encoder
			shared.ERR(fmt.Sprintf("encodeJsObject: unsupported slice element kind for field %s: %s", jsName, field.Type().Elem().Kind()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			obj.Set(jsName, int(field.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			obj.Set(jsName, int(field.Uint()))
		case reflect.Float32, reflect.Float64:
			obj.Set(jsName, field.Float())
		case reflect.Bool:
			obj.Set(jsName, field.Bool())
		case reflect.String:
			obj.Set(jsName, field.String())
		default:
			shared.ERR(fmt.Sprintf("encodeJsObject: unsupported field kind %s for %s", field.Kind(), jsName))
		}
	}

	return obj
}

func parseJsObject[T any](jsObj js.Value, strct *T) {
	shared.Warn(jsObj)

	val := reflect.ValueOf(strct).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)

		// Get the JSON field name from the tag or fall back to the struct field name
		var jsName string
		if jsTag, ok := typeField.Tag.Lookup("gasm"); ok {
			jsName, _ = parseJsTag(jsTag)

		} else {
			jsName = typeField.Type.Name()
		}

		// Log the JSON name for debugging purposes
		shared.Info(fmt.Sprintf("Parsing field: %s", jsName))

		// Retrieve the value from the JavaScript object
		fieldValue := jsObj.Get(jsName)

		if !fieldValue.IsUndefined() { // Check if the field exists in the jsObj
			if err := setFieldValue(field, fieldValue); err != nil {
				shared.ERR(fmt.Sprintf("Error setting field %s: %v", jsName, err))
			}
		} else {
			shared.Warn(fmt.Sprintf("Field %s does not exist in the provided JavaScript object.", jsName))
		}
	}
}

// Helper function to set the field value based on its kind
func setFieldValue(field reflect.Value, fieldValue js.Value) error {
	if !field.CanSet() {
		return fmt.Errorf("cannot set field value")
	}

	switch field.Kind() {
	case reflect.Int:
		field.SetInt(int64(fieldValue.Int()))
	case reflect.Float64:
		field.SetFloat(fieldValue.Float())
	case reflect.String:
		field.SetString(fieldValue.String())
	case reflect.Slice:
		// Special case for []byte - use optimized CopyBytesToGo
		if field.Type().Elem().Kind() == reflect.Uint8 || field.Type().Elem().Kind() == reflect.Int8 {
			length := fieldValue.Length()
			byteSlice := make([]byte, length)
			js.CopyBytesToGo(byteSlice, fieldValue)
			field.SetBytes(byteSlice)
			return nil
		}

		return fmt.Errorf("parsing other than Uint8Array or Uint8ClampedArra not supported")
	}
	return nil
}

// parseJsTag splits `gasm` struct tag into field name and option flags (e.g., "data,clamped").
func parseJsTag(tag string) (name string, opts map[string]bool) {
	opts = make(map[string]bool)
	if tag == "" {
		return
	}
	parts := strings.Split(tag, ",")
	shared.Info(parts)

	if len(parts) > 0 {
		name = parts[0]
	}

	for _, p := range parts[1:] {
		p = strings.TrimSpace(p)
		if p != "" {
			opts[p] = true
		}
	}
	return
}
