package pokeapi

import (
	"encoding/json"
	"testing"
)

func TestLocationAreaSerialization(t *testing.T) {
	jsonData := `{
		"count": 2,
		"next": "http://example.com/next",
		"previous": "http://example.com/prev",
		"results": [
			{"name": "area1", "url": "http://example.com/area1"},
			{"name": "area2", "url": "http://example.com/area2"}
		]
	}`

	var locationArea LocationArea

	err := json.Unmarshal([]byte(jsonData), &locationArea)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if locationArea.Count != 2 {
		t.Errorf("Expected count 2, got %d", locationArea.Count)
	}

	if locationArea.Next == nil || *locationArea.Next != "http://example.com/next" {
		t.Errorf("Expected 'next' URL, got %v", locationArea.Next)
	}

	if locationArea.Previous == nil || *locationArea.Previous != "http://example.com/prev" {
		t.Errorf("Expected 'previous' URL, got %v", locationArea.Previous)
	}
}