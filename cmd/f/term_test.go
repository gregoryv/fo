package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/gregoryv/fox"
)

func TestTerm(t *testing.T) {
	m := NewTerm()
	ok, _k := assert(t)
	ok(silentLog(m))

	m.Logger = t
	m.exit = func(int) {}
	m.Verbose = true

	_k(silentLog(m))
	m.Shf("%s", "touch term_test.go")
	_k(m.Shf("%s", "hubladuble"))
	// output is trimmed
	m.Sh("echo '  hello '")
}

func TestColor(t *testing.T) {
	line := "/home/john"
	ok, _k := assert(t)
	ok(Color(&line, "/home"))
	_k(Color(&line, "/etc"))
}

func TestStrip(t *testing.T) {
	line := "/home/john"
	ok, _k := assert(t)
	ok(Strip(&line, "/home"))

	line2 := "/home/john"
	_k(Strip(&line2, "/etc"))
}

func TestDefaultTerm(t *testing.T) {
	SetOutput(ioutil.Discard)
	Verbose()
	NoExit()
	DefaultTerm.Verbose = false //so below output doesn't show
	Sh("whohooo ")
	Sh("touch x")
	Shf("%s %s", "touch", "x")
	Sh("rm x")
}

// ----------------------------------------

func TestRunCmd(t *testing.T) {
	ok, _k := assert(t)
	ok(RunCmd(exec.Command("echo")))
	_k(RunCmd(exec.Command("")))
}

// ----------------------------------------

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
