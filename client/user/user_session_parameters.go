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

// NewUserSessionParams creates a new UserSessionParams object
// with the default values initialized.
func NewUserSessionParams() *UserSessionParams {
	var ()
	return &UserSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserSessionParamsWithTimeout creates a new UserSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserSessionParamsWithTimeout(timeout time.Duration) *UserSessionParams {
	var ()
	return &UserSessionParams{

		timeout: timeout,
	}
}

/*UserSessionParams contains all the parameters to send to the API endpoint
for the user session operation typically these are written to a http.Request
*/
type UserSessionParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*SessionID
	  Id of Web login session

	*/
	SessionID int64
	/*UserID
	  Id of user

	*/
	UserID int64

	timeout time.Duration
}

// WithFields adds the fields to the user session params
func (o *UserSessionParams) WithFields(Fields *string) *UserSessionParams {
	o.Fields = Fields
	return o
}

// WithSessionID adds the sessionId to the user session params
func (o *UserSessionParams) WithSessionID(SessionID int64) *UserSessionParams {
	o.SessionID = SessionID
	return o
}

// WithUserID adds the userId to the user session params
func (o *UserSessionParams) WithUserID(UserID int64) *UserSessionParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UserSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param session_id
	if err := r.SetPathParam("session_id", swag.FormatInt64(o.SessionID)); err != nil {
		return err
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
