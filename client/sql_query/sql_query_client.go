package sql_query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sql query API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sql query API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSqlQuery creates s q l runner query

create a sql runner query
*/
func (a *Client) CreateSqlQuery(params *CreateSqlQueryParams) (*CreateSqlQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSqlQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_sql_query",
		Method:             "POST",
		PathPattern:        "/sql_queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSqlQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSqlQueryOK), nil
}

/*
SqlQuery gets s q l runner query

get a sql runner query
*/
func (a *Client) SqlQuery(params *SqlQueryParams) (*SqlQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSqlQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sql_query",
		Method:             "GET",
		PathPattern:        "/sql_queries/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SqlQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SqlQueryOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
