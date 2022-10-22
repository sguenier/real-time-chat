// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"rc-client/models"
)

// DeleteRoomIDReader is a Reader for the DeleteRoomID structure.
type DeleteRoomIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoomIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoomIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteRoomIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoomIDOK creates a DeleteRoomIDOK with default headers values
func NewDeleteRoomIDOK() *DeleteRoomIDOK {
	return &DeleteRoomIDOK{}
}

/*
DeleteRoomIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRoomIDOK struct {
	Payload *models.AdapterResponse
}

// IsSuccess returns true when this delete room Id o k response has a 2xx status code
func (o *DeleteRoomIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete room Id o k response has a 3xx status code
func (o *DeleteRoomIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete room Id o k response has a 4xx status code
func (o *DeleteRoomIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete room Id o k response has a 5xx status code
func (o *DeleteRoomIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete room Id o k response a status code equal to that given
func (o *DeleteRoomIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRoomIDOK) Error() string {
	return fmt.Sprintf("[DELETE /room/{id}][%d] deleteRoomIdOK  %+v", 200, o.Payload)
}

func (o *DeleteRoomIDOK) String() string {
	return fmt.Sprintf("[DELETE /room/{id}][%d] deleteRoomIdOK  %+v", 200, o.Payload)
}

func (o *DeleteRoomIDOK) GetPayload() *models.AdapterResponse {
	return o.Payload
}

func (o *DeleteRoomIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoomIDInternalServerError creates a DeleteRoomIDInternalServerError with default headers values
func NewDeleteRoomIDInternalServerError() *DeleteRoomIDInternalServerError {
	return &DeleteRoomIDInternalServerError{}
}

/*
DeleteRoomIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteRoomIDInternalServerError struct {
	Payload *models.AdapterHTTPError
}

// IsSuccess returns true when this delete room Id internal server error response has a 2xx status code
func (o *DeleteRoomIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete room Id internal server error response has a 3xx status code
func (o *DeleteRoomIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete room Id internal server error response has a 4xx status code
func (o *DeleteRoomIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete room Id internal server error response has a 5xx status code
func (o *DeleteRoomIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete room Id internal server error response a status code equal to that given
func (o *DeleteRoomIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoomIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /room/{id}][%d] deleteRoomIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoomIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /room/{id}][%d] deleteRoomIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoomIDInternalServerError) GetPayload() *models.AdapterHTTPError {
	return o.Payload
}

func (o *DeleteRoomIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}