package schema

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type UserErrors struct {
	UsernameError    error
	PasswordError    error
	CredentialsError error
	EmailError       error
	PhoneError       error
}

func NewUserErrors() *UserErrors {
	return &UserErrors{
		UsernameError:    nil,
		PasswordError:    nil,
		CredentialsError: nil,
		EmailError:       nil,
		PhoneError:       nil,
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique().MinLen(3),
		field.String("password").MinLen(7),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("role"),
		field.String("phonenumber").Optional().Unique().MinLen(4),
		field.String("email").Optional().Unique().Validate(ValidateUserEmail),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Optional(),
		field.String("token").Default(""),
		field.String("refresh_token").Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("settings", UserSettings.Type).Unique(),
	}
}

func ValidateUserEmail(email string) error {
	// _, err := mail.ParseAddress(email)
	// if err != nil {
	// 	log.Panic("email is not valid")
	// }

	return nil
}

func ValidateUserUsername(username string) error {
	var err error
	if len(username) < 3 {
		err = errors.New("username less then 3 symbols")
		return err
	}

	if len(username) > 20 {
		err = errors.New("username less more than 20 symbos")
		return err
	}

	username = strings.ToLower(username)
	isUsernameCorrect, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username)
	if !isUsernameCorrect {
		err = errors.New("username must by only latin and nums")
		return err
	}

	return nil
}

func ValidateUserPassword(password string) error {
	var err error
	if len(password) < 6 {
		err = errors.New("password fields is less then 6 symbols")
		return err
	}

	if len(password) > 20 {
		err = errors.New("password fields is too big (max20)")
		return err
	}

	secure := true
	tests := [][]string{
		{".{8,}", "8 symbols minimum"},
		{"[a-z]", "must contain lower case letter"},
		{"[A-Z]", "must contain upper case letter"},
		{"[0-9]", "must contain number"},
		{"[^\\d\\w]", "must contain special symbol"},
	}
	btest := []string{"", "password don't match criteria"}

	for _, test := range tests {
		t, _ := regexp.MatchString(test[0], password)
		if !t {
			btest = test
			secure = false
			break
		}
	}

	if !secure {
		err = errors.New(btest[1])
		return err
	}

	return nil
}
