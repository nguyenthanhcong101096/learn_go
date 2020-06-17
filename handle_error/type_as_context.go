package main

import (
	"fmt"
	"reflect"
)

// An UnmarshalTypeError describes a JSON value that was not appropriate for
// a value of a specific Go type.
type UnmarshalTypeError struct {
	Value string       // description of JSON value
	Type  reflect.Type // loại giá trị Go không thể gán cho
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// In the implementation, we are validating all the fields are being used in the error message
func (e *UnmarshalTypeError) Error() string {
	return "json: cannot urmarshal" + e.Value + "into Go value of type" + e.Type.String()
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unrmarshal(non-pointer" + e.Type.String() + ")"
	}

	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

type user struct {
	Name int
}

func Unmarshal(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}

func main() {
	var u user
	err := Unmarshal(u) // Run with a value and pointer.
	if err != nil {
		// This is a special type assertion that only works on the switch.
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		return
	}

	fmt.Println("Name:", u.Name)
}
