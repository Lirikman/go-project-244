package parsers

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestReadFiles(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     []byte
		err      error
	}{
		{"file_yaml", "./testdata/file1.yaml", []byte("host: hexlet.io\ntimeout: 50\n"), nil},
		{"file_json", "./testdata/file3.json", []byte("{\n  \"login\": \"admin\",\n  \"password\": \"123Ivan456\",\n}\n"), nil},
		{"empty_yaml_file", "./testdata/empty.yaml", []byte{}, nil},
		{"read_error_file", "./testdata/notFile.yaml", []byte{}, errors.New("file read error")},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			got, err := ReadFiles(ts.filePath)
			if ts.err != nil {
				if err == nil || err.Error() != ts.err.Error() {
					t.Errorf("expected to get error: %q, but got: %q", ts.err, err)
				}
			}
			if !bytes.Equal(got, ts.want) {
				t.Errorf("expected to get: %v,\nbut got: %v", ts.want, got)
			}
		})
	}
}

func TestReadData(t *testing.T) {
	tests := []struct {
		name string
		path string
		want map[string]map[string]any
		err  error
	}{
		{"empty_json_file", "./testdata/empty.json", nil, errors.New("unable to deserialize JSON file")},
		{"empty_yaml_file", "./testdata/empty.yaml", map[string]map[string]any{"empty.yaml": nil}, nil},
		{"file_yaml", "./testdata/file2.yaml", map[string]map[string]any{"file2.yaml": {"timeout": 20, "verbose": true, "host": "hexlet.io"}, "empty.yaml": nil}, nil},
		{"file_json", "./testdata/file1.json", map[string]map[string]any{"file1.json": {"user_id": "100", "username": "ivan", "is_admin": false}, "file2.yaml": {"timeout": 20, "verbose": true, "host": "hexlet.io"}, "empty.yaml": nil}, nil},
		{"read_error_json_file", "./testdata/noFile.json", nil, errors.New("file read error")},
		{"read_error_yaml_file", "./testdata/noFile.yaml", nil, errors.New("file read error")},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			got, err := ReadData(ts.path)
			// fmt.Printf("ОШИБКА: %v", err)
			if ts.err != nil {
				if err == nil || err.Error() != ts.err.Error() {
					t.Errorf("expected to get error: %q, but got: %q", ts.err, err)
				}
			}
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("expected to get: %v,\nbut got: %v", ts.want, got)
			}
		})
	}
}
