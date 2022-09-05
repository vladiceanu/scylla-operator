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

// NewFindConfigCompactionLargeRowWarningThresholdMbParams creates a new FindConfigCompactionLargeRowWarningThresholdMbParams object
// with the default values initialized.
func NewFindConfigCompactionLargeRowWarningThresholdMbParams() *FindConfigCompactionLargeRowWarningThresholdMbParams {

	return &FindConfigCompactionLargeRowWarningThresholdMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithTimeout creates a new FindConfigCompactionLargeRowWarningThresholdMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithTimeout(timeout time.Duration) *FindConfigCompactionLargeRowWarningThresholdMbParams {

	return &FindConfigCompactionLargeRowWarningThresholdMbParams{

		timeout: timeout,
	}
}

// NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithContext creates a new FindConfigCompactionLargeRowWarningThresholdMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithContext(ctx context.Context) *FindConfigCompactionLargeRowWarningThresholdMbParams {

	return &FindConfigCompactionLargeRowWarningThresholdMbParams{

		Context: ctx,
	}
}

// NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithHTTPClient creates a new FindConfigCompactionLargeRowWarningThresholdMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCompactionLargeRowWarningThresholdMbParamsWithHTTPClient(client *http.Client) *FindConfigCompactionLargeRowWarningThresholdMbParams {

	return &FindConfigCompactionLargeRowWarningThresholdMbParams{
		HTTPClient: client,
	}
}

/*
FindConfigCompactionLargeRowWarningThresholdMbParams contains all the parameters to send to the API endpoint
for the find config compaction large row warning threshold mb operation typically these are written to a http.Request
*/
type FindConfigCompactionLargeRowWarningThresholdMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) WithTimeout(timeout time.Duration) *FindConfigCompactionLargeRowWarningThresholdMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) WithContext(ctx context.Context) *FindConfigCompactionLargeRowWarningThresholdMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) WithHTTPClient(client *http.Client) *FindConfigCompactionLargeRowWarningThresholdMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config compaction large row warning threshold mb params
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCompactionLargeRowWarningThresholdMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
