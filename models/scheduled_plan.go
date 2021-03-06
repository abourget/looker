package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*ScheduledPlan scheduled plan

swagger:model ScheduledPlan
*/
type ScheduledPlan struct {

	/* Vixie-Style crontab specification when to run
	 */
	Crontab string `json:"crontab,omitempty"`

	/* Id of a dashboard
	 */
	DashboardID int64 `json:"dashboard_id,omitempty"`

	/* Whether the ScheduledPlan is enabled
	 */
	Enabled bool `json:"enabled,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* When the ScheduledPlan was last run

	Read Only: true
	*/
	LastRunAt strfmt.DateTime `json:"last_run_at,omitempty"`

	/* Id of a look
	 */
	LookID int64 `json:"look_id,omitempty"`

	/* Id of a LookML dashboard
	 */
	LookmlDashboardID string `json:"lookml_dashboard_id,omitempty"`

	/* Name
	 */
	Name string `json:"name,omitempty"`

	/* When the ScheduledPlan will next run
	 */
	NextRunAt strfmt.DateTime `json:"next_run_at,omitempty"`

	/* Delivery should occur if data have changed since the last run
	 */
	RequireChange bool `json:"require_change,omitempty"`

	/* Delivery should occur if the dashboard look does not return results
	 */
	RequireNoResults bool `json:"require_no_results,omitempty"`

	/* Delivery should occur if running the dashboard or look returns results
	 */
	RequireResults bool `json:"require_results,omitempty"`

	/* Scheduled plan destinations

	Read Only: true
	*/
	ScheduledPlanDestination []ScheduledPlanDestination `json:"scheduled_plan_destination,omitempty"`

	/* Timezone for interpreting the specified crontab (default is Looker instance timezone)
	 */
	Timezone string `json:"timezone,omitempty"`

	/* Title

	Read Only: true
	*/
	Title string `json:"title,omitempty"`

	/* User who owns this ScheduledPlan

	Read Only: true
	*/
	User `json:"user,omitempty"`
}

// Validate validates this scheduled plan
func (m *ScheduledPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScheduledPlanDestination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduledPlan) validateScheduledPlanDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledPlanDestination) { // not required
		return nil
	}

	return nil
}
