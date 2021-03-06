package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*CredentialsAPI credentials Api

swagger:model CredentialsApi
*/
type CredentialsAPI struct {

	/* Timestamp for the creation of this credential

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* Has this credential been disabled?

	Read Only: true
	*/
	IsDisabled *bool `json:"is_disabled,omitempty"`

	/* API key token

	Read Only: true
	*/
	Token string `json:"token,omitempty"`

	/* Short name for the type of this kind of credential

	Read Only: true
	*/
	Type string `json:"type,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this credentials Api
func (m *CredentialsAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
