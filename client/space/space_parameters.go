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

// NewSpaceParams creates a new SpaceParams object
// with the default values initialized.
func NewSpaceParams() *SpaceParams {
	var ()
	return &SpaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceParamsWithTimeout creates a new SpaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpaceParamsWithTimeout(timeout time.Duration) *SpaceParams {
	var ()
	return &SpaceParams{

		timeout: timeout,
	}
}

/*SpaceParams contains all the parameters to send to the API endpoint
for the space operation typically these are written to a http.Request
*/
type SpaceParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*SpaceID
	  Id of space

	*/
	SpaceID string

	timeout time.Duration
}

// WithFields adds the fields to the space params
func (o *SpaceParams) WithFields(Fields *string) *SpaceParams {
	o.Fields = Fields
	return o
}

// WithSpaceID adds the spaceId to the space params
func (o *SpaceParams) WithSpaceID(SpaceID string) *SpaceParams {
	o.SpaceID = SpaceID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
