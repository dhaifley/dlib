package dauth

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/dhaifley/dlib/ptypes"
)

// User represensts a single API user.
type User struct {
	ID    int64  `json:"id,omitempty"`
	User  string `json:"user,omitempty"`
	Pass  string `json:"pass,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// UserRow values represent a single row in the user table.
type UserRow struct {
	ID    int64
	User  string
	Pass  string
	Name  sql.NullString
	Email sql.NullString
}

// UserFind values are used to find user records in the database.
type UserFind struct {
	ID    *int64  `json:"id,omitempty"`
	User  *string `json:"user,omitempty"`
	Pass  *string `json:"pass,omitempty"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewUser initializes and returns a pointer to a new user value.
func NewUser(id int64, user, pass, name, email string) *User {
	return &User{
		ID:    id,
		User:  user,
		Pass:  pass,
		Name:  name,
		Email: email,
	}
}

// Equals tests for deep equality between user values
func (u *User) Equals(b *User) bool {
	switch {
	case u == nil || b == nil:
		return false
	case u == b:
		return true
	case u.ID != b.ID:
		return false
	case u.User != b.User:
		return false
	case u.Pass != b.Pass:
		return false
	case u.Name != b.Name:
		return false
	case u.Email != b.Email:
		return false
	default:
		return true
	}
}

// Copy returns an exact copy of the value.
func (u *User) Copy() User {
	var b User
	b.ID = u.ID
	b.User = u.User
	b.Pass = u.Pass
	b.Name = u.Name
	b.Email = u.Email
	return b
}

// String formats a user value as a JSON format string.
func (u *User) String() string {
	str, err := json.Marshal(u)
	if err != nil {
		return ""
	}

	return string(str)
}

// FromRequest populates a user value from a protobuf user message.
func (u *User) FromRequest(req *ptypes.UserRequest) error {
	if req.ID != 0 {
		u.ID = req.ID
	}

	if req.User != "" {
		u.User = req.User
	}

	if req.Pass != "" {
		u.Pass = req.Pass
	}

	if req.Name != "" {
		u.Name = req.Name
	}

	if req.Email != "" {
		u.Email = req.Email
	}

	return nil
}

// ToRequest returns a user protobuf message created from this value.
func (u *User) ToRequest() ptypes.UserRequest {
	req := ptypes.UserRequest{}
	req.ID = u.ID
	req.User = u.User
	req.Pass = u.Pass
	req.Name = u.Name
	req.Email = u.Email
	return req
}

// FromResponse populates a user value from a protobuf user message.
func (u *User) FromResponse(res *ptypes.UserResponse) error {
	if res.ID != 0 {
		u.ID = res.ID
	}

	if res.User != "" {
		u.User = res.User
	}

	if res.Pass != "" {
		u.Pass = res.Pass
	}

	if res.Name != "" {
		u.Name = res.Name
	}

	if res.Email != "" {
		u.Email = res.Email
	}

	return nil
}

// ToResponse returns a user protobuf message created from this value.
func (u *User) ToResponse() ptypes.UserResponse {
	res := ptypes.UserResponse{}
	res.ID = u.ID
	res.User = u.User
	res.Pass = u.Pass
	res.Name = u.Name
	res.Email = u.Email
	return res
}

// FromQueryValues populates a value from a query string values map.
func (u *User) FromQueryValues(vals url.Values) error {
	var err error
	if vals.Get("id") != "" {
		u.ID, err = strconv.ParseInt(vals.Get("id"), 10, 64)
		if err != nil {
			return err
		}
	}

	if vals.Get("user") != "" {
		u.User = vals.Get("user")
	}

	if vals.Get("pass") != "" {
		u.Pass = vals.Get("pass")
	}

	if vals.Get("name") != "" {
		u.Name = vals.Get("name")
	}

	if vals.Get("email") != "" {
		u.Email = vals.Get("email")
	}

	return nil
}

// ToUser returns a token value created from this UserRow value.
func (r UserRow) ToUser() User {
	u := User{}
	u.ID = r.ID
	u.User = r.User
	u.Pass = r.Pass
	if r.Name.Valid {
		u.Name = r.Name.String
	}

	if r.Email.Valid {
		u.Email = r.Email.String
	}

	return u
}

// FromUser populates a user find value from a user value.
func (f *UserFind) FromUser(u *User) error {
	if u.ID != 0 {
		f.ID = &u.ID
	}

	if u.User != "" {
		f.User = &u.User
	}

	if u.Pass != "" {
		f.Pass = &u.Pass
	}

	if u.Name != "" {
		f.Name = &u.Name
	}

	if u.Email != "" {
		f.Email = &u.Email
	}

	return nil
}

// FromUserRequest populates a user find value from a user protobuf request.
func (f *UserFind) FromUserRequest(r *ptypes.UserRequest) error {
	if r.ID != 0 {
		f.ID = &r.ID
	}

	if r.User != "" {
		f.User = &r.User
	}

	if r.Pass != "" {
		f.Pass = &r.Pass
	}

	if r.Name != "" {
		f.Name = &r.Name
	}

	if r.Email != "" {
		f.Email = &r.Email
	}

	return nil
}
