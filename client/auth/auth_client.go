package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSamlTestConfig creates saml test configuration

### Create a SAML test configuration.

*/
func (a *Client) CreateSamlTestConfig(params *CreateSamlTestConfigParams) (*CreateSamlTestConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSamlTestConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_saml_test_config",
		Method:             "POST",
		PathPattern:        "/saml_test_configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSamlTestConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSamlTestConfigOK), nil
}

/*
DeleteSamlTestConfig deletes saml test configuration

### Delete a SAML test configuration.

*/
func (a *Client) DeleteSamlTestConfig(params *DeleteSamlTestConfigParams) (*DeleteSamlTestConfigNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSamlTestConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_saml_test_config",
		Method:             "DELETE",
		PathPattern:        "/saml_test_configs/{test_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSamlTestConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSamlTestConfigNoContent), nil
}

/*
FetchAndParseSamlIdpMetadata fetches and parse saml idp metadata xml

### Fetch the given url and parse it as a Saml IdP metadata document and return the result.
Note that this requires that the url be public or at least at a location where the Looker instance
can fetch it without requiring any special authentication.

*/
func (a *Client) FetchAndParseSamlIdpMetadata(params *FetchAndParseSamlIdpMetadataParams) (*FetchAndParseSamlIdpMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAndParseSamlIdpMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fetch_and_parse_saml_idp_metadata",
		Method:             "POST",
		PathPattern:        "/fetch_and_parse_saml_idp_metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchAndParseSamlIdpMetadataReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*FetchAndParseSamlIdpMetadataOK), nil
}

/*
LdapConfig gets ldap configuration

### Get the LDAP configuration.

Looker can be optionally configured to authenticate users against an Active Directory or other LDAP directory server.
LDAP setup requires coordination with an administrator of that directory server.

Only Looker administrators can read and update the LDAP configuration.

Configuring LDAP impacts authentication for all users. This configuration should be done carefully.

Looker maintains a single LDAP configuation. It can be read and updated.       Updates only succeed if the new state will be valid (in the sense that all required fields are populated);       it is up to you to ensure that the configuration is appropriate and correct).

LDAP is enabled or disabled for Looker using the **enabled** field.

Looker will never return an **auth_password** field. That value can be set, but never retreived.

See the [Looker LDAP docs]( http://www.looker.com/docs/admin/security/ldap-setup) for additional information.

*/
func (a *Client) LdapConfig(params *LdapConfigParams) (*LdapConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ldap_config",
		Method:             "GET",
		PathPattern:        "/ldap_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*LdapConfigOK), nil
}

/*
ParseSamlIdpMetadata parses saml idp metadata xml

### Parse the given xml as a Saml IdP metadata document and return the result.

*/
func (a *Client) ParseSamlIdpMetadata(params *ParseSamlIdpMetadataParams) (*ParseSamlIdpMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewParseSamlIdpMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "parse_saml_idp_metadata",
		Method:             "POST",
		PathPattern:        "/parse_saml_idp_metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ParseSamlIdpMetadataReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*ParseSamlIdpMetadataOK), nil
}

/*
SamlConfig gets saml configuration

### Get the SAML configuration.

Looker can be optionally configured to authenticate users against a SAML authentication server.
SAML setup requires coordination with an administrator of that server.

Only Looker administrators can read and update the SAML configuration.

Configuring SAML impacts authentication for all users. This configuration should be done carefully.

Looker maintains a single SAML configuation. It can be read and updated.       Updates only succeed if the new state will be valid (in the sense that all required fields are populated);       it is up to you to ensure that the configuration is appropriate and correct).

SAML is enabled or disabled for Looker using the **enabled** field.

*/
func (a *Client) SamlConfig(params *SamlConfigParams) (*SamlConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saml_config",
		Method:             "GET",
		PathPattern:        "/saml_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SamlConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SamlConfigOK), nil
}

/*
SamlTestConfig gets saml test configuration

### Get a SAML test configuration by test_slug.

*/
func (a *Client) SamlTestConfig(params *SamlTestConfigParams) (*SamlTestConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlTestConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saml_test_config",
		Method:             "GET",
		PathPattern:        "/saml_test_configs/{test_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SamlTestConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SamlTestConfigOK), nil
}

/*
TestLdapConfigAuth tests ldap auth config

### Test the connection authentication settings for an LDAP configuration.

This tests that the connection is possible and that a 'server' account to be used by Looker can       authenticate to the LDAP server given connection and authentication information.

**connection_host**, **connection_port**, and **auth_username**, are required.       **connection_tls** and **auth_password** are optional.

Example:
```json
{
  "connection_host": "ldap.example.com",
  "connection_port": "636",
  "connection_tls": true,
  "auth_username": "cn=looker,dc=example,dc=com",
  "auth_password": "secret"
}
```

Looker will never return an **auth_password**. If this request omits the **auth_password** field, then       the **auth_password** value from the active config (if present) will be used for the test.

The active LDAP settings are not modified.


*/
func (a *Client) TestLdapConfigAuth(params *TestLdapConfigAuthParams) (*TestLdapConfigAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestLdapConfigAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test_ldap_config_auth",
		Method:             "PUT",
		PathPattern:        "/ldap_config/test_auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestLdapConfigAuthReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestLdapConfigAuthOK), nil
}

/*
TestLdapConfigConnection tests ldap connection config

### Test the connection settings for an LDAP configuration.

This tests that the connection is possible given a connection_host and connection_port.

**connection_host** and **connection_port** are required. **connection_tls** is optional.

Example:
```json
{
  "connection_host": "ldap.example.com",
  "connection_port": "636",
  "connection_tls": true
}
```

No authentication to the LDAP server is attempted.

The active LDAP settings are not modified.

*/
func (a *Client) TestLdapConfigConnection(params *TestLdapConfigConnectionParams) (*TestLdapConfigConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestLdapConfigConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test_ldap_config_connection",
		Method:             "PUT",
		PathPattern:        "/ldap_config/test_connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestLdapConfigConnectionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestLdapConfigConnectionOK), nil
}

/*
TestLdapConfigUserAuth tests ldap user auth config

### Test the user authentication settings for an LDAP configuration.

This test accepts a full LDAP configuration along with a username/password pair and attempts to       authenticate the user with the LDAP server. The configuration is validated before attempting the       authentication.

Looker will never return an **auth_password**. If this request omits the **auth_password** field, then       the **auth_password** value from the active config (if present) will be used for the test.

**test_ldap_user** and **test_ldap_password** are required.

The active LDAP settings are not modified.


*/
func (a *Client) TestLdapConfigUserAuth(params *TestLdapConfigUserAuthParams) (*TestLdapConfigUserAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestLdapConfigUserAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test_ldap_config_user_auth",
		Method:             "PUT",
		PathPattern:        "/ldap_config/test_user_auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestLdapConfigUserAuthReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestLdapConfigUserAuthOK), nil
}

/*
TestLdapConfigUserInfo tests ldap user info config

### Test the user authentication settings for an LDAP configuration without authenticating the user.

This test will let you easily test the mapping for user properties and roles for any user without      needing to authenticate as that user.

This test accepts a full LDAP configuration along with a username and attempts to find the full info      for the user from the LDAP server without actually authenticating the user. So, user password is not      required.The configuration is validated before attempting to contact the server.

**test_ldap_user** is required.

The active LDAP settings are not modified.


*/
func (a *Client) TestLdapConfigUserInfo(params *TestLdapConfigUserInfoParams) (*TestLdapConfigUserInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestLdapConfigUserInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test_ldap_config_user_info",
		Method:             "PUT",
		PathPattern:        "/ldap_config/test_user_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestLdapConfigUserInfoReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestLdapConfigUserInfoOK), nil
}

/*
UpdateLdapConfig updates ldap configuration

### Update the LDAP configuration.

Configuring LDAP impacts authentication for all users. This configuration should be done carefully.

Only Looker administrators can read and update the LDAP configuration.

LDAP is enabled or disabled for Looker using the **enabled** field.

It is **highly** recommended that any LDAP setting changes be tested using the APIs below before being set globally.

See the [Looker LDAP docs]( http://www.looker.com/docs/admin/security/ldap-setup) for additional information.

*/
func (a *Client) UpdateLdapConfig(params *UpdateLdapConfigParams) (*UpdateLdapConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLdapConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_ldap_config",
		Method:             "PATCH",
		PathPattern:        "/ldap_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLdapConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateLdapConfigOK), nil
}

/*
UpdateSamlConfig updates saml configuration

### Update the SAML configuration.

Configuring SAML impacts authentication for all users. This configuration should be done carefully.

Only Looker administrators can read and update the SAML configuration.

SAML is enabled or disabled for Looker using the **enabled** field.

It is **highly** recommended that any SAML setting changes be tested using the APIs below before being set globally.

*/
func (a *Client) UpdateSamlConfig(params *UpdateSamlConfigParams) (*UpdateSamlConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSamlConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_saml_config",
		Method:             "PATCH",
		PathPattern:        "/saml_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSamlConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSamlConfigOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
