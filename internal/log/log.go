// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package log

import (
	"fmt"
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
)

// Level is used to configure the logs persisted.
type Level int32

var _ Logger = (*logger)(nil)
var _ ConditionalLogger = (*logger)(nil)
var _ ComponentLogger = (*logger)(nil)

const (
	// Trace is the most verbose level. Intended to be used for the tracing
	// of actions in code, such as function enters/exits, etc.
	Trace Level = iota + 1

	// Debug information for programmer lowlevel analysis.
	Debug

	// Info information about steady state operations.
	Info

	// Warn information about rare but handled events.
	Warn

	// Error information about unrecoverable events.
	Error

	// Off disables all logging output.
	Off
)

// Logger is a basic logging implementation that accepts a simple msg and any
// additional attributes to attach to the log output.
type Logger interface {
	BasicLogger
	ComponentLogger
	ConditionalLogger

	SetLevel(level Level)
	StandardLogger() *log.Logger
	StandardWriter() io.Writer
}

// BasicLogger provides basic logging and attribute-attaching functionality.
type BasicLogger interface {
	Info(msg string, a ...Attributes)
	Debug(msg string, a ...Attributes)
	Error(msg string, a ...Attributes)
	Trace(msg string, a ...Attributes)
	Warn(msg string, a ...Attributes)
}

// ComponentLogger allows a component to name its own logger that is attached to
// the parent. This prefixes the logs with the component name.
type ComponentLogger interface {
	Component(name string, a ...Attributes) Logger
}

// ConditionalLogger allows a caller to check if the log entry would be processed
// to avoid expensive log messages if they won't be processed anyway.
type ConditionalLogger interface {
	IsInfo() bool
	IsDebug() bool
	IsError() bool
	IsTrace() bool
	IsWarn() bool
}

// log is an implementation of ConditionalLogger
type logger struct {
	log hclog.Logger
}

// New creates a new Logger implementation and returns it. The Logger will be preconfigured
// with the given Level.
func New(opts ...OptFunc) (Logger, error) {
	var logger = &logger{
		log: hclog.New(&hclog.LoggerOptions{}),
	}

	for i, opt := range opts {
		if err := opt(logger); err != nil {
			return nil, fmt.Errorf("logger option %d: %w", i, err)
		}
	}

	return logger, nil
}

// Info will log an message if the appropriate level is set.
func (l *logger) Info(msg string, a ...Attributes) {
	l.log.Info(msg, mergeManyAttributes(a...)...)
}

// Debug will log an message if the appropriate level is set.
func (l *logger) Debug(msg string, a ...Attributes) {
	l.log.Debug(msg, mergeManyAttributes(a...)...)
}

// Error will log an message if the appropriate level is set.
func (l *logger) Error(msg string, a ...Attributes) {
	l.log.Error(msg, mergeManyAttributes(a...)...)
}

// Trace will log an message if the appropriate level is set.
func (l *logger) Trace(msg string, a ...Attributes) {
	l.log.Trace(msg, mergeManyAttributes(a...)...)
}

// Warn will log an message if the appropriate level is set.
func (l *logger) Warn(msg string, a ...Attributes) {
	l.log.Warn(msg, mergeManyAttributes(a...)...)
}

// Component creates a new named logger with the given attributes (if any) and returns
// it to be used by the component for its logging purposes.
func (l *logger) Component(name string, a ...Attributes) Logger {
	if len(a) > 0 {
		return &logger{
			log: l.log.With(mergeManyAttributes(a...)...).Named(name),
		}
	}

	return &logger{
		log: l.log.Named(name),
	}
}

// SetLevel sets the level of this logger and all children of this logger.
func (l *logger) SetLevel(level Level) {
	l.setLevel(level)
}

// setLevel sets the level of the underlying hclog instance.
func (l *logger) setLevel(level Level) {
	switch level {
	case Info:
		l.log.SetLevel(hclog.Info)
	case Debug:
		l.log.SetLevel(hclog.Debug)
	case Error:
		l.log.SetLevel(hclog.Error)
	case Trace:
		l.log.SetLevel(hclog.Trace)
	case Warn:
		l.log.SetLevel(hclog.Warn)
	}
}

// IsDebug returns true if a debug message would be processed. This should be used
// to avoid an expensive logging message if possible.
func (l *logger) IsDebug() bool {
	return l.log.IsDebug()
}

// IsError returns true if an error message would be processed. This should be used
// to avoid an expensive logging message if possible.
func (l *logger) IsError() bool {
	return l.log.IsError()
}

// IsInfo returns true if an info message would be processed. This should be used
// to avoid an expensive logging message if possible.
func (l *logger) IsInfo() bool {
	return l.log.IsInfo()
}

// IsTrace returns true if a trace message would be processed. This should be used
// to avoid an expensive logging message if possible.
func (l *logger) IsTrace() bool {
	return l.log.IsTrace()
}

// IsWarn returns true if a warn message would be processed. This should be used
// to avoid an expensive logging message if possible.
func (l *logger) IsWarn() bool {
	return l.log.IsWarn()
}

// StandardLogger returns a standard logger instance that marshals the data into
// our custom logger format.
func (l *logger) StandardLogger() *log.Logger {
	return l.log.StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: true,
	})
}

func (l *logger) StandardWriter() io.Writer {
	return l.log.StandardWriter(&hclog.StandardLoggerOptions{})
}
