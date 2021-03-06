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

// NewSpaceAncestorsParams creates a new SpaceAncestorsParams object
// with the default values initialized.
func NewSpaceAncestorsParams() *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceAncestorsParamsWithTimeout creates a new SpaceAncestorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpaceAncestorsParamsWithTimeout(timeout time.Duration) *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{

		timeout: timeout,
	}
}

/*SpaceAncestorsParams contains all the parameters to send to the API endpoint
for the space ancestors operation typically these are written to a http.Request
*/
type SpaceAncestorsParams struct {

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

// WithFields adds the fields to the space ancestors params
func (o *SpaceAncestorsParams) WithFields(Fields *string) *SpaceAncestorsParams {
	o.Fields = Fields
	return o
}

// WithSpaceID adds the spaceId to the space ancestors params
func (o *SpaceAncestorsParams) WithSpaceID(SpaceID string) *SpaceAncestorsParams {
	o.SpaceID = SpaceID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceAncestorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
