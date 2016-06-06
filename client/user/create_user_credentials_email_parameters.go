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

// NewCreateUserCredentialsEmailParams creates a new CreateUserCredentialsEmailParams object
// with the default values initialized.
func NewCreateUserCredentialsEmailParams() *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsEmailParamsWithTimeout creates a new CreateUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsEmailParams {
	var ()
	return &CreateUserCredentialsEmailParams{

		timeout: timeout,
	}
}

/*CreateUserCredentialsEmailParams contains all the parameters to send to the API endpoint
for the create user credentials email operation typically these are written to a http.Request
*/
type CreateUserCredentialsEmailParams struct {

	/*Body
	  email/password credential

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

// WithBody adds the body to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithBody(Body) *CreateUserCredentialsEmailParams {
	o.Body = Body
	return o
}

// WithFields adds the fields to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithFields(Fields *string) *CreateUserCredentialsEmailParams {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the create user credentials email params
func (o *CreateUserCredentialsEmailParams) WithUserID(UserID int64) *CreateUserCredentialsEmailParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
