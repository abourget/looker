package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*DashboardBase dashboard base

swagger:model DashboardBase
*/
type DashboardBase struct {

	/* Description

	Read Only: true
	*/
	Description string `json:"description,omitempty"`

	/* Is Hidden

	Read Only: true
	*/
	Hidden *bool `json:"hidden,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID string `json:"id,omitempty"`

	/* Model

	Read Only: true
	*/
	Model string `json:"model,omitempty"`

	/* Is Read-only

	Read Only: true
	*/
	Readonly *bool `json:"readonly,omitempty"`

	/* Refresh Interval

	Read Only: true
	*/
	RefreshInterval string `json:"refresh_interval,omitempty"`

	/* Refresh Interval as Integer

	Read Only: true
	*/
	RefreshIntervalToI int64 `json:"refresh_interval_to_i,omitempty"`

	/* ScheduledPlan

	Read Only: true
	*/
	ScheduledPlan `json:"scheduled_plan,omitempty"`

	/* Space

	Read Only: true
	*/
	Space `json:"space,omitempty"`

	/* Look Title

	Read Only: true
	*/
	Title string `json:"title,omitempty"`

	/* Id of User

	Read Only: true
	*/
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this dashboard base
func (m *DashboardBase) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
