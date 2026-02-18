package inventory

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	var buf bytes.Buffer
	notifier := WriteTo{&buf}
	notifier.Notify("testing")

	got := buf.String()
	want := "testing"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
