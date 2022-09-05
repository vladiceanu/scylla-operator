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

// FindConfigReplaceTokenReader is a Reader for the FindConfigReplaceToken structure.
type FindConfigReplaceTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigReplaceTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigReplaceTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigReplaceTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigReplaceTokenOK creates a FindConfigReplaceTokenOK with default headers values
func NewFindConfigReplaceTokenOK() *FindConfigReplaceTokenOK {
	return &FindConfigReplaceTokenOK{}
}

/*
FindConfigReplaceTokenOK handles this case with default header values.

Config value
*/
type FindConfigReplaceTokenOK struct {
	Payload string
}

func (o *FindConfigReplaceTokenOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigReplaceTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigReplaceTokenDefault creates a FindConfigReplaceTokenDefault with default headers values
func NewFindConfigReplaceTokenDefault(code int) *FindConfigReplaceTokenDefault {
	return &FindConfigReplaceTokenDefault{
		_statusCode: code,
	}
}

/*
FindConfigReplaceTokenDefault handles this case with default header values.

unexpected error
*/
type FindConfigReplaceTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config replace token default response
func (o *FindConfigReplaceTokenDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigReplaceTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigReplaceTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigReplaceTokenDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
