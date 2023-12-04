package main

import "testing"

func TestIntFinder(t *testing.T) {
	input := "fourzvkqhdninetwoftscrmsd64nxsgx"
	wanted := 64
	actual := intFinder(input)
	if actual != wanted {
		t.Errorf("got %d, wanted %d\n", actual, wanted)
	}
}

type Test struct {
	Input  string
	Wanted int
}

var table = []Test{
	{Input: "two1nine", Wanted: 29},
	{Input: "eightwothree", Wanted: 83},
	{Input: "abcone2threexyz", Wanted: 13},
	{Input: "xtwone3four", Wanted: 24},
	{Input: "4nineeightseven2", Wanted: 42},
	{Input: "zoneight234", Wanted: 14},
	{Input: "7pqrstsixteen", Wanted: 76},
	{Input: "sixsix5qmbqd", Wanted: 65},
	{Input: "5eight82sixtwonev", Wanted: 51},
}

func TestIntFinderV2_a(t *testing.T) {
	for _, x := range table {
		actual := intFinderV2(x.Input)
		if actual != x.Wanted {
			t.Errorf("got %d, wanted %d\n", actual, x.Wanted)
		}
	}
}
