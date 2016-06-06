package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateSpaceReader is a Reader for the CreateSpace structure.
type CreateSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSpaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSpaceOK creates a CreateSpaceOK with default headers values
func NewCreateSpaceOK() *CreateSpaceOK {
	return &CreateSpaceOK{}
}

/*CreateSpaceOK handles this case with default header values.

space
*/
type CreateSpaceOK struct {
	Payload
}

func (o *CreateSpaceOK) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceOK  %+v", 200, o.Payload)
}

func (o *CreateSpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceBadRequest creates a CreateSpaceBadRequest with default headers values
func NewCreateSpaceBadRequest() *CreateSpaceBadRequest {
	return &CreateSpaceBadRequest{}
}

/*CreateSpaceBadRequest handles this case with default header values.

Bad Request
*/
type CreateSpaceBadRequest struct {
	Payload
}

func (o *CreateSpaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceNotFound creates a CreateSpaceNotFound with default headers values
func NewCreateSpaceNotFound() *CreateSpaceNotFound {
	return &CreateSpaceNotFound{}
}

/*CreateSpaceNotFound handles this case with default header values.

Not Found
*/
type CreateSpaceNotFound struct {
	Payload
}

func (o *CreateSpaceNotFound) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceNotFound  %+v", 404, o.Payload)
}

func (o *CreateSpaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}