// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerStatus server status
//
// swagger:model server_status
type ServerStatus struct {

	// status
	// Enum: [running upgrading]
	Status string `json:"status,omitempty"`
}

// Validate validates this server status
func (m *ServerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serverStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["running","upgrading"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverStatusTypeStatusPropEnum = append(serverStatusTypeStatusPropEnum, v)
	}
}

const (

	// ServerStatusStatusRunning captures enum value "running"
	ServerStatusStatusRunning string = "running"

	// ServerStatusStatusUpgrading captures enum value "upgrading"
	ServerStatusStatusUpgrading string = "upgrading"
)

// prop value enum
func (m *ServerStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server status based on context it is used
func (m *ServerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerStatus) UnmarshalBinary(b []byte) error {
	var res ServerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
