// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trace

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Span is a mock span used in association with Tracer for
// testing purpose only.
type Span struct {
	sc            trace.SpanContext
	tracer        *Tracer
	Name          string
	Attributes    map[label.Key]label.Value
	Kind          trace.SpanKind
	Status        codes.Code
	StatusMessage string
	ParentSpanID  trace.SpanID
	Links         map[trace.SpanContext][]label.KeyValue
}

var _ trace.Span = (*Span)(nil)

// SpanContext returns associated trace.SpanContext.
//
// If the receiver is nil it returns an empty trace.SpanContext.
func (ms *Span) SpanContext() trace.SpanContext {
	if ms == nil {
		return trace.SpanContext{}
	}
	return ms.sc
}

// IsRecording always returns false for Span.
func (ms *Span) IsRecording() bool {
	return false
}

// SetStatus sets the Status member.
func (ms *Span) SetStatus(status codes.Code, msg string) {
	ms.Status = status
	ms.StatusMessage = msg
}

// SetAttribute adds a single inferred attribute.
func (ms *Span) SetAttribute(key string, value interface{}) {
	ms.SetAttributes(label.Any(key, value))
}

// SetAttributes adds an attribute to Attributes member.
func (ms *Span) SetAttributes(attributes ...label.KeyValue) {
	if ms.Attributes == nil {
		ms.Attributes = make(map[label.Key]label.Value)
	}
	for _, kv := range attributes {
		ms.Attributes[kv.Key] = kv.Value
	}
}

// End puts the span into tracers ended spans.
func (ms *Span) End(options ...trace.SpanOption) {
	ms.tracer.addEndedSpan(ms)
}

// RecordError does nothing.
func (ms *Span) RecordError(err error, opts ...trace.EventOption) {
}

// SetName sets the span name.
func (ms *Span) SetName(name string) {
	ms.Name = name
}

// Tracer returns the mock tracer implementation of Tracer.
func (ms *Span) Tracer() trace.Tracer {
	return ms.tracer
}

// AddEvent does nothing.
func (ms *Span) AddEvent(name string, opts ...trace.EventOption) {
}

// AddEvent does nothing.
func (ms *Span) AddEventWithTimestamp(ctx context.Context, timestamp time.Time, name string, attrs ...label.KeyValue) {
}
