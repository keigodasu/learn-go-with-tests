package os_exec

import (
	"strings"
	"testing"
)

func TestGetData(t *testing.T) {
	input := strings.NewReader(`<payload>
    <message>Cats</message>
</payload>`)
	got := GetData(input)
	want := "CATS"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
