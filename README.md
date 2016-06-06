Auto-generated Looker API v3 bindings for Go
============================================

Based on the
https://yourinstance.looker.com:19999/api/3.0/swagger.json file
retrieved from
https://yourinstance.looker.com:19999/api-docs/index.html

Used https://goswagger.io/ to generate.

See documentation at
https://godoc.org/github.com/abourget/looker/client

WARN: this is API v3, it will not work with API v2.


Setting up your client
----------------------

Get the code and its dependencies:

    go get -v github.com/abourget/looker/client

In your code, get a client:

	import (
        looker "github.com/abourget/looker/client"
	    "github.com/abourget/looker/client/api_auth"
        "github.com/go-openapi/strfmt"
    )

    ...

	lookerClient := looker.New(httptransport.New("your-shiny-instance.looker.com", "/api/3.0", []string{"https"}), strfmt.Default)

Then login and do stuff:

	ok, err := lookerClient.APIAuth.Login(&api_auth.LoginParams{
		ClientID:     *myClientID,
		ClientSecret: *myClientSecret,
	})


License
-------

MIT
