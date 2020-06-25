// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/seanrmurphy/go-vecty-swagger/models"
)

// GetTodoReader is a Reader for the GetTodo structure.
type GetTodoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTodoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTodoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTodoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTodoOK creates a GetTodoOK with default headers values
func NewGetTodoOK() *GetTodoOK {
	return &GetTodoOK{}
}

/*GetTodoOK handles this case with default header values.

get given todo
*/
type GetTodoOK struct {
	Payload *models.Todo
}

func (o *GetTodoOK) Error() string {
	return fmt.Sprintf("[GET /todo/{todoid}][%d] getTodoOK  %+v", 200, o.Payload)
}

func (o *GetTodoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Todo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTodoNotFound creates a GetTodoNotFound with default headers values
func NewGetTodoNotFound() *GetTodoNotFound {
	return &GetTodoNotFound{}
}

/*GetTodoNotFound handles this case with default header values.

item not found
*/
type GetTodoNotFound struct {
}

func (o *GetTodoNotFound) Error() string {
	return fmt.Sprintf("[GET /todo/{todoid}][%d] getTodoNotFound ", 404)
}

func (o *GetTodoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
