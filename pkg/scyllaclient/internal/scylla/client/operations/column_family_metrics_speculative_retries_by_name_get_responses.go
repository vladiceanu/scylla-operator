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

// ColumnFamilyMetricsSpeculativeRetriesByNameGetReader is a Reader for the ColumnFamilyMetricsSpeculativeRetriesByNameGet structure.
type ColumnFamilyMetricsSpeculativeRetriesByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsSpeculativeRetriesByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK creates a ColumnFamilyMetricsSpeculativeRetriesByNameGetOK with default headers values
func NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK() *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK {
	return &ColumnFamilyMetricsSpeculativeRetriesByNameGetOK{}
}

/*
ColumnFamilyMetricsSpeculativeRetriesByNameGetOK handles this case with default header values.

ColumnFamilyMetricsSpeculativeRetriesByNameGetOK column family metrics speculative retries by name get o k
*/
type ColumnFamilyMetricsSpeculativeRetriesByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsSpeculativeRetriesByNameGetDefault creates a ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault with default headers values
func NewColumnFamilyMetricsSpeculativeRetriesByNameGetDefault(code int) *ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault {
	return &ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics speculative retries by name get default response
func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
