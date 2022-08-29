// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Access access
//
// swagger:model access
type Access string

func NewAccess(value Access) *Access {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Access.
func (m Access) Pointer() *Access {
	return &m
}

const (

	// AccessR captures enum value "r"
	AccessR Access = "r"

	// AccessRw captures enum value "rw"
	AccessRw Access = "rw"
)

// for schema
var accessEnum []interface{}

func init() {
	var res []Access
	if err := json.Unmarshal([]byte(`["r","rw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessEnum = append(accessEnum, v)
	}
}

func (m Access) validateAccessEnum(path, location string, value Access) error {
	if err := validate.EnumCase(path, location, value, accessEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this access
func (m Access) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this access based on context it is used
func (m Access) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
