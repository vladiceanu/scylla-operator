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

// NewStorageServiceKeyspaceFlushByKeyspacePostParams creates a new StorageServiceKeyspaceFlushByKeyspacePostParams object
// with the default values initialized.
func NewStorageServiceKeyspaceFlushByKeyspacePostParams() *StorageServiceKeyspaceFlushByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceFlushByKeyspacePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithTimeout creates a new StorageServiceKeyspaceFlushByKeyspacePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithTimeout(timeout time.Duration) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceFlushByKeyspacePostParams{

		timeout: timeout,
	}
}

// NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithContext creates a new StorageServiceKeyspaceFlushByKeyspacePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithContext(ctx context.Context) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceFlushByKeyspacePostParams{

		Context: ctx,
	}
}

// NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithHTTPClient creates a new StorageServiceKeyspaceFlushByKeyspacePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceKeyspaceFlushByKeyspacePostParamsWithHTTPClient(client *http.Client) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceFlushByKeyspacePostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceKeyspaceFlushByKeyspacePostParams contains all the parameters to send to the API endpoint
for the storage service keyspace flush by keyspace post operation typically these are written to a http.Request
*/
type StorageServiceKeyspaceFlushByKeyspacePostParams struct {

	/*Cf
	  Comma seperated column family names

	*/
	Cf *string
	/*Keyspace
	  The keyspace to flush

	*/
	Keyspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WithTimeout(timeout time.Duration) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WithContext(ctx context.Context) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WithHTTPClient(client *http.Client) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WithCf(cf *string) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) SetCf(cf *string) {
	o.Cf = cf
}

// WithKeyspace adds the keyspace to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WithKeyspace(keyspace string) *StorageServiceKeyspaceFlushByKeyspacePostParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service keyspace flush by keyspace post params
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceKeyspaceFlushByKeyspacePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cf != nil {

		// query param cf
		var qrCf string
		if o.Cf != nil {
			qrCf = *o.Cf
		}
		qCf := qrCf
		if qCf != "" {
			if err := r.SetQueryParam("cf", qCf); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
