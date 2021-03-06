package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*DBConnectionBase d b connection base

swagger:model DBConnectionBase
*/
type DBConnectionBase struct {

	/* (Read-only) SQL Dialect details

	Read Only: true
	*/
	Dialect `json:"dialect,omitempty"`

	/* Name of the connection. Also used as the unique identifier

	Read Only: true
	*/
	Name string `json:"name,omitempty"`

	/* SQL Runner snippets for this connection

	Read Only: true
	*/
	Snippets string `json:"snippets,omitempty"`
}

// Validate validates this d b connection base
func (m *DBConnectionBase) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
