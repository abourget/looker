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

// NewTestLdapConfigAuthParams creates a new TestLdapConfigAuthParams object
// with the default values initialized.
func NewTestLdapConfigAuthParams() *TestLdapConfigAuthParams {
	var ()
	return &TestLdapConfigAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigAuthParamsWithTimeout creates a new TestLdapConfigAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestLdapConfigAuthParamsWithTimeout(timeout time.Duration) *TestLdapConfigAuthParams {
	var ()
	return &TestLdapConfigAuthParams{

		timeout: timeout,
	}
}

/*TestLdapConfigAuthParams contains all the parameters to send to the API endpoint
for the test ldap config auth operation typically these are written to a http.Request
*/
type TestLdapConfigAuthParams struct {

	/*Body
	  LDAP Config

	*/
	Body

	timeout time.Duration
}

// WithBody adds the body to the test ldap config auth params
func (o *TestLdapConfigAuthParams) WithBody(Body) *TestLdapConfigAuthParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
