// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RequestDeniedError The request payload used to accept a login or consent request.
// swagger:model RequestDeniedError
type RequestDeniedError struct {

	// code
	Code int64 `json:"status_code,omitempty"`

	// debug
	Debug string `json:"error_debug,omitempty"`

	// description
	Description string `json:"error_description,omitempty"`

	// hint
	Hint string `json:"error_hint,omitempty"`

	// name
	Name string `json:"error,omitempty"`
}

// Validate validates this request denied error
func (m *RequestDeniedError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestDeniedError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestDeniedError) UnmarshalBinary(b []byte) error {
	var res RequestDeniedError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
