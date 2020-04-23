package worker

import (
	"encoding/json"
	"testing"
)

func TestFilterAttachedVolumes(t *testing.T) {
	rawAttribete1 := `
	{
		"InstanceId": "i-1234567890abcdef0",
		"ebsOptimized": true
	}`

	rawAttribete2 := `
	{
		"InstanceId": "i-1234567890abcdef0",
		"ebsOptimized": false
	}`

	attribete1 := &InstanceAtribute{}
	var _ = json.Unmarshal([]byte(rawAttribete1), &attribete1)

	attribete2 := &InstanceAtribute{}
	_ = json.Unmarshal([]byte(rawAttribete2), &attribete2)

	tables := []struct {
		x *InstanceAtribute
		y bool
	}{
		{attribete1, false},
		{attribete2, true},
	}

	for _, table := range tables {
		optimized := checkEbsNotOptimized([]*InstanceAtribute{table.x})
		if optimized != table.y {
			t.Errorf("checkEbsNotOptimized does not return correct value, got: %t, want: %t.", optimized, table.y)
		}
	}
}
