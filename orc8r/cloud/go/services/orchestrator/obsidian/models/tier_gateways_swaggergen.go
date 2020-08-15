// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	models1 "github.com/go-magma/magma/orc8r/cloud/go/models"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TierGateways tier gateways
// swagger:model tier_gateways
type TierGateways []models1.GatewayID

// Validate validates this tier gateways
func (m TierGateways) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.UniqueItems("", "body", m); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if err := m[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(strconv.Itoa(i))
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
