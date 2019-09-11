package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	arrayTestFilePath  = "sample/github-repo-contributors.json"
	singleTestFilePath = "sample/github-repo-contributor.json"
)

func TestFindHandler(t *testing.T) {

	testFile, err := os.Open(arrayTestFilePath)
	if err != nil {
		t.Fatalf("test file not found: %s", arrayTestFilePath)
	}
	defer testFile.Close()

	content, _ := ioutil.ReadAll(testFile)

	r, err := http.NewRequest("POST", "/", bytes.NewBuffer(content))
	if err != nil {
		t.Fatalf("error on post: %s", err.Error())
	}

	r.Header.Set("Select-query", "login")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(findHandler)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Result: %s", string(body))

}

func TestSelectHandler(t *testing.T) {

	testFile, err := os.Open(singleTestFilePath)
	if err != nil {
		t.Fatalf("test file not found: %s", singleTestFilePath)
	}
	defer testFile.Close()

	content, _ := ioutil.ReadAll(testFile)

	r, err := http.NewRequest("POST", "/", bytes.NewBuffer(content))
	if err != nil {
		t.Fatalf("error on post: %s", err.Error())
	}

	r.Header.Set("Select-query", "login")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(selectHandler)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Result: %s", string(body))

}
