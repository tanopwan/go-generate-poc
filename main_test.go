package main_test

import (
	"bytes"
	"net/http"
	"testing"
)

// Project


// Test role Owner, test_owner

// get my projects by user_id success
func TestGetMyProjectsOKByOwner(t *testing.T) {
	expectedBody := `{"body":{"projects":[{"id":"p1","name":"Fizbone","status":"ACTIVE"}]},"code":"success"}`
	req, _ := http.NewRequest("GET", "/projects/", nil)
	req.Header.Set("X-Forwarded-User", "test_owner;2")
	response := executeRequest(req)

	checkResponseCode(t, 200, response.Code)
	checkResponseBody(t, expectedBody, response)
}

// create a new project success
func TestCreateProjectOKByAdmin(t *testing.T) {
	expectedBody := `{"body":{"project":{"name":"new_project_name","created_by":"test_admin","status":"ACTIVE"}},"code":"success"}`
	payload := []byte(`{"name":"new_project_name"}`)
	req, _ := http.NewRequest("POST", "/projects/", bytes.NewBuffer(payload))
	req.Header.Set("X-Forwarded-User", "test_admin;2")
	response := executeRequest(req)

	checkResponseCode(t, 200, response.Code)
	checkResponseBody(t, expectedBody, response)
}


// Test role Admin, test_admin

// get my projects by user_id success
func TestGetMyProjectsOKByOwner(t *testing.T) {
	expectedBody := `{"body":{"projects":[{"id":"p1","name":"Fizbone","status":"ACTIVE"}]},"code":"success"}`
	req, _ := http.NewRequest("GET", "/projects/", nil)
	req.Header.Set("X-Forwarded-User", "test_owner;2")
	response := executeRequest(req)

	checkResponseCode(t, 200, response.Code)
	checkResponseBody(t, expectedBody, response)
}

// create a new project success
func TestCreateProjectOKByAdmin(t *testing.T) {
	expectedBody := `{"body":{"project":{"name":"new_project_name","created_by":"test_admin","status":"ACTIVE"}},"code":"success"}`
	payload := []byte(`{"name":"new_project_name"}`)
	req, _ := http.NewRequest("POST", "/projects/", bytes.NewBuffer(payload))
	req.Header.Set("X-Forwarded-User", "test_admin;2")
	response := executeRequest(req)

	checkResponseCode(t, 200, response.Code)
	checkResponseBody(t, expectedBody, response)
}


