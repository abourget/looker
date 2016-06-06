package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSpaceReader is a Reader for the DeleteSpace structure.
type DeleteSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSpaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSpaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSpaceNoContent creates a DeleteSpaceNoContent with default headers values
func NewDeleteSpaceNoContent() *DeleteSpaceNoContent {
	return &DeleteSpaceNoContent{}
}

/*DeleteSpaceNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteSpaceNoContent struct {
	Payload string
}

func (o *DeleteSpaceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /spaces/{space_id}][%d] deleteSpaceNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSpaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpaceBadRequest creates a DeleteSpaceBadRequest with default headers values
func NewDeleteSpaceBadRequest() *DeleteSpaceBadRequest {
	return &DeleteSpaceBadRequest{}
}

/*DeleteSpaceBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSpaceBadRequest struct {
	Payload
}

func (o *DeleteSpaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /spaces/{space_id}][%d] deleteSpaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpaceNotFound creates a DeleteSpaceNotFound with default headers values
func NewDeleteSpaceNotFound() *DeleteSpaceNotFound {
	return &DeleteSpaceNotFound{}
}

/*DeleteSpaceNotFound handles this case with default header values.

Not Found
*/
type DeleteSpaceNotFound struct {
	Payload
}

func (o *DeleteSpaceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /spaces/{space_id}][%d] deleteSpaceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
