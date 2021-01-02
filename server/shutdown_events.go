package server

import "context"

// shutdownEventsService closes the pubsub channels and shuts down the event bus
func (s *Server) shutdownEventsService(ctx context.Context) error {
	s.events.Close(ctx)
	return nil
}
