package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*DashboardFilter dashboard filter

swagger:model DashboardFilter
*/
type DashboardFilter struct {

	/* Default value of filter
	 */
	DefaultValue string `json:"default_value,omitempty"`

	/* Dimension of filter (required if type = field)
	 */
	Dimension string `json:"dimension,omitempty"`

	/* Explore of filter (required if type = field)
	 */
	Explore string `json:"explore,omitempty"`

	/* Field information

	Read Only: true
	*/
	Field string `json:"field,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Array of listeners for faceted filters
	 */
	ListensToFilters []string `json:"listens_to_filters,omitempty"`

	/* Model of filter (required if type = field)
	 */
	Model string `json:"model,omitempty"`

	/* Name of filter
	 */
	Name string `json:"name,omitempty"`

	/* Name of title
	 */
	Title string `json:"title,omitempty"`

	/* Type of filter: one of date, number, string, or field
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this dashboard filter
func (m *DashboardFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListensToFilters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardFilter) validateListensToFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ListensToFilters) { // not required
		return nil
	}

	return nil
}
