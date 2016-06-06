package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateDashboardPrefetchReader is a Reader for the CreateDashboardPrefetch structure.
type CreateDashboardPrefetchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateDashboardPrefetchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDashboardPrefetchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDashboardPrefetchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateDashboardPrefetchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateDashboardPrefetchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateDashboardPrefetchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardPrefetchOK creates a CreateDashboardPrefetchOK with default headers values
func NewCreateDashboardPrefetchOK() *CreateDashboardPrefetchOK {
	return &CreateDashboardPrefetchOK{}
}

/*CreateDashboardPrefetchOK handles this case with default header values.

Newly created Prefetch
*/
type CreateDashboardPrefetchOK struct {
	Payload
}

func (o *CreateDashboardPrefetchOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/{dashboard_id}/prefetch][%d] createDashboardPrefetchOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardPrefetchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardPrefetchBadRequest creates a CreateDashboardPrefetchBadRequest with default headers values
func NewCreateDashboardPrefetchBadRequest() *CreateDashboardPrefetchBadRequest {
	return &CreateDashboardPrefetchBadRequest{}
}

/*CreateDashboardPrefetchBadRequest handles this case with default header values.

Bad Request
*/
type CreateDashboardPrefetchBadRequest struct {
	Payload
}

func (o *CreateDashboardPrefetchBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/{dashboard_id}/prefetch][%d] createDashboardPrefetchBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDashboardPrefetchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardPrefetchNotFound creates a CreateDashboardPrefetchNotFound with default headers values
func NewCreateDashboardPrefetchNotFound() *CreateDashboardPrefetchNotFound {
	return &CreateDashboardPrefetchNotFound{}
}

/*CreateDashboardPrefetchNotFound handles this case with default header values.

Not Found
*/
type CreateDashboardPrefetchNotFound struct {
	Payload
}

func (o *CreateDashboardPrefetchNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards/{dashboard_id}/prefetch][%d] createDashboardPrefetchNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardPrefetchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardPrefetchConflict creates a CreateDashboardPrefetchConflict with default headers values
func NewCreateDashboardPrefetchConflict() *CreateDashboardPrefetchConflict {
	return &CreateDashboardPrefetchConflict{}
}

/*CreateDashboardPrefetchConflict handles this case with default header values.

Resource Already Exists
*/
type CreateDashboardPrefetchConflict struct {
	Payload
}

func (o *CreateDashboardPrefetchConflict) Error() string {
	return fmt.Sprintf("[POST /dashboards/{dashboard_id}/prefetch][%d] createDashboardPrefetchConflict  %+v", 409, o.Payload)
}

func (o *CreateDashboardPrefetchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardPrefetchUnprocessableEntity creates a CreateDashboardPrefetchUnprocessableEntity with default headers values
func NewCreateDashboardPrefetchUnprocessableEntity() *CreateDashboardPrefetchUnprocessableEntity {
	return &CreateDashboardPrefetchUnprocessableEntity{}
}

/*CreateDashboardPrefetchUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateDashboardPrefetchUnprocessableEntity struct {
	Payload
}

func (o *CreateDashboardPrefetchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboards/{dashboard_id}/prefetch][%d] createDashboardPrefetchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateDashboardPrefetchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
