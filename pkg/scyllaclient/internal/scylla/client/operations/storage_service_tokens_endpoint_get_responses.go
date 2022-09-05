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

// StorageServiceTokensEndpointGetReader is a Reader for the StorageServiceTokensEndpointGet structure.
type StorageServiceTokensEndpointGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTokensEndpointGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTokensEndpointGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTokensEndpointGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTokensEndpointGetOK creates a StorageServiceTokensEndpointGetOK with default headers values
func NewStorageServiceTokensEndpointGetOK() *StorageServiceTokensEndpointGetOK {
	return &StorageServiceTokensEndpointGetOK{}
}

/*
StorageServiceTokensEndpointGetOK handles this case with default header values.

StorageServiceTokensEndpointGetOK storage service tokens endpoint get o k
*/
type StorageServiceTokensEndpointGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceTokensEndpointGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *StorageServiceTokensEndpointGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceTokensEndpointGetDefault creates a StorageServiceTokensEndpointGetDefault with default headers values
func NewStorageServiceTokensEndpointGetDefault(code int) *StorageServiceTokensEndpointGetDefault {
	return &StorageServiceTokensEndpointGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceTokensEndpointGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceTokensEndpointGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service tokens endpoint get default response
func (o *StorageServiceTokensEndpointGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTokensEndpointGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTokensEndpointGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTokensEndpointGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
