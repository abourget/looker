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

// SearchUsersNamesReader is a Reader for the SearchUsersNames structure.
type SearchUsersNamesReader struct {
  formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SearchUsersNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
  switch response.Code() {
  
    case 200:
      result := NewSearchUsersNamesOK()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return result, nil
  
    case 400:
      result := NewSearchUsersNamesBadRequest()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    case 404:
      result := NewSearchUsersNamesNotFound()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return nil, result
  
    default:
      return nil, runtime.NewAPIError("unknown error", response, response.Code())
  }
}


// NewSearchUsersNamesOK creates a SearchUsersNamesOK with default headers values
func NewSearchUsersNamesOK() *SearchUsersNamesOK {
  return &SearchUsersNamesOK{
    }
}

/*SearchUsersNamesOK handles this case with default header values.

Matching users.
*/
type SearchUsersNamesOK struct {
  
  
  Payload []
  
}


func (o *SearchUsersNamesOK) Error() string {
	return fmt.Sprintf("[GET /users/search/names/{pattern}][%d] searchUsersNamesOK  %+v", 200, o.Payload)
}


func (o *SearchUsersNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSearchUsersNamesBadRequest creates a SearchUsersNamesBadRequest with default headers values
func NewSearchUsersNamesBadRequest() *SearchUsersNamesBadRequest {
  return &SearchUsersNamesBadRequest{
    }
}

/*SearchUsersNamesBadRequest handles this case with default header values.

Bad Request
*/
type SearchUsersNamesBadRequest struct {
  
  
  Payload 
  
}


func (o *SearchUsersNamesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/search/names/{pattern}][%d] searchUsersNamesBadRequest  %+v", 400, o.Payload)
}


func (o *SearchUsersNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}


// NewSearchUsersNamesNotFound creates a SearchUsersNamesNotFound with default headers values
func NewSearchUsersNamesNotFound() *SearchUsersNamesNotFound {
  return &SearchUsersNamesNotFound{
    }
}

/*SearchUsersNamesNotFound handles this case with default header values.

Not Found
*/
type SearchUsersNamesNotFound struct {
  
  
  Payload 
  
}


func (o *SearchUsersNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /users/search/names/{pattern}][%d] searchUsersNamesNotFound  %+v", 404, o.Payload)
}


func (o *SearchUsersNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
  
  
  
  // response payload
  if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
    return err
  }
  
  return nil
}




