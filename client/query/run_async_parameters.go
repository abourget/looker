package query

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

// NewRunAsyncParams creates a new RunAsyncParams object
// with the default values initialized.
func NewRunAsyncParams() *RunAsyncParams {
	var ()
	return &RunAsyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunAsyncParamsWithTimeout creates a new RunAsyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunAsyncParamsWithTimeout(timeout time.Duration) *RunAsyncParams {
	var ()
	return &RunAsyncParams{

		timeout: timeout,
	}
}

/*RunAsyncParams contains all the parameters to send to the API endpoint
for the run async operation typically these are written to a http.Request
*/
type RunAsyncParams struct {

	/*Body
	  async query run

	*/
	Body
	/*QueryID
	  ID of query

	*/
	QueryID int64

	timeout time.Duration
}

// WithBody adds the body to the run async params
func (o *RunAsyncParams) WithBody(Body) *RunAsyncParams {
	o.Body = Body
	return o
}

// WithQueryID adds the queryId to the run async params
func (o *RunAsyncParams) WithQueryID(QueryID int64) *RunAsyncParams {
	o.QueryID = QueryID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RunAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param query_id
	if err := r.SetPathParam("query_id", swag.FormatInt64(o.QueryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
