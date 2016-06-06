package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserCredentialsLookerOpenidReader is a Reader for the DeleteUserCredentialsLookerOpenid structure.
type DeleteUserCredentialsLookerOpenidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserCredentialsLookerOpenidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsLookerOpenidNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsLookerOpenidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsLookerOpenidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsLookerOpenidNoContent creates a DeleteUserCredentialsLookerOpenidNoContent with default headers values
func NewDeleteUserCredentialsLookerOpenidNoContent() *DeleteUserCredentialsLookerOpenidNoContent {
	return &DeleteUserCredentialsLookerOpenidNoContent{}
}

/*DeleteUserCredentialsLookerOpenidNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsLookerOpenidNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsLookerOpenidNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_looker_openid][%d] deleteUserCredentialsLookerOpenidNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsLookerOpenidNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLookerOpenidBadRequest creates a DeleteUserCredentialsLookerOpenidBadRequest with default headers values
func NewDeleteUserCredentialsLookerOpenidBadRequest() *DeleteUserCredentialsLookerOpenidBadRequest {
	return &DeleteUserCredentialsLookerOpenidBadRequest{}
}

/*DeleteUserCredentialsLookerOpenidBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsLookerOpenidBadRequest struct {
	Payload
}

func (o *DeleteUserCredentialsLookerOpenidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_looker_openid][%d] deleteUserCredentialsLookerOpenidBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsLookerOpenidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLookerOpenidNotFound creates a DeleteUserCredentialsLookerOpenidNotFound with default headers values
func NewDeleteUserCredentialsLookerOpenidNotFound() *DeleteUserCredentialsLookerOpenidNotFound {
	return &DeleteUserCredentialsLookerOpenidNotFound{}
}

/*DeleteUserCredentialsLookerOpenidNotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsLookerOpenidNotFound struct {
	Payload
}

func (o *DeleteUserCredentialsLookerOpenidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_looker_openid][%d] deleteUserCredentialsLookerOpenidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsLookerOpenidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
