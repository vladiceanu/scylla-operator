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

// ColumnFamilyMetricsWriteLatencyHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyHistogramByNameGet structure.
type ColumnFamilyMetricsWriteLatencyHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyHistogramByNameGetOK creates a ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyHistogramByNameGetOK() *ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK{}
}

/*
ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK column family metrics write latency histogram by name get o k
*/
type ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault creates a ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault(code int) *ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault {
	return &ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency histogram by name get default response
func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
