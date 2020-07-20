// Code generated by go-swagger; DO NOT EDIT.

package repository_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// CreateRepositoryReader is a Reader for the CreateRepository structure.
type CreateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRepositoryOK creates a CreateRepositoryOK with default headers values
func NewCreateRepositoryOK() *CreateRepositoryOK {
	return &CreateRepositoryOK{}
}

/*CreateRepositoryOK handles this case with default header values.

(empty)
*/
type CreateRepositoryOK struct {
	Payload *models.V1alpha1Repository
}

func (o *CreateRepositoryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repositories][%d] createRepositoryOK  %+v", 200, o.Payload)
}

func (o *CreateRepositoryOK) GetPayload() *models.V1alpha1Repository {
	return o.Payload
}

func (o *CreateRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}