package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandlers(t *testing.T) {
	user := []struct {
		name string
		id   string
	}{
		{name: "Mustafo", id: "1"},
		{name: "", id: "2"},
		{name: "Something", id: "string"},
	}

	for _, test := range user {
		// Creating own Request
		request, err := http.NewRequest("GET", "localhost:8080/user/"+test.name, nil)

		if err != nil {
			t.Fatalf("Failed connection ->%q", err)
		}

		// Creating MUX request
		vars := map[string]string{
			"account": test.name,
			"id":      test.id,
		}
		request = mux.SetURLVars(request, vars)
		// Creating type like ResponseWriter
		recorder := httptest.NewRecorder()
		Delete(recorder, request)
		// Getting Final Result after writing
		res := recorder.Result()
		defer res.Body.Close()

		_, err = ioutil.ReadAll(res.Body)

		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			if res.StatusCode == http.StatusBadRequest {
				t.Errorf("expected status %v, got -> %v", http.StatusOK, res.StatusCode)
			}

			t.Errorf("expected status OK; got %v", res.Status)
		}

	}

}
