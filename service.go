package dlib

import "encoding/json"

// ServiceInfo valies provide information about this service.
type ServiceInfo struct {
	// Name is the short name of this service.
	Name string `json:"name,omitempty"`

	// Short is a brief description of this service.
	Short string `json:"short,omitempty"`

	// Long is a full description of this service.
	Long string `json:"long,omitempty"`

	// Version is the semantic version number of this service.
	Version string `json:"version,omitempty"`
}

// NewServiceInfo creates a new service info value.
func NewServiceInfo(name string, short string,
	long string, version string) *ServiceInfo {
	return &ServiceInfo{
		Name:    name,
		Short:   short,
		Long:    long,
		Version: version,
	}
}

// String formats ServiceInfo value as a JSON format string.
func (si *ServiceInfo) String() string {
	str, err := json.Marshal(si)
	if err != nil {
		return ""
	}

	return string(str)
}
