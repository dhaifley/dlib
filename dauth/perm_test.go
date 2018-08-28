package dauth

import (
	"net/url"
	"testing"

	"github.com/dhaifley/dlib/ptypes"
)

func TestPermEquals(t *testing.T) {
	cases := []struct {
		a        *Perm
		b        *Perm
		expected bool
	}{
		{
			a:        NewPerm(1, "test", "test"),
			b:        NewPerm(1, "test", "test"),
			expected: true,
		},
		{
			a:        NewPerm(1, "test", "test"),
			b:        NewPerm(1, "test2", "test"),
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

func TestPermCopy(t *testing.T) {
	a := Perm{
		ID:   1,
		Name: "test",
	}

	result := a.Copy()
	expected := Perm{
		ID:   1,
		Name: "test",
	}

	if !result.Equals(&expected) {
		t.Errorf("Expected perm: %v, got: %v", expected, result)
	}
}

func TestPermString(t *testing.T) {
	a := Perm{
		ID:   1,
		Name: "test",
	}

	expected := `{"id":1,"name":"test"}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %v, got: %v", expected, result)
	}
}

func TestPermFromRequest(t *testing.T) {
	req := ptypes.PermRequest{
		ID:   1,
		Name: "test",
	}

	dv := Perm{}
	if err := dv.FromRequest(&req); err != nil {
		t.Error(err)
	}

	if dv.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", dv.ID)
	}
}

func TestPermToRequest(t *testing.T) {
	dv := Perm{
		ID:   1,
		Name: "test",
	}

	msg := dv.ToRequest()
	if msg.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", msg.ID)
	}
}

func TestPermFromResponse(t *testing.T) {
	res := ptypes.PermResponse{
		ID:   1,
		Name: "test",
	}

	dv := Perm{}
	if err := dv.FromResponse(&res); err != nil {
		t.Error(err)
	}

	if dv.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", dv.ID)
	}
}

func TestPermToResponse(t *testing.T) {
	dv := Perm{
		ID:   1,
		Name: "test",
	}

	msg := dv.ToResponse()
	if msg.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", msg.ID)
	}
}

func TestPermRowToPerm(t *testing.T) {
	ur := PermRow{
		ID:   1,
		Name: "test",
	}

	u := ur.ToPerm()
	if u.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", u.Name)
	}
}

func TestPermFromQueryValues(t *testing.T) {
	vals := url.Values{}
	vals.Add("id", "1")
	vals.Add("name", "test")
	dv := Perm{}
	dv.FromQueryValues(vals)
	if dv.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", dv.ID)
	}
}

func TestPermFindFromPerm(t *testing.T) {
	u := Perm{
		ID:   1,
		Name: "test",
	}

	uf := PermFind{}
	err := uf.FromPerm(&u)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}
}

func TestPermFindFromPermRequest(t *testing.T) {
	ur := ptypes.PermRequest{
		ID:   1,
		Name: "test",
	}

	uf := PermFind{}
	err := uf.FromPermRequest(&ur)
	if err != nil {
		t.Error(err)
	}

	if *uf.ID != 1 {
		t.Errorf("ID expected: 1, got: %v", *uf.ID)
	}
}
