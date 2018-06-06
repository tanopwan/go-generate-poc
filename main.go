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

	tmpl, err := template.ParseFiles("./template/file.tmpl", "./template/funcs.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(os.Stdout, fileData)
}
