package storage

import (
	"bytes"
	"encoding/gob"

	"github.com/CaptainFallaway/XDH/internal"
)

func encodeObj(obj *internal.Session) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeObj(data []byte) (*internal.Session, error) {
	obj := new(internal.Session)
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
