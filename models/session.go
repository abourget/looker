package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Session session

swagger:model Session
*/
type Session struct {

	/* User's browser type

	Read Only: true
	*/
	Browser string `json:"browser,omitempty"`

	/* City component of user location (derived from IP address)

	Read Only: true
	*/
	City string `json:"city,omitempty"`

	/* Country component of user location (derived from IP address)

	Read Only: true
	*/
	Country string `json:"country,omitempty"`

	/* Time when this session was initiated

	Read Only: true
	*/
	CreatedAt string `json:"created_at,omitempty"`

	/* Type of credentials used for logging in this session

	Read Only: true
	*/
	CredentialsType string `json:"credentials_type,omitempty"`

	/* Time when this session will expire

	Read Only: true
	*/
	ExpiresAt string `json:"expires_at,omitempty"`

	/* Time when this session was last extended by the user

	Read Only: true
	*/
	ExtendedAt string `json:"extended_at,omitempty"`

	/* Number of times this session was extended

	Read Only: true
	*/
	ExtendedCount int64 `json:"extended_count,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* IP address of user when this session was initiated

	Read Only: true
	*/
	IPAddress string `json:"ip_address,omitempty"`

	/* User's Operating System

	Read Only: true
	*/
	OperatingSystem string `json:"operating_system,omitempty"`

	/* State component of user location (derived from IP address)

	Read Only: true
	*/
	State string `json:"state,omitempty"`

	/* Actual user in the case when this session represents one user sudo'ing as another

	Read Only: true
	*/
	SudoUserID int64 `json:"sudo_user_id,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
