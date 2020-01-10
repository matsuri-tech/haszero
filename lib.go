package haszero

import (
	"errors"
	"fmt"
	"reflect"
)

func HasZero(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind().String() != "struct" {
		name := reflect.TypeOf(s).Name()
		return errors.New(fmt.Sprintf("%s is not struct", name))
	}
	return hasZeroSub(val)
}

func hasZeroSub(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		valueField := func() reflect.Value {
			val := v.Field(i)
			if val.Kind() == reflect.Ptr {
				return v.Field(i).Elem()
			}
			return val
		}()

		if valueField.Kind() == reflect.Struct {
			if err := hasZeroSub(valueField); err != nil {
				return err
			}
		}
		if !valueField.IsValid() {
			return errors.New(fmt.Sprintf("%s is invalid value", typeField.Name))
		}
		if valueField.IsZero() {
			return errors.New(fmt.Sprintf("%s is zero value", typeField.Name))
		}
	}
	return nil
}
