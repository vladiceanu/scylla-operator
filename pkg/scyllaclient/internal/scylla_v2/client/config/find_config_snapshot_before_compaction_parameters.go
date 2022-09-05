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

// NewFindConfigSnapshotBeforeCompactionParams creates a new FindConfigSnapshotBeforeCompactionParams object
// with the default values initialized.
func NewFindConfigSnapshotBeforeCompactionParams() *FindConfigSnapshotBeforeCompactionParams {

	return &FindConfigSnapshotBeforeCompactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigSnapshotBeforeCompactionParamsWithTimeout creates a new FindConfigSnapshotBeforeCompactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigSnapshotBeforeCompactionParamsWithTimeout(timeout time.Duration) *FindConfigSnapshotBeforeCompactionParams {

	return &FindConfigSnapshotBeforeCompactionParams{

		timeout: timeout,
	}
}

// NewFindConfigSnapshotBeforeCompactionParamsWithContext creates a new FindConfigSnapshotBeforeCompactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigSnapshotBeforeCompactionParamsWithContext(ctx context.Context) *FindConfigSnapshotBeforeCompactionParams {

	return &FindConfigSnapshotBeforeCompactionParams{

		Context: ctx,
	}
}

// NewFindConfigSnapshotBeforeCompactionParamsWithHTTPClient creates a new FindConfigSnapshotBeforeCompactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigSnapshotBeforeCompactionParamsWithHTTPClient(client *http.Client) *FindConfigSnapshotBeforeCompactionParams {

	return &FindConfigSnapshotBeforeCompactionParams{
		HTTPClient: client,
	}
}

/*
FindConfigSnapshotBeforeCompactionParams contains all the parameters to send to the API endpoint
for the find config snapshot before compaction operation typically these are written to a http.Request
*/
type FindConfigSnapshotBeforeCompactionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) WithTimeout(timeout time.Duration) *FindConfigSnapshotBeforeCompactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) WithContext(ctx context.Context) *FindConfigSnapshotBeforeCompactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) WithHTTPClient(client *http.Client) *FindConfigSnapshotBeforeCompactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config snapshot before compaction params
func (o *FindConfigSnapshotBeforeCompactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigSnapshotBeforeCompactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
