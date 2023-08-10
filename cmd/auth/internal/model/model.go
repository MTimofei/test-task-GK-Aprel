package model

type Audit struct {
	Event_time string `json:"event_timestamp"`
	Event_type string `json:"event_type"`
}
