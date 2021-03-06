package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateDashboardParams creates a new UpdateDashboardParams object
// with the default values initialized.
func NewUpdateDashboardParams() *UpdateDashboardParams {
	var ()
	return &UpdateDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardParamsWithTimeout creates a new UpdateDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDashboardParamsWithTimeout(timeout time.Duration) *UpdateDashboardParams {
	var ()
	return &UpdateDashboardParams{

		timeout: timeout,
	}
}

/*UpdateDashboardParams contains all the parameters to send to the API endpoint
for the update dashboard operation typically these are written to a http.Request
*/
type UpdateDashboardParams struct {

	/*Body
	  dashboard

	*/
	Body
	/*DashboardID
	  Id of dashboard

	*/
	DashboardID string

	timeout time.Duration
}

// WithBody adds the body to the update dashboard params
func (o *UpdateDashboardParams) WithBody(Body) *UpdateDashboardParams {
	o.Body = Body
	return o
}

// WithDashboardID adds the dashboardId to the update dashboard params
func (o *UpdateDashboardParams) WithDashboardID(DashboardID string) *UpdateDashboardParams {
	o.DashboardID = DashboardID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param dashboard_id
	if err := r.SetPathParam("dashboard_id", o.DashboardID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
