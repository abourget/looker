package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateSamlTestConfigParams creates a new CreateSamlTestConfigParams object
// with the default values initialized.
func NewCreateSamlTestConfigParams() *CreateSamlTestConfigParams {
	var ()
	return &CreateSamlTestConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSamlTestConfigParamsWithTimeout creates a new CreateSamlTestConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSamlTestConfigParamsWithTimeout(timeout time.Duration) *CreateSamlTestConfigParams {
	var ()
	return &CreateSamlTestConfigParams{

		timeout: timeout,
	}
}

/*CreateSamlTestConfigParams contains all the parameters to send to the API endpoint
for the create saml test config operation typically these are written to a http.Request
*/
type CreateSamlTestConfigParams struct {

	/*Body
	  SAML test config

	*/
	Body

	timeout time.Duration
}

// WithBody adds the body to the create saml test config params
func (o *CreateSamlTestConfigParams) WithBody(Body) *CreateSamlTestConfigParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSamlTestConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
