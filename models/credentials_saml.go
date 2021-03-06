package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*CredentialsSaml credentials saml

swagger:model CredentialsSaml
*/
type CredentialsSaml struct {

	/* Timestamp for the creation of this credential

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* EMail address

	Read Only: true
	*/
	Email string `json:"email,omitempty"`

	/* Has this credential been disabled?

	Read Only: true
	*/
	IsDisabled *bool `json:"is_disabled,omitempty"`

	/* Timestamp for most recent login using credential

	Read Only: true
	*/
	LoggedInAt string `json:"logged_in_at,omitempty"`

	/* Saml IdP's Unique ID for this user

	Read Only: true
	*/
	SamlUserID string `json:"saml_user_id,omitempty"`

	/* Short name for the type of this kind of credential

	Read Only: true
	*/
	Type string `json:"type,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this credentials saml
func (m *CredentialsSaml) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
