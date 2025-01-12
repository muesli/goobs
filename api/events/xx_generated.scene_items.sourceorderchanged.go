// This file has been automatically generated. Don't edit it.

package events

/*
SourceOrderChanged represents the event body for the "SourceOrderChanged" event.
Since v4.0.0.
*/
type SourceOrderChanged struct {
	EventBasic

	SceneItems []struct {
		// Scene item unique ID
		ItemId int `json:"item-id"`

		// Item source name
		SourceName string `json:"source-name"`
	} `json:"scene-items"`

	// Name of the scene where items have been reordered.
	SceneName string `json:"scene-name"`
}
