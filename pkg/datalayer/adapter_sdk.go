package datalayer

import (
	"encoding/json"

	"github.com/getlantern/deepcopy"
)

type AdapterSDK struct {
	SDKObject interface{}
}

func (a *AdapterSDK) Get(output interface{}) error {
	return deepcopy.Copy(output, a.SDKObject)
}
func (a *AdapterSDK) GetJSON() ([]byte, error) {
	return json.Marshal(a.SDKObject)
}

func NewAdapterSDK(object interface{}) *AdapterSDK {
	return &AdapterSDK{
		SDKObject: object,
	}
}
