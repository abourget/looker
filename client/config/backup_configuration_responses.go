package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// BackupConfigurationReader is a Reader for the BackupConfiguration structure.
type BackupConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *BackupConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBackupConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBackupConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBackupConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupConfigurationOK creates a BackupConfigurationOK with default headers values
func NewBackupConfigurationOK() *BackupConfigurationOK {
	return &BackupConfigurationOK{}
}

/*BackupConfigurationOK handles this case with default header values.

Current Backup Configuration
*/
type BackupConfigurationOK struct {
	Payload
}

func (o *BackupConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /backup_configuration][%d] backupConfigurationOK  %+v", 200, o.Payload)
}

func (o *BackupConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupConfigurationBadRequest creates a BackupConfigurationBadRequest with default headers values
func NewBackupConfigurationBadRequest() *BackupConfigurationBadRequest {
	return &BackupConfigurationBadRequest{}
}

/*BackupConfigurationBadRequest handles this case with default header values.

Bad Request
*/
type BackupConfigurationBadRequest struct {
	Payload
}

func (o *BackupConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /backup_configuration][%d] backupConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *BackupConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupConfigurationNotFound creates a BackupConfigurationNotFound with default headers values
func NewBackupConfigurationNotFound() *BackupConfigurationNotFound {
	return &BackupConfigurationNotFound{}
}

/*BackupConfigurationNotFound handles this case with default header values.

Not Found
*/
type BackupConfigurationNotFound struct {
	Payload
}

func (o *BackupConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /backup_configuration][%d] backupConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *BackupConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
