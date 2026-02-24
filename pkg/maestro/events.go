package maestro

import (
	"encoding/json"
	"sync"
)

// EventHandler is a callback for push events from the device driver.
type EventHandler func(params json.RawMessage)

// OnEvent registers a handler for the named event.
// Only one handler per event name is supported; a second call replaces the previous.
func (c *Client) OnEvent(event string, handler EventHandler) {
	c.events.Store(event, handler)
}

// RemoveEvent removes the handler for the named event.
func (c *Client) RemoveEvent(event string) {
	c.events.Delete(event)
}

// KeyboardTracker tracks keyboard state from push events.
// Attach it to a Client to get instant keyboard state without polling.
type KeyboardTracker struct {
	mu    sync.RWMutex
	info  KeyboardInfo
	ready bool // true after first event received
}

// NewKeyboardTracker creates a tracker and wires it to the client's
// Input.keyboardStateChanged event.
func NewKeyboardTracker(c *Client) *KeyboardTracker {
	kt := &KeyboardTracker{}
	c.OnEvent("Input.keyboardStateChanged", func(params json.RawMessage) {
		var info KeyboardInfo
		if err := json.Unmarshal(params, &info); err != nil {
			return
		}
		kt.mu.Lock()
		kt.info = info
		kt.ready = true
		kt.mu.Unlock()
	})
	return kt
}

// GetKeyboardInfo returns the latest keyboard state.
// Returns nil if no event has been received yet.
func (kt *KeyboardTracker) GetKeyboardInfo() *KeyboardInfo {
	kt.mu.RLock()
	defer kt.mu.RUnlock()
	if !kt.ready {
		return nil
	}
	info := kt.info
	return &info
}
