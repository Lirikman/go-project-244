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

func TestFormatterStylishFlatFiles(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"first_name": {"type": "changed", "value1": "Ivan", "value2": "Maksim"},
		"last_name":  {"type": "unchanged", "value1": "Sidorov"},
		"age":        {"type": "deleted", "value1": 20},
		"job":        {"type": "added", "value2": "go developer"},
	}

	got := code.FormatterStylish(diffTreeTest)
	want := "{\n..- age: 20\n..- first_name: Ivan\n..+ first_name: Maksim\n..+ job: go developer\n..  last_name: Sidorov\n}"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestFormatterStylishNestedFiles(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"user": {"type": "nested", "children": map[string]map[string]any{"first_name": {"type": "unchanged", "value1": "Ivan"},
			"last_name": {"type": "changed", "value1": "Ivanov", "value2": "Petrov"}, "age": {"type": "added", "value2": 20}}},
		"email": {"type": "changed", "value1": "tester@ya.ru", "value2": "developer@mail.ru"},
		"pc": {"type": "nested", "children": map[string]map[string]any{"cpu": {"type": "unchanged", "value1": "intel core i5 11400"},
			"ram": {"type": "changed", "value1": "16Gb DDR5", "value2": "32Gb DDR5"}, "os": {"type": "deleted", "value1": "windows 10 PRO x64"},
			"storage": {"type": "nested", "children": map[string]map[string]any{"ssd": {"type": "unchanged", "value1": "Hitachi 1TB"},
				"m2_nvme": {"type": "added", "value2": "Samsung 870EVO"}, "hdd": {"type": "changed", "value1": "Toshiba 2Tb", "value2": "Seagate 4 Tb"}}}}},
		"games": {"type": "added", "value2": map[string]any{"offline": map[string]any{"rpg": "elden ring"}}}}

	got := code.FormatterStylish(diffTreeTest)
	want := "{\n..- email: tester@ya.ru\n..+ email: developer@mail.ru\n..+ games: {\n........offline: {\n............rpg: elden ring\n........}\n....}\n..  pc: {\n......  cpu: intel core i5 11400\n......- os: windows 10 PRO x64\n......- ram: 16Gb DDR5\n......+ ram: 32Gb DDR5\n......  storage: {\n..........- hdd: Toshiba 2Tb\n..........+ hdd: Seagate 4 Tb\n..........+ m2_nvme: Samsung 870EVO\n..........  ssd: Hitachi 1TB\n......  }\n..  }\n..  user: {\n......+ age: 20\n......  first_name: Ivan\n......- last_name: Ivanov\n......+ last_name: Petrov\n..  }\n}"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
