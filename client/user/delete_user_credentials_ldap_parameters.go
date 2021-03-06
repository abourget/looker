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

// NewDeleteUserCredentialsLdapParams creates a new DeleteUserCredentialsLdapParams object
// with the default values initialized.
func NewDeleteUserCredentialsLdapParams() *DeleteUserCredentialsLdapParams {
	var ()
	return &DeleteUserCredentialsLdapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsLdapParamsWithTimeout creates a new DeleteUserCredentialsLdapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserCredentialsLdapParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsLdapParams {
	var ()
	return &DeleteUserCredentialsLdapParams{

		timeout: timeout,
	}
}

/*DeleteUserCredentialsLdapParams contains all the parameters to send to the API endpoint
for the delete user credentials ldap operation typically these are written to a http.Request
*/
type DeleteUserCredentialsLdapParams struct {

	/*UserID
	  id of user

	*/
	UserID int64

	timeout time.Duration
}

// WithUserID adds the userId to the delete user credentials ldap params
func (o *DeleteUserCredentialsLdapParams) WithUserID(UserID int64) *DeleteUserCredentialsLdapParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsLdapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
