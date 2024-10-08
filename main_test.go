package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bytes"
	"io"
	"os"
)

func TestMain(t *testing.T) {
	mainStdout := captureStdout(main)

	assert.Equal(t, "This is the best Golang program, ever!\n", mainStdout)
}

// Source https://gist.github.com/mindscratch/0faa78bd3c0005d080bf
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
