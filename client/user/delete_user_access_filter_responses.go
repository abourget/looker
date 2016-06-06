package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserAccessFilterReader is a Reader for the DeleteUserAccessFilter structure.
type DeleteUserAccessFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserAccessFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserAccessFilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserAccessFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserAccessFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserAccessFilterNoContent creates a DeleteUserAccessFilterNoContent with default headers values
func NewDeleteUserAccessFilterNoContent() *DeleteUserAccessFilterNoContent {
	return &DeleteUserAccessFilterNoContent{}
}

/*DeleteUserAccessFilterNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserAccessFilterNoContent struct {
	Payload string
}

func (o *DeleteUserAccessFilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/access_filters/{access_filter_id}][%d] deleteUserAccessFilterNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserAccessFilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAccessFilterBadRequest creates a DeleteUserAccessFilterBadRequest with default headers values
func NewDeleteUserAccessFilterBadRequest() *DeleteUserAccessFilterBadRequest {
	return &DeleteUserAccessFilterBadRequest{}
}

/*DeleteUserAccessFilterBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserAccessFilterBadRequest struct {
	Payload
}

func (o *DeleteUserAccessFilterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/access_filters/{access_filter_id}][%d] deleteUserAccessFilterBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserAccessFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAccessFilterNotFound creates a DeleteUserAccessFilterNotFound with default headers values
func NewDeleteUserAccessFilterNotFound() *DeleteUserAccessFilterNotFound {
	return &DeleteUserAccessFilterNotFound{}
}

/*DeleteUserAccessFilterNotFound handles this case with default header values.

Not Found
*/
type DeleteUserAccessFilterNotFound struct {
	Payload
}

func (o *DeleteUserAccessFilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/access_filters/{access_filter_id}][%d] deleteUserAccessFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserAccessFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}