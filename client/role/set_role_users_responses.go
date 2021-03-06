package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command


import (
  "io"
  "net/http"

  "github.com/go-openapi/runtime"
  "github.com/go-openapi/swag"
  "github.com/go-openapi/errors"
  "github.com/go-openapi/validate"
  "github.com/go-openapi/runtime"

  strfmt "github.com/go-openapi/strfmt"

  "github.com/abourget/looker/models"
  
  
)

// SetRoleUsersReader is a Reader for the SetRoleUsers structure.
type SetRoleUsersReader struct {
  formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SetRoleUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
  switch response.Code() {
  
    case 200:
      result := NewSetRoleUsersOK()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return result, nil
  
    case 400:
      result := NewSetRoleUsersBadRequest()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    case 404:
      result := NewSetRoleUsersNotFound()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    case 405:
      result := NewSetRoleUsersMethodNotAllowed()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    case 500:
      result := NewSetRoleUsersInternalServerError()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    default:
      return nil, runtime.NewAPIError("unknown error", response, response.Code())
  }
}


// NewSetRoleUsersOK creates a SetRoleUsersOK with default headers values
func NewSetRoleUsersOK() *SetRoleUsersOK {
  return &SetRoleUsersOK{
    }
}

/*SetRoleUsersOK handles this case with default header values.

Users wiith role.
*/
type SetRoleUsersOK struct {
  
  
  Payload []
  
}


func (o *SetRoleUsersOK) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/users][%d] setRoleUsersOK  %+v", 200, o.Payload)
}


func (o *SetRoleUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSetRoleUsersBadRequest creates a SetRoleUsersBadRequest with default headers values
func NewSetRoleUsersBadRequest() *SetRoleUsersBadRequest {
  return &SetRoleUsersBadRequest{
    }
}

/*SetRoleUsersBadRequest handles this case with default header values.

Bad Request
*/
type SetRoleUsersBadRequest struct {
  
  
  Payload 
  
}


func (o *SetRoleUsersBadRequest) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/users][%d] setRoleUsersBadRequest  %+v", 400, o.Payload)
}


func (o *SetRoleUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSetRoleUsersNotFound creates a SetRoleUsersNotFound with default headers values
func NewSetRoleUsersNotFound() *SetRoleUsersNotFound {
  return &SetRoleUsersNotFound{
    }
}

/*SetRoleUsersNotFound handles this case with default header values.

Not Found
*/
type SetRoleUsersNotFound struct {
  
  
  Payload 
  
}


func (o *SetRoleUsersNotFound) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/users][%d] setRoleUsersNotFound  %+v", 404, o.Payload)
}


func (o *SetRoleUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSetRoleUsersMethodNotAllowed creates a SetRoleUsersMethodNotAllowed with default headers values
func NewSetRoleUsersMethodNotAllowed() *SetRoleUsersMethodNotAllowed {
  return &SetRoleUsersMethodNotAllowed{
    }
}

/*SetRoleUsersMethodNotAllowed handles this case with default header values.

Resource Can't Be Modified
*/
type SetRoleUsersMethodNotAllowed struct {
  
  
  Payload 
  
}


func (o *SetRoleUsersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/users][%d] setRoleUsersMethodNotAllowed  %+v", 405, o.Payload)
}


func (o *SetRoleUsersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSetRoleUsersInternalServerError creates a SetRoleUsersInternalServerError with default headers values
func NewSetRoleUsersInternalServerError() *SetRoleUsersInternalServerError {
  return &SetRoleUsersInternalServerError{
    }
}

/*SetRoleUsersInternalServerError handles this case with default header values.

Server Error
*/
type SetRoleUsersInternalServerError struct {
  
  
  Payload 
  
}


func (o *SetRoleUsersInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/users][%d] setRoleUsersInternalServerError  %+v", 500, o.Payload)
}


func (o *SetRoleUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}





