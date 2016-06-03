package main

import "testing"

func TestReverseFunction(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHash2(t *testing.T) {
	cases := []struct {
		in1, in2 string
		want string
	}{
		{"hey", "you", "9ba5af89904a2a261dffabd61556860915da1e557cb9a56371140af118f2212f"},
		{"1", "2", "6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443"},
	}
	for _, c := range cases {
		got := Hash2(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Hash2(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

func TestMerkleRoot(t *testing.T) {
	cases := []struct {
		in []string
		want string
	}{
		{[]string{"hey", "you"}, "9ba5af89904a2a261dffabd61556860915da1e557cb9a56371140af118f2212f"},
		{[]string{"1", "2"}, "6f4b6612125fb3a0daecd2799dfd6c9c299424fd920f9b308110a2c1fbd8f443"},
	}
	for _, c := range cases {
		got := MerkleRoot(c.in)
		if got != c.want {
			t.Errorf("MerkleRoot(%q, %q) == %q, want %q", c.in, got, c.want)
		}
	}
}