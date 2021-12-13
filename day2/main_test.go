package main

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	const ok = "forward 5"

	cmd, err := parsecmd(ok)
	if err != nil {
		t.Fatalf("failed to parse cmdline: %s", err)
	}

	if cmd.dir != cmdFWD {
		t.Fatalf("expected dir = forward, got %s", cmd.dir)
	}

	if cmd.change != 5 {
		t.Fatalf("expected change = 5, got %d", cmd.change)
	}

	const bad = "forward a"
	cmd, err = parsecmd(bad)
	if err == nil {
		t.Fatal("expected error from bad input")
	} else {

		if cmd.dir != cmdUNK {
			t.Fatalf("expected cmdUNK from bad input, got %s", cmd.dir)
		}

		if cmd.change != 0 {
			t.Fatalf("expected 0 from bad input, got %d", cmd.change)
		}
	}

	const badcmd = "doot 999"
	cmd, err = parsecmd(bad)
	if err == nil {
		t.Fatal("expected error from bad input")
	}

	if cmd.dir != cmdUNK {
		t.Fatalf("expected cmdUNK from bad input, got %s", cmd.dir)
	}

	if cmd.change != 0 {
		t.Fatalf("expected 0 from bad input, got %d", cmd.change)
	}
}

func TestExample(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/example.input")
	if err != nil {
		t.Fatalf("failed to open example.input: %s", err)
	}

	commands, err := parse(input)
	if err != nil {
		t.Fatalf("expected err = nil, got %s", err)
	}

	if len(commands) != 6 {
		t.Fatalf("expected len(commands) = 6, got %d", len(commands))
	}

	x, y := simulate(0, 0, commands)
	if x != 15 || y != 10 {
		t.Fatalf("expected final pos (15,10), got (%d,%d)", x, y)
	}
}
