package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateRoleReader is a Reader for the CreateRole structure.
type CreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateRoleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRoleOK creates a CreateRoleOK with default headers values
func NewCreateRoleOK() *CreateRoleOK {
	return &CreateRoleOK{}
}

/*CreateRoleOK handles this case with default header values.

Newly created Role
*/
type CreateRoleOK struct {
	Payload
}

func (o *CreateRoleOK) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleOK  %+v", 200, o.Payload)
}

func (o *CreateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleBadRequest creates a CreateRoleBadRequest with default headers values
func NewCreateRoleBadRequest() *CreateRoleBadRequest {
	return &CreateRoleBadRequest{}
}

/*CreateRoleBadRequest handles this case with default header values.

Bad Request
*/
type CreateRoleBadRequest struct {
	Payload
}

func (o *CreateRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleNotFound creates a CreateRoleNotFound with default headers values
func NewCreateRoleNotFound() *CreateRoleNotFound {
	return &CreateRoleNotFound{}
}

/*CreateRoleNotFound handles this case with default header values.

Not Found
*/
type CreateRoleNotFound struct {
	Payload
}

func (o *CreateRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleNotFound  %+v", 404, o.Payload)
}

func (o *CreateRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleConflict creates a CreateRoleConflict with default headers values
func NewCreateRoleConflict() *CreateRoleConflict {
	return &CreateRoleConflict{}
}

/*CreateRoleConflict handles this case with default header values.

Resource Already Exists
*/
type CreateRoleConflict struct {
	Payload
}

func (o *CreateRoleConflict) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleConflict  %+v", 409, o.Payload)
}

func (o *CreateRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleUnprocessableEntity creates a CreateRoleUnprocessableEntity with default headers values
func NewCreateRoleUnprocessableEntity() *CreateRoleUnprocessableEntity {
	return &CreateRoleUnprocessableEntity{}
}

/*CreateRoleUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateRoleUnprocessableEntity struct {
	Payload
}

func (o *CreateRoleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateRoleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleInternalServerError creates a CreateRoleInternalServerError with default headers values
func NewCreateRoleInternalServerError() *CreateRoleInternalServerError {
	return &CreateRoleInternalServerError{}
}

/*CreateRoleInternalServerError handles this case with default header values.

Server Error
*/
type CreateRoleInternalServerError struct {
	Payload
}

func (o *CreateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
