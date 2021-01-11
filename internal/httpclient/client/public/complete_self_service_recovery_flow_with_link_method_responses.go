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

// CompleteSelfServiceRecoveryFlowWithLinkMethodReader is a Reader for the CompleteSelfServiceRecoveryFlowWithLinkMethod structure.
type CompleteSelfServiceRecoveryFlowWithLinkMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceRecoveryFlowWithLinkMethodFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewCompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodFound creates a CompleteSelfServiceRecoveryFlowWithLinkMethodFound with default headers values
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodFound() *CompleteSelfServiceRecoveryFlowWithLinkMethodFound {
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodFound{}
}

/*CompleteSelfServiceRecoveryFlowWithLinkMethodFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceRecoveryFlowWithLinkMethodFound struct {
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodFound) Error() string {
	return fmt.Sprintf("[POST /self-service/recovery/methods/link][%d] completeSelfServiceRecoveryFlowWithLinkMethodFound ", 302)
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest creates a CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest with default headers values
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest() *CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest {
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest{}
}

/*CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest handles this case with default header values.

recoveryFlow
*/
type CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest struct {
	Payload *models.RecoveryFlow
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest) Error() string {
	return fmt.Sprintf("[POST /self-service/recovery/methods/link][%d] completeSelfServiceRecoveryFlowWithLinkMethodBadRequest  %+v", 400, o.Payload)
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest) GetPayload() *models.RecoveryFlow {
	return o.Payload
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecoveryFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError creates a CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError with default headers values
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError() *CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError {
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError{}
}

/*CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/recovery/methods/link][%d] completeSelfServiceRecoveryFlowWithLinkMethodInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
