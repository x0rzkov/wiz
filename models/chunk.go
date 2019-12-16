// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Chunk chunk
// swagger:model Chunk
type Chunk struct {
	ChunkState

	// the ID of the chunk
	// Required: true
	ChunkID *string `json:"chunk_id"`

	// the error which caused the chunk to fail if it is in the Failed state
	Error *Error `json:"error,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Chunk) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ChunkState
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ChunkState = aO0

	// AO1
	var dataAO1 struct {
		ChunkID *string `json:"chunk_id"`

		Error *Error `json:"error,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ChunkID = dataAO1.ChunkID

	m.Error = dataAO1.Error

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Chunk) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ChunkState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ChunkID *string `json:"chunk_id"`

		Error *Error `json:"error,omitempty"`
	}

	dataAO1.ChunkID = m.ChunkID

	dataAO1.Error = m.Error

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this chunk
func (m *Chunk) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChunkState
	if err := m.ChunkState.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChunkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chunk) validateChunkID(formats strfmt.Registry) error {

	if err := validate.Required("chunk_id", "body", m.ChunkID); err != nil {
		return err
	}

	return nil
}

func (m *Chunk) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Chunk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chunk) UnmarshalBinary(b []byte) error {
	var res Chunk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
