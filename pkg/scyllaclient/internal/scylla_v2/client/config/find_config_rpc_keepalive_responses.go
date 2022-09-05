// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigRPCKeepaliveReader is a Reader for the FindConfigRPCKeepalive structure.
type FindConfigRPCKeepaliveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRPCKeepaliveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRPCKeepaliveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRPCKeepaliveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRPCKeepaliveOK creates a FindConfigRPCKeepaliveOK with default headers values
func NewFindConfigRPCKeepaliveOK() *FindConfigRPCKeepaliveOK {
	return &FindConfigRPCKeepaliveOK{}
}

/*
FindConfigRPCKeepaliveOK handles this case with default header values.

Config value
*/
type FindConfigRPCKeepaliveOK struct {
	Payload bool
}

func (o *FindConfigRPCKeepaliveOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigRPCKeepaliveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRPCKeepaliveDefault creates a FindConfigRPCKeepaliveDefault with default headers values
func NewFindConfigRPCKeepaliveDefault(code int) *FindConfigRPCKeepaliveDefault {
	return &FindConfigRPCKeepaliveDefault{
		_statusCode: code,
	}
}

/*
FindConfigRPCKeepaliveDefault handles this case with default header values.

unexpected error
*/
type FindConfigRPCKeepaliveDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config rpc keepalive default response
func (o *FindConfigRPCKeepaliveDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRPCKeepaliveDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRPCKeepaliveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRPCKeepaliveDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
