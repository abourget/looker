package user

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

// AllUserCredentialsEmbedsReader is a Reader for the AllUserCredentialsEmbeds structure.
type AllUserCredentialsEmbedsReader struct {
  formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *AllUserCredentialsEmbedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
  switch response.Code() {
  
    case 200:
      result := NewAllUserCredentialsEmbedsOK()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return result, nil
  
    case 400:
      result := NewAllUserCredentialsEmbedsBadRequest()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    case 404:
      result := NewAllUserCredentialsEmbedsNotFound()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    default:
      return nil, runtime.NewAPIError("unknown error", response, response.Code())
  }
}


// NewAllUserCredentialsEmbedsOK creates a AllUserCredentialsEmbedsOK with default headers values
func NewAllUserCredentialsEmbedsOK() *AllUserCredentialsEmbedsOK {
  return &AllUserCredentialsEmbedsOK{
    }
}

/*AllUserCredentialsEmbedsOK handles this case with default header values.

Embedding credential
*/
type AllUserCredentialsEmbedsOK struct {
  
  
  Payload []
  
}


func (o *AllUserCredentialsEmbedsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed][%d] allUserCredentialsEmbedsOK  %+v", 200, o.Payload)
}


func (o *AllUserCredentialsEmbedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewAllUserCredentialsEmbedsBadRequest creates a AllUserCredentialsEmbedsBadRequest with default headers values
func NewAllUserCredentialsEmbedsBadRequest() *AllUserCredentialsEmbedsBadRequest {
  return &AllUserCredentialsEmbedsBadRequest{
    }
}

/*AllUserCredentialsEmbedsBadRequest handles this case with default header values.

Bad Request
*/
type AllUserCredentialsEmbedsBadRequest struct {
  
  
  Payload 
  
}


func (o *AllUserCredentialsEmbedsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed][%d] allUserCredentialsEmbedsBadRequest  %+v", 400, o.Payload)
}


func (o *AllUserCredentialsEmbedsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewAllUserCredentialsEmbedsNotFound creates a AllUserCredentialsEmbedsNotFound with default headers values
func NewAllUserCredentialsEmbedsNotFound() *AllUserCredentialsEmbedsNotFound {
  return &AllUserCredentialsEmbedsNotFound{
    }
}

/*AllUserCredentialsEmbedsNotFound handles this case with default header values.

Not Found
*/
type AllUserCredentialsEmbedsNotFound struct {
  
  
  Payload 
  
}


func (o *AllUserCredentialsEmbedsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed][%d] allUserCredentialsEmbedsNotFound  %+v", 404, o.Payload)
}


func (o *AllUserCredentialsEmbedsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}




