package hypem

import (
	"testing"
)

func Test_Stream(t *testing.T) {
	url, err := Stream("2a7n7")
	if err != nil {
		t.Error(err)
		return
	}
	if url == "" {
		t.Error("Unknown error: stream url was empty.")
		return
	}
}
