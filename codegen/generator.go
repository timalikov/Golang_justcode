package codegen

import (
	"errors"
	"fmt"
	"regexp"
)

//easyjson:json
type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Emails []string `json:"emails"`
}

//easyjson:json
type Profile struct {
	UserID      string `json:"user_id"`
	Occupation  string `json:"occupation"`
	Description string `json:"description"`
}

// String returns a string representation of the User.
func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Emails: %v", u.Name, u.Age, u.Emails)
}

// Validate checks the fields of User and returns an error if any are invalid.
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("user name cannot be empty")
	}
	if u.Age <= 0 {
		return errors.New("user age must be positive")
	}
	for _, email := range u.Emails {
		if !isValidEmail(email) {
			return fmt.Errorf("invalid email: %s", email)
		}
	}
	return nil
}

// isValidEmail validates the email format.
func isValidEmail(email string) bool {
	// regex pattern for checking email
	const emailPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}
