// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigHintsDirectoryParams creates a new FindConfigHintsDirectoryParams object
// with the default values initialized.
func NewFindConfigHintsDirectoryParams() *FindConfigHintsDirectoryParams {

	return &FindConfigHintsDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigHintsDirectoryParamsWithTimeout creates a new FindConfigHintsDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigHintsDirectoryParamsWithTimeout(timeout time.Duration) *FindConfigHintsDirectoryParams {

	return &FindConfigHintsDirectoryParams{

		timeout: timeout,
	}
}

// NewFindConfigHintsDirectoryParamsWithContext creates a new FindConfigHintsDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigHintsDirectoryParamsWithContext(ctx context.Context) *FindConfigHintsDirectoryParams {

	return &FindConfigHintsDirectoryParams{

		Context: ctx,
	}
}

// NewFindConfigHintsDirectoryParamsWithHTTPClient creates a new FindConfigHintsDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigHintsDirectoryParamsWithHTTPClient(client *http.Client) *FindConfigHintsDirectoryParams {

	return &FindConfigHintsDirectoryParams{
		HTTPClient: client,
	}
}

/*
FindConfigHintsDirectoryParams contains all the parameters to send to the API endpoint
for the find config hints directory operation typically these are written to a http.Request
*/
type FindConfigHintsDirectoryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) WithTimeout(timeout time.Duration) *FindConfigHintsDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) WithContext(ctx context.Context) *FindConfigHintsDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) WithHTTPClient(client *http.Client) *FindConfigHintsDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config hints directory params
func (o *FindConfigHintsDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigHintsDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
