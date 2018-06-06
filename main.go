package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

//go:generate go run main.go

type file struct {
	Name  string
	Tests []roleTest
}

type testResult struct {
	Status int
	Body   string
}

type testcase struct {
	Name   string
	Method string
	Desc   string
	Path   string
	Body   string
}

type roleTest struct {
	UserID    string
	RoleName  string
	Testcases []testcase
	Results   []testResult
}

var page = `
import (
	"bytes"
	"net/http"
	"testing"
)

// {{.Name}}
{{template "funcs" .Tests}}
`

var funcs = `
{{define "funcs"}}
{{range .}}
// Test role {{.RoleName}}, {{.UserID}}
{{range $i, $e := .Testcases}}{{with $roleTest := index $ $i}}{{with $result := index $roleTest.Results $i}}
// {{$e.Desc}}
func Test{{$e.Name}}By{{$roleTest.RoleName}}(t *test.T) {
	expectedBody := ` + "`{{$result.Body}}`" + `
	{{$bodyLength := len $e.Body}}{{if gt $bodyLength 0}}payload := []byte("{{$e.Body}}")
	req, _ := http.NewRequest("{{$e.Method}}", "{{$e.Path}}", bytes.NewBuffer(payload)){{else}}req, _ := http.NewRequest("{{$e.Method}}", "{{$e.Path}}", nil){{end}}
	req.Header.Set("X-Forwarded-User", "{{$roleTest.UserID}};2")
	response := executeRequest(req)

	checkResponseCode(t, {{$result.Status}}, response.Code)
	checkResponseBody(t, expectedBody, response)
}
{{end}}{{end}}{{end}}
{{end}}
{{end}}
`

func main() {
	var projectTests = []testcase{
		{"GetMyProjectsOK", "GET", "get my projects by user_id success", "/projects/", ``},
		{"CreateProjectOK", "POST", "create a new project success", "/projects/", `{"name":"new_project_name"}`},
	}

	var roleTestData = []roleTest{
		roleTest{
			UserID:    "test_owner",
			RoleName:  "Owner",
			Testcases: projectTests,
			Results: []testResult{
				{http.StatusOK, `{"body":{"projects":[{"id":"p1","name":"Fizbone","status":"ACTIVE"}]},"code":"success"}`},
				{http.StatusOK, `{"body":{"project":{"name":"new_project_name","created_by":"test_owner","status":"ACTIVE"}},"code":"success"}`},
			},
		},
		roleTest{
			UserID:    "test_admin",
			RoleName:  "Admin",
			Testcases: projectTests,
			Results: []testResult{
				{http.StatusOK, `{"body":{"projects":[{"id":"p2","name":"Fizbone","status":"ACTIVE"}]},"code":"success"}`},
				{http.StatusOK, `{"body":{"project":{"name":"new_project_name","created_by":"test_admin","status":"ACTIVE"}},"code":"success"}`},
			},
		},
	}

	fileData := &file{Name: "Project", Tests: roleTestData}

	tmpl := template.New("page")
	var err error
	if tmpl, err = tmpl.Parse(page); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(funcs); err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, fileData)
}
