package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllPermissionsParams creates a new AllPermissionsParams object
// with the default values initialized.
func NewAllPermissionsParams() *AllPermissionsParams {

	return &AllPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllPermissionsParamsWithTimeout creates a new AllPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllPermissionsParamsWithTimeout(timeout time.Duration) *AllPermissionsParams {

	return &AllPermissionsParams{

		timeout: timeout,
	}
}

/*AllPermissionsParams contains all the parameters to send to the API endpoint
for the all permissions operation typically these are written to a http.Request
*/
type AllPermissionsParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *AllPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
