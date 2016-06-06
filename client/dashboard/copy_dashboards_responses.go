package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CopyDashboardsReader is a Reader for the CopyDashboards structure.
type CopyDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CopyDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCopyDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCopyDashboardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCopyDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopyDashboardsOK creates a CopyDashboardsOK with default headers values
func NewCopyDashboardsOK() *CopyDashboardsOK {
	return &CopyDashboardsOK{}
}

/*CopyDashboardsOK handles this case with default header values.

Empty string response.
*/
type CopyDashboardsOK struct {
	Payload string
}

func (o *CopyDashboardsOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/copy][%d] copyDashboardsOK  %+v", 200, o.Payload)
}

func (o *CopyDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyDashboardsBadRequest creates a CopyDashboardsBadRequest with default headers values
func NewCopyDashboardsBadRequest() *CopyDashboardsBadRequest {
	return &CopyDashboardsBadRequest{}
}

/*CopyDashboardsBadRequest handles this case with default header values.

Bad Request
*/
type CopyDashboardsBadRequest struct {
	Payload
}

func (o *CopyDashboardsBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/copy][%d] copyDashboardsBadRequest  %+v", 400, o.Payload)
}

func (o *CopyDashboardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyDashboardsNotFound creates a CopyDashboardsNotFound with default headers values
func NewCopyDashboardsNotFound() *CopyDashboardsNotFound {
	return &CopyDashboardsNotFound{}
}

/*CopyDashboardsNotFound handles this case with default header values.

Not Found
*/
type CopyDashboardsNotFound struct {
	Payload
}

func (o *CopyDashboardsNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards/copy][%d] copyDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *CopyDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
