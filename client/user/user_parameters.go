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

// NewUserParams creates a new UserParams object
// with the default values initialized.
func NewUserParams() *UserParams {
	var ()
	return &UserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserParamsWithTimeout creates a new UserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserParamsWithTimeout(timeout time.Duration) *UserParams {
	var ()
	return &UserParams{

		timeout: timeout,
	}
}

/*UserParams contains all the parameters to send to the API endpoint
for the user operation typically these are written to a http.Request
*/
type UserParams struct {

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

// WithFields adds the fields to the user params
func (o *UserParams) WithFields(Fields *string) *UserParams {
	o.Fields = Fields
	return o
}

// WithUserID adds the userId to the user params
func (o *UserParams) WithUserID(UserID int64) *UserParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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
