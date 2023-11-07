package codegen

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUserString(t *testing.T) {
	user := User{
		Name:   "John Doe",
		Age:    30,
		Emails: []string{"john.doe@example.com", "j.doe@company.com"},
	}

	expected := "Name: John Doe, Age: 30, Emails: [john.doe@example.com j.doe@company.com]"
	if user.String() != expected {
		t.Errorf("User.String() = %v, want %v", user.String(), expected)
	}
}

func TestUserValidate(t *testing.T) {
	tests := []struct {
		user    User
		wantErr bool
	}{
		{User{"John Doe", 30, []string{"john.doe@example.com"}}, false},
		{User{"", 30, []string{"john.doe@example.com"}}, true},
		{User{"John Doe", -1, []string{"john.doe@example.com"}}, true},
		{User{"John Doe", 30, []string{"invalid-email"}}, true},
	}

	for _, tt := range tests {
		err := tt.user.Validate()
		if (err != nil) != tt.wantErr {
			t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"john.doe@example.com", true},
		{"j.doe@company.com", true},
		{"invalid-email", false},
		{"john.doe@example", false},
	}

	for _, tt := range tests {
		if isValidEmail(tt.email) != tt.valid {
			t.Errorf("isValidEmail(%v) = %v, want %v", tt.email, !tt.valid, tt.valid)
		}
	}
}

func TestUserJSONEncodingDecoding(t *testing.T) {
	user := User{
		Name:   "John Doe",
		Age:    30,
		Emails: []string{"john.doe@example.com", "j.doe@company.com"},
	}

	// Marshal the User struct to JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("json.Marshal() failed with '%s'", err)
	}

	// Unmarshal the JSON back into a User struct
	var decodedUser User
	err = json.Unmarshal(jsonBytes, &decodedUser)
	if err != nil {
		t.Fatalf("json.Unmarshal() failed with '%s'", err)
	}

	// Use reflect.DeepEqual to compare the structs with slices
	if !reflect.DeepEqual(decodedUser, user) {
		t.Errorf("Decoded User does not match the original User struct")
	}
}
