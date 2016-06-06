package sql_query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSqlQueryParams creates a new SqlQueryParams object
// with the default values initialized.
func NewSqlQueryParams() *SqlQueryParams {
	var ()
	return &SqlQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSqlQueryParamsWithTimeout creates a new SqlQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSqlQueryParamsWithTimeout(timeout time.Duration) *SqlQueryParams {
	var ()
	return &SqlQueryParams{

		timeout: timeout,
	}
}

/*SqlQueryParams contains all the parameters to send to the API endpoint
for the sql query operation typically these are written to a http.Request
*/
type SqlQueryParams struct {

	/*Slug
	  slug of query

	*/
	Slug string

	timeout time.Duration
}

// WithSlug adds the slug to the sql query params
func (o *SqlQueryParams) WithSlug(Slug string) *SqlQueryParams {
	o.Slug = Slug
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SqlQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
