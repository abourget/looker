package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserCredentialsEmbedReader is a Reader for the DeleteUserCredentialsEmbed structure.
type DeleteUserCredentialsEmbedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserCredentialsEmbedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsEmbedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsEmbedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsEmbedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsEmbedNoContent creates a DeleteUserCredentialsEmbedNoContent with default headers values
func NewDeleteUserCredentialsEmbedNoContent() *DeleteUserCredentialsEmbedNoContent {
	return &DeleteUserCredentialsEmbedNoContent{}
}

/*DeleteUserCredentialsEmbedNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsEmbedNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsEmbedNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] deleteUserCredentialsEmbedNoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsEmbedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsEmbedBadRequest creates a DeleteUserCredentialsEmbedBadRequest with default headers values
func NewDeleteUserCredentialsEmbedBadRequest() *DeleteUserCredentialsEmbedBadRequest {
	return &DeleteUserCredentialsEmbedBadRequest{}
}

/*DeleteUserCredentialsEmbedBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsEmbedBadRequest struct {
	Payload
}

func (o *DeleteUserCredentialsEmbedBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] deleteUserCredentialsEmbedBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsEmbedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsEmbedNotFound creates a DeleteUserCredentialsEmbedNotFound with default headers values
func NewDeleteUserCredentialsEmbedNotFound() *DeleteUserCredentialsEmbedNotFound {
	return &DeleteUserCredentialsEmbedNotFound{}
}

/*DeleteUserCredentialsEmbedNotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsEmbedNotFound struct {
	Payload
}

func (o *DeleteUserCredentialsEmbedNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] deleteUserCredentialsEmbedNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsEmbedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
