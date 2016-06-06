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
        "github.com/abourget/looker"
    )

    ...

	lookerClient := looker.New("your-shiny-instance.looker.com", nil)

Then login and do stuff:

	import (
	    "github.com/abourget/looker/client/api_auth"
    )

    ...

	ok, err := lookerClient.APIAuth.Login(&api_auth.LoginParams{
		ClientID:     *myClientID,
		ClientSecret: *myClientSecret,
	})


License
-------

MIT
