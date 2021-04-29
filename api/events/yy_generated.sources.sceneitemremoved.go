// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemRemoved represents the event body for the "SceneItemRemoved" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SceneItemRemoved.
*/
type SceneItemRemoved struct {
	EventBasic

	// Name of the item removed from the scene.
	ItemName string `json:"item-name"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
