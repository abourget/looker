package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllLegacyFeatures gets all legacy features

### Get all legacy features.
*/
func (a *Client) AllLegacyFeatures(params *AllLegacyFeaturesParams) (*AllLegacyFeaturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllLegacyFeaturesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_legacy_features",
		Method:             "GET",
		PathPattern:        "/legacy_features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllLegacyFeaturesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllLegacyFeaturesOK), nil
}

/*
AllTimezones gets all timezones

### Get a list of timezones that Looker supports (e.g. useful for scheduling tasks).

*/
func (a *Client) AllTimezones(params *AllTimezonesParams) (*AllTimezonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllTimezonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_timezones",
		Method:             "GET",
		PathPattern:        "/timezones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllTimezonesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllTimezonesOK), nil
}

/*
BackupConfiguration gets backup configuration

### Get the current backup configuration

*/
func (a *Client) BackupConfiguration(params *BackupConfigurationParams) (*BackupConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backup_configuration",
		Method:             "GET",
		PathPattern:        "/backup_configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupConfigurationReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackupConfigurationOK), nil
}

/*
LegacyFeature gets legacy feature

### Get information about the legacy feature with a specific id.
*/
func (a *Client) LegacyFeature(params *LegacyFeatureParams) (*LegacyFeatureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLegacyFeatureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "legacy_feature",
		Method:             "GET",
		PathPattern:        "/legacy_features/{legacy_feature_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LegacyFeatureReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*LegacyFeatureOK), nil
}

/*
UpdateBackupConfiguration updates backup configuration

### Update Looker Backup Configuration

*/
func (a *Client) UpdateBackupConfiguration(params *UpdateBackupConfigurationParams) (*UpdateBackupConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBackupConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_backup_configuration",
		Method:             "PATCH",
		PathPattern:        "/backup_configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBackupConfigurationReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateBackupConfigurationOK), nil
}

/*
UpdateLegacyFeature updates legacy feature

### Update information about the legacy feature with a specific id.
*/
func (a *Client) UpdateLegacyFeature(params *UpdateLegacyFeatureParams) (*UpdateLegacyFeatureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLegacyFeatureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_legacy_feature",
		Method:             "PATCH",
		PathPattern:        "/legacy_features/{legacy_feature_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLegacyFeatureReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateLegacyFeatureOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
