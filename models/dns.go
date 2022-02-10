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

// DNS dns
//
// swagger:model dns
type DNS struct {

	// domainname
	// Required: true
	// Min Length: 1
	Domainname *string `json:"domainname"`

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// zoneid
	// Required: true
	// Min Length: 1
	Zoneid *string `json:"zoneid"`
}

// Validate validates this dns
func (m *DNS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNS) validateDomainname(formats strfmt.Registry) error {

	if err := validate.Required("domainname", "body", m.Domainname); err != nil {
		return err
	}

	if err := validate.MinLength("domainname", "body", *m.Domainname, 1); err != nil {
		return err
	}

	return nil
}

func (m *DNS) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *DNS) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *DNS) validateZoneid(formats strfmt.Registry) error {

	if err := validate.Required("zoneid", "body", m.Zoneid); err != nil {
		return err
	}

	if err := validate.MinLength("zoneid", "body", *m.Zoneid, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dns based on context it is used
func (m *DNS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNS) UnmarshalBinary(b []byte) error {
	var res DNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
