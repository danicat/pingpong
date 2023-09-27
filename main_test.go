package main

import (
	"context"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	err := run(ctx, ":7777")
	if err != nil && err != context.DeadlineExceeded {
		t.Fatalf("expected no errors, but got: %s", err)
	}
}
