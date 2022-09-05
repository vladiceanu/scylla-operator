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

// NewFindConfigEnableCommitlogParams creates a new FindConfigEnableCommitlogParams object
// with the default values initialized.
func NewFindConfigEnableCommitlogParams() *FindConfigEnableCommitlogParams {

	return &FindConfigEnableCommitlogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigEnableCommitlogParamsWithTimeout creates a new FindConfigEnableCommitlogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigEnableCommitlogParamsWithTimeout(timeout time.Duration) *FindConfigEnableCommitlogParams {

	return &FindConfigEnableCommitlogParams{

		timeout: timeout,
	}
}

// NewFindConfigEnableCommitlogParamsWithContext creates a new FindConfigEnableCommitlogParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigEnableCommitlogParamsWithContext(ctx context.Context) *FindConfigEnableCommitlogParams {

	return &FindConfigEnableCommitlogParams{

		Context: ctx,
	}
}

// NewFindConfigEnableCommitlogParamsWithHTTPClient creates a new FindConfigEnableCommitlogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigEnableCommitlogParamsWithHTTPClient(client *http.Client) *FindConfigEnableCommitlogParams {

	return &FindConfigEnableCommitlogParams{
		HTTPClient: client,
	}
}

/*
FindConfigEnableCommitlogParams contains all the parameters to send to the API endpoint
for the find config enable commitlog operation typically these are written to a http.Request
*/
type FindConfigEnableCommitlogParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) WithTimeout(timeout time.Duration) *FindConfigEnableCommitlogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) WithContext(ctx context.Context) *FindConfigEnableCommitlogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) WithHTTPClient(client *http.Client) *FindConfigEnableCommitlogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config enable commitlog params
func (o *FindConfigEnableCommitlogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigEnableCommitlogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
