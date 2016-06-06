package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SpaceChildrenReader is a Reader for the SpaceChildren structure.
type SpaceChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SpaceChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpaceChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpaceChildrenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpaceChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSpaceChildrenOK creates a SpaceChildrenOK with default headers values
func NewSpaceChildrenOK() *SpaceChildrenOK {
	return &SpaceChildrenOK{}
}

/*SpaceChildrenOK handles this case with default header values.

space
*/
type SpaceChildrenOK struct {
	Payload
}

func (o *SpaceChildrenOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenOK  %+v", 200, o.Payload)
}

func (o *SpaceChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceChildrenBadRequest creates a SpaceChildrenBadRequest with default headers values
func NewSpaceChildrenBadRequest() *SpaceChildrenBadRequest {
	return &SpaceChildrenBadRequest{}
}

/*SpaceChildrenBadRequest handles this case with default header values.

Bad Request
*/
type SpaceChildrenBadRequest struct {
	Payload
}

func (o *SpaceChildrenBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenBadRequest  %+v", 400, o.Payload)
}

func (o *SpaceChildrenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceChildrenNotFound creates a SpaceChildrenNotFound with default headers values
func NewSpaceChildrenNotFound() *SpaceChildrenNotFound {
	return &SpaceChildrenNotFound{}
}

/*SpaceChildrenNotFound handles this case with default header values.

Not Found
*/
type SpaceChildrenNotFound struct {
	Payload
}

func (o *SpaceChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/children][%d] spaceChildrenNotFound  %+v", 404, o.Payload)
}

func (o *SpaceChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
