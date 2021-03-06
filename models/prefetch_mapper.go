package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*PrefetchMapper prefetch mapper

swagger:model PrefetchMapper
*/
type PrefetchMapper struct {

	/* Number of seconds it took to compute results for prefetch.

	Read Only: true
	*/
	ComputationTime float32 `json:"computation_time,omitempty"`

	/* Time when prefetch was created.

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* Number of times prefetch has been accessed.

	Read Only: true
	*/
	HitCount int64 `json:"hit_count,omitempty"`

	/* Time when prefetch was last accessed.

	Read Only: true
	*/
	TouchedAt strfmt.DateTime `json:"touched_at,omitempty"`

	/* Number of seconds prefetch will live for.

	Read Only: true
	*/
	TTL int64 `json:"ttl,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`

	/* data associated with the queries stored by prefetching the data

	Read Only: true
	*/
	Value []QueryResult `json:"value,omitempty"`
}

// Validate validates this prefetch mapper
func (m *PrefetchMapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrefetchMapper) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	return nil
}
