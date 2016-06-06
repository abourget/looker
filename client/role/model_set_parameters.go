package role

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

// NewModelSetParams creates a new ModelSetParams object
// with the default values initialized.
func NewModelSetParams() *ModelSetParams {
	var ()
	return &ModelSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModelSetParamsWithTimeout creates a new ModelSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModelSetParamsWithTimeout(timeout time.Duration) *ModelSetParams {
	var ()
	return &ModelSetParams{

		timeout: timeout,
	}
}

/*ModelSetParams contains all the parameters to send to the API endpoint
for the model set operation typically these are written to a http.Request
*/
type ModelSetParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*ModelSetID
	  Id of model set

	*/
	ModelSetID int64

	timeout time.Duration
}

// WithFields adds the fields to the model set params
func (o *ModelSetParams) WithFields(Fields *string) *ModelSetParams {
	o.Fields = Fields
	return o
}

// WithModelSetID adds the modelSetId to the model set params
func (o *ModelSetParams) WithModelSetID(ModelSetID int64) *ModelSetParams {
	o.ModelSetID = ModelSetID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ModelSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param model_set_id
	if err := r.SetPathParam("model_set_id", swag.FormatInt64(o.ModelSetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
