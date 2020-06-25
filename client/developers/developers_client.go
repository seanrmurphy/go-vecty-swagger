// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the developers client
type API interface {
	/*
	   AddTodo adds an item to the todo list
	   Adds an item to the system
	*/
	AddTodo(ctx context.Context, params *AddTodoParams) (*AddTodoCreated, error)
	/*
	   DeleteTodo deletes a specific todo
	   delete a given todo
	*/
	DeleteTodo(ctx context.Context, params *DeleteTodoParams) (*DeleteTodoOK, error)
	/*
	   GetAllTodos gets todos
	   Gets a list of todos - currently this is universal for all users...
	*/
	GetAllTodos(ctx context.Context, params *GetAllTodosParams) (*GetAllTodosOK, error)
	/*
	   GetTodo gets given todo
	   Gets a specific todo
	*/
	GetTodo(ctx context.Context, params *GetTodoParams) (*GetTodoOK, error)
	/*
	   UpdateTodo updates an item in the todo list
	   updates a given todo
	*/
	UpdateTodo(ctx context.Context, params *UpdateTodoParams) (*UpdateTodoOK, error)
}

// New creates a new developers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for developers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddTodo adds an item to the todo list

Adds an item to the system
*/
func (a *Client) AddTodo(ctx context.Context, params *AddTodoParams) (*AddTodoCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addTodo",
		Method:             "POST",
		PathPattern:        "/todo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddTodoReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddTodoCreated), nil

}

/*
DeleteTodo deletes a specific todo

delete a given todo
*/
func (a *Client) DeleteTodo(ctx context.Context, params *DeleteTodoParams) (*DeleteTodoOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTodo",
		Method:             "DELETE",
		PathPattern:        "/todo/{todoid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTodoReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTodoOK), nil

}

/*
GetAllTodos gets todos

Gets a list of todos - currently this is universal for all users...
*/
func (a *Client) GetAllTodos(ctx context.Context, params *GetAllTodosParams) (*GetAllTodosOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllTodos",
		Method:             "GET",
		PathPattern:        "/todo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllTodosReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllTodosOK), nil

}

/*
GetTodo gets given todo

Gets a specific todo
*/
func (a *Client) GetTodo(ctx context.Context, params *GetTodoParams) (*GetTodoOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTodo",
		Method:             "GET",
		PathPattern:        "/todo/{todoid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTodoReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTodoOK), nil

}

/*
UpdateTodo updates an item in the todo list

updates a given todo
*/
func (a *Client) UpdateTodo(ctx context.Context, params *UpdateTodoParams) (*UpdateTodoOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTodo",
		Method:             "PUT",
		PathPattern:        "/todo/{todoid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTodoReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTodoOK), nil

}
