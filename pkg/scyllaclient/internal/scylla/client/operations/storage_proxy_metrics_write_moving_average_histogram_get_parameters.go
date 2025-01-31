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

// NewStorageProxyMetricsWriteMovingAverageHistogramGetParams creates a new StorageProxyMetricsWriteMovingAverageHistogramGetParams object
// with the default values initialized.
func NewStorageProxyMetricsWriteMovingAverageHistogramGetParams() *StorageProxyMetricsWriteMovingAverageHistogramGetParams {

	return &StorageProxyMetricsWriteMovingAverageHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithTimeout creates a new StorageProxyMetricsWriteMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {

	return &StorageProxyMetricsWriteMovingAverageHistogramGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithContext creates a new StorageProxyMetricsWriteMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithContext(ctx context.Context) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {

	return &StorageProxyMetricsWriteMovingAverageHistogramGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithHTTPClient creates a new StorageProxyMetricsWriteMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsWriteMovingAverageHistogramGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {

	return &StorageProxyMetricsWriteMovingAverageHistogramGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyMetricsWriteMovingAverageHistogramGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics write moving average histogram get operation typically these are written to a http.Request
*/
type StorageProxyMetricsWriteMovingAverageHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) WithContext(ctx context.Context) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsWriteMovingAverageHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics write moving average histogram get params
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
