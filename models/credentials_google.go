package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*CredentialsGoogle credentials google

swagger:model CredentialsGoogle
*/
type CredentialsGoogle struct {

	/* Timestamp for the creation of this credential

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* Google domain

	Read Only: true
	*/
	Domain string `json:"domain,omitempty"`

	/* EMail address

	Read Only: true
	*/
	Email string `json:"email,omitempty"`

	/* Google's Unique ID for this user

	Read Only: true
	*/
	GoogleUserID string `json:"google_user_id,omitempty"`

	/* Has this credential been disabled?

	Read Only: true
	*/
	IsDisabled *bool `json:"is_disabled,omitempty"`

	/* Timestamp for most recent login using credential

	Read Only: true
	*/
	LoggedInAt string `json:"logged_in_at,omitempty"`

	/* Short name for the type of this kind of credential

	Read Only: true
	*/
	Type string `json:"type,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this credentials google
func (m *CredentialsGoogle) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
