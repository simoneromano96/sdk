// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRelationTuplesParams creates a new GetRelationTuplesParams object
// with the default values initialized.
func NewGetRelationTuplesParams() *GetRelationTuplesParams {
	var ()
	return &GetRelationTuplesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRelationTuplesParamsWithTimeout creates a new GetRelationTuplesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRelationTuplesParamsWithTimeout(timeout time.Duration) *GetRelationTuplesParams {
	var ()
	return &GetRelationTuplesParams{

		timeout: timeout,
	}
}

// NewGetRelationTuplesParamsWithContext creates a new GetRelationTuplesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRelationTuplesParamsWithContext(ctx context.Context) *GetRelationTuplesParams {
	var ()
	return &GetRelationTuplesParams{

		Context: ctx,
	}
}

// NewGetRelationTuplesParamsWithHTTPClient creates a new GetRelationTuplesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRelationTuplesParamsWithHTTPClient(client *http.Client) *GetRelationTuplesParams {
	var ()
	return &GetRelationTuplesParams{
		HTTPClient: client,
	}
}

/*GetRelationTuplesParams contains all the parameters to send to the API endpoint
for the get relation tuples operation typically these are written to a http.Request
*/
type GetRelationTuplesParams struct {

	/*Namespace*/
	Namespace string
	/*Object*/
	Object *string
	/*PageSize*/
	PageSize *int64
	/*PageToken*/
	PageToken *string
	/*Relation*/
	Relation *string
	/*Subject*/
	Subject *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get relation tuples params
func (o *GetRelationTuplesParams) WithTimeout(timeout time.Duration) *GetRelationTuplesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get relation tuples params
func (o *GetRelationTuplesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get relation tuples params
func (o *GetRelationTuplesParams) WithContext(ctx context.Context) *GetRelationTuplesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get relation tuples params
func (o *GetRelationTuplesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get relation tuples params
func (o *GetRelationTuplesParams) WithHTTPClient(client *http.Client) *GetRelationTuplesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get relation tuples params
func (o *GetRelationTuplesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get relation tuples params
func (o *GetRelationTuplesParams) WithNamespace(namespace string) *GetRelationTuplesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get relation tuples params
func (o *GetRelationTuplesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithObject adds the object to the get relation tuples params
func (o *GetRelationTuplesParams) WithObject(object *string) *GetRelationTuplesParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the get relation tuples params
func (o *GetRelationTuplesParams) SetObject(object *string) {
	o.Object = object
}

// WithPageSize adds the pageSize to the get relation tuples params
func (o *GetRelationTuplesParams) WithPageSize(pageSize *int64) *GetRelationTuplesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get relation tuples params
func (o *GetRelationTuplesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the get relation tuples params
func (o *GetRelationTuplesParams) WithPageToken(pageToken *string) *GetRelationTuplesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the get relation tuples params
func (o *GetRelationTuplesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithRelation adds the relation to the get relation tuples params
func (o *GetRelationTuplesParams) WithRelation(relation *string) *GetRelationTuplesParams {
	o.SetRelation(relation)
	return o
}

// SetRelation adds the relation to the get relation tuples params
func (o *GetRelationTuplesParams) SetRelation(relation *string) {
	o.Relation = relation
}

// WithSubject adds the subject to the get relation tuples params
func (o *GetRelationTuplesParams) WithSubject(subject *string) *GetRelationTuplesParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the get relation tuples params
func (o *GetRelationTuplesParams) SetSubject(subject *string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *GetRelationTuplesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {
		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	if o.Object != nil {

		// query param object
		var qrObject string
		if o.Object != nil {
			qrObject = *o.Object
		}
		qObject := qrObject
		if qObject != "" {
			if err := r.SetQueryParam("object", qObject); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string
		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {
			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}

	}

	if o.Relation != nil {

		// query param relation
		var qrRelation string
		if o.Relation != nil {
			qrRelation = *o.Relation
		}
		qRelation := qrRelation
		if qRelation != "" {
			if err := r.SetQueryParam("relation", qRelation); err != nil {
				return err
			}
		}

	}

	if o.Subject != nil {

		// query param subject
		var qrSubject string
		if o.Subject != nil {
			qrSubject = *o.Subject
		}
		qSubject := qrSubject
		if qSubject != "" {
			if err := r.SetQueryParam("subject", qSubject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
