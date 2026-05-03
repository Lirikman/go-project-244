package formatters

import "testing"

func TestFormatterStylishFlatFiles(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"first_name": {"type": "changed", "value1": "Ivan", "value2": "Maksim"},
		"last_name":  {"type": "unchanged", "value1": "Sidorov"},
		"age":        {"type": "deleted", "value1": 20},
		"job":        {"type": "added", "value2": "go developer"},
	}

	got := FormatterStylish(diffTreeTest)
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

	got := FormatterStylish(diffTreeTest)
	want := "{\n..- email: tester@ya.ru\n..+ email: developer@mail.ru\n..+ games: {\n........offline: {\n............rpg: elden ring\n........}\n....}\n..  pc: {\n......  cpu: intel core i5 11400\n......- os: windows 10 PRO x64\n......- ram: 16Gb DDR5\n......+ ram: 32Gb DDR5\n......  storage: {\n..........- hdd: Toshiba 2Tb\n..........+ hdd: Seagate 4 Tb\n..........+ m2_nvme: Samsung 870EVO\n..........  ssd: Hitachi 1TB\n......  }\n..  }\n..  user: {\n......+ age: 20\n......  first_name: Ivan\n......- last_name: Ivanov\n......+ last_name: Petrov\n..  }\n}"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestFormatterPlainFlatFiles(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"first_name": {"type": "changed", "value1": "Ivan", "value2": "Maksim"},
		"last_name":  {"type": "unchanged", "value1": "Sidorov"},
		"age":        {"type": "deleted", "value1": 20},
		"job":        {"type": "added", "value2": "go developer"},
	}

	got := FormmaterPlain(diffTreeTest)
	want := "Property 'age' was removed\nProperty 'first_name' was updated. From 'Ivan' to 'Maksim'\nProperty 'job' was added with value: go developer\n"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestFormatterPlainNestedFiles(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"user": {"type": "nested", "children": map[string]map[string]any{"first_name": {"type": "unchanged", "value1": "Maksim"},
			"last_name": {"type": "changed", "value1": "Ivanov", "value2": "Petrov"}, "age": {"type": "added", "value2": 30},
			"job": {"type": "deleted", "value1": "designer"}}},
		"games": {"type": "added", "value2": "the wither 3 wild hunt"},
	}

	got := FormmaterPlain(diffTreeTest)
	want := "Property 'games' was added with value: the wither 3 wild hunt\nProperty 'user.age' was added with value: '30'\nProperty 'user.job' was removed\nProperty 'user.last_name' was updated. From 'Ivanov' to 'Petrov'\n"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestFormatterPlainNestedChildNode(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"pc": {"type": "nested", "children": map[string]map[string]any{"cpu": {"type": "changed", "value1": "intel core i5 12400", "value2": "intel core i7 12700K"},
			"ram": {"type": "uncahnged", "value1": "64Gb DDR5"}, "VGA": {"type": "added", "value2": "NVIDIA RTX5080"},
			"storage": {"type": "nested", "children": map[string]map[string]any{"hdd": {"type": "added", "value2": "Hitachi 4Tb"},
				"ssd": {"type": "deleted", "value1": "Crucial 256Gb"}, "m2_nvme": {"type": "added", "value2": "samsung evo 870 1Tb"}}}}}}

	got := FormmaterPlain(diffTreeTest)
	want := "Property 'pc.VGA' was added with value: 'NVIDIA RTX5080'\nProperty 'pc.cpu' was updated. From 'intel core i5 12400' to 'intel core i7 12700K'\nProperty 'pc.storage.hdd' was added with value: 'Hitachi 4Tb'\nProperty 'pc.storage.m2_nvme' was added with value: 'samsung evo 870 1Tb'\nProperty 'pc.storage.ssd' was removed\n"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestFormatterPlainComplexValue(t *testing.T) {
	diffTreeTest := map[string]map[string]any{
		"pc": {"type": "nested", "children": map[string]map[string]any{"cpu": {"type": "changed", "value1": "intel celerone", "value2": "intel pentium"},
			"hdd": {"type": "added", "value2": map[string]any{"hdd": "samsung 1TB"}}}},
		"monitor": {"type": "added", "value2": map[string]any{"samsung": "165MHz"}},
		"games":   {"type": "changed", "value1": map[string]any{"rpg": "gothic 2"}, "value2": "gothic 1 remake"},
	}

	got := FormmaterPlain(diffTreeTest)
	want := "Property 'games' was updated. From [complex value] to 'gothic 1 remake'\nProperty 'monitor' was added with value: [complex value]\nProperty 'pc.cpu' was updated. From 'intel celerone' to 'intel pentium'\nProperty 'pc.hdd' was added with value: [complex value]\n"

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
