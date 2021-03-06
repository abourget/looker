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

// NewSpaceChildrenSearchParams creates a new SpaceChildrenSearchParams object
// with the default values initialized.
func NewSpaceChildrenSearchParams() *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceChildrenSearchParamsWithTimeout creates a new SpaceChildrenSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpaceChildrenSearchParamsWithTimeout(timeout time.Duration) *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{

		timeout: timeout,
	}
}

/*SpaceChildrenSearchParams contains all the parameters to send to the API endpoint
for the space children search operation typically these are written to a http.Request
*/
type SpaceChildrenSearchParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Name
	  Match Space name.

	*/
	Name *string
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string
	/*SpaceID
	  Id of space

	*/
	SpaceID string

	timeout time.Duration
}

// WithFields adds the fields to the space children search params
func (o *SpaceChildrenSearchParams) WithFields(Fields *string) *SpaceChildrenSearchParams {
	o.Fields = Fields
	return o
}

// WithName adds the name to the space children search params
func (o *SpaceChildrenSearchParams) WithName(Name *string) *SpaceChildrenSearchParams {
	o.Name = Name
	return o
}

// WithSorts adds the sorts to the space children search params
func (o *SpaceChildrenSearchParams) WithSorts(Sorts *string) *SpaceChildrenSearchParams {
	o.Sorts = Sorts
	return o
}

// WithSpaceID adds the spaceId to the space children search params
func (o *SpaceChildrenSearchParams) WithSpaceID(SpaceID string) *SpaceChildrenSearchParams {
	o.SpaceID = SpaceID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceChildrenSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string
		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {
			if err := r.SetQueryParam("sorts", qSorts); err != nil {
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
