package http_server

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContesnts, _ := ioutil.ReadAll(file)

	got := string(newFileContesnts)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}