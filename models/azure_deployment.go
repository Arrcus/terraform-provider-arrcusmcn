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

// AzureDeployment azure deployment
//
// swagger:model azure_deployment
type AzureDeployment struct {

	// accelerated networking enabled
	AcceleratedNetworkingEnabled bool `json:"accelerated_networking_enabled,omitempty"`

	// byoip
	// Min Length: 1
	Byoip string `json:"byoip,omitempty"`

	// enable accelerated networking
	EnableAcceleratedNetworking bool `json:"enable_accelerated_networking,omitempty"`

	// instance key
	InstanceKey *InstanceKey `json:"instance_key,omitempty"`

	// instance type
	// Min Length: 1
	InstanceType string `json:"instance_type,omitempty"`

	// private subnet
	// Min Length: 1
	PrivateSubnet string `json:"private_subnet,omitempty"`

	// public subnet
	// Min Length: 1
	PublicSubnet string `json:"public_subnet,omitempty"`

	// region
	// Min Length: 1
	Region string `json:"region,omitempty"`

	// resource group
	// Min Length: 1
	ResourceGroup string `json:"resource_group,omitempty"`

	// vnet
	// Min Length: 1
	Vnet string `json:"vnet,omitempty"`
}

// Validate validates this azure deployment
func (m *AzureDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByoip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDeployment) validateByoip(formats strfmt.Registry) error {
	if swag.IsZero(m.Byoip) { // not required
		return nil
	}

	if err := validate.MinLength("byoip", "body", m.Byoip, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateInstanceKey(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceKey) { // not required
		return nil
	}

	if m.InstanceKey != nil {
		if err := m.InstanceKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instance_key")
			}
			return err
		}
	}

	return nil
}

func (m *AzureDeployment) validateInstanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	if err := validate.MinLength("instance_type", "body", m.InstanceType, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validatePrivateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateSubnet) { // not required
		return nil
	}

	if err := validate.MinLength("private_subnet", "body", m.PrivateSubnet, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validatePublicSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicSubnet) { // not required
		return nil
	}

	if err := validate.MinLength("public_subnet", "body", m.PublicSubnet, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := validate.MinLength("region", "body", m.Region, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateResourceGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if err := validate.MinLength("resource_group", "body", m.ResourceGroup, 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureDeployment) validateVnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Vnet) { // not required
		return nil
	}

	if err := validate.MinLength("vnet", "body", m.Vnet, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this azure deployment based on the context it is used
func (m *AzureDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDeployment) contextValidateInstanceKey(ctx context.Context, formats strfmt.Registry) error {

	if m.InstanceKey != nil {
		if err := m.InstanceKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instance_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureDeployment) UnmarshalBinary(b []byte) error {
	var res AzureDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
