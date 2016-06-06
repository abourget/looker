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

// NewRunInlineQueryParams creates a new RunInlineQueryParams object
// with the default values initialized.
func NewRunInlineQueryParams() *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunInlineQueryParamsWithTimeout creates a new RunInlineQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunInlineQueryParamsWithTimeout(timeout time.Duration) *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{

		timeout: timeout,
	}
}

/*RunInlineQueryParams contains all the parameters to send to the API endpoint
for the run inline query operation typically these are written to a http.Request
*/
type RunInlineQueryParams struct {

	/*ApplyFormatting
	  Apply model-specified formatting to each result.

	*/
	ApplyFormatting *bool
	/*Body
	  inline query

	*/
	Body
	/*Cache
	  Get results from cache if available.

	*/
	Cache *bool
	/*ForceProduction
	  Force use of production models even if the user is in developer mode.

	*/
	ForceProduction *bool
	/*Format
	  Format of result

	*/
	Format string
	/*GenerateDrillLinks
	  Generate drill links (only applicable to 'unified' format.

	*/
	GenerateDrillLinks *bool
	/*ImageHeight
	  Render height for image formats.

	*/
	ImageHeight *int64
	/*ImageWidth
	  Render width for image formats.

	*/
	ImageWidth *int64
	/*Limit
	  Row limit (may override the limit in the saved query).

	*/
	Limit *int64

	timeout time.Duration
}

// WithApplyFormatting adds the applyFormatting to the run inline query params
func (o *RunInlineQueryParams) WithApplyFormatting(ApplyFormatting *bool) *RunInlineQueryParams {
	o.ApplyFormatting = ApplyFormatting
	return o
}

// WithBody adds the body to the run inline query params
func (o *RunInlineQueryParams) WithBody(Body) *RunInlineQueryParams {
	o.Body = Body
	return o
}

// WithCache adds the cache to the run inline query params
func (o *RunInlineQueryParams) WithCache(Cache *bool) *RunInlineQueryParams {
	o.Cache = Cache
	return o
}

// WithForceProduction adds the forceProduction to the run inline query params
func (o *RunInlineQueryParams) WithForceProduction(ForceProduction *bool) *RunInlineQueryParams {
	o.ForceProduction = ForceProduction
	return o
}

// WithFormat adds the format to the run inline query params
func (o *RunInlineQueryParams) WithFormat(Format string) *RunInlineQueryParams {
	o.Format = Format
	return o
}

// WithGenerateDrillLinks adds the generateDrillLinks to the run inline query params
func (o *RunInlineQueryParams) WithGenerateDrillLinks(GenerateDrillLinks *bool) *RunInlineQueryParams {
	o.GenerateDrillLinks = GenerateDrillLinks
	return o
}

// WithImageHeight adds the imageHeight to the run inline query params
func (o *RunInlineQueryParams) WithImageHeight(ImageHeight *int64) *RunInlineQueryParams {
	o.ImageHeight = ImageHeight
	return o
}

// WithImageWidth adds the imageWidth to the run inline query params
func (o *RunInlineQueryParams) WithImageWidth(ImageWidth *int64) *RunInlineQueryParams {
	o.ImageWidth = ImageWidth
	return o
}

// WithLimit adds the limit to the run inline query params
func (o *RunInlineQueryParams) WithLimit(Limit *int64) *RunInlineQueryParams {
	o.Limit = Limit
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RunInlineQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ApplyFormatting != nil {

		// query param apply_formatting
		var qrApplyFormatting bool
		if o.ApplyFormatting != nil {
			qrApplyFormatting = *o.ApplyFormatting
		}
		qApplyFormatting := swag.FormatBool(qrApplyFormatting)
		if qApplyFormatting != "" {
			if err := r.SetQueryParam("apply_formatting", qApplyFormatting); err != nil {
				return err
			}
		}

	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Cache != nil {

		// query param cache
		var qrCache bool
		if o.Cache != nil {
			qrCache = *o.Cache
		}
		qCache := swag.FormatBool(qrCache)
		if qCache != "" {
			if err := r.SetQueryParam("cache", qCache); err != nil {
				return err
			}
		}

	}

	if o.ForceProduction != nil {

		// query param force_production
		var qrForceProduction bool
		if o.ForceProduction != nil {
			qrForceProduction = *o.ForceProduction
		}
		qForceProduction := swag.FormatBool(qrForceProduction)
		if qForceProduction != "" {
			if err := r.SetQueryParam("force_production", qForceProduction); err != nil {
				return err
			}
		}

	}

	// path param format
	if err := r.SetPathParam("format", o.Format); err != nil {
		return err
	}

	if o.GenerateDrillLinks != nil {

		// query param generate_drill_links
		var qrGenerateDrillLinks bool
		if o.GenerateDrillLinks != nil {
			qrGenerateDrillLinks = *o.GenerateDrillLinks
		}
		qGenerateDrillLinks := swag.FormatBool(qrGenerateDrillLinks)
		if qGenerateDrillLinks != "" {
			if err := r.SetQueryParam("generate_drill_links", qGenerateDrillLinks); err != nil {
				return err
			}
		}

	}

	if o.ImageHeight != nil {

		// query param image_height
		var qrImageHeight int64
		if o.ImageHeight != nil {
			qrImageHeight = *o.ImageHeight
		}
		qImageHeight := swag.FormatInt64(qrImageHeight)
		if qImageHeight != "" {
			if err := r.SetQueryParam("image_height", qImageHeight); err != nil {
				return err
			}
		}

	}

	if o.ImageWidth != nil {

		// query param image_width
		var qrImageWidth int64
		if o.ImageWidth != nil {
			qrImageWidth = *o.ImageWidth
		}
		qImageWidth := swag.FormatInt64(qrImageWidth)
		if qImageWidth != "" {
			if err := r.SetQueryParam("image_width", qImageWidth); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
