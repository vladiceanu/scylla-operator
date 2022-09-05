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

// StorageServiceTraceProbabilityPostReader is a Reader for the StorageServiceTraceProbabilityPost structure.
type StorageServiceTraceProbabilityPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTraceProbabilityPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTraceProbabilityPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTraceProbabilityPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTraceProbabilityPostOK creates a StorageServiceTraceProbabilityPostOK with default headers values
func NewStorageServiceTraceProbabilityPostOK() *StorageServiceTraceProbabilityPostOK {
	return &StorageServiceTraceProbabilityPostOK{}
}

/*
StorageServiceTraceProbabilityPostOK handles this case with default header values.

StorageServiceTraceProbabilityPostOK storage service trace probability post o k
*/
type StorageServiceTraceProbabilityPostOK struct {
}

func (o *StorageServiceTraceProbabilityPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceTraceProbabilityPostDefault creates a StorageServiceTraceProbabilityPostDefault with default headers values
func NewStorageServiceTraceProbabilityPostDefault(code int) *StorageServiceTraceProbabilityPostDefault {
	return &StorageServiceTraceProbabilityPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceTraceProbabilityPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceTraceProbabilityPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service trace probability post default response
func (o *StorageServiceTraceProbabilityPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTraceProbabilityPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTraceProbabilityPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTraceProbabilityPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
