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

// GetStreamIDReader is a Reader for the GetStreamID structure.
type GetStreamIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStreamIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetStreamIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStreamIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStreamIDBadRequest creates a GetStreamIDBadRequest with default headers values
func NewGetStreamIDBadRequest() *GetStreamIDBadRequest {
	return &GetStreamIDBadRequest{}
}

/*
GetStreamIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetStreamIDBadRequest struct {
	Payload *models.AdapterHTTPError
}

// IsSuccess returns true when this get stream Id bad request response has a 2xx status code
func (o *GetStreamIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stream Id bad request response has a 3xx status code
func (o *GetStreamIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stream Id bad request response has a 4xx status code
func (o *GetStreamIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stream Id bad request response has a 5xx status code
func (o *GetStreamIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get stream Id bad request response a status code equal to that given
func (o *GetStreamIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetStreamIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /stream/{id}][%d] getStreamIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetStreamIDBadRequest) String() string {
	return fmt.Sprintf("[GET /stream/{id}][%d] getStreamIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetStreamIDBadRequest) GetPayload() *models.AdapterHTTPError {
	return o.Payload
}

func (o *GetStreamIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStreamIDInternalServerError creates a GetStreamIDInternalServerError with default headers values
func NewGetStreamIDInternalServerError() *GetStreamIDInternalServerError {
	return &GetStreamIDInternalServerError{}
}

/*
GetStreamIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStreamIDInternalServerError struct {
	Payload *models.AdapterHTTPError
}

// IsSuccess returns true when this get stream Id internal server error response has a 2xx status code
func (o *GetStreamIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stream Id internal server error response has a 3xx status code
func (o *GetStreamIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stream Id internal server error response has a 4xx status code
func (o *GetStreamIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stream Id internal server error response has a 5xx status code
func (o *GetStreamIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get stream Id internal server error response a status code equal to that given
func (o *GetStreamIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetStreamIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stream/{id}][%d] getStreamIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStreamIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /stream/{id}][%d] getStreamIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStreamIDInternalServerError) GetPayload() *models.AdapterHTTPError {
	return o.Payload
}

func (o *GetStreamIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdapterHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
