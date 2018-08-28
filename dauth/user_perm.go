package dauth

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/dhaifley/dlib/ptypes"
)

// UserPerm values represenst a single API user permission assignment.
type UserPerm struct {
	ID     int64 `json:"id,omitempty"`
	UserID int64 `json:"user_id,omitempty"`
	PermID int64 `json:"perm_id,omitempty"`
}

// UserPermRow values represent a single row in the user_perm table.
type UserPermRow struct {
	ID     int64
	UserID int64
	PermID int64
}

// UserPermFind values are used to find user_perm records in the database.
type UserPermFind struct {
	ID     *int64 `json:"id,omitempty"`
	UserID *int64 `json:"user_id,omitempty"`
	PermID *int64 `json:"perm_id,omitempty"`
}

// NewUserPerm initializes and returns a pointer to a new
// user permission value.
func NewUserPerm(id, userID, permID int64) *UserPerm {
	return &UserPerm{
		ID:     id,
		UserID: userID,
		PermID: permID,
	}
}

// Equals tests for deep equality between user perm assignments.
func (up *UserPerm) Equals(b *UserPerm) bool {
	switch {
	case up == nil || b == nil:
		return false
	case *up != *b:
		return false
	default:
		return true
	}
}

// Copy returns an exact deep copy of the value.
func (up *UserPerm) Copy() UserPerm {
	return UserPerm{
		ID:     up.ID,
		UserID: up.UserID,
		PermID: up.PermID,
	}
}

// String formats a user_perm value as a JSON format string.
func (up *UserPerm) String() string {
	str, err := json.Marshal(up)
	if err != nil {
		return ""
	}

	return string(str)
}

// FromRequest populates this value from a protobuf request.
func (up *UserPerm) FromRequest(req *ptypes.UserPermRequest) error {
	up.ID = req.ID
	up.UserID = req.UserID
	up.PermID = req.PermID
	return nil

}

// ToRequest returns a protobuf request created from this value.
func (up *UserPerm) ToRequest() ptypes.UserPermRequest {
	return ptypes.UserPermRequest{
		ID:     up.ID,
		UserID: up.UserID,
		PermID: up.PermID,
	}
}

// FromResponse populates this value from a protobuf response.
func (up *UserPerm) FromResponse(res *ptypes.UserPermResponse) error {
	up.ID = res.ID
	up.UserID = res.UserID
	up.PermID = res.PermID
	return nil
}

// ToResponse returns a protobuf response created from this value.
func (up *UserPerm) ToResponse() ptypes.UserPermResponse {
	return ptypes.UserPermResponse{
		ID:     up.ID,
		UserID: up.UserID,
		PermID: up.PermID,
	}
}

// FromQueryValues populates this value from a query string map.
func (up *UserPerm) FromQueryValues(vals url.Values) error {
	var err error
	up.ID = 0
	if vals.Get("id") != "" {
		up.ID, err = strconv.ParseInt(vals.Get("id"), 10, 64)
		if err != nil {
			return err
		}
	}

	up.UserID = 0
	if vals.Get("user_id") != "" {
		up.UserID, err = strconv.ParseInt(vals.Get("user_id"), 10, 64)
		if err != nil {
			return err
		}
	}

	up.PermID = 0
	if vals.Get("perm_id") != "" {
		up.PermID, err = strconv.ParseInt(vals.Get("perm_id"), 10, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

// ToUserPerm returns a value created from this row value.
func (r UserPermRow) ToUserPerm() UserPerm {
	return UserPerm{
		ID:     r.ID,
		UserID: r.UserID,
		PermID: r.PermID,
	}
}

// FromUserPerm populates a user_perm find value from a user_perm value.
func (f *UserPermFind) FromUserPerm(up *UserPerm) error {
	f.ID = nil
	if up.ID != 0 {
		f.ID = &up.ID
	}

	f.UserID = nil
	if up.UserID != 0 {
		f.UserID = &up.UserID
	}

	f.PermID = nil
	if up.PermID != 0 {
		f.PermID = &up.PermID
	}

	return nil
}

// FromUserPermRequest populates a user_perm find value from a user_perm
// protobuf request.
func (f *UserPermFind) FromUserPermRequest(r *ptypes.UserPermRequest) error {
	f.ID = nil
	if r.ID != 0 {
		f.ID = &r.ID
	}

	f.UserID = nil
	if r.UserID != 0 {
		f.UserID = &r.UserID
	}

	f.PermID = nil
	if r.PermID != 0 {
		f.PermID = &r.PermID
	}

	return nil
}
