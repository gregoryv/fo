package f

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/fox"
)

func TestTerm(t *testing.T) {
	m := NewTerm()
	ok(t, loggerSet(m))
	ok(t, silentLog(m))

	m.Verbose = true
	bad(t, silentLog(m))

	m.Shf("%s", "touch term_test.go")
	bad(t, unknownCommand(m))
	// output is trimmed
	m.Sh("echo '  hello '")
}

func TestColor(t *testing.T) {
	line := "/home/john"
	ok(t, Color(&line, "/home"))
	bad(t, Color(&line, "/etc"))
}

func TestStrip(t *testing.T) {
	line := "/home/john"
	ok(t, Strip(&line, "/home"))
	line2 := "/home/john"
	bad(t, Strip(&line2, "/etc"))
}

func loggerSet(m *Term) error {
	if m.Logger == nil {
		return fmt.Errorf("Logger is nil")
	}
	return nil
}

func silentLog(m *Term) error {
	var buf bytes.Buffer
	l := fox.NewSyncLog(&buf)
	m.Logger = l
	m.Log("x")
	got := buf.String()
	if got != "" {
		return fmt.Errorf("Default Verbose should be silent")
	}
	return nil
}

func unknownCommand(m *Term) error {
	var failed bool
	m.exit = func(x int) { failed = true }
	m.Shf("%s", "hubladuble")
	if failed {
		return fmt.Errorf("did not fail when executing unknown command")
	}
	return nil
}