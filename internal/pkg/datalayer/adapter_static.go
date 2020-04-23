package datalayer

import "encoding/json"

type AdapterStatic struct {
	Payload []byte
}

func (a *AdapterStatic) Get(output interface{}) error {
	return json.Unmarshal(a.Payload, output)
}

func (a *AdapterStatic) GetJSON() ([]byte, error) {
	return a.Payload, nil
}

func NewAdapterStatic(payload []byte) *AdapterStatic {
	return &AdapterStatic{
		Payload: payload,
	}
}
