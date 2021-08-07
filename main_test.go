package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestIsHelp(t *testing.T) {
	if !isHelp([]string{"-h"}) {
		t.Fatal("Expected isHelp to return true but got false")
	}
	if !isHelp([]string{"--help"}) {
		t.Fatal("Expected isHelp to return true but got false")
	}
	if isHelp([]string{"12"}) {
		t.Fatal("Expected isHelp to return false but got true")
	}
}

func TestGetHelp(t *testing.T) {
	h := getHelp("rand-str")
	r := "usage: rand-str <string_len> <string_set>\n"
	if h != r {
		t.Fatalf("getHelp expected %q but got %q", r, h)
	}
}

func TestParseArgs(t *testing.T) {
	siz, str, err := parseArgs([]string{})
	if err != nil {
		t.Fatal(err)
	}
	if siz != size {
		t.Fatalf("parseArgs expected %d but got %d", size, siz)
	}
	if str != strSet {
		t.Fatalf("parseArgs expected %q but got %q", strSet, str)
	}

	_, _, err = parseArgs([]string{"sdasdas"})
	if err == nil {
		t.Fatal("parseArgs should return error but got nil")
	}

	siz, _, err = parseArgs([]string{"5"})
	if err != nil {
		t.Fatal(err)
	}
	if siz != 5 {
		t.Fatalf("parseArgs expected %d but got %d", 5, siz)
	}

	siz, str, err = parseArgs([]string{"6", "ABC"})
	if err != nil {
		t.Fatal(err)
	}
	if siz != 6 {
		t.Fatalf("parseArgs expected %d but got %d", 6, siz)
	}
	if str != "ABC" {
		t.Fatalf("parseArgs expected %q but got %q", "ABC", str)
	}
}

func TestMainExitError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "adasda", "", "-test.run=TestMainExitError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMainSuccess(t *testing.T) {
	if os.Getenv("NO_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "", "", "-test.run=TestMainSuccess")
	cmd.Env = append(os.Environ(), "NO_CRASHER=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("process ran with err %#v, want exit status 0", err)
	}
}
