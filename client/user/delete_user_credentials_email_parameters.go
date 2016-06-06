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

// NewDeleteUserCredentialsEmailParams creates a new DeleteUserCredentialsEmailParams object
// with the default values initialized.
func NewDeleteUserCredentialsEmailParams() *DeleteUserCredentialsEmailParams {
	var ()
	return &DeleteUserCredentialsEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsEmailParamsWithTimeout creates a new DeleteUserCredentialsEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsEmailParams {
	var ()
	return &DeleteUserCredentialsEmailParams{

		timeout: timeout,
	}
}

/*DeleteUserCredentialsEmailParams contains all the parameters to send to the API endpoint
for the delete user credentials email operation typically these are written to a http.Request
*/
type DeleteUserCredentialsEmailParams struct {

	/*UserID
	  id of user

	*/
	UserID int64

	timeout time.Duration
}

// WithUserID adds the userId to the delete user credentials email params
func (o *DeleteUserCredentialsEmailParams) WithUserID(UserID int64) *DeleteUserCredentialsEmailParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
