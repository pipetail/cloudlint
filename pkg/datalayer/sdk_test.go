package datalayer

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

func TestSimpleSDKObject(t *testing.T) {
	payload := `
	{
		"Arn":"arn:aws:iam::aws:policy/AWSLambdaExecute",
		"AttachmentCount":0,
		"CreateDate":"2015-02-06T18:40:46Z",
		"DefaultVersionId":"v1",
		"Description":"Provides Put, Get access to S3 and full access to CloudWatch Logs.",
		"IsAttachable":true,
		"Path":"/",
		"PermissionsBoundaryUsageCount":0,
		"PolicyId":"ANPAJE5FX7FQZSU5XAKGO",
		"PolicyName":"AWSLambdaExecute",
		"UpdateDate":"2015-02-06T18:40:46Z"
	}
	`

	policy := iam.Policy{}
	err := json.Unmarshal([]byte(payload), &policy)
	if err != nil {
		t.Errorf("could not parse input json: %s", err)
	}

	a := NewAdapterSDK(policy)
	policyNew := iam.Policy{}

	err = a.Get(&policyNew)
	if err != nil {
		t.Errorf("could not obtain struct: %s", err)
	}

	if !reflect.DeepEqual(policy, policyNew) {
		t.Error("object from adapter does not match the input")
	}
}
