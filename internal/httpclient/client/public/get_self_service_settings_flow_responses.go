// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/client/models"
)

// GetSelfServiceSettingsFlowReader is a Reader for the GetSelfServiceSettingsFlow structure.
type GetSelfServiceSettingsFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceSettingsFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceSettingsFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceSettingsFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceSettingsFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceSettingsFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceSettingsFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSelfServiceSettingsFlowOK creates a GetSelfServiceSettingsFlowOK with default headers values
func NewGetSelfServiceSettingsFlowOK() *GetSelfServiceSettingsFlowOK {
	return &GetSelfServiceSettingsFlowOK{}
}

/*GetSelfServiceSettingsFlowOK handles this case with default header values.

settingsFlow
*/
type GetSelfServiceSettingsFlowOK struct {
	Payload *models.SettingsFlow
}

func (o *GetSelfServiceSettingsFlowOK) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/flows][%d] getSelfServiceSettingsFlowOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceSettingsFlowOK) GetPayload() *models.SettingsFlow {
	return o.Payload
}

func (o *GetSelfServiceSettingsFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceSettingsFlowForbidden creates a GetSelfServiceSettingsFlowForbidden with default headers values
func NewGetSelfServiceSettingsFlowForbidden() *GetSelfServiceSettingsFlowForbidden {
	return &GetSelfServiceSettingsFlowForbidden{}
}

/*GetSelfServiceSettingsFlowForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceSettingsFlowForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceSettingsFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/flows][%d] getSelfServiceSettingsFlowForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceSettingsFlowForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceSettingsFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceSettingsFlowNotFound creates a GetSelfServiceSettingsFlowNotFound with default headers values
func NewGetSelfServiceSettingsFlowNotFound() *GetSelfServiceSettingsFlowNotFound {
	return &GetSelfServiceSettingsFlowNotFound{}
}

/*GetSelfServiceSettingsFlowNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceSettingsFlowNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceSettingsFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/flows][%d] getSelfServiceSettingsFlowNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceSettingsFlowNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceSettingsFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceSettingsFlowGone creates a GetSelfServiceSettingsFlowGone with default headers values
func NewGetSelfServiceSettingsFlowGone() *GetSelfServiceSettingsFlowGone {
	return &GetSelfServiceSettingsFlowGone{}
}

/*GetSelfServiceSettingsFlowGone handles this case with default header values.

genericError
*/
type GetSelfServiceSettingsFlowGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceSettingsFlowGone) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/flows][%d] getSelfServiceSettingsFlowGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceSettingsFlowGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceSettingsFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceSettingsFlowInternalServerError creates a GetSelfServiceSettingsFlowInternalServerError with default headers values
func NewGetSelfServiceSettingsFlowInternalServerError() *GetSelfServiceSettingsFlowInternalServerError {
	return &GetSelfServiceSettingsFlowInternalServerError{}
}

/*GetSelfServiceSettingsFlowInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceSettingsFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceSettingsFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/settings/flows][%d] getSelfServiceSettingsFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceSettingsFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceSettingsFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
