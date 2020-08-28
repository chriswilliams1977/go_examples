package types

import (
	"reflect"
	"strconv"
)

// Any formats any value as a string.
func CheckType(value interface{}) string {
	return typeName(reflect.ValueOf(value))
}

func typeName(v reflect.Value) string {
	return v.Type().String()
}

// formatAtom formats a value without inspecting its internal structure.
func formatType(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
