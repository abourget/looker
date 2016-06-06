package api_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLogoutParams creates a new LogoutParams object
// with the default values initialized.
func NewLogoutParams() *LogoutParams {

	return &LogoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogoutParamsWithTimeout creates a new LogoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogoutParamsWithTimeout(timeout time.Duration) *LogoutParams {

	return &LogoutParams{

		timeout: timeout,
	}
}

/*LogoutParams contains all the parameters to send to the API endpoint
for the logout operation typically these are written to a http.Request
*/
type LogoutParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *LogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}