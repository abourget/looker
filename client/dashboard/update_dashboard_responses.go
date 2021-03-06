package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateDashboardReader is a Reader for the UpdateDashboard structure.
type UpdateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDashboardOK creates a UpdateDashboardOK with default headers values
func NewUpdateDashboardOK() *UpdateDashboardOK {
	return &UpdateDashboardOK{}
}

/*UpdateDashboardOK handles this case with default header values.

dashboard
*/
type UpdateDashboardOK struct {
	Payload
}

func (o *UpdateDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardBadRequest creates a UpdateDashboardBadRequest with default headers values
func NewUpdateDashboardBadRequest() *UpdateDashboardBadRequest {
	return &UpdateDashboardBadRequest{}
}

/*UpdateDashboardBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDashboardBadRequest struct {
	Payload
}

func (o *UpdateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardNotFound creates a UpdateDashboardNotFound with default headers values
func NewUpdateDashboardNotFound() *UpdateDashboardNotFound {
	return &UpdateDashboardNotFound{}
}

/*UpdateDashboardNotFound handles this case with default header values.

Not Found
*/
type UpdateDashboardNotFound struct {
	Payload
}

func (o *UpdateDashboardNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
