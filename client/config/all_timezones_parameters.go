package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllTimezonesParams creates a new AllTimezonesParams object
// with the default values initialized.
func NewAllTimezonesParams() *AllTimezonesParams {

	return &AllTimezonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllTimezonesParamsWithTimeout creates a new AllTimezonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllTimezonesParamsWithTimeout(timeout time.Duration) *AllTimezonesParams {

	return &AllTimezonesParams{

		timeout: timeout,
	}
}

/*AllTimezonesParams contains all the parameters to send to the API endpoint
for the all timezones operation typically these are written to a http.Request
*/
type AllTimezonesParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *AllTimezonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
