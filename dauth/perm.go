package dauth

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/dhaifley/dlib/ptypes"
)

// Perm values represenst a single API permissions.
type Perm struct {
	ID      int64  `json:"id,omitempty"`
	Service string `json:"service,omitempty"`
	Name    string `json:"name,omitempty"`
}

// PermRow values represent a single row in the perm table.
type PermRow struct {
	ID      int64
	Service string
	Name    string
}

// PermFind values are used to find perm records in the database.
type PermFind struct {
	ID      *int64  `json:"id,omitempty"`
	Service *string `json:"service,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// NewPerm initializes and returns a pointer to a new permission value.
func NewPerm(id int64, service, name string) *Perm {
	return &Perm{
		ID:      id,
		Service: service,
		Name:    name,
	}
}

// Equals tests for deep equality between perm values.
func (p *Perm) Equals(b *Perm) bool {
	switch {
	case p == nil || b == nil:
		return false
	case *p != *b:
		return false
	default:
		return true
	}
}

// Copy returns an exact copy of the value.
func (p *Perm) Copy() Perm {
	return Perm{
		ID:      p.ID,
		Service: p.Service,
		Name:    p.Name,
	}
}

// String formats a perm value as a JSON format string.
func (p *Perm) String() string {
	str, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(str)
}

// FromRequest populates this value from a protobuf request.
func (p *Perm) FromRequest(req *ptypes.PermRequest) error {
	p.ID = req.ID
	p.Service = req.Service
	p.Name = req.Name
	return nil
}

// ToRequest returns a protobuf request created from this value.
func (p *Perm) ToRequest() ptypes.PermRequest {
	return ptypes.PermRequest{
		ID:      p.ID,
		Service: p.Service,
		Name:    p.Name,
	}
}

// FromResponse populates this value from a protobuf response.
func (p *Perm) FromResponse(res *ptypes.PermResponse) error {
	p.ID = res.ID
	p.Service = res.Service
	p.Name = res.Name
	return nil
}

// ToResponse returns a protobuf response created from this value.
func (p *Perm) ToResponse() ptypes.PermResponse {
	return ptypes.PermResponse{
		ID:      p.ID,
		Service: p.Service,
		Name:    p.Name,
	}
}

// FromQueryValues populates this value from a query string map.
func (p *Perm) FromQueryValues(vals url.Values) error {
	var err error
	p.ID = 0
	if vals.Get("id") != "" {
		p.ID, err = strconv.ParseInt(vals.Get("id"), 10, 64)
		if err != nil {
			return err
		}
	}

	p.Service = vals.Get("service")
	p.Name = vals.Get("name")
	return nil
}

// ToPerm returns a value created from this row value.
func (r *PermRow) ToPerm() Perm {
	return Perm{
		ID:      r.ID,
		Service: r.Service,
		Name:    r.Name,
	}
}

// FromPerm populates a perm find value from a perm value.
func (f *PermFind) FromPerm(p *Perm) error {
	f.ID = nil
	if p.ID != 0 {
		f.ID = &p.ID
	}

	f.Service = nil
	if p.Service != "" {
		f.Service = &p.Service
	}

	f.Name = nil
	if p.Name != "" {
		f.Name = &p.Name
	}

	return nil
}

// FromPermRequest populates a perm find value from a perm protobuf request.
func (f *PermFind) FromPermRequest(r *ptypes.PermRequest) error {
	f.ID = nil
	if r.ID != 0 {
		f.ID = &r.ID
	}

	f.Service = nil
	if r.Service != "" {
		f.Service = &r.Service
	}

	f.Name = nil
	if r.Name != "" {
		f.Name = &r.Name
	}

	return nil
}
