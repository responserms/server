package event

import (
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/internal/services/events"
)

// AnySessionStarted is emitted when any session has been started.
func AnySessionStarted() events.Event {
	return events.Event("session.started")
}

// SessionTerminated is emitted when a specific session is terminated.
func SessionTerminated(session *ent.Session) events.Event {
	return events.Event(fmt.Sprintf("session.terminated.%d", session.ID))
}

// AnySessionTerminated is emitted when any session has been terminated.
func AnySessionTerminated() events.Event {
	return events.Event("session.terminated")
}
