package rand

import (
	"strings"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	b, err := randomBytes(8)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) != 8 {
		t.Fatalf("Expected randomBytes to return size 8 but got %d", len(b))
	}
}

func TestGenerateRandomString(t *testing.T) {
	b, err := randomBytes(8)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != 8 {
		t.Fatalf("Expected randomBytes to return size 8 but got %d", len(b))
	}

	strSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str := generateRandomString(8, strSet, b)
	if len(str) != 8 {
		t.Fatalf("Expected generateRandomString to return size 8 but got %d", len(str))
	}
	for _, s := range str {
		if !strings.Contains(strSet, string(s)) {
			t.Fatalf("Expected string to be a subset of %q but got %q", strSet, s)
		}
	}
}

func TestRandomString(t *testing.T) {
	strSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str, err := RandomString(8, strSet)
	if err != nil {
		t.Fatal(err)
	}
	if len(str) != 8 {
		t.Fatalf("Expected generateRandomString to return size 8 but got %d", len(str))
	}
	for _, s := range str {
		if !strings.Contains(strSet, string(s)) {
			t.Fatalf("Expected string to be a subset of %q but got %q", strSet, s)
		}
	}
}
