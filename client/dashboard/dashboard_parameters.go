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

// NewDashboardParams creates a new DashboardParams object
// with the default values initialized.
func NewDashboardParams() *DashboardParams {
	var ()
	return &DashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardParamsWithTimeout creates a new DashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDashboardParamsWithTimeout(timeout time.Duration) *DashboardParams {
	var ()
	return &DashboardParams{

		timeout: timeout,
	}
}

/*DashboardParams contains all the parameters to send to the API endpoint
for the dashboard operation typically these are written to a http.Request
*/
type DashboardParams struct {

	/*DashboardID
	  Id of dashboard

	*/
	DashboardID string
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout time.Duration
}

// WithDashboardID adds the dashboardId to the dashboard params
func (o *DashboardParams) WithDashboardID(DashboardID string) *DashboardParams {
	o.DashboardID = DashboardID
	return o
}

// WithFields adds the fields to the dashboard params
func (o *DashboardParams) WithFields(Fields *string) *DashboardParams {
	o.Fields = Fields
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param dashboard_id
	if err := r.SetPathParam("dashboard_id", o.DashboardID); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
