package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new query API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for query API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateQuery creates query

### Create a query.

This allows you to create a new query that you can later run. Looker queries are immutable once created
and are not deleted. If you create a query that is exactly like an existing query then the existing query
will be returned and no new query will be created. Whether a new query is created or not, you can use
the 'id' in the returned query with the 'run' method.

The query parameters are passed as json in the body of the request.


*/
func (a *Client) CreateQuery(params *CreateQueryParams) (*CreateQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_query",
		Method:             "POST",
		PathPattern:        "/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateQueryOK), nil
}

/*
CreateQueryAndRunAsync creates a query and run it asynchronously

Given a set of properties for a Query, fetches or creates the Query and starts running it asynchronously.
The Query Task results endpoint can be used to fetch the query.

*/
func (a *Client) CreateQueryAndRunAsync(params *CreateQueryAndRunAsyncParams) (*CreateQueryAndRunAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQueryAndRunAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_query_and_run_async",
		Method:             "POST",
		PathPattern:        "/queries/run_async",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQueryAndRunAsyncReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateQueryAndRunAsyncOK), nil
}

/*
Query gets query

### Get a previously created query by id.

A Looker query object includes the various parameters that define a database query that has been run or
could be run in the future. These parameters include: model, view, fields, filters, pivots, etc.
Query *results* are not part of the query object.

Query objects are unique and immutable. Query objects are created automatically in Looker as users explore data.
Looker does not delete them; they become part of the query history. When asked to create a query for
any given set of parameters, Looker will first try to find an existing query object with matching
parameters and will only create a new object when an appropriate object can not be found.

This 'get' method is used to get the details about a query for a given id. See the other methods here
to 'create' and 'run' queries.

Note that some fields like 'filter_config' and 'vis_config' etc are specific to how the Looker UI
builds queries and visualizations and are not generally useful for API use. They are not required when
creating new queries and can usually just be ignored.


*/
func (a *Client) Query(params *QueryParams) (*QueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query",
		Method:             "GET",
		PathPattern:        "/queries/{query_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryOK), nil
}

/*
QueryForSlug gets query for slug

### Get the query for a given query slug.

This returns the query for the 'slug' in a query share URL.

The 'slug' is a randomly chosen short string that is used as an alternative to the query's id value
for use in URLs etc. This method exists as a convenience to help you use the API to 'find' queries that
have been created using the Looker UI.

You can use the Looker explore page to build a query and then choose the 'Share' option to
show the share url for the query. Share urls generally look something like 'https://looker.yourcompany/x/vwGSbfc'.
The trailing 'vwGSbfc' is the share slug. You can pass that string to this api method to get details about the query.
Those details include the 'id' that you can use to run the query. Or, you can copy the query body
(perhaps with your own modification) and use that as the basis to make/run new queries.


*/
func (a *Client) QueryForSlug(params *QueryForSlugParams) (*QueryForSlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryForSlugParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query_for_slug",
		Method:             "GET",
		PathPattern:        "/queries/slug/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryForSlugReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryForSlugOK), nil
}

/*
QueryTask gets a query task

Fetch a Query Task.

Query Tasks are generated by running queries asynchronously. They are represented by a GUID returned
from one of the async query endpoints.

*/
func (a *Client) QueryTask(params *QueryTaskParams) (*QueryTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query_task",
		Method:             "GET",
		PathPattern:        "/query_tasks/{query_task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryTaskReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryTaskOK), nil
}

/*
QueryTaskMultiResults gets multiple query task results in one request

Fetch a multiple Query Task results at once.

*/
func (a *Client) QueryTaskMultiResults(params *QueryTaskMultiResultsParams) (*QueryTaskMultiResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTaskMultiResultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query_task_multi_results",
		Method:             "GET",
		PathPattern:        "/query_tasks/multi_results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryTaskMultiResultsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryTaskMultiResultsOK), nil
}

/*
QueryTaskResults gets a query task s results

Fetch a Query Task's results.

*/
func (a *Client) QueryTaskResults(params *QueryTaskResultsParams) (*QueryTaskResultsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTaskResultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query_task_results",
		Method:             "GET",
		PathPattern:        "/query_tasks/{query_task_id}/results",
		ProducesMediaTypes: []string{"application/json", "text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryTaskResultsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryTaskResultsNoContent), nil
}

/*
RunAsync runs a query asynchronously

### Run a saved query asynchronously.

Runs a previously created query asynchronously. Returns a Query Task ID
which can be used to fetch the results from the Query Tasks results endpoint.


*/
func (a *Client) RunAsync(params *RunAsyncParams) (*RunAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "run_async",
		Method:             "POST",
		PathPattern:        "/queries/{query_id}/run_async",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunAsyncReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunAsyncOK), nil
}

/*
RunInlineQuery runs inline query

### Run the query that is specified inline in the posted body.

This allows running a query as defined in json in the posted body. This combines
the two actions of posting & running a query into one step.

Here is an example body in json:
```
{
  "model":"thelook",
  "view":"inventory_items",
  "fields":["category.name","inventory_items.days_in_inventory_tier","products.count"],
  "filters":{"category.name":"socks"},
  "sorts":["products.count desc 0"],
  "limit":"500",
  "query_timezone":"America/Los_Angeles"
}
```

When using the Ruby SDK this would be passed as a Ruby hash like:
```
{
 :model=>"thelook",
 :view=>"inventory_items",
 :fields=>
  ["category.name",
   "inventory_items.days_in_inventory_tier",
   "products.count"],
 :filters=>{:"category.name"=>"socks"},
 :sorts=>["products.count desc 0"],
 :limit=>"500",
 :query_timezone=>"America/Los_Angeles",
}
```

This will return the result of running the query in the format specified by the 'format' paramater.

Suported formats:
- json - plain json
- csv - comma separated values with a header
- txt - tab separated values with a header
- html - simple html
- md - simple markdown
- sql - shows the generated SQL rather than running the query
- png - a PNG image of the visualization of the query
- jpg - a JPG image of the visualization of the query
- unified - json that is annotated with additional metadata as used by the Looker web application



*/
func (a *Client) RunInlineQuery(params *RunInlineQueryParams) (*RunInlineQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunInlineQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "run_inline_query",
		Method:             "POST",
		PathPattern:        "/queries/run/{format}",
		ProducesMediaTypes: []string{"application/json", "image/jpg", "image/png", "text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunInlineQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunInlineQueryOK), nil
}

/*
RunQuery runs query

### Run a saved query.

This runs a previously saved query. You can use this on a query that was generated in the Looker UI
or one that you have explicitly created using the API. You can also use a query 'id' from a saved 'Look'.

This will return the result of running the query in the format specified by the 'format' paramater.

Suported formats:
- json - plain json
- csv - comma separated values with a header
- txt - tab separated values with a header
- html - simple html
- md - simple markdown
- sql - shows the generated SQL rather than running the query
- png - a PNG image of the visualization of the query
- jpg - a JPG image of the visualization of the query
- unified - json that is annotated with additional metadata as used by the Looker web application



*/
func (a *Client) RunQuery(params *RunQueryParams) (*RunQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "run_query",
		Method:             "GET",
		PathPattern:        "/queries/{query_id}/run/{format}",
		ProducesMediaTypes: []string{"application/json", "image/jpg", "image/png", "text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunQueryOK), nil
}

/*
RunURLEncodedQuery runs url encoded query

### Run an URL encoded query.

This requires the caller to encode the specifiers for the query into the URL query part using
Looker-specific syntax as explained below.

Generally, you would want to use one of the methods that takes the parameters as json in the POST body
for creating and/or running queries. This method exists for cases where one really needs to encode the
paramaters into the URL of a single 'GET' request. This matches the way that the Looker UI formats
'explore' URLs etc.

The parameters here are very similar to the json body formatting except that the filter syntax is
tricky. Unfortunately, this format makes this method not currently callible via the 'Try it out!' button
in this documentation page. But, this is callable  when creating URLs manually or when using the Looker SDK.

Here is an example inline query URL:

"https://looker.mycompany.com:19999/api/3.0/queries/models/thelook/views/inventory_items/run/json?fields=category.name,inventory_items.days_in_inventory_tier,products.count&f[category.name]=socks&sorts=products.count+desc+0&limit=500&query_timezone=America/Los_Angeles"

When calling the Ruby SDK, one passes the query parameter parts as a hash. The hash to match the above would like like:

```
{
  :fields => "category.name,inventory_items.days_in_inventory_tier,products.count",
  :"f[category.name]" => "socks",
  :sorts => "products.count desc 0",
  :limit => "500",
  :query_timezone => "America/Los_Angeles"
}
```

Again, it is generally easier to use the variant of this method that passes the full query in the POST body.
This method is available for cases where other alternatives won't fit the need.

Suported formats:
- json - plain json
- csv - comma separated values with a header
- txt - tab separated values with a header
- html - simple html
- md - simple markdown
- sql - shows the generated SQL rather than running the query
- png - a PNG image of the visualization of the query
- jpg - a JPG image of the visualization of the query
- unified - json that is annotated with additional metadata as used by the Looker web application



*/
func (a *Client) RunURLEncodedQuery(params *RunURLEncodedQueryParams) (*RunURLEncodedQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunURLEncodedQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "run_url_encoded_query",
		Method:             "GET",
		PathPattern:        "/queries/models/{model_name}/views/{view_name}/run/{format}",
		ProducesMediaTypes: []string{"application/json", "image/jpg", "image/png", "text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunURLEncodedQueryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunURLEncodedQueryOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
