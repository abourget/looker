package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*CredentialsTotp credentials totp

swagger:model CredentialsTotp
*/
type CredentialsTotp struct {

	/* Timestamp for the creation of this credential

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* Has this credential been disabled?

	Read Only: true
	*/
	IsDisabled *bool `json:"is_disabled,omitempty"`

	/* Short name for the type of this kind of credential

	Read Only: true
	*/
	Type string `json:"type,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`

	/* User has verified

	Read Only: true
	*/
	Verified *bool `json:"verified,omitempty"`
}

// Validate validates this credentials totp
func (m *CredentialsTotp) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
