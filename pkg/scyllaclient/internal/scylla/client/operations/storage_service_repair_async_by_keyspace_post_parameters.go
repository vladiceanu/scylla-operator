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

// NewStorageServiceRepairAsyncByKeyspacePostParams creates a new StorageServiceRepairAsyncByKeyspacePostParams object
// with the default values initialized.
func NewStorageServiceRepairAsyncByKeyspacePostParams() *StorageServiceRepairAsyncByKeyspacePostParams {
	var ()
	return &StorageServiceRepairAsyncByKeyspacePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceRepairAsyncByKeyspacePostParamsWithTimeout creates a new StorageServiceRepairAsyncByKeyspacePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceRepairAsyncByKeyspacePostParamsWithTimeout(timeout time.Duration) *StorageServiceRepairAsyncByKeyspacePostParams {
	var ()
	return &StorageServiceRepairAsyncByKeyspacePostParams{

		timeout: timeout,
	}
}

// NewStorageServiceRepairAsyncByKeyspacePostParamsWithContext creates a new StorageServiceRepairAsyncByKeyspacePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceRepairAsyncByKeyspacePostParamsWithContext(ctx context.Context) *StorageServiceRepairAsyncByKeyspacePostParams {
	var ()
	return &StorageServiceRepairAsyncByKeyspacePostParams{

		Context: ctx,
	}
}

// NewStorageServiceRepairAsyncByKeyspacePostParamsWithHTTPClient creates a new StorageServiceRepairAsyncByKeyspacePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceRepairAsyncByKeyspacePostParamsWithHTTPClient(client *http.Client) *StorageServiceRepairAsyncByKeyspacePostParams {
	var ()
	return &StorageServiceRepairAsyncByKeyspacePostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceRepairAsyncByKeyspacePostParams contains all the parameters to send to the API endpoint
for the storage service repair async by keyspace post operation typically these are written to a http.Request
*/
type StorageServiceRepairAsyncByKeyspacePostParams struct {

	/*ColumnFamilies
	  Which column families to repair in the given keyspace. Multiple columns families can be named separated by commas. If this option is missing, all column families in the keyspace are repaired.

	*/
	ColumnFamilies *string
	/*DataCenters
	  Which data centers are to participate in this repair. Multiple data centers can be listed separated by commas.

	*/
	DataCenters *string
	/*EndToken
	  Token on which to end repair

	*/
	EndToken *string
	/*Hosts
	  Which hosts are to participate in this repair. Multiple hosts can be listed separated by commas.

	*/
	Hosts *string
	/*Incremental
	  If the value is the string 'true' with any capitalization, perform incremental repair.

	*/
	Incremental *string
	/*JobThreads
	  An integer specifying the parallelism on each node.

	*/
	JobThreads *string
	/*Keyspace
	  The keyspace to repair

	*/
	Keyspace string
	/*Parallelism
	  Repair parallelism, can be 0 (sequential), 1 (parallel) or 2 (datacenter-aware).

	*/
	Parallelism *string
	/*PrimaryRange
	  If the value is the string 'true' with any capitalization, repair only the first range returned by the partitioner.

	*/
	PrimaryRange *string
	/*Ranges
	  An explicit list of ranges to repair, overriding the default choice. Each range is expressed as token1:token2, and multiple ranges can be given as a comma separated list.

	*/
	Ranges *string
	/*StartToken
	  Token on which to begin repair

	*/
	StartToken *string
	/*Trace
	  If the value is the string 'true' with any capitalization, enable tracing of the repair.

	*/
	Trace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithTimeout(timeout time.Duration) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithContext(ctx context.Context) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithHTTPClient(client *http.Client) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnFamilies adds the columnFamilies to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithColumnFamilies(columnFamilies *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetColumnFamilies(columnFamilies)
	return o
}

// SetColumnFamilies adds the columnFamilies to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetColumnFamilies(columnFamilies *string) {
	o.ColumnFamilies = columnFamilies
}

// WithDataCenters adds the dataCenters to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithDataCenters(dataCenters *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetDataCenters(dataCenters)
	return o
}

// SetDataCenters adds the dataCenters to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetDataCenters(dataCenters *string) {
	o.DataCenters = dataCenters
}

// WithEndToken adds the endToken to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithEndToken(endToken *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetEndToken(endToken)
	return o
}

// SetEndToken adds the endToken to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetEndToken(endToken *string) {
	o.EndToken = endToken
}

// WithHosts adds the hosts to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithHosts(hosts *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetHosts(hosts)
	return o
}

// SetHosts adds the hosts to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetHosts(hosts *string) {
	o.Hosts = hosts
}

// WithIncremental adds the incremental to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithIncremental(incremental *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetIncremental(incremental)
	return o
}

// SetIncremental adds the incremental to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetIncremental(incremental *string) {
	o.Incremental = incremental
}

// WithJobThreads adds the jobThreads to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithJobThreads(jobThreads *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetJobThreads(jobThreads)
	return o
}

// SetJobThreads adds the jobThreads to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetJobThreads(jobThreads *string) {
	o.JobThreads = jobThreads
}

// WithKeyspace adds the keyspace to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithKeyspace(keyspace string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WithParallelism adds the parallelism to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithParallelism(parallelism *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetParallelism(parallelism)
	return o
}

// SetParallelism adds the parallelism to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetParallelism(parallelism *string) {
	o.Parallelism = parallelism
}

// WithPrimaryRange adds the primaryRange to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithPrimaryRange(primaryRange *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetPrimaryRange(primaryRange)
	return o
}

// SetPrimaryRange adds the primaryRange to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetPrimaryRange(primaryRange *string) {
	o.PrimaryRange = primaryRange
}

// WithRanges adds the ranges to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithRanges(ranges *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetRanges(ranges)
	return o
}

// SetRanges adds the ranges to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetRanges(ranges *string) {
	o.Ranges = ranges
}

// WithStartToken adds the startToken to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithStartToken(startToken *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetStartToken(startToken)
	return o
}

// SetStartToken adds the startToken to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetStartToken(startToken *string) {
	o.StartToken = startToken
}

// WithTrace adds the trace to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WithTrace(trace *string) *StorageServiceRepairAsyncByKeyspacePostParams {
	o.SetTrace(trace)
	return o
}

// SetTrace adds the trace to the storage service repair async by keyspace post params
func (o *StorageServiceRepairAsyncByKeyspacePostParams) SetTrace(trace *string) {
	o.Trace = trace
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceRepairAsyncByKeyspacePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ColumnFamilies != nil {

		// query param columnFamilies
		var qrColumnFamilies string
		if o.ColumnFamilies != nil {
			qrColumnFamilies = *o.ColumnFamilies
		}
		qColumnFamilies := qrColumnFamilies
		if qColumnFamilies != "" {
			if err := r.SetQueryParam("columnFamilies", qColumnFamilies); err != nil {
				return err
			}
		}

	}

	if o.DataCenters != nil {

		// query param dataCenters
		var qrDataCenters string
		if o.DataCenters != nil {
			qrDataCenters = *o.DataCenters
		}
		qDataCenters := qrDataCenters
		if qDataCenters != "" {
			if err := r.SetQueryParam("dataCenters", qDataCenters); err != nil {
				return err
			}
		}

	}

	if o.EndToken != nil {

		// query param endToken
		var qrEndToken string
		if o.EndToken != nil {
			qrEndToken = *o.EndToken
		}
		qEndToken := qrEndToken
		if qEndToken != "" {
			if err := r.SetQueryParam("endToken", qEndToken); err != nil {
				return err
			}
		}

	}

	if o.Hosts != nil {

		// query param hosts
		var qrHosts string
		if o.Hosts != nil {
			qrHosts = *o.Hosts
		}
		qHosts := qrHosts
		if qHosts != "" {
			if err := r.SetQueryParam("hosts", qHosts); err != nil {
				return err
			}
		}

	}

	if o.Incremental != nil {

		// query param incremental
		var qrIncremental string
		if o.Incremental != nil {
			qrIncremental = *o.Incremental
		}
		qIncremental := qrIncremental
		if qIncremental != "" {
			if err := r.SetQueryParam("incremental", qIncremental); err != nil {
				return err
			}
		}

	}

	if o.JobThreads != nil {

		// query param jobThreads
		var qrJobThreads string
		if o.JobThreads != nil {
			qrJobThreads = *o.JobThreads
		}
		qJobThreads := qrJobThreads
		if qJobThreads != "" {
			if err := r.SetQueryParam("jobThreads", qJobThreads); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if o.Parallelism != nil {

		// query param parallelism
		var qrParallelism string
		if o.Parallelism != nil {
			qrParallelism = *o.Parallelism
		}
		qParallelism := qrParallelism
		if qParallelism != "" {
			if err := r.SetQueryParam("parallelism", qParallelism); err != nil {
				return err
			}
		}

	}

	if o.PrimaryRange != nil {

		// query param primaryRange
		var qrPrimaryRange string
		if o.PrimaryRange != nil {
			qrPrimaryRange = *o.PrimaryRange
		}
		qPrimaryRange := qrPrimaryRange
		if qPrimaryRange != "" {
			if err := r.SetQueryParam("primaryRange", qPrimaryRange); err != nil {
				return err
			}
		}

	}

	if o.Ranges != nil {

		// query param ranges
		var qrRanges string
		if o.Ranges != nil {
			qrRanges = *o.Ranges
		}
		qRanges := qrRanges
		if qRanges != "" {
			if err := r.SetQueryParam("ranges", qRanges); err != nil {
				return err
			}
		}

	}

	if o.StartToken != nil {

		// query param startToken
		var qrStartToken string
		if o.StartToken != nil {
			qrStartToken = *o.StartToken
		}
		qStartToken := qrStartToken
		if qStartToken != "" {
			if err := r.SetQueryParam("startToken", qStartToken); err != nil {
				return err
			}
		}

	}

	if o.Trace != nil {

		// query param trace
		var qrTrace string
		if o.Trace != nil {
			qrTrace = *o.Trace
		}
		qTrace := qrTrace
		if qTrace != "" {
			if err := r.SetQueryParam("trace", qTrace); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
