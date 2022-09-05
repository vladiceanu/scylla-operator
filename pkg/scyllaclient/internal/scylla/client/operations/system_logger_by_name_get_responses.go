// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// SystemLoggerByNameGetReader is a Reader for the SystemLoggerByNameGet structure.
type SystemLoggerByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemLoggerByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemLoggerByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSystemLoggerByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSystemLoggerByNameGetOK creates a SystemLoggerByNameGetOK with default headers values
func NewSystemLoggerByNameGetOK() *SystemLoggerByNameGetOK {
	return &SystemLoggerByNameGetOK{}
}

/*
SystemLoggerByNameGetOK handles this case with default header values.

SystemLoggerByNameGetOK system logger by name get o k
*/
type SystemLoggerByNameGetOK struct {
	Payload string
}

func (o *SystemLoggerByNameGetOK) GetPayload() string {
	return o.Payload
}

func (o *SystemLoggerByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemLoggerByNameGetDefault creates a SystemLoggerByNameGetDefault with default headers values
func NewSystemLoggerByNameGetDefault(code int) *SystemLoggerByNameGetDefault {
	return &SystemLoggerByNameGetDefault{
		_statusCode: code,
	}
}

/*
SystemLoggerByNameGetDefault handles this case with default header values.

internal server error
*/
type SystemLoggerByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the system logger by name get default response
func (o *SystemLoggerByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *SystemLoggerByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *SystemLoggerByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *SystemLoggerByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
