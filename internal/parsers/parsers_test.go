package parsers

import (
	"bytes"
	"errors"
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
		{"file_json", "./testdata/file3.json", []byte(`{\n  "login": "admin",\n  "password": "123Ivan456",\n}`), nil},
		{"empty_yaml_file", "./testdata/empty.yaml", []byte{}, nil},
		{"no_file", "./testdata/notFile.yaml", []byte{}, errors.New("file read error")},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			got, err := ReadFiles(ts.filePath)
			if ts.err != nil {
				if err == nil || err.Error() != ts.err.Error() {
					t.Errorf("expected to get error %q, but got %v", ts.err, err)
				}
			}
			if !bytes.Equal(got, ts.want) {
				t.Errorf("expected to get %v, but got %v", ts.want, got)
			}
		})
	}
}
