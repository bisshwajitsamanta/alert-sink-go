package accept

import (
	"fmt"
	"strings"
	"testing"
)

// This is to test various type of incoming source to see the workability of the function
func TestParseAlerts(t *testing.T) {
	fmt.Println("Testing Fake Json!! ")
	fakeJson := `{
		"version": "4",
		"status": "firing",
		"receiver": "alert-sink",
		"groupLabels": {
			"alertname": "HighCPUUsage"
		},
		"commonLabels": {
			"alertname": "HighCPUUsage",
			"service": "payment-api",
			"severity": "critical",
			"env": "prod"
		},
		"commonAnnotations": {
			"summary": "CPU usage above 90%",
			"description": "The payment-api service is using too much CPU on node prod-node-12"
		},
		"externalURL": "http://alertmanager.example.com",
		"alerts": [
			{
			"status": "firing",
			"labels": {
				"alertname": "HighCPUUsage",
				"instance": "10.10.0.5:9100",
				"job": "node_exporter",
				"service": "payment-api",
				"severity": "critical"
			},
			"annotations": {
				"summary": "CPU usage above 90%",
				"description": "Instance 10.10.0.5 is using >90% CPU for 5m"
			},
			"startsAt": "2025-10-21T10:00:00Z",
			"endsAt": "0001-01-01T00:00:00Z",
			"generatorURL": "http://prometheus.example.com/graph?g0.expr=node_cpu_seconds_total"
			},
			{
			"status": "resolved",
			"labels": {
				"alertname": "HighMemoryUsage",
				"instance": "10.10.0.6:9100",
				"job": "node_exporter",
				"service": "inventory-api",
				"severity": "warning"
			},
			"annotations": {
				"summary": "Memory usage dropped below 75%",
				"description": "Instance 10.10.0.6 memory usage normalized"
			},
			"startsAt": "2025-10-21T08:30:00Z",
			"endsAt": "2025-10-21T09:15:00Z",
			"generatorURL": "http://prometheus.example.com/graph?g0.expr=node_memory_usage_bytes"
			}
		]
	}`
	events, err := ParseAlerts(strings.NewReader(fakeJson))
	t.Run("Not Empty Json File", func(t *testing.T) {
		if err != nil {
			t.Errorf("Valid Json File")
			return
		}
	})
	t.Run("Alerts Not Empty", func(t *testing.T) {
		if len(events) == 0 {
			t.Errorf("Expected Non empty Alerts, but got empty")
		}
	})
	for _, ev := range events {
		ev := ev
		t.Run("Status Check", func(t *testing.T) {
			if ev.Status != "resolved" && ev.Status != "firing" {
				t.Errorf("Expected Status as Resolved or Firing but got '%v'", ev.Status)
			}
		})
		t.Run("Generator URL Check", func(t *testing.T) {
			if len(ev.GeneratorURL) == 0 {
				t.Errorf("Expected Non Empty String, but it is empty '%v'", ev.GeneratorURL)
			}
		})
		t.Run("End Time Stamp is not empty", func(t *testing.T) {
			if ev.Status == "resolved" && ev.EndsAt.IsZero() {
				t.Errorf("Expected End Timestamp to be present, but it is empty")
			}
		})
		t.Run("Start Time Stamp is not empty", func(t *testing.T) {
			if ev.Status == "resolved" && ev.StartsAt.IsZero() {
				t.Errorf("Expected Start Timestamp to be present, but it is empty")
			}
		})
		t.Run("Labels is not empty", func(t *testing.T) {
			if len(ev.Labels) == 0 {
				t.Errorf("Expected Start Timestamp to be present, but it is empty")
			}
		})
	}

}
