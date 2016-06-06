package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ScheduledJobStage scheduled job stage

swagger:model ScheduledJobStage
*/
type ScheduledJobStage struct {

	/* Time whent the stage was completed

	Read Only: true
	*/
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Node Id stage was run on

	Read Only: true
	*/
	NodeID int64 `json:"node_id,omitempty"`

	/* Job this stage describes

	Read Only: true
	*/
	ScheduledJobID int64 `json:"scheduled_job_id,omitempty"`

	/* Stage

	Read Only: true
	*/
	Stage string `json:"stage,omitempty"`

	/* Time when the stage was started

	Read Only: true
	*/
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`
}

// Validate validates this scheduled job stage
func (m *ScheduledJobStage) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}