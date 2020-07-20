// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// UpdateSpecReader is a Reader for the UpdateSpec structure.
type UpdateSpecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSpecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSpecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSpecOK creates a UpdateSpecOK with default headers values
func NewUpdateSpecOK() *UpdateSpecOK {
	return &UpdateSpecOK{}
}

/*UpdateSpecOK handles this case with default header values.

(empty)
*/
type UpdateSpecOK struct {
	Payload *models.V1alpha1ApplicationSpec
}

func (o *UpdateSpecOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/applications/{name}/spec][%d] updateSpecOK  %+v", 200, o.Payload)
}

func (o *UpdateSpecOK) GetPayload() *models.V1alpha1ApplicationSpec {
	return o.Payload
}

func (o *UpdateSpecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1ApplicationSpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}