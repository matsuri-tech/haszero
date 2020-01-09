package haszero

import (
	"errors"
	"fmt"
	"reflect"
)

func HasZero(s interface{}) error {
	val := reflect.ValueOf(s)
	name := reflect.TypeOf(s).Name()
	if val.Kind().String() != "struct" {
		return errors.New(fmt.Sprintf("%s is not struct", name))
	}
	return hasZeroSub(val)
}

func hasZeroSub(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		typeField := v.Type().Field(i)
		if typeField.Type.Kind().String() == "struct" {
			if err := hasZeroSub(valueField); err != nil {
				return err
			}
		}
		if valueField.IsZero() {
			return errors.New(fmt.Sprintf("%s is zero value", typeField.Name))
		}
	}
	return nil
}
