package main

import (
	"code"
	"reflect"
	"testing"
)

func TestGenDiffFlatFiles(t *testing.T) {
	testMap1 := map[string]any{
		"first_name": "Ivan",
		"last_name":  "Petrov",
		"age":        20,
		"job":        "php_developer",
	}
	testMap2 := map[string]any{
		"first_name": "Ivan",
		"last_name":  "Petrov",
		"age":        20,
		"job":        "php_developer",
	}
	got := code.GenDiff(testMap1, testMap2)
	want := map[string]map[string]any{
		"first_name": {"type": "unchanged", "value1": "Ivan"},
		"last_name":  {"type": "unchanged", "value1": "Petrov"},
		"age":        {"type": "unchanged", "value1": 20},
		"job":        {"type": "unchanged", "value1": "php_developer"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGenDiffNestedFiles(t *testing.T) {
	testMap1 := map[string]any{
		"user":      map[string]any{"name": "Petr", "last_name": "Ivanov"},
		"nick_name": "Designer2000",
		"age":       20,
		"hobby":     "music",
		"job":       "designer",
	}
	testMap2 := map[string]any{
		"user":      map[string]any{"name": "Boris", "last_name": "Ivanov"},
		"nick_name": "Designer2002",
		"age":       22,
		"job":       "designer",
		"city":      "Kaliningrad",
	}
	got := code.GenDiff(testMap1, testMap2)
	want := map[string]map[string]any{
		"user": {"type": "nested", "children": map[string]map[string]any{"name": {"type": "changed", "value1": "Petr", "value2": "Boris"},
			"last_name": {"type": "unchanged", "value1": "Ivanov"}}},
		"nick_name": {"type": "changed", "value1": "Designer2000", "value2": "Designer2002"},
		"age":       {"type": "changed", "value1": 20, "value2": 22},
		"job":       {"type": "unchanged", "value1": "designer"},
		"hobby":     {"type": "deleted", "value1": "music"},
		"city":      {"type": "added", "value2": "Kaliningrad"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGenDiffEmptyFirstFile(t *testing.T) {
	testMap1 := map[string]any{}
	testMap2 := map[string]any{
		"first_name": "Ivan",
		"last_name":  "Petrov",
		"age":        20,
		"job":        "php_developer",
	}

	got := code.GenDiff(testMap1, testMap2)
	want := map[string]map[string]any{
		"first_name": {"type": "added", "value2": "Ivan"},
		"last_name":  {"type": "added", "value2": "Petrov"},
		"age":        {"type": "added", "value2": 20},
		"job":        {"type": "added", "value2": "php_developer"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGenDiffEmptySecondFile(t *testing.T) {
	testMap1 := map[string]any{
		"first_name": "Ivan",
		"last_name":  "Petrov",
		"age":        20,
		"job":        "php_developer",
	}
	testMap2 := map[string]any{}

	got := code.GenDiff(testMap1, testMap2)
	want := map[string]map[string]any{
		"first_name": {"type": "deleted", "value1": "Ivan"},
		"last_name":  {"type": "deleted", "value1": "Petrov"},
		"age":        {"type": "deleted", "value1": 20},
		"job":        {"type": "deleted", "value1": "php_developer"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
