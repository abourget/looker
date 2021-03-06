package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*SqlQuery sql query

swagger:model SqlQuery
*/
type SqlQuery struct {

	/* Maximum number of rows this query will display on the SQL Runner page

	Read Only: true
	*/
	BrowserLimit int64 `json:"browser_limit,omitempty"`

	/* Connection this query uses

	Read Only: true
	*/
	Connection `json:"connection,omitempty"`

	/* User who created this SQL query

	Read Only: true
	*/
	Creator `json:"creator,omitempty"`

	/* Explore page URL for this SQL query

	Read Only: true
	*/
	ExploreURL string `json:"explore_url,omitempty"`

	/* The most recent time this query was run

	Read Only: true
	*/
	LastRunAt string `json:"last_run_at,omitempty"`

	/* Number of seconds this query took to run the most recent time it was run

	Read Only: true
	*/
	LastRuntime float32 `json:"last_runtime,omitempty"`

	/* Should this query be rendered as plain text

	Read Only: true
	*/
	Plaintext *bool `json:"plaintext,omitempty"`

	/* Number of times this query has been run

	Read Only: true
	*/
	RunCount int64 `json:"run_count,omitempty"`

	/* The identifier of the SQL query

	Read Only: true
	*/
	Slug string `json:"slug,omitempty"`

	/* SQL query text

	Read Only: true
	*/
	Sql string `json:"sql,omitempty"`
}

// Validate validates this sql query
func (m *SqlQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
