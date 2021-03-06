package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*LookWithQuery look with query

swagger:model LookWithQuery
*/
type LookWithQuery struct {

	/* Description
	 */
	Description string `json:"description,omitempty"`

	/* Embed Url

	Read Only: true
	*/
	EmbedURL string `json:"embed_url,omitempty"`

	/* Excel File Url

	Read Only: true
	*/
	ExcelFileURL string `json:"excel_file_url,omitempty"`

	/* Google Spreadsheet Formula

	Read Only: true
	*/
	GoogleSpreadsheetFormula string `json:"google_spreadsheet_formula,omitempty"`

	/* Unique Id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Model

	Read Only: true
	*/
	Model string `json:"model,omitempty"`

	/* Is Public

	Read Only: true
	*/
	Public *bool `json:"public,omitempty"`

	/* Public Slug

	Read Only: true
	*/
	PublicSlug string `json:"public_slug,omitempty"`

	/* Public Url

	Read Only: true
	*/
	PublicURL string `json:"public_url,omitempty"`

	/* Query

	Read Only: true
	*/
	Query `json:"query,omitempty"`

	/* Query Id
	 */
	QueryID int64 `json:"query_id,omitempty"`

	/* Short Url

	Read Only: true
	*/
	ShortURL string `json:"short_url,omitempty"`

	/* Space of this Look

	Read Only: true
	*/
	Space `json:"space,omitempty"`

	/* (Write-only) Space Id
	 */
	SpaceID int64 `json:"space_id,omitempty"`

	/* Look Title
	 */
	Title string `json:"title,omitempty"`

	/* Url

	Read Only: true
	*/
	URL string `json:"url,omitempty"`

	/* User

	Read Only: true
	*/
	User `json:"user,omitempty"`

	/* (Write-only) User Id
	 */
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this look with query
func (m *LookWithQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
