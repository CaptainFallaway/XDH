package storage

import (
	"bytes"
	"encoding/gob"
)

func encodeObj(obj *Session) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeObj(data []byte) (*Session, error) {
	obj := new(Session)
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
