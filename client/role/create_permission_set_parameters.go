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

// NewCreatePermissionSetParams creates a new CreatePermissionSetParams object
// with the default values initialized.
func NewCreatePermissionSetParams() *CreatePermissionSetParams {
	var ()
	return &CreatePermissionSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePermissionSetParamsWithTimeout creates a new CreatePermissionSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePermissionSetParamsWithTimeout(timeout time.Duration) *CreatePermissionSetParams {
	var ()
	return &CreatePermissionSetParams{

		timeout: timeout,
	}
}

/*CreatePermissionSetParams contains all the parameters to send to the API endpoint
for the create permission set operation typically these are written to a http.Request
*/
type CreatePermissionSetParams struct {

	/*Body
	  PermissionSet

	*/
	Body

	timeout time.Duration
}

// WithBody adds the body to the create permission set params
func (o *CreatePermissionSetParams) WithBody(Body) *CreatePermissionSetParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePermissionSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
