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

// NewStorageProxyHintedHandoffEnabledByDcPostParams creates a new StorageProxyHintedHandoffEnabledByDcPostParams object
// with the default values initialized.
func NewStorageProxyHintedHandoffEnabledByDcPostParams() *StorageProxyHintedHandoffEnabledByDcPostParams {
	var ()
	return &StorageProxyHintedHandoffEnabledByDcPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcPostParamsWithTimeout creates a new StorageProxyHintedHandoffEnabledByDcPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyHintedHandoffEnabledByDcPostParamsWithTimeout(timeout time.Duration) *StorageProxyHintedHandoffEnabledByDcPostParams {
	var ()
	return &StorageProxyHintedHandoffEnabledByDcPostParams{

		timeout: timeout,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcPostParamsWithContext creates a new StorageProxyHintedHandoffEnabledByDcPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyHintedHandoffEnabledByDcPostParamsWithContext(ctx context.Context) *StorageProxyHintedHandoffEnabledByDcPostParams {
	var ()
	return &StorageProxyHintedHandoffEnabledByDcPostParams{

		Context: ctx,
	}
}

// NewStorageProxyHintedHandoffEnabledByDcPostParamsWithHTTPClient creates a new StorageProxyHintedHandoffEnabledByDcPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyHintedHandoffEnabledByDcPostParamsWithHTTPClient(client *http.Client) *StorageProxyHintedHandoffEnabledByDcPostParams {
	var ()
	return &StorageProxyHintedHandoffEnabledByDcPostParams{
		HTTPClient: client,
	}
}

/*
StorageProxyHintedHandoffEnabledByDcPostParams contains all the parameters to send to the API endpoint
for the storage proxy hinted handoff enabled by dc post operation typically these are written to a http.Request
*/
type StorageProxyHintedHandoffEnabledByDcPostParams struct {

	/*Dcs
	  The dcs to enable in the CSV format

	*/
	Dcs string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) WithTimeout(timeout time.Duration) *StorageProxyHintedHandoffEnabledByDcPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) WithContext(ctx context.Context) *StorageProxyHintedHandoffEnabledByDcPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) WithHTTPClient(client *http.Client) *StorageProxyHintedHandoffEnabledByDcPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDcs adds the dcs to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) WithDcs(dcs string) *StorageProxyHintedHandoffEnabledByDcPostParams {
	o.SetDcs(dcs)
	return o
}

// SetDcs adds the dcs to the storage proxy hinted handoff enabled by dc post params
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) SetDcs(dcs string) {
	o.Dcs = dcs
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyHintedHandoffEnabledByDcPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param dcs
	qrDcs := o.Dcs
	qDcs := qrDcs
	if qDcs != "" {
		if err := r.SetQueryParam("dcs", qDcs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
