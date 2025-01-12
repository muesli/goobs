// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DeleteSceneItemParams represents the params body for the "DeleteSceneItem" request.
Deletes a scene item.
Since 4.5.0.
*/
type DeleteSceneItemParams struct {
	requests.ParamsBasic

	Item struct {
		// Scene Item ID.
		Id int `json:"id"`

		// Scene Item name (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the scene item belongs to. Defaults to the current scene.
	Scene string `json:"scene"`
}

// GetSelfName just returns "DeleteSceneItem".
func (o *DeleteSceneItemParams) GetSelfName() string {
	return "DeleteSceneItem"
}

/*
DeleteSceneItemResponse represents the response body for the "DeleteSceneItem" request.
Deletes a scene item.
Since v4.5.0.
*/
type DeleteSceneItemResponse struct {
	requests.ResponseBasic
}

// DeleteSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DeleteSceneItem(params *DeleteSceneItemParams) (*DeleteSceneItemResponse, error) {
	data := &DeleteSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
