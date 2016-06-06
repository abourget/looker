package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserCredentialsSamlReader is a Reader for the DeleteUserCredentialsSaml structure.
type DeleteUserCredentialsSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserCredentialsSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsSamlNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsSamlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsSamlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsSamlNoContent creates a DeleteUserCredentialsSamlNoContent with default headers values
func NewDeleteUserCredentialsSamlNoContent() *DeleteUserCredentialsSamlNoContent {
	return &DeleteUserCredentialsSamlNoContent{}
}

/*DeleteUserCredentialsSamlNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsSamlNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsSamlNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_saml][%d] deleteUserCredentialsSamlNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsSamlNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsSamlBadRequest creates a DeleteUserCredentialsSamlBadRequest with default headers values
func NewDeleteUserCredentialsSamlBadRequest() *DeleteUserCredentialsSamlBadRequest {
	return &DeleteUserCredentialsSamlBadRequest{}
}

/*DeleteUserCredentialsSamlBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsSamlBadRequest struct {
	Payload
}

func (o *DeleteUserCredentialsSamlBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_saml][%d] deleteUserCredentialsSamlBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsSamlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsSamlNotFound creates a DeleteUserCredentialsSamlNotFound with default headers values
func NewDeleteUserCredentialsSamlNotFound() *DeleteUserCredentialsSamlNotFound {
	return &DeleteUserCredentialsSamlNotFound{}
}

/*DeleteUserCredentialsSamlNotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsSamlNotFound struct {
	Payload
}

func (o *DeleteUserCredentialsSamlNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_saml][%d] deleteUserCredentialsSamlNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsSamlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
