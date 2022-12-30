package engagebay

import (
	"testing"
)

var client *API

func init() {
	client = New("s67u2j6mou7ovrbgoqfkcq0l8n")
}

func TestAPI_CreateContact(t *testing.T) {
	params := &CreateContactParams{}
	params.SetName("unit_test")
	params.SetEmail("unit_test@test.com")
	if _, err := client.CreateContact(params); err != nil {
		t.Fatal(err)
	}
}
