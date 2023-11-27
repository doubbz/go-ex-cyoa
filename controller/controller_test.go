package controller_test

import (
	"gophercise-03-cyoa/controller"
	"gophercise-03-cyoa/model"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type mockDataclient struct {
	mockGetByPath func() (model.Chapter, bool)
}

func (m mockDataclient) GetByPath(path string) (model.Chapter, bool) {
	return m.mockGetByPath()
}

func TestRedirectServeHTTP(t *testing.T) {
	mockedGetByPath := func () (model.Chapter, bool) {
		return model.Chapter{}, false
	}

	dc := mockDataclient{mockedGetByPath}
	tmpl := template.New("mock")
	s := controller.StoryHandler{dc, tmpl}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	
	s.ServeHTTP(writer, req)

	response := writer.Result()
	if response.StatusCode != http.StatusFound {
		t.Errorf("expected redirection to %v, but got %v", http.StatusFound, response.StatusCode)
	}

	body, _ := io.ReadAll(response.Body) // test package, thus no need to worry about performance
	if res, _ := regexp.MatchString("/intro", string(body)); res == false {
		t.Errorf("body: expected to contain /intro, got %s", string(body))
	}
}

func TestTemplateServeHTTP(t *testing.T) {
	mockedGetByPath := func () (model.Chapter, bool) {
		return model.Chapter{
			Name: "/mock-intro",
			Title: "My mock intro",
			Story: []string{"Super story"},
			Options: []model.Options{},
		}, true
	}

	dc := mockDataclient{mockedGetByPath}
	tmpl, _ := template.ParseFiles("../templates/layout.html")
	s := controller.StoryHandler{dc, tmpl}
	req := httptest.NewRequest(http.MethodGet, "/mock-intro", nil)
	writer := httptest.NewRecorder()
	
	s.ServeHTTP(writer, req)

	response := writer.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected OK status %v, but got %v", http.StatusOK, response.StatusCode)
	}

	body, _ := io.ReadAll(response.Body)
	if res, _ := regexp.MatchString("My mock intro", string(body)); res == false {
		t.Errorf("body: expected to contain %s, got %s", "My mock intro", string(body))
	}
	if res, _ := regexp.MatchString("Super story", string(body)); res == false {
		t.Errorf("body: expected to contain %s, got %s", "Super story", string(body))
	}
}
