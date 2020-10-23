package logging

import (
	"fmt"
	"strings"
	"testing"
)

func TestPanicRecorder(t *testing.T) {
	rec := panics.registerPlugin("test")

	output := []string{
		"panic: test",
		"  stack info",
	}

	for _, line := range output {
		rec(line)
	}

	expected := fmt.Sprintf(pluginPanicOutput, "test", strings.Join(output, "\n"))

	res := PluginPanics()
	if len(res) == 0 {
		t.Fatal("no output")
	}

	if res[0] != expected {
		t.Fatalf("expected: %q\ngot: %q", expected, res[0])
	}
}

func TestPanicLimit(t *testing.T) {
	rec := panics.registerPlugin("test")

	rec("panic: test")

	for i := 0; i < 200; i++ {
		rec(fmt.Sprintf("LINE: %d", i))
	}

	res := PluginPanics()
	// take the extra formatting into account
	if len(res) > panics.maxLines+15 {
		t.Fatalf("expected no more than %d lines, got: %d", panics.maxLines+50, len(res))
	}
}
