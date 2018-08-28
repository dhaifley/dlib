package dauth

import (
	"net/url"
	"testing"
	"time"

	"github.com/dhaifley/dlib/ptypes"
)

func TestTokenEquals(t *testing.T) {
	cases := []struct {
		a        *Token
		b        *Token
		expected bool
	}{
		{
			a:        NewToken(1, "test", 1, nil, nil),
			b:        NewToken(1, "test", 1, nil, nil),
			expected: true,
		},
		{
			a:        NewToken(1, "test", 1, nil, nil),
			b:        NewToken(1, "test2", 1, nil, nil),
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

func TestTokenCopy(t *testing.T) {
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	a := Token{
		ID:      int64(1),
		Token:   "test",
		Created: &d,
		Expires: &d,
	}

	result := a.Copy()
	expected := Token{
		ID:      int64(1),
		Token:   "test",
		Created: &d,
		Expires: &d,
	}

	if !result.Equals(&expected) {
		t.Errorf("Expected user: %v, got: %v", expected, result)
	}
}

func TestTokenString(t *testing.T) {
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	a := Token{
		ID:      int64(1),
		Token:   "test",
		UserID:  int64(1),
		Created: &d,
		Expires: &d,
	}

	expected := `{"id":1,"token":"test","user_id":1,"created":"1983-02-02T00:00:00-05:00","expires":"1983-02-02T00:00:00-05:00"}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %v, got: %v", expected, result)
	}
}

func TestTokenFromRequest(t *testing.T) {
	req := ptypes.TokenRequest{
		ID:    int64(1),
		Token: "test",
	}

	dv := Token{}
	if err := dv.FromRequest(&req); err != nil {
		t.Error(err)
	}

	expected := int64(1)
	if dv.ID != expected {
		t.Errorf("ID expected: %v, got: %v", expected, dv.ID)
	}

	exp := "test"
	if dv.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, dv.Token)
	}
}

func TestTokenToRequest(t *testing.T) {
	dv := Token{
		ID:    int64(1),
		Token: "test",
	}

	msg := dv.ToRequest()
	expected := int64(1)
	if msg.ID != expected {
		t.Errorf("ID expected: %v, got: %v", expected, msg.ID)
	}

	exp := "test"
	if msg.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, msg.Token)
	}
}

func TestTokenFromResponse(t *testing.T) {
	res := ptypes.TokenResponse{
		ID:    int64(1),
		Token: "test",
	}

	dv := Token{}
	if err := dv.FromResponse(&res); err != nil {
		t.Error(err)
	}

	expected := int64(1)
	if dv.ID != expected {
		t.Errorf("ID expected: %v, got: %v", expected, dv.ID)
	}

	exp := "test"
	if dv.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, dv.Token)
	}
}

func TestTokenToResponse(t *testing.T) {
	dv := Token{
		ID:    int64(1),
		Token: "test",
	}

	msg := dv.ToResponse()
	expected := int64(1)
	if msg.ID != expected {
		t.Errorf("ID expected: %v, got: %v", expected, msg.ID)
	}

	exp := "test"
	if msg.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, msg.Token)
	}
}

func TestTokenRowFromToken(t *testing.T) {
	tk := Token{
		ID:    1,
		Token: "test",
	}

	tr := TokenRow{}
	err := tr.FromToken(&tk)
	if err != nil {
		t.Error(err)
	}

	if tr.ID != int64(1) {
		t.Errorf("ID expected: 1, got: %v", tk.ID)
	}

	exp := "test"
	if tr.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, tk.Token)
	}
}

func TestTokenRowToToken(t *testing.T) {
	tr := TokenRow{
		ID:    1,
		Token: "test",
	}

	tk := tr.ToToken()
	if tk.ID != int64(1) {
		t.Errorf("ID expected: 1, got: %v", tk.ID)
	}

	exp := "test"
	if tk.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, tk.Token)
	}
}

func TestTokenFromQueryValues(t *testing.T) {
	vals := url.Values{}
	vals.Add("id", "1")
	vals.Add("token", "test")
	dv := Token{}
	dv.FromQueryValues(vals)
	expected := int64(1)
	if dv.ID != expected {
		t.Errorf("ID expected: %v, got: %v", expected, dv.ID)
	}

	exp := "test"
	if dv.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, dv.Token)
	}
}

func TestTokenFindFromToken(t *testing.T) {
	tk := Token{
		ID:    1,
		Token: "test",
	}

	tf := TokenFind{}
	err := tf.FromToken(&tk)
	if err != nil {
		t.Error(err)
	}

	if *tf.ID != int64(1) {
		t.Errorf("ID expected: 1, got: %v", *tf.ID)
	}

	exp := "test"
	if *tf.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, *tf.Token)
	}
}

func TestTokenFindFromTokenRequest(t *testing.T) {
	tr := ptypes.TokenRequest{
		ID:    1,
		Token: "test",
	}

	tf := TokenFind{}
	err := tf.FromTokenRequest(&tr)
	if err != nil {
		t.Error(err)
	}

	if *tf.ID != int64(1) {
		t.Errorf("ID expected: 1, got: %v", *tf.ID)
	}

	exp := "test"
	if *tf.Token != exp {
		t.Errorf("Value expected: %v, got: %v", exp, *tf.Token)
	}
}
