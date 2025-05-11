package storage

import "reflect"

// convertFields copies the fields from src to dest if they have the same name and type.
func convertFields(src, dest any) {
	sVal := reflect.ValueOf(src).Elem()
	dVal := reflect.ValueOf(dest).Elem()

	for i := 0; i < sVal.NumField(); i++ {
		sField := sVal.Type().Field(i)
		dField, exists := dVal.Type().FieldByName(sField.Name)
		if !exists {
			continue
		}

		if sField.Type == dField.Type {
			dVal.FieldByName(sField.Name).Set(sVal.Field(i))
		}
	}
}
