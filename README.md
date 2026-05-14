## Difference generator


### Hexlet tests and linter status:
[![Actions Status](https://github.com/Lirikman/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Lirikman/go-project-244/actions)

### SonarQube tests status:
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Lirikman_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Lirikman_go-project-244)

### Description

Compares two configuration files and shows a difference

USAGE:
   gendiff [global options] [arguments...]

GLOBAL OPTIONS:
   --format string, -f string  supported formats: stylish, plain, json (default: "stylish")
   --help, -h                  show help

### Requirements

* Go 1.25
* Make
* urfave/cli v3

### Setup

```bash
git clone https://github.com/Lirikman/go-project-244.git
cd go-project-244
```

### Run build

```bash
make build
```

### Run golangci-lint 

```bash
make lint
```

### Run tests

```bash
make test
```

### Example

```bash
./bin/gendiff/ -f ./testdata/file1.yml ./testdata/file2.yml
{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }

bin/gendiff/ -f stylish ./file1.json ./file2.json
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}

bin/gendiff/ -f plain ./file1.json ./file2.json
Property 'follow' was removed
Property 'proxy' was removed
Property 'timeout' was updated. From '50' to '20'
Property 'verbose' was added with value: true

bin/gendiff/ -f json ./file1.json ./file2.json
{
  "follow": {
    "type": "deleted",
    "value1": false
  },
  "host": {
    "type": "unchanged",
    "value1": "hexlet.io"
  },
  "proxy": {
    "type": "deleted",
    "value1": "123.234.53.22"
  },
  "timeout": {
    "type": "changed",
    "value1": 50,
    "value2": 20
  },
  "verbose": {
    "type": "added",
    "value2": true
  }
}
```

### Asciinema сomparison of flat JSON files

https://asciinema.org/a/8FBjj6DRgtVRE5W2

### Asciinema сomparison of flat YAML files

https://asciinema.org/a/GthGBNFWEzH3mCsZ

### Asciinema comparison nested JSON, YAML files format 'STYLISH'

https://asciinema.org/a/Vw5hQ4WYXjtX6bhP

### Asciinema comparison nested JSON, YAML files format 'PLAIN'
https://asciinema.org/a/MzGoPOq17EdcJBSb

### Asciinema comparison nested JSON, YAML files format 'JSON'
https://asciinema.org/a/8Tjlr0z5Y99TkUEe