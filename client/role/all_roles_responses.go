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

// AllRolesReader is a Reader for the AllRoles structure.
type AllRolesReader struct {
  formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *AllRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
  switch response.Code() {
  
    case 200:
      result := NewAllRolesOK()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return result, nil
  
    case 404:
      result := NewAllRolesNotFound()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    default:
      return nil, runtime.NewAPIError("unknown error", response, response.Code())
  }
}


// NewAllRolesOK creates a AllRolesOK with default headers values
func NewAllRolesOK() *AllRolesOK {
  return &AllRolesOK{
    }
}

/*AllRolesOK handles this case with default header values.

All roles.
*/
type AllRolesOK struct {
  
  
  Payload []
  
}


func (o *AllRolesOK) Error() string {
	return fmt.Sprintf("[GET /roles][%d] allRolesOK  %+v", 200, o.Payload)
}


func (o *AllRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewAllRolesNotFound creates a AllRolesNotFound with default headers values
func NewAllRolesNotFound() *AllRolesNotFound {
  return &AllRolesNotFound{
    }
}

/*AllRolesNotFound handles this case with default header values.

Not Found
*/
type AllRolesNotFound struct {
  
  
  Payload 
  
}


func (o *AllRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /roles][%d] allRolesNotFound  %+v", 404, o.Payload)
}


func (o *AllRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}





