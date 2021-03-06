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

// NewCreateUserCredentialsTotpParams creates a new CreateUserCredentialsTotpParams object
// with the default values initialized.
func NewCreateUserCredentialsTotpParams() *CreateUserCredentialsTotpParams {
	var ()
	return &CreateUserCredentialsTotpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsTotpParamsWithTimeout creates a new CreateUserCredentialsTotpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserCredentialsTotpParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsTotpParams {
	var ()
	return &CreateUserCredentialsTotpParams{

		timeout: timeout,
	}
}

/*CreateUserCredentialsTotpParams contains all the parameters to send to the API endpoint
for the create user credentials totp operation typically these are written to a http.Request
*/
type CreateUserCredentialsTotpParams struct {

	/*Body
	  Two-factor credential

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

// WithBody adds the body to the create user credentials totp params
func (o *CreateUserCredentialsTotpParams) WithBody(Body) *CreateUserCredentialsTotpParams {
	o.Body = Body
	return o
}

// WithFields adds the fields to the create user credentials totp params
func (o *CreateUserCredentialsTotpParams) WithFields(Fields *string) *CreateUserCredentialsTotpParams {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the create user credentials totp params
func (o *CreateUserCredentialsTotpParams) WithUserID(UserID int64) *CreateUserCredentialsTotpParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsTotpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
