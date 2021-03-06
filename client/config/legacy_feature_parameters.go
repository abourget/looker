package config

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

// NewLegacyFeatureParams creates a new LegacyFeatureParams object
// with the default values initialized.
func NewLegacyFeatureParams() *LegacyFeatureParams {
	var ()
	return &LegacyFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLegacyFeatureParamsWithTimeout creates a new LegacyFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLegacyFeatureParamsWithTimeout(timeout time.Duration) *LegacyFeatureParams {
	var ()
	return &LegacyFeatureParams{

		timeout: timeout,
	}
}

/*LegacyFeatureParams contains all the parameters to send to the API endpoint
for the legacy feature operation typically these are written to a http.Request
*/
type LegacyFeatureParams struct {

	/*LegacyFeatureID
	  id of legacy feature

	*/
	LegacyFeatureID int64

	timeout time.Duration
}

// WithLegacyFeatureID adds the legacyFeatureId to the legacy feature params
func (o *LegacyFeatureParams) WithLegacyFeatureID(LegacyFeatureID int64) *LegacyFeatureParams {
	o.LegacyFeatureID = LegacyFeatureID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *LegacyFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param legacy_feature_id
	if err := r.SetPathParam("legacy_feature_id", swag.FormatInt64(o.LegacyFeatureID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
