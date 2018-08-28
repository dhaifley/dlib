package dauth

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/dhaifley/dlib"
	"github.com/dhaifley/dlib/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Token values represenst a single API token.
type Token struct {
	ID      int64      `json:"id,omitempty"`
	Token   string     `json:"token,omitempty"`
	UserID  int64      `json:"user_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
}

// TokenRow values represent a single row in the token table.
type TokenRow struct {
	ID      int64
	Token   string
	UserID  int64
	Created dlib.NullTime
	Expires dlib.NullTime
}

// TokenFind values are used to find token records in the database.
type TokenFind struct {
	ID      *int64     `json:"id,omitempty"`
	Token   *string    `json:"token,omitempty"`
	UserID  *int64     `json:"user_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
	Start   *time.Time `json:"start,omitempty"`
	End     *time.Time `json:"end,omitempty"`
	Old     *time.Time `json:"old,omitempty"`
}

// NewToken initializes and returns a pointer to a new token value.
func NewToken(id int64, token string, userID int64, created, expires *time.Time) *Token {
	return &Token{
		ID:      id,
		Token:   token,
		UserID:  userID,
		Created: created,
		Expires: expires,
	}
}

// Equals tests for deep equality between Token values
func (t *Token) Equals(b *Token) bool {
	switch {
	case t == nil || b == nil:
		return false
	case t == b:
		return true
	case t.ID != b.ID:
		return false
	case t.Token != b.Token:
		return false
	case t.UserID != b.UserID:
		return false
	case (t.Created == nil && b.Created != nil) || (t.Created != nil && b.Created == nil):
		return false
	case t.Created != nil && b.Created != nil && *t.Created != *b.Created:
		return false
	case (t.Expires == nil && b.Expires != nil) || (t.Expires != nil && b.Expires == nil):
		return false
	case t.Expires != nil && b.Expires != nil && *t.Expires != *b.Expires:
		return false
	default:
		return true
	}
}

// Copy returns an exact copy of the value.
func (t *Token) Copy() Token {
	var b Token
	b.ID = t.ID
	b.Token = t.Token
	b.UserID = t.UserID
	if t.Created != nil {
		d := *t.Created
		b.Created = &d
	}

	if t.Expires != nil {
		d := *t.Expires
		b.Expires = &d
	}

	return b
}

// String formats an Token value as a JSON format string.
func (t *Token) String() string {
	str, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(str)
}

// FromRequest populates a token value from a protobuf token message.
func (t *Token) FromRequest(req *ptypes.TokenRequest) error {
	t.ID = req.ID
	t.Token = req.Token
	t.UserID = req.UserID
	if req.Created != nil {
		tt := time.Unix(req.Created.Seconds, 0)
		t.Created = &tt
	}

	if req.Expires != nil {
		tt := time.Unix(req.Expires.Seconds, 0)
		t.Expires = &tt
	}

	return nil
}

// ToRequest returns a token protobuf message created from this value.
func (t *Token) ToRequest() ptypes.TokenRequest {
	req := ptypes.TokenRequest{}
	req.ID = t.ID
	req.Token = t.Token
	req.UserID = t.UserID
	if t.Created != nil {
		req.Created = &timestamp.Timestamp{Seconds: t.Created.Unix(), Nanos: 0}
	}

	if t.Expires != nil {
		req.Expires = &timestamp.Timestamp{Seconds: t.Expires.Unix(), Nanos: 0}
	}

	return req
}

// FromResponse populates a token value from a protobuf token message.
func (t *Token) FromResponse(res *ptypes.TokenResponse) error {
	t.ID = res.ID
	t.Token = res.Token
	t.UserID = res.UserID
	if res.Created != nil {
		tt := time.Unix(res.Created.Seconds, 0)
		t.Created = &tt
	}

	if res.Expires != nil {
		tt := time.Unix(res.Expires.Seconds, 0)
		t.Expires = &tt
	}

	return nil
}

// ToResponse returns a token protobuf message created from this value.
func (t *Token) ToResponse() ptypes.TokenResponse {
	res := ptypes.TokenResponse{}
	res.ID = t.ID
	res.Token = t.Token
	res.UserID = t.UserID
	if t.Created != nil {
		res.Created = &timestamp.Timestamp{Seconds: t.Created.Unix(), Nanos: 0}
	}

	if t.Expires != nil {
		res.Expires = &timestamp.Timestamp{Seconds: t.Expires.Unix(), Nanos: 0}
	}

	return res
}

// FromQueryValues populates a value from a query string values map.
func (t *Token) FromQueryValues(vals url.Values) error {
	var err error
	t.ID = 0
	if vals.Get("id") != "" {
		if t.ID, err = strconv.ParseInt(vals.Get("id"), 10, 64); err != nil {
			return err
		}
	}

	t.Token = ""
	if vals.Get("token") != "" {
		t.Token = vals.Get("token")
	}

	t.UserID = 0
	if vals.Get("user") != "" {
		if t.UserID, err = strconv.ParseInt(vals.Get("id"), 10, 64); err != nil {
			return err
		}
	}

	if vals.Get("created") != "" {
		pt, err := time.ParseInLocation("2006-01-02T15:04:05-0700",
			vals.Get("created"), time.Local)
		if err != nil {
			return err
		}

		t.Created = &pt
	}

	if vals.Get("expires") != "" {
		pt, err := time.ParseInLocation("2006-01-02T15:04:05-0700",
			vals.Get("expires"), time.Local)
		if err != nil {
			return err
		}

		t.Expires = &pt
	}

	return nil
}

// FromToken populates a token row value from a token value.
func (r *TokenRow) FromToken(t *Token) error {
	r.ID = t.ID
	r.Token = t.Token
	r.UserID = t.UserID
	if t.Created != nil {
		r.Created = dlib.NullTime{Valid: true, Time: *t.Created}
	}

	if t.Expires != nil {
		r.Expires = dlib.NullTime{Valid: true, Time: *t.Expires}
	}

	return nil
}

// ToToken returns a token value created from this TokenRow value.
func (r *TokenRow) ToToken() Token {
	t := Token{}
	t.ID = r.ID
	t.Token = r.Token
	t.UserID = r.UserID
	if r.Created.Valid {
		tt := r.Created.Time
		t.Created = &tt
	}

	if r.Expires.Valid {
		tt := r.Expires.Time
		t.Expires = &tt
	}

	return t
}

// FromToken populates a token find value from a token value.
func (f *TokenFind) FromToken(t *Token) error {
	f.ID = nil
	if t.ID != 0 {
		f.ID = &t.ID
	}

	f.Token = nil
	if t.Token != "" {
		f.Token = &t.Token
	}

	f.UserID = nil
	if t.UserID != 0 {
		f.UserID = &t.UserID
	}

	f.Created = nil
	if t.Created != nil {
		dt := time.Unix(t.Created.Unix(), 0)
		f.Created = &dt
	}

	f.Expires = nil
	if t.Expires != nil {
		dt := time.Unix(t.Expires.Unix(), 0)
		f.Expires = &dt
	}

	return nil
}

// FromTokenRequest populates a token find value from a token protobuf request.
func (f *TokenFind) FromTokenRequest(r *ptypes.TokenRequest) error {
	f.ID = nil
	if r.ID != 0 {
		f.ID = &r.ID
	}

	f.Token = nil
	if r.Token != "" {
		f.Token = &r.Token
	}

	f.UserID = nil
	if r.UserID != 0 {
		f.UserID = &r.UserID
	}

	f.Created = nil
	if r.Created != nil {
		dt := time.Unix(r.Created.Seconds, 0)
		f.Created = &dt
	}

	f.Expires = nil
	if r.Expires != nil {
		dt := time.Unix(r.Expires.Seconds, 0)
		f.Expires = &dt
	}

	f.Start = nil
	if r.Start != nil {
		dt := time.Unix(r.Start.Seconds, 0)
		f.Start = &dt
	}

	f.End = nil
	if r.End != nil {
		dt := time.Unix(r.End.Seconds, 0)
		f.End = &dt
	}

	f.Old = nil
	if r.Old != nil {
		dt := time.Unix(r.Old.Seconds, 0)
		f.Old = &dt
	}

	return nil
}
