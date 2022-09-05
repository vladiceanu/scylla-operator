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

// ColumnFamilyMetricsMemtableOffHeapSizeGetReader is a Reader for the ColumnFamilyMetricsMemtableOffHeapSizeGet structure.
type ColumnFamilyMetricsMemtableOffHeapSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableOffHeapSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableOffHeapSizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableOffHeapSizeGetOK creates a ColumnFamilyMetricsMemtableOffHeapSizeGetOK with default headers values
func NewColumnFamilyMetricsMemtableOffHeapSizeGetOK() *ColumnFamilyMetricsMemtableOffHeapSizeGetOK {
	return &ColumnFamilyMetricsMemtableOffHeapSizeGetOK{}
}

/*
ColumnFamilyMetricsMemtableOffHeapSizeGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableOffHeapSizeGetOK column family metrics memtable off heap size get o k
*/
type ColumnFamilyMetricsMemtableOffHeapSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableOffHeapSizeGetDefault creates a ColumnFamilyMetricsMemtableOffHeapSizeGetDefault with default headers values
func NewColumnFamilyMetricsMemtableOffHeapSizeGetDefault(code int) *ColumnFamilyMetricsMemtableOffHeapSizeGetDefault {
	return &ColumnFamilyMetricsMemtableOffHeapSizeGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsMemtableOffHeapSizeGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableOffHeapSizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable off heap size get default response
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
