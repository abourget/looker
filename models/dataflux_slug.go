package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*DatafluxSlug dataflux slug

swagger:model DatafluxSlug
*/
type DatafluxSlug struct {

	/* Time that the slug was created

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* GUID used to identify slug
	 */
	GUID string `json:"guid,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Time that the slug was last fetched

	Read Only: true
	*/
	LastFetched strfmt.DateTime `json:"last_fetched,omitempty"`

	/* ID of associated query
	 */
	QueryID int64 `json:"query_id,omitempty"`
}

// Validate validates this dataflux slug
func (m *DatafluxSlug) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}