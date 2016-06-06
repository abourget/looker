package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateLdapConfigReader is a Reader for the UpdateLdapConfig structure.
type UpdateLdapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateLdapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateLdapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateLdapConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateLdapConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateLdapConfigOK creates a UpdateLdapConfigOK with default headers values
func NewUpdateLdapConfigOK() *UpdateLdapConfigOK {
	return &UpdateLdapConfigOK{}
}

/*UpdateLdapConfigOK handles this case with default header values.

New state for LDAP Configuration.
*/
type UpdateLdapConfigOK struct {
	Payload
}

func (o *UpdateLdapConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateLdapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigBadRequest creates a UpdateLdapConfigBadRequest with default headers values
func NewUpdateLdapConfigBadRequest() *UpdateLdapConfigBadRequest {
	return &UpdateLdapConfigBadRequest{}
}

/*UpdateLdapConfigBadRequest handles this case with default header values.

Bad Request
*/
type UpdateLdapConfigBadRequest struct {
	Payload
}

func (o *UpdateLdapConfigBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLdapConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigNotFound creates a UpdateLdapConfigNotFound with default headers values
func NewUpdateLdapConfigNotFound() *UpdateLdapConfigNotFound {
	return &UpdateLdapConfigNotFound{}
}

/*UpdateLdapConfigNotFound handles this case with default header values.

Not Found
*/
type UpdateLdapConfigNotFound struct {
	Payload
}

func (o *UpdateLdapConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLdapConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}