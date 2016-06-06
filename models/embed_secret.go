package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*EmbedSecret embed secret

swagger:model EmbedSecret
*/
type EmbedSecret struct {

	/* Signing algorithm to use with this secret
	 */
	Algorithm string `json:"algorithm,omitempty"`

	/* When secret was created

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* Is this secret currently enabled
	 */
	Enabled bool `json:"enabled,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Secret for use with SSO embedding

	Read Only: true
	*/
	Secret string `json:"secret,omitempty"`

	/* Id of user who created this secret

	Read Only: true
	*/
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this embed secret
func (m *EmbedSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}