// This file has been automatically generated. Don't edit it.

package events

/*
StreamStatus represents the event body for the "StreamStatus" event.
Since v0.3.
*/
type StreamStatus struct {
	EventBasic

	// Average frame time (in milliseconds)
	AverageFrameTime float64 `json:"average-frame-time"`

	// Amount of data per second (in bytes) transmitted by the stream encoder.
	BytesPerSec int `json:"bytes-per-sec"`

	// Current CPU usage (percentage)
	CpuUsage float64 `json:"cpu-usage"`

	// Current framerate.
	Fps float64 `json:"fps"`

	// Free recording disk space (in megabytes)
	FreeDiskSpace float64 `json:"free-disk-space"`

	// Amount of data per second (in kilobits) transmitted by the stream encoder.
	KbitsPerSec int `json:"kbits-per-sec"`

	// Current RAM usage (in megabytes)
	MemoryUsage float64 `json:"memory-usage"`

	// Number of frames dropped by the encoder since the stream started.
	NumDroppedFrames int `json:"num-dropped-frames"`

	// Total number of frames transmitted since the stream started.
	NumTotalFrames int `json:"num-total-frames"`

	// Number of frames skipped due to encoding lag
	OutputSkippedFrames int `json:"output-skipped-frames"`

	// Number of frames outputted
	OutputTotalFrames int `json:"output-total-frames"`

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`

	// Current recording state.
	Recording bool `json:"recording"`

	// Number of frames missed due to rendering lag
	RenderMissedFrames int `json:"render-missed-frames"`

	// Number of frames rendered
	RenderTotalFrames int `json:"render-total-frames"`

	// Replay Buffer status
	ReplayBufferActive bool `json:"replay-buffer-active"`

	// Percentage of dropped frames.
	Strain float64 `json:"strain"`

	// Current streaming state.
	Streaming bool `json:"streaming"`

	// Total time (in seconds) since the stream started.
	TotalStreamTime int `json:"total-stream-time"`
}
