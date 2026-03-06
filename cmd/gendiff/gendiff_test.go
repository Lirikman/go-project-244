package main

import (
	"code"
	"testing"
)

func TestGenDiffNormal(t *testing.T) {
	testMap := map[string]map[string]any{
		"file1": {
			"first_name": "Ivan",
			"last_name":  "Petrov",
			"age":        20,
			"job":        "php_developer",
		},
		"file2": {
			"first_name": "Ivan",
			"last_name":  "Petrov",
			"age":        20,
			"job":        "php_developer",
		},
	}
	got := code.GenDiff(testMap)
	want := "{\n    age: 20\n    first_name: Ivan\n    job: php_developer\n    last_name: Petrov\n}\n"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGenDiffEmptyJson1(t *testing.T) {
	testMap := map[string]map[string]any{
		"file1": {},
		"file2": {
			"first_name": "Ivan",
			"last_name":  "Petrov",
			"age":        20,
			"job":        "php_developer",
		},
	}
	got := code.GenDiff(testMap)
	want := "{\n  + age: 20\n  + first_name: Ivan\n  + job: php_developer\n  + last_name: Petrov\n}\n"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGenDiffEmptyJson2(t *testing.T) {
	testMap := map[string]map[string]any{
		"file1": {
			"first_name": "Ivan",
			"last_name":  "Petrov",
			"age":        20,
			"job":        "php_developer",
		},
		"file2": {},
	}
	got := code.GenDiff(testMap)
	want := "{\n  - age: 20\n  - first_name: Ivan\n  - job: php_developer\n  - last_name: Petrov\n}\n"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
