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

// StorageProxyMetricsReadGetReader is a Reader for the StorageProxyMetricsReadGet structure.
type StorageProxyMetricsReadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsReadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsReadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsReadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsReadGetOK creates a StorageProxyMetricsReadGetOK with default headers values
func NewStorageProxyMetricsReadGetOK() *StorageProxyMetricsReadGetOK {
	return &StorageProxyMetricsReadGetOK{}
}

/*
StorageProxyMetricsReadGetOK handles this case with default header values.

StorageProxyMetricsReadGetOK storage proxy metrics read get o k
*/
type StorageProxyMetricsReadGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsReadGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsReadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsReadGetDefault creates a StorageProxyMetricsReadGetDefault with default headers values
func NewStorageProxyMetricsReadGetDefault(code int) *StorageProxyMetricsReadGetDefault {
	return &StorageProxyMetricsReadGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsReadGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsReadGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics read get default response
func (o *StorageProxyMetricsReadGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsReadGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsReadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsReadGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
