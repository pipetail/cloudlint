package datalayer

import (
	"log"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

// This particular test is testing if we can pass
// all implemented Adapters as Adapter interface

func someDummyCheck(a Adapter) {
	log.Printf("received adapter %s", reflect.TypeOf(a))
}

func TestPassStaticAsInterface(t *testing.T) {
	payload := `{}`
	a := NewAdapterStatic([]byte(payload))
	someDummyCheck(a)
}

func TestPassSDKAsInterface(t *testing.T) {
	policy := iam.Policy{}
	a := NewAdapterSDK(policy)
	someDummyCheck(a)
}
