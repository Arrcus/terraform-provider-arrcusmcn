// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Arcedgecredentials arcedgecredentials
//
// swagger:model arcedgecredentials
type Arcedgecredentials struct {

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// password
	// Required: true
	// Min Length: 1
	Password *string `json:"password"`

	// username
	// Required: true
	// Min Length: 1
	Username *string `json:"username"`
}

// Validate validates this arcedgecredentials
func (m *Arcedgecredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Arcedgecredentials) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Arcedgecredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 1); err != nil {
		return err
	}

	return nil
}

func (m *Arcedgecredentials) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this arcedgecredentials based on context it is used
func (m *Arcedgecredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Arcedgecredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Arcedgecredentials) UnmarshalBinary(b []byte) error {
	var res Arcedgecredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
