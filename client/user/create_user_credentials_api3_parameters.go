package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateUserCredentialsApi3Params creates a new CreateUserCredentialsApi3Params object
// with the default values initialized.
func NewCreateUserCredentialsApi3Params() *CreateUserCredentialsApi3Params {
	var ()
	return &CreateUserCredentialsApi3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsApi3ParamsWithTimeout creates a new CreateUserCredentialsApi3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserCredentialsApi3ParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsApi3Params {
	var ()
	return &CreateUserCredentialsApi3Params{

		timeout: timeout,
	}
}

/*CreateUserCredentialsApi3Params contains all the parameters to send to the API endpoint
for the create user credentials api3 operation typically these are written to a http.Request
*/
type CreateUserCredentialsApi3Params struct {

	/*Body
	  API 3 credential

	*/
	Body
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*UserID
	  id of user

	*/
	UserID int64

	timeout time.Duration
}

// WithBody adds the body to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithBody(Body) *CreateUserCredentialsApi3Params {
	o.Body = Body
	return o
}

// WithFields adds the fields to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithFields(Fields *string) *CreateUserCredentialsApi3Params {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithUserID(UserID int64) *CreateUserCredentialsApi3Params {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsApi3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}