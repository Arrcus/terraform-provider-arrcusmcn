// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArcorchVersion arcorch version
//
// swagger:model arcorch_version
type ArcorchVersion struct {

	// platform
	Platform string `json:"platform,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this arcorch version
func (m *ArcorchVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this arcorch version based on context it is used
func (m *ArcorchVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArcorchVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArcorchVersion) UnmarshalBinary(b []byte) error {
	var res ArcorchVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
