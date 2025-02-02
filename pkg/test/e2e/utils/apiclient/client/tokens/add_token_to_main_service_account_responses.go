// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// AddTokenToMainServiceAccountReader is a Reader for the AddTokenToMainServiceAccount structure.
type AddTokenToMainServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTokenToMainServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddTokenToMainServiceAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddTokenToMainServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTokenToMainServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddTokenToMainServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddTokenToMainServiceAccountCreated creates a AddTokenToMainServiceAccountCreated with default headers values
func NewAddTokenToMainServiceAccountCreated() *AddTokenToMainServiceAccountCreated {
	return &AddTokenToMainServiceAccountCreated{}
}

/*AddTokenToMainServiceAccountCreated handles this case with default header values.

ServiceAccountToken
*/
type AddTokenToMainServiceAccountCreated struct {
	Payload *models.ServiceAccountToken
}

func (o *AddTokenToMainServiceAccountCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/serviceaccounts/{serviceaccount_id}/tokens][%d] addTokenToMainServiceAccountCreated  %+v", 201, o.Payload)
}

func (o *AddTokenToMainServiceAccountCreated) GetPayload() *models.ServiceAccountToken {
	return o.Payload
}

func (o *AddTokenToMainServiceAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccountToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTokenToMainServiceAccountUnauthorized creates a AddTokenToMainServiceAccountUnauthorized with default headers values
func NewAddTokenToMainServiceAccountUnauthorized() *AddTokenToMainServiceAccountUnauthorized {
	return &AddTokenToMainServiceAccountUnauthorized{}
}

/*AddTokenToMainServiceAccountUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type AddTokenToMainServiceAccountUnauthorized struct {
}

func (o *AddTokenToMainServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/serviceaccounts/{serviceaccount_id}/tokens][%d] addTokenToMainServiceAccountUnauthorized ", 401)
}

func (o *AddTokenToMainServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddTokenToMainServiceAccountForbidden creates a AddTokenToMainServiceAccountForbidden with default headers values
func NewAddTokenToMainServiceAccountForbidden() *AddTokenToMainServiceAccountForbidden {
	return &AddTokenToMainServiceAccountForbidden{}
}

/*AddTokenToMainServiceAccountForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type AddTokenToMainServiceAccountForbidden struct {
}

func (o *AddTokenToMainServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/serviceaccounts/{serviceaccount_id}/tokens][%d] addTokenToMainServiceAccountForbidden ", 403)
}

func (o *AddTokenToMainServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddTokenToMainServiceAccountDefault creates a AddTokenToMainServiceAccountDefault with default headers values
func NewAddTokenToMainServiceAccountDefault(code int) *AddTokenToMainServiceAccountDefault {
	return &AddTokenToMainServiceAccountDefault{
		_statusCode: code,
	}
}

/*AddTokenToMainServiceAccountDefault handles this case with default header values.

errorResponse
*/
type AddTokenToMainServiceAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add token to main service account default response
func (o *AddTokenToMainServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *AddTokenToMainServiceAccountDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/serviceaccounts/{serviceaccount_id}/tokens][%d] addTokenToMainServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *AddTokenToMainServiceAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddTokenToMainServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
