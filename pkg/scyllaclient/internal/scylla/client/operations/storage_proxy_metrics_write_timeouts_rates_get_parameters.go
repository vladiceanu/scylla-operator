// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewStorageProxyMetricsWriteTimeoutsRatesGetParams creates a new StorageProxyMetricsWriteTimeoutsRatesGetParams object
// with the default values initialized.
func NewStorageProxyMetricsWriteTimeoutsRatesGetParams() *StorageProxyMetricsWriteTimeoutsRatesGetParams {

	return &StorageProxyMetricsWriteTimeoutsRatesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithTimeout creates a new StorageProxyMetricsWriteTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsWriteTimeoutsRatesGetParams {

	return &StorageProxyMetricsWriteTimeoutsRatesGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithContext creates a new StorageProxyMetricsWriteTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithContext(ctx context.Context) *StorageProxyMetricsWriteTimeoutsRatesGetParams {

	return &StorageProxyMetricsWriteTimeoutsRatesGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithHTTPClient creates a new StorageProxyMetricsWriteTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsWriteTimeoutsRatesGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsWriteTimeoutsRatesGetParams {

	return &StorageProxyMetricsWriteTimeoutsRatesGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyMetricsWriteTimeoutsRatesGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics write timeouts rates get operation typically these are written to a http.Request
*/
type StorageProxyMetricsWriteTimeoutsRatesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsWriteTimeoutsRatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) WithContext(ctx context.Context) *StorageProxyMetricsWriteTimeoutsRatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsWriteTimeoutsRatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics write timeouts rates get params
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsWriteTimeoutsRatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
