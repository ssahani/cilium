// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDHealthzReader is a Reader for the GetEndpointIDHealthz structure.
type GetEndpointIDHealthzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointIDHealthzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointIDHealthzOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointIDHealthzInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointIDHealthzNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointIDHealthzOK creates a GetEndpointIDHealthzOK with default headers values
func NewGetEndpointIDHealthzOK() *GetEndpointIDHealthzOK {
	return &GetEndpointIDHealthzOK{}
}

/*GetEndpointIDHealthzOK handles this case with default header values.

Success
*/
type GetEndpointIDHealthzOK struct {
	Payload *models.EndpointHealth
}

func (o *GetEndpointIDHealthzOK) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzOK  %+v", 200, o.Payload)
}

func (o *GetEndpointIDHealthzOK) GetPayload() *models.EndpointHealth {
	return o.Payload
}

func (o *GetEndpointIDHealthzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDHealthzInvalid creates a GetEndpointIDHealthzInvalid with default headers values
func NewGetEndpointIDHealthzInvalid() *GetEndpointIDHealthzInvalid {
	return &GetEndpointIDHealthzInvalid{}
}

/*GetEndpointIDHealthzInvalid handles this case with default header values.

Invalid identity provided
*/
type GetEndpointIDHealthzInvalid struct {
}

func (o *GetEndpointIDHealthzInvalid) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzInvalid ", 400)
}

func (o *GetEndpointIDHealthzInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDHealthzNotFound creates a GetEndpointIDHealthzNotFound with default headers values
func NewGetEndpointIDHealthzNotFound() *GetEndpointIDHealthzNotFound {
	return &GetEndpointIDHealthzNotFound{}
}

/*GetEndpointIDHealthzNotFound handles this case with default header values.

Endpoint not found
*/
type GetEndpointIDHealthzNotFound struct {
}

func (o *GetEndpointIDHealthzNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzNotFound ", 404)
}

func (o *GetEndpointIDHealthzNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}