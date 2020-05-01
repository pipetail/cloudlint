package datalayer

import "testing"

type simpleJSONSNameSurname struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func TestSimpleJSONObject(t *testing.T) {
	payload := `
	{
		"name": "Stepan",
		"surname": "Vrany"
	}
	`

	a := NewAdapterStatic([]byte(payload))
	nameSurname := simpleJSONSNameSurname{}
	err := a.Get(&nameSurname)
	if err != nil {
		t.Errorf("could not obtain struct: %s", err)
	}

	if nameSurname.Name != "Stepan" || nameSurname.Surname != "Vrany" {
		t.Errorf(
			"Name (%s) or Surname (%s) does not match expected values",
			nameSurname.Name,
			nameSurname.Surname,
		)
	}
}
