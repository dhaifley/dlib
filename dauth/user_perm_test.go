package dauth

import (
	"net/url"
	"testing"

	"github.com/dhaifley/dlib/ptypes"
)

func TestUserPermEquals(t *testing.T) {
	cases := []struct {
		a        *UserPerm
		b        *UserPerm
		expected bool
	}{
		{
			a:        NewUserPerm(1, 1, 1),
			b:        NewUserPerm(1, 1, 1),
			expected: true,
		},
		{
			a:        NewUserPerm(1, 1, 1),
			b:        NewUserPerm(1, 1, 2),
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

func TestUserPermCopy(t *testing.T) {
	a := UserPerm{
		ID:     1,
		UserID: 1,
	}

	result := a.Copy()
	expected := UserPerm{
		ID:     1,
		UserID: 1,
	}

	if !result.Equals(&expected) {
		t.Errorf("Expected perm: %v, got: %v", expected, result)
	}
}

func TestUserPermString(t *testing.T) {
	a := UserPerm{
		ID:     1,
		UserID: 1,
	}

	expected := `{"id":1,"user_id":1}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %v, got: %v", expected, result)
	}
}

func TestUserPermFromRequest(t *testing.T) {
	req := ptypes.UserPermRequest{
		ID:     1,
		UserID: 1,
	}

	dv := UserPerm{}
	if err := dv.FromRequest(&req); err != nil {
		t.Error(err)
	}

	if dv.UserID != 1 {
		t.Errorf("UserID expected: 1, got: %v", dv.UserID)
	}
}

func TestUserPermToRequest(t *testing.T) {
	dv := UserPerm{
		ID:     1,
		UserID: 1,
	}

	msg := dv.ToRequest()
	if msg.UserID != 1 {
		t.Errorf("UserID expected: 1, got: %v", msg.UserID)
	}
}

func TestUserPermFromResponse(t *testing.T) {
	res := ptypes.UserPermResponse{
		ID:     1,
		UserID: 1,
	}

	dv := UserPerm{}
	if err := dv.FromResponse(&res); err != nil {
		t.Error(err)
	}

	if dv.UserID != 1 {
		t.Errorf("UserID expected: 1, got: %v", dv.UserID)
	}
}

func TestUserPermToResponse(t *testing.T) {
	dv := UserPerm{
		ID:     1,
		UserID: 1,
	}

	msg := dv.ToResponse()
	if msg.UserID != 1 {
		t.Errorf("UserID expected: 1, got: %v", msg.UserID)
	}
}

func TestUserPermRowToUserPerm(t *testing.T) {
	ur := UserPermRow{
		ID:     1,
		UserID: 1,
	}

	u := ur.ToUserPerm()
	if u.UserID != 1 {
		t.Errorf("UserID expected: 1, got: %v", u.UserID)
	}
}

func TestUserPermFromQueryValues(t *testing.T) {
	vals := url.Values{}
	vals.Add("id", "1")
	vals.Add("user_id", "1")
	dv := UserPerm{}
	dv.FromQueryValues(vals)
	if dv.UserID != int64(1) {
		t.Errorf("UserID expected: 1, got: %v", dv.UserID)
	}

	if dv.ID != int64(1) {
		t.Errorf("ID expected: 1, got: %v", dv.ID)
	}
}

func TestUserPermFindFromUserPerm(t *testing.T) {
	u := UserPerm{
		ID:     1,
		UserID: 1,
	}

	uf := UserPermFind{}
	err := uf.FromUserPerm(&u)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}

	if *uf.UserID != 1 {
		t.Errorf("Value expected: 1, got: %v", *uf.UserID)
	}
}

func TestUserPermFindFromUserPermRequest(t *testing.T) {
	r := ptypes.UserPermRequest{
		ID:     1,
		UserID: 1,
	}

	uf := UserPermFind{}
	err := uf.FromUserPermRequest(&r)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}

	if *uf.UserID != 1 {
		t.Errorf("Value expected: 1, got: %v", *uf.UserID)
	}
}
