// Package eventlog is the showcase's shared event feed: every form reports
// the events it receives here, and MainForm displays them live.
//
// It exists so each demo form can stay focused on the controls it shows
// while still proving that events actually fire - which is the whole point
// of a showcase.
package eventlog

import (
	"fmt"
	"sync"
)

// Entry is one recorded event.
type Entry struct {
	// Source is the form and control it came from, e.g. "Controls/btnGreet".
	Source string
	// Event is the event name, e.g. "Click".
	Event string
	// Detail is whatever the handler thought was worth showing (coordinates,
	// the new value, ...). It may be empty.
	Detail string
}

// String renders an entry the way MainForm's list shows it.
func (e Entry) String() string {
	if e.Detail == "" {
		return fmt.Sprintf("%-26s %s", e.Source, e.Event)
	}
	return fmt.Sprintf("%-26s %-20s %s", e.Source, e.Event, e.Detail)
}

var (
	mu        sync.Mutex
	listeners []func(Entry)
	motion    bool
)

// MotionEnabled reports whether the high-frequency pointer events are being
// logged; see SetMotionEnabled.
func MotionEnabled() bool {
	mu.Lock()
	defer mu.Unlock()
	return motion
}

// SetMotionEnabled turns MouseMove/MouseWheel logging on. It is off by
// default because those two events fire continuously, and every entry costs
// a repaint of the log.
func SetMotionEnabled(on bool) {
	mu.Lock()
	motion = on
	mu.Unlock()
}

// Subscribe registers a listener. Handlers run on the UI goroutine (that is
// where GoForms raises events), so a listener may touch controls directly.
func Subscribe(fn func(Entry)) {
	mu.Lock()
	defer mu.Unlock()
	listeners = append(listeners, fn)
}

// Log records an event and notifies every listener.
func Log(source, event, detail string) {
	mu.Lock()
	subs := make([]func(Entry), len(listeners))
	copy(subs, listeners)
	mu.Unlock()

	e := Entry{Source: source, Event: event, Detail: detail}
	for _, fn := range subs {
		fn(e)
	}
}

// Logf is Log with a formatted detail.
func Logf(source, event, format string, args ...any) {
	Log(source, event, fmt.Sprintf(format, args...))
}
