package main

import (
	"context"
	"log"
)

// ContextLog makes a log with the trace ID as a prefix.
func ContextLog(ctx context.Context, f string, args ...interface{}) {

	traceID, ok := ctx.Value(TraceKey).(string)

	if ok && traceID != "" {
		f = traceID + ": " + f
	}

	log.Printf(f, args...)
}

// We pass the TraceID in the program through the context from privatecontextkeytype.go and look it up like a map, check if it's string convertible before adding the traceID to the rest of the string and prrinting it
