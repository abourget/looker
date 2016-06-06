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

// NewUserCredentialsApi3Params creates a new UserCredentialsApi3Params object
// with the default values initialized.
func NewUserCredentialsApi3Params() *UserCredentialsApi3Params {
	var ()
	return &UserCredentialsApi3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCredentialsApi3ParamsWithTimeout creates a new UserCredentialsApi3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCredentialsApi3ParamsWithTimeout(timeout time.Duration) *UserCredentialsApi3Params {
	var ()
	return &UserCredentialsApi3Params{

		timeout: timeout,
	}
}

/*UserCredentialsApi3Params contains all the parameters to send to the API endpoint
for the user credentials api3 operation typically these are written to a http.Request
*/
type UserCredentialsApi3Params struct {

	/*CredentialsApi3ID
	  Id of API 3 credential

	*/
	CredentialsApi3ID int64
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*UserID
	  Id of user

	*/
	UserID int64

	timeout time.Duration
}

// WithCredentialsApi3ID adds the credentialsApi3Id to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithCredentialsApi3ID(CredentialsApi3ID int64) *UserCredentialsApi3Params {
	o.CredentialsApi3ID = CredentialsApi3ID
	return o
}

// WithFields adds the fields to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithFields(Fields *string) *UserCredentialsApi3Params {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithUserID(UserID int64) *UserCredentialsApi3Params {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UserCredentialsApi3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param credentials_api3_id
	if err := r.SetPathParam("credentials_api3_id", swag.FormatInt64(o.CredentialsApi3ID)); err != nil {
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
