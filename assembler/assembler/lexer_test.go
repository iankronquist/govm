package assembler

import "testing"

func TestTokenize(t *testing.T) {
	testString := `a b c	d  e
		f g
		hh j `
	expectedSlice := []string{"a", "b", "c", "d", "e", "f", "g", "hh", "j"}
	receivedSlice := []string{}
	out := make(chan string)
	go Tokenize(testString, out)
	for token := range out {
		receivedSlice = append(receivedSlice, token)
	}
	if len(expectedSlice) != len(receivedSlice) {
		t.Error("Received slice is the wrong length")
	}
	for index := range expectedSlice {
		if expectedSlice[index] != receivedSlice[index] {
			t.Error("Wrong output from SplitOnWhiteSpace.")
			t.Error("Expected: ", expectedSlice, " Got: ", receivedSlice)
		}
	}
}
