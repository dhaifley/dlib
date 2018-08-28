package dauth

import (
	"net/url"
	"testing"

	"github.com/dhaifley/dlib/ptypes"
)

func TestUserEquals(t *testing.T) {
	cases := []struct {
		a        *User
		b        *User
		expected bool
	}{
		{
			a:        NewUser(1, "test", "test", "", ""),
			b:        NewUser(1, "test", "test", "", ""),
			expected: true,
		},
		{
			a:        NewUser(1, "test", "test", "", ""),
			b:        NewUser(1, "test", "test2", "", ""),
			expected: false,
		},
	}

	for _, c := range cases {
		result := c.a.Equals(c.b)
		if result != c.expected {
			t.Errorf("Expected bool: %v, got: %v", c.expected, result)
		}
	}
}

func TestUserCopy(t *testing.T) {
	a := User{
		ID:   1,
		User: "test",
		Pass: "test",
	}

	result := a.Copy()
	expected := User{
		ID:   1,
		User: "test",
		Pass: "test",
	}

	if !result.Equals(&expected) {
		t.Errorf("Expected user: %v, got: %v", expected, result)
	}
}

func TestUserString(t *testing.T) {
	a := User{
		ID:   1,
		User: "test",
		Pass: "test",
	}

	expected := `{"id":1,"user":"test","pass":"test"}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %v, got: %v", expected, result)
	}
}

func TestUserFromRequest(t *testing.T) {
	req := ptypes.UserRequest{
		ID:   1,
		User: "test",
	}

	dv := User{}
	if err := dv.FromRequest(&req); err != nil {
		t.Error(err)
	}

	if dv.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", dv.User)
	}
}

func TestUserToRequest(t *testing.T) {
	dv := User{
		ID:   1,
		User: "test",
	}

	msg := dv.ToRequest()
	if msg.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", msg.ID)
	}
}

func TestUserFromResponse(t *testing.T) {
	res := ptypes.UserResponse{
		ID:   1,
		User: "test",
	}

	dv := User{}
	if err := dv.FromResponse(&res); err != nil {
		t.Error(err)
	}

	if dv.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", dv.User)
	}
}

func TestUserToResponse(t *testing.T) {
	dv := User{
		ID:   1,
		User: "test",
	}

	msg := dv.ToResponse()
	if msg.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", msg.User)
	}
}

func TestUserRowToUser(t *testing.T) {
	ur := UserRow{
		ID:   1,
		User: "test",
	}

	u := ur.ToUser()
	if u.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", u.ID)
	}
}

func TestUserFromQueryValues(t *testing.T) {
	vals := url.Values{}
	vals.Add("user", "test")
	vals.Add("pass", "test")
	dv := User{}
	dv.FromQueryValues(vals)
	exp := "test"
	if dv.User != exp {
		t.Errorf("Value expected: %v, got: %v", exp, dv.User)
	}

	if dv.Pass != exp {
		t.Errorf("Value expected: %v, got: %v", exp, dv.Pass)
	}
}

func TestUserFindFromUser(t *testing.T) {
	u := User{
		ID:   1,
		User: "test",
	}

	uf := UserFind{}
	err := uf.FromUser(&u)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}

	exp := "test"
	if *uf.User != exp {
		t.Errorf("Value expected: %v, got: %v", exp, *uf.User)
	}
}

func TestUserFindFromUserRequest(t *testing.T) {
	r := ptypes.UserRequest{
		ID:   1,
		User: "test",
	}

	uf := UserFind{}
	err := uf.FromUserRequest(&r)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}

	exp := "test"
	if *uf.User != exp {
		t.Errorf("Value expected: %v, got: %v", exp, *uf.User)
	}
}
