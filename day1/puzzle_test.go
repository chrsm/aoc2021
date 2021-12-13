package main

import (
	"io/ioutil"
	"testing"
)

func _assert(t *testing.T) func(bool, string) {
	return func(b bool, msg string) {
		t.Helper()

		if !b {
			t.Fatal(msg)
		}
	}
}

func TestParse(t *testing.T) {
	simple := []byte("1\n2\n3\n4")
	assert := _assert(t)

	r, err := parse(simple)
	assert(err == nil, "failed to parse input simple input")

	assert(r[0] == 1, "rec[0] should be 1")
	assert(r[1] == 2, "rec[1] should be 2")
	assert(r[2] == 3, "rec[2] should be 3")
	assert(r[3] == 4, "rec[3] should be 4")

	annoying := []byte("33\n\n4\n\n\n5\n\n")
	r, err = parse(annoying)
	assert(err == nil, "failed to parse annoying input")
	assert(r[0] == 33, "rec[0] should be 33")
	assert(r[1] == 4, "rec[1] should be 4")
	assert(r[2] == 5, "rec[2] should be 5")

	bad := []byte("a\n\n5")
	_, err = parse(bad)
	assert(err != nil, "should have failed to parse bad input")
}

func TestExample(t *testing.T) {
	assert := _assert(t)

	input, err := ioutil.ReadFile("testdata/example.input")
	if err != nil {
		t.Fatalf("failed to open example.input: %s", err)
	}

	r, err := parse(input)
	assert(err == nil, "failed to parse sample input")

	assert(len(r) == 10, "len(r) should be 10")
	assert(r[0] == 199, "")
	assert(r[1] == 200, "")
	assert(r[2] == 208, "")
	assert(r[3] == 210, "")
	assert(r[4] == 200, "")
	assert(r[5] == 207, "")
	assert(r[6] == 240, "")
	assert(r[7] == 269, "")
	assert(r[8] == 260, "")
	assert(r[9] == 263, "")

	assert(countInc(r) == 7, "should have 7 increases")
}
