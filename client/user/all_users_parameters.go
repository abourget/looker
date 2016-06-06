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

// NewAllUsersParams creates a new AllUsersParams object
// with the default values initialized.
func NewAllUsersParams() *AllUsersParams {
	var ()
	return &AllUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllUsersParamsWithTimeout creates a new AllUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllUsersParamsWithTimeout(timeout time.Duration) *AllUsersParams {
	var ()
	return &AllUsersParams{

		timeout: timeout,
	}
}

/*AllUsersParams contains all the parameters to send to the API endpoint
for the all users operation typically these are written to a http.Request
*/
type AllUsersParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Page
	  Requested page.

	*/
	Page *int64
	/*PerPage
	  Results per page.

	*/
	PerPage *int64
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string

	timeout time.Duration
}

// WithFields adds the fields to the all users params
func (o *AllUsersParams) WithFields(Fields *string) *AllUsersParams {
	o.Fields = Fields
	return o
}

// WithPage adds the page to the all users params
func (o *AllUsersParams) WithPage(Page *int64) *AllUsersParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the all users params
func (o *AllUsersParams) WithPerPage(PerPage *int64) *AllUsersParams {
	o.PerPage = PerPage
	return o
}

// WithSorts adds the sorts to the all users params
func (o *AllUsersParams) WithSorts(Sorts *string) *AllUsersParams {
	o.Sorts = Sorts
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AllUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
