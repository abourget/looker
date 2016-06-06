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

// NewUpdateUserCredentialsEmailParams creates a new UpdateUserCredentialsEmailParams object
// with the default values initialized.
func NewUpdateUserCredentialsEmailParams() *UpdateUserCredentialsEmailParams {
	var ()
	return &UpdateUserCredentialsEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserCredentialsEmailParamsWithTimeout creates a new UpdateUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *UpdateUserCredentialsEmailParams {
	var ()
	return &UpdateUserCredentialsEmailParams{

		timeout: timeout,
	}
}

/*UpdateUserCredentialsEmailParams contains all the parameters to send to the API endpoint
for the update user credentials email operation typically these are written to a http.Request
*/
type UpdateUserCredentialsEmailParams struct {

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

// WithBody adds the body to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithBody(Body) *UpdateUserCredentialsEmailParams {
	o.Body = Body
	return o
}

// WithFields adds the fields to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithFields(Fields *string) *UpdateUserCredentialsEmailParams {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the update user credentials email params
func (o *UpdateUserCredentialsEmailParams) WithUserID(UserID int64) *UpdateUserCredentialsEmailParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
