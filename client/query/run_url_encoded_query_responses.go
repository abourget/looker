package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RunURLEncodedQueryReader is a Reader for the RunURLEncodedQuery structure.
type RunURLEncodedQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *RunURLEncodedQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunURLEncodedQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRunURLEncodedQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRunURLEncodedQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunURLEncodedQueryOK creates a RunURLEncodedQueryOK with default headers values
func NewRunURLEncodedQueryOK() *RunURLEncodedQueryOK {
	return &RunURLEncodedQueryOK{}
}

/*RunURLEncodedQueryOK handles this case with default header values.

query
*/
type RunURLEncodedQueryOK struct {
	Payload string
}

func (o *RunURLEncodedQueryOK) Error() string {
	return fmt.Sprintf("[GET /queries/models/{model_name}/views/{view_name}/run/{format}][%d] runUrlEncodedQueryOK  %+v", 200, o.Payload)
}

func (o *RunURLEncodedQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunURLEncodedQueryBadRequest creates a RunURLEncodedQueryBadRequest with default headers values
func NewRunURLEncodedQueryBadRequest() *RunURLEncodedQueryBadRequest {
	return &RunURLEncodedQueryBadRequest{}
}

/*RunURLEncodedQueryBadRequest handles this case with default header values.

Bad Request
*/
type RunURLEncodedQueryBadRequest struct {
	Payload
}

func (o *RunURLEncodedQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /queries/models/{model_name}/views/{view_name}/run/{format}][%d] runUrlEncodedQueryBadRequest  %+v", 400, o.Payload)
}

func (o *RunURLEncodedQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunURLEncodedQueryNotFound creates a RunURLEncodedQueryNotFound with default headers values
func NewRunURLEncodedQueryNotFound() *RunURLEncodedQueryNotFound {
	return &RunURLEncodedQueryNotFound{}
}

/*RunURLEncodedQueryNotFound handles this case with default header values.

Not Found
*/
type RunURLEncodedQueryNotFound struct {
	Payload
}

func (o *RunURLEncodedQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /queries/models/{model_name}/views/{view_name}/run/{format}][%d] runUrlEncodedQueryNotFound  %+v", 404, o.Payload)
}

func (o *RunURLEncodedQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
