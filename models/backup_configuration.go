package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*BackupConfiguration backup configuration

swagger:model BackupConfiguration
*/
type BackupConfiguration struct {

	/* Name of bucket for custom-s3 backups
	 */
	CustomS3Bucket string `json:"custom_s3_bucket,omitempty"`

	/* Name of region where the bucket is located
	 */
	CustomS3BucketRegion string `json:"custom_s3_bucket_region,omitempty"`

	/* AWS S3 key used for custom-s3 backups
	 */
	CustomS3Key string `json:"custom_s3_key,omitempty"`

	/* AWS S3 secret used for custom-s3 backups
	 */
	CustomS3Secret string `json:"custom_s3_secret,omitempty"`

	/* Type of backup: looker-s3 or custom-s3
	 */
	Type string `json:"type,omitempty"`

	/* Link to get this item

	Read Only: true
	*/
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this backup configuration
func (m *BackupConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
