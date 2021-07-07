package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_GetTechnologies(t *testing.T) {
	app := App{}

	r, _ := http.NewRequest("GET", "/spiral?rows=3&cols=3", nil)
	w := httptest.NewRecorder()

	app.GetSpiral(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	want := `{"rows": [[1, 1, 2], [3, 5, 8], [13, 21, 34]]}` + "\n"
	if got := w.Body.String(); got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}
