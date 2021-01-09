package error_types

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDumbGetter(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusTeapot)
	}))
	defer svr.Close()

	_, err := DumbGetter(svr.URL)

	if err != nil {
		t.Fatal("expected an error")
	}

	var got BadStatusError
	//instead using reflection, applying errors.As to extract error type😄
	//err.(BadStatueError)
	isBadStatusError := errors.As(err, &got)
	want := BadStatusError{
		URL:    svr.URL,
		Status: http.StatusTeapot,
	}

	if !isBadStatusError {
		t.Fatalf("was not a BadStatusError, got %T", err)
	}

	if got != want {
		t.Errorf(`got "%v", want "%v"`, got, want)
	}
}
