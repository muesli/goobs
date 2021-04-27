// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileParams struct {
	requests.Params

	// Name of the desired profile.
	ProfileName string `json:"profile-name"`
}

// GetRequestType returns the RequestType of SetCurrentProfileParams
func (o *SetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentProfileParams
func (o *SetCurrentProfileParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentProfileParams
func (o *SetCurrentProfileParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentProfile.
*/
type SetCurrentProfileResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentProfileResponse
func (o *SetCurrentProfileResponse) GetError() string {
	return o.Error
}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(
	params *SetCurrentProfileParams,
) (*SetCurrentProfileResponse, error) {
	params.RequestType = "SetCurrentProfile"
	data := &SetCurrentProfileResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetCurrentProfileParams represents the params body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentProfileParams
func (o *GetCurrentProfileParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentProfileParams
func (o *GetCurrentProfileParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentProfileParams
func (o *GetCurrentProfileParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentProfileResponse represents the response body for the "GetCurrentProfile" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentProfile.
*/
type GetCurrentProfileResponse struct {
	requests.Response

	// Name of the currently active profile.
	ProfileName string `json:"profile-name"`
}

// GetMessageID returns the MessageID of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentProfileResponse
func (o *GetCurrentProfileResponse) GetError() string {
	return o.Error
}

// GetCurrentProfile sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentProfile(
	paramss ...*GetCurrentProfileParams,
) (*GetCurrentProfileResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProfileParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentProfile"
	data := &GetCurrentProfileResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
ListProfilesParams represents the params body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ListProfilesParams
func (o *ListProfilesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ListProfilesParams
func (o *ListProfilesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ListProfilesParams
func (o *ListProfilesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ListProfilesResponse represents the response body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesResponse struct {
	requests.Response

	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}

// GetMessageID returns the MessageID of ListProfilesResponse
func (o *ListProfilesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ListProfilesResponse
func (o *ListProfilesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ListProfilesResponse
func (o *ListProfilesResponse) GetError() string {
	return o.Error
}

// ListProfiles sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ListProfiles(paramss ...*ListProfilesParams) (*ListProfilesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListProfilesParams{{}}
	}
	params := paramss[0]
	params.RequestType = "ListProfiles"
	data := &ListProfilesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}