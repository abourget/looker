package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardOK creates a CreateDashboardOK with default headers values
func NewCreateDashboardOK() *CreateDashboardOK {
	return &CreateDashboardOK{}
}

/*CreateDashboardOK handles this case with default header values.

dashboard
*/
type CreateDashboardOK struct {
	Payload
}

func (o *CreateDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardBadRequest creates a CreateDashboardBadRequest with default headers values
func NewCreateDashboardBadRequest() *CreateDashboardBadRequest {
	return &CreateDashboardBadRequest{}
}

/*CreateDashboardBadRequest handles this case with default header values.

Bad Request
*/
type CreateDashboardBadRequest struct {
	Payload
}

func (o *CreateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNotFound creates a CreateDashboardNotFound with default headers values
func NewCreateDashboardNotFound() *CreateDashboardNotFound {
	return &CreateDashboardNotFound{}
}

/*CreateDashboardNotFound handles this case with default header values.

Not Found
*/
type CreateDashboardNotFound struct {
	Payload
}

func (o *CreateDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
