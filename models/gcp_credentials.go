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

// GcpCredentials gcp credentials
//
// swagger:model gcp_credentials
type GcpCredentials struct {

	// auth provider x509 cert url
	// Min Length: 1
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`

	// auth uri
	// Min Length: 1
	AuthURI string `json:"auth_uri,omitempty"`

	// client Id
	// Min Length: 1
	ClientID string `json:"clientId,omitempty"`

	// client email
	// Min Length: 1
	ClientEmail string `json:"client_email,omitempty"`

	// client x509 cert url
	// Min Length: 1
	ClientX509CertURL string `json:"client_x509_cert_url,omitempty"`

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// private key
	// Min Length: 1
	PrivateKey string `json:"private_key,omitempty"`

	// private key id
	// Min Length: 1
	PrivateKeyID string `json:"private_key_id,omitempty"`

	// project id
	// Min Length: 1
	ProjectID string `json:"project_id,omitempty"`

	// token uri
	// Min Length: 1
	TokenURI string `json:"token_uri,omitempty"`

	// type
	// Min Length: 1
	Type string `json:"type,omitempty"`
}

// Validate validates this gcp credentials
func (m *GcpCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthProviderX509CertURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientX509CertURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCredentials) validateAuthProviderX509CertURL(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProviderX509CertURL) { // not required
		return nil
	}

	if err := validate.MinLength("auth_provider_x509_cert_url", "body", m.AuthProviderX509CertURL, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateAuthURI(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthURI) { // not required
		return nil
	}

	if err := validate.MinLength("auth_uri", "body", m.AuthURI, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateClientID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientID) { // not required
		return nil
	}

	if err := validate.MinLength("clientId", "body", m.ClientID, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateClientEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientEmail) { // not required
		return nil
	}

	if err := validate.MinLength("client_email", "body", m.ClientEmail, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateClientX509CertURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientX509CertURL) { // not required
		return nil
	}

	if err := validate.MinLength("client_x509_cert_url", "body", m.ClientX509CertURL, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validatePrivateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateKey) { // not required
		return nil
	}

	if err := validate.MinLength("private_key", "body", m.PrivateKey, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validatePrivateKeyID(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateKeyID) { // not required
		return nil
	}

	if err := validate.MinLength("private_key_id", "body", m.PrivateKeyID, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinLength("project_id", "body", m.ProjectID, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateTokenURI(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenURI) { // not required
		return nil
	}

	if err := validate.MinLength("token_uri", "body", m.TokenURI, 1); err != nil {
		return err
	}

	return nil
}

func (m *GcpCredentials) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.MinLength("type", "body", m.Type, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gcp credentials based on context it is used
func (m *GcpCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpCredentials) UnmarshalBinary(b []byte) error {
	var res GcpCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
