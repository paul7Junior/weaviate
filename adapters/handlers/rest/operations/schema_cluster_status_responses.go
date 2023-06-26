//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaClusterStatusOKCode is the HTTP code returned for type SchemaClusterStatusOK
const SchemaClusterStatusOKCode int = 200

/*
SchemaClusterStatusOK The schema in the cluster is in sync.

swagger:response schemaClusterStatusOK
*/
type SchemaClusterStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.SchemaClusterStatus `json:"body,omitempty"`
}

// NewSchemaClusterStatusOK creates SchemaClusterStatusOK with default headers values
func NewSchemaClusterStatusOK() *SchemaClusterStatusOK {

	return &SchemaClusterStatusOK{}
}

// WithPayload adds the payload to the schema cluster status o k response
func (o *SchemaClusterStatusOK) WithPayload(payload *models.SchemaClusterStatus) *SchemaClusterStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema cluster status o k response
func (o *SchemaClusterStatusOK) SetPayload(payload *models.SchemaClusterStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaClusterStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaClusterStatusInternalServerErrorCode is the HTTP code returned for type SchemaClusterStatusInternalServerError
const SchemaClusterStatusInternalServerErrorCode int = 500

/*
SchemaClusterStatusInternalServerError The schema is either out of sync (see response body) or the sync check could not be completed.

swagger:response schemaClusterStatusInternalServerError
*/
type SchemaClusterStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.SchemaClusterStatus `json:"body,omitempty"`
}

// NewSchemaClusterStatusInternalServerError creates SchemaClusterStatusInternalServerError with default headers values
func NewSchemaClusterStatusInternalServerError() *SchemaClusterStatusInternalServerError {

	return &SchemaClusterStatusInternalServerError{}
}

// WithPayload adds the payload to the schema cluster status internal server error response
func (o *SchemaClusterStatusInternalServerError) WithPayload(payload *models.SchemaClusterStatus) *SchemaClusterStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema cluster status internal server error response
func (o *SchemaClusterStatusInternalServerError) SetPayload(payload *models.SchemaClusterStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaClusterStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
