package httpserv_test

import (
	"gophercise-03-cyoa/infra/httpserv"
	"gophercise-03-cyoa/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDataclient struct {}

func (m mockDataclient) GetByPath(path string) (model.Chapter, bool) {
	return model.Chapter{}, false
}

func TestServeHTTP(t *testing.T) {
	dc := mockDataclient{}
	s := httpserv.StoryHandler{dc}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	
	s.ServeHTTP(writer, req)

	resCode := writer.Result().StatusCode
	if resCode != http.StatusFound {
		t.Errorf("expected redirection to %v, but got %v", http.StatusFound, resCode)
	}
}
