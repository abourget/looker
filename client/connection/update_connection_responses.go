package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateConnectionReader is a Reader for the UpdateConnection structure.
type UpdateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConnectionOK creates a UpdateConnectionOK with default headers values
func NewUpdateConnectionOK() *UpdateConnectionOK {
	return &UpdateConnectionOK{}
}

/*UpdateConnectionOK handles this case with default header values.

connection
*/
type UpdateConnectionOK struct {
	Payload
}

func (o *UpdateConnectionOK) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionBadRequest creates a UpdateConnectionBadRequest with default headers values
func NewUpdateConnectionBadRequest() *UpdateConnectionBadRequest {
	return &UpdateConnectionBadRequest{}
}

/*UpdateConnectionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateConnectionBadRequest struct {
	Payload
}

func (o *UpdateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionNotFound creates a UpdateConnectionNotFound with default headers values
func NewUpdateConnectionNotFound() *UpdateConnectionNotFound {
	return &UpdateConnectionNotFound{}
}

/*UpdateConnectionNotFound handles this case with default header values.

Not Found
*/
type UpdateConnectionNotFound struct {
	Payload
}

func (o *UpdateConnectionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
