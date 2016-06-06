package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateSpaceParams creates a new UpdateSpaceParams object
// with the default values initialized.
func NewUpdateSpaceParams() *UpdateSpaceParams {
	var ()
	return &UpdateSpaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSpaceParamsWithTimeout creates a new UpdateSpaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSpaceParamsWithTimeout(timeout time.Duration) *UpdateSpaceParams {
	var ()
	return &UpdateSpaceParams{

		timeout: timeout,
	}
}

/*UpdateSpaceParams contains all the parameters to send to the API endpoint
for the update space operation typically these are written to a http.Request
*/
type UpdateSpaceParams struct {

	/*Body
	  space

	*/
	Body
	/*SpaceID
	  Id of space

	*/
	SpaceID string

	timeout time.Duration
}

// WithBody adds the body to the update space params
func (o *UpdateSpaceParams) WithBody(Body) *UpdateSpaceParams {
	o.Body = Body
	return o
}

// WithSpaceID adds the spaceId to the update space params
func (o *UpdateSpaceParams) WithSpaceID(SpaceID string) *UpdateSpaceParams {
	o.SpaceID = SpaceID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSpaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}