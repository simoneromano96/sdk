// Code generated by go-swagger; DO NOT EDIT.

package write

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchRelationTuplesReader is a Reader for the PatchRelationTuples structure.
type PatchRelationTuplesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRelationTuplesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchRelationTuplesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRelationTuplesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRelationTuplesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRelationTuplesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRelationTuplesNoContent creates a PatchRelationTuplesNoContent with default headers values
func NewPatchRelationTuplesNoContent() *PatchRelationTuplesNoContent {
	return &PatchRelationTuplesNoContent{}
}

/*PatchRelationTuplesNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 201.
*/
type PatchRelationTuplesNoContent struct {
}

func (o *PatchRelationTuplesNoContent) Error() string {
	return fmt.Sprintf("[PATCH /relation-tuples][%d] patchRelationTuplesNoContent ", 204)
}

func (o *PatchRelationTuplesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRelationTuplesBadRequest creates a PatchRelationTuplesBadRequest with default headers values
func NewPatchRelationTuplesBadRequest() *PatchRelationTuplesBadRequest {
	return &PatchRelationTuplesBadRequest{}
}

/*PatchRelationTuplesBadRequest handles this case with default header values.

The standard error format
*/
type PatchRelationTuplesBadRequest struct {
	Payload *PatchRelationTuplesBadRequestBody
}

func (o *PatchRelationTuplesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /relation-tuples][%d] patchRelationTuplesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRelationTuplesBadRequest) GetPayload() *PatchRelationTuplesBadRequestBody {
	return o.Payload
}

func (o *PatchRelationTuplesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchRelationTuplesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRelationTuplesNotFound creates a PatchRelationTuplesNotFound with default headers values
func NewPatchRelationTuplesNotFound() *PatchRelationTuplesNotFound {
	return &PatchRelationTuplesNotFound{}
}

/*PatchRelationTuplesNotFound handles this case with default header values.

The standard error format
*/
type PatchRelationTuplesNotFound struct {
	Payload *PatchRelationTuplesNotFoundBody
}

func (o *PatchRelationTuplesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /relation-tuples][%d] patchRelationTuplesNotFound  %+v", 404, o.Payload)
}

func (o *PatchRelationTuplesNotFound) GetPayload() *PatchRelationTuplesNotFoundBody {
	return o.Payload
}

func (o *PatchRelationTuplesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchRelationTuplesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRelationTuplesInternalServerError creates a PatchRelationTuplesInternalServerError with default headers values
func NewPatchRelationTuplesInternalServerError() *PatchRelationTuplesInternalServerError {
	return &PatchRelationTuplesInternalServerError{}
}

/*PatchRelationTuplesInternalServerError handles this case with default header values.

The standard error format
*/
type PatchRelationTuplesInternalServerError struct {
	Payload *PatchRelationTuplesInternalServerErrorBody
}

func (o *PatchRelationTuplesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /relation-tuples][%d] patchRelationTuplesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRelationTuplesInternalServerError) GetPayload() *PatchRelationTuplesInternalServerErrorBody {
	return o.Payload
}

func (o *PatchRelationTuplesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchRelationTuplesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchRelationTuplesBadRequestBody patch relation tuples bad request body
swagger:model PatchRelationTuplesBadRequestBody
*/
type PatchRelationTuplesBadRequestBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this patch relation tuples bad request body
func (o *PatchRelationTuplesBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchRelationTuplesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchRelationTuplesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PatchRelationTuplesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchRelationTuplesInternalServerErrorBody patch relation tuples internal server error body
swagger:model PatchRelationTuplesInternalServerErrorBody
*/
type PatchRelationTuplesInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this patch relation tuples internal server error body
func (o *PatchRelationTuplesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchRelationTuplesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchRelationTuplesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PatchRelationTuplesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchRelationTuplesNotFoundBody patch relation tuples not found body
swagger:model PatchRelationTuplesNotFoundBody
*/
type PatchRelationTuplesNotFoundBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this patch relation tuples not found body
func (o *PatchRelationTuplesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchRelationTuplesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchRelationTuplesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PatchRelationTuplesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
