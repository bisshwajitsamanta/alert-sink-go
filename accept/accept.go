package accept

import (
	"encoding/json"
	"io"
	"time"
)

type Events struct {
	Receiver string   `json:"receiver"`
	Alerts   []Alerts `json:"alerts"`
}

type Alerts struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	GeneratorURL string            `json:"generatorURL"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
}

type SingleEvent struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	GeneratorURL string            `json:"generatorURL"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	Received     time.Time
}

// Parse Alerts - Job is to Parse the incoming alerts from webhook
func ParseAlerts(jsonStream io.Reader) ([]SingleEvent, error) {
	// Decode Json Message
	decodeIncomingMessage := json.NewDecoder(jsonStream)
	var decodedMessage Events
	err := decodeIncomingMessage.Decode(&decodedMessage)

	//Empty Slice of Events
	events := []SingleEvent{}

	// Loop over Decoded Message Alerts
	for _, msg := range decodedMessage.Alerts {
		e := SingleEvent{
			Status:       msg.Status,
			Labels:       msg.Labels,
			GeneratorURL: msg.GeneratorURL,
			StartsAt:     msg.StartsAt,
			EndsAt:       msg.EndsAt,
			Received:     time.Now(),
		}
		events = append(events, e)
	}
	return events, err
}

// Handle Alerts - Accepts the alerts and routes to appropiate Channel.
// func HandleAlerts() {

// }
