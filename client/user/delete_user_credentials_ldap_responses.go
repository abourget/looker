package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserCredentialsLdapReader is a Reader for the DeleteUserCredentialsLdap structure.
type DeleteUserCredentialsLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserCredentialsLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsLdapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsLdapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsLdapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsLdapNoContent creates a DeleteUserCredentialsLdapNoContent with default headers values
func NewDeleteUserCredentialsLdapNoContent() *DeleteUserCredentialsLdapNoContent {
	return &DeleteUserCredentialsLdapNoContent{}
}

/*DeleteUserCredentialsLdapNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsLdapNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsLdapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsLdapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLdapBadRequest creates a DeleteUserCredentialsLdapBadRequest with default headers values
func NewDeleteUserCredentialsLdapBadRequest() *DeleteUserCredentialsLdapBadRequest {
	return &DeleteUserCredentialsLdapBadRequest{}
}

/*DeleteUserCredentialsLdapBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsLdapBadRequest struct {
	Payload
}

func (o *DeleteUserCredentialsLdapBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsLdapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLdapNotFound creates a DeleteUserCredentialsLdapNotFound with default headers values
func NewDeleteUserCredentialsLdapNotFound() *DeleteUserCredentialsLdapNotFound {
	return &DeleteUserCredentialsLdapNotFound{}
}

/*DeleteUserCredentialsLdapNotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsLdapNotFound struct {
	Payload
}

func (o *DeleteUserCredentialsLdapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsLdapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
