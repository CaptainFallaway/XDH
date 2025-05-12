package storage

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

func encodeObj[T any](obj *T) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeObj[T any](data []byte) (*T, error) {
	obj := new(T)
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

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
