// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterRemoved represents the event body for the "SourceFilterRemoved" event.
Since v4.6.0.
*/
type SourceFilterRemoved struct {
	EventBasic

	// Filter name
	FilterName string `json:"filterName"`

	// Filter type
	FilterType string `json:"filterType"`

	// Source name
	SourceName string `json:"sourceName"`
}
