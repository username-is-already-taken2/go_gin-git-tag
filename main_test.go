package main

import (
	"encoding/json"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestRootGet(t *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})
	t.Run("Returns expected reposonse", func(t *testing.T) {
		if recorder.Body.String() != "Hello, Gin!" {
			t.Error("Expected '\"Hello, Gin!\"', got ", recorder.Body.String())
		}
	})
}

func TestHealthGet(t *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/health", nil))
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})
	t.Run("Returns expected reposonse", func(t *testing.T) {
		var body map[string]interface{}
		json.Unmarshal(recorder.Body.Bytes(), &body)
		matched, _ := regexp.MatchString("^OK", body["status"].(string))
		if !matched {
			t.Error("Expected object with status OK, got ", body)
		}
	})
}

func TestVersionGet(t *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/version", nil))
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})
	t.Run("Returns expected reposonse", func(t *testing.T) {
		var body map[string]interface{}
		json.Unmarshal(recorder.Body.Bytes(), &body)
		matched, _ := regexp.MatchString("^development", body["version"].(string))
		if !matched {
			t.Error("Expected object with status, got ", body)
		}
	})
}
