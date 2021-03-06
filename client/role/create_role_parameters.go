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

// NewCreateRoleParams creates a new CreateRoleParams object
// with the default values initialized.
func NewCreateRoleParams() *CreateRoleParams {
	var ()
	return &CreateRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleParamsWithTimeout creates a new CreateRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRoleParamsWithTimeout(timeout time.Duration) *CreateRoleParams {
	var ()
	return &CreateRoleParams{

		timeout: timeout,
	}
}

/*CreateRoleParams contains all the parameters to send to the API endpoint
for the create role operation typically these are written to a http.Request
*/
type CreateRoleParams struct {

	/*Body
	  Role

	*/
	Body

	timeout time.Duration
}

// WithBody adds the body to the create role params
func (o *CreateRoleParams) WithBody(Body) *CreateRoleParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
