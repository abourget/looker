package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DashboardPrefetchReader is a Reader for the DashboardPrefetch structure.
type DashboardPrefetchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DashboardPrefetchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDashboardPrefetchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDashboardPrefetchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDashboardPrefetchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDashboardPrefetchOK creates a DashboardPrefetchOK with default headers values
func NewDashboardPrefetchOK() *DashboardPrefetchOK {
	return &DashboardPrefetchOK{}
}

/*DashboardPrefetchOK handles this case with default header values.

Prefetch
*/
type DashboardPrefetchOK struct {
	Payload
}

func (o *DashboardPrefetchOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/prefetch][%d] dashboardPrefetchOK  %+v", 200, o.Payload)
}

func (o *DashboardPrefetchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardPrefetchBadRequest creates a DashboardPrefetchBadRequest with default headers values
func NewDashboardPrefetchBadRequest() *DashboardPrefetchBadRequest {
	return &DashboardPrefetchBadRequest{}
}

/*DashboardPrefetchBadRequest handles this case with default header values.

Bad Request
*/
type DashboardPrefetchBadRequest struct {
	Payload
}

func (o *DashboardPrefetchBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/prefetch][%d] dashboardPrefetchBadRequest  %+v", 400, o.Payload)
}

func (o *DashboardPrefetchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardPrefetchNotFound creates a DashboardPrefetchNotFound with default headers values
func NewDashboardPrefetchNotFound() *DashboardPrefetchNotFound {
	return &DashboardPrefetchNotFound{}
}

/*DashboardPrefetchNotFound handles this case with default header values.

Not Found
*/
type DashboardPrefetchNotFound struct {
	Payload
}

func (o *DashboardPrefetchNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/prefetch][%d] dashboardPrefetchNotFound  %+v", 404, o.Payload)
}

func (o *DashboardPrefetchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
