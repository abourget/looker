package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCredentialsEmailReader is a Reader for the UserCredentialsEmail structure.
type UserCredentialsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UserCredentialsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCredentialsEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserCredentialsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserCredentialsEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCredentialsEmailOK creates a UserCredentialsEmailOK with default headers values
func NewUserCredentialsEmailOK() *UserCredentialsEmailOK {
	return &UserCredentialsEmailOK{}
}

/*UserCredentialsEmailOK handles this case with default header values.

email/password credential
*/
type UserCredentialsEmailOK struct {
	Payload
}

func (o *UserCredentialsEmailOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_email][%d] userCredentialsEmailOK  %+v", 200, o.Payload)
}

func (o *UserCredentialsEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsEmailBadRequest creates a UserCredentialsEmailBadRequest with default headers values
func NewUserCredentialsEmailBadRequest() *UserCredentialsEmailBadRequest {
	return &UserCredentialsEmailBadRequest{}
}

/*UserCredentialsEmailBadRequest handles this case with default header values.

Bad Request
*/
type UserCredentialsEmailBadRequest struct {
	Payload
}

func (o *UserCredentialsEmailBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_email][%d] userCredentialsEmailBadRequest  %+v", 400, o.Payload)
}

func (o *UserCredentialsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsEmailNotFound creates a UserCredentialsEmailNotFound with default headers values
func NewUserCredentialsEmailNotFound() *UserCredentialsEmailNotFound {
	return &UserCredentialsEmailNotFound{}
}

/*UserCredentialsEmailNotFound handles this case with default header values.

Not Found
*/
type UserCredentialsEmailNotFound struct {
	Payload
}

func (o *UserCredentialsEmailNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_email][%d] userCredentialsEmailNotFound  %+v", 404, o.Payload)
}

func (o *UserCredentialsEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
