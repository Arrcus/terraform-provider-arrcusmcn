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

// Providers providers
//
// swagger:model Providers
type Providers string

func NewProviders(value Providers) *Providers {
	v := value
	return &v
}

const (

	// ProvidersAws captures enum value "aws"
	ProvidersAws Providers = "aws"

	// ProvidersGcp captures enum value "gcp"
	ProvidersGcp Providers = "gcp"

	// ProvidersAzure captures enum value "azure"
	ProvidersAzure Providers = "azure"

	// ProvidersOnpremise captures enum value "onpremise"
	ProvidersOnpremise Providers = "onpremise"
)

// for schema
var providersEnum []interface{}

func init() {
	var res []Providers
	if err := json.Unmarshal([]byte(`["aws","gcp","azure","onpremise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersEnum = append(providersEnum, v)
	}
}

func (m Providers) validateProvidersEnum(path, location string, value Providers) error {
	if err := validate.EnumCase(path, location, value, providersEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this providers
func (m Providers) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProvidersEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this providers based on context it is used
func (m Providers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
