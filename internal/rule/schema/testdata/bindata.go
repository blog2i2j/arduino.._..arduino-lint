// Package testdata Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// internal/rule/schema/testdata/input/invalid-schema.json
// internal/rule/schema/testdata/input/referenced-schema-1.json
// internal/rule/schema/testdata/input/referenced-schema-2.json
// internal/rule/schema/testdata/input/schema-without-id.json
// internal/rule/schema/testdata/input/valid-schema-with-references.json
// internal/rule/schema/testdata/input/valid-schema.json
package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _invalidSchemaJson = []byte(`{
  "foo": "bar"
  "baz": "bat"
}
`)

func invalidSchemaJsonBytes() ([]byte, error) {
	return _invalidSchemaJson, nil
}

func invalidSchemaJson() (*asset, error) {
	bytes, err := invalidSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "invalid-schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _referencedSchema1Json = []byte(`{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://raw.githubusercontent.com/arduino/arduino-lint/main/internal/rule/schema/testdata/input/referenced-schema-1.json",
  "title": "Schema for use in unit tests",
  "definitions": {
    "patternObject": {
      "pattern": "^[a-z]+$"
    },
    "requiredObject": {
      "required": ["property1"]
    }
  }
}
`)

func referencedSchema1JsonBytes() ([]byte, error) {
	return _referencedSchema1Json, nil
}

func referencedSchema1Json() (*asset, error) {
	bytes, err := referencedSchema1JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "referenced-schema-1.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _referencedSchema2Json = []byte(`{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://raw.githubusercontent.com/arduino/arduino-lint/main/internal/rule/schema/testdata/input/referenced-schema-2.json",
  "title": "Schema for use in unit tests",
  "definitions": {
    "dependenciesObject": {
      "dependencies": {
        "dependentProperty": ["dependencyProperty"]
      }
    },
    "minLengthObject": {
      "minLength": 2
    },
    "maxLengthObject": {
      "maxLength": 4
    },
    "enumObject": {
      "enum": ["baz"]
    },
    "notPatternObject": {
      "not": {
        "allOf": [
          {
            "pattern": "[A-Z]"
          }
        ]
      }
    },
    "misspelledOptionalProperties": {
      "propertyNames": {
        "not": {
          "pattern": "porpert([y]|[ies])"
        }
      }
    }
  }
}
`)

func referencedSchema2JsonBytes() ([]byte, error) {
	return _referencedSchema2Json, nil
}

func referencedSchema2Json() (*asset, error) {
	bytes, err := referencedSchema2JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "referenced-schema-2.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaWithoutIdJson = []byte(`{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Schema without $id keyword for use in unit tests"
}
`)

func schemaWithoutIdJsonBytes() ([]byte, error) {
	return _schemaWithoutIdJson, nil
}

func schemaWithoutIdJson() (*asset, error) {
	bytes, err := schemaWithoutIdJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema-without-id.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _validSchemaWithReferencesJson = []byte(`{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://raw.githubusercontent.com/arduino/arduino-lint/main/internal/rule/schema/testdata/input/valid-schema-with-references.json",
  "title": "Schema for use in unit tests",
  "type": "object",
  "properties": {
    "property1": {
      "allOf": [
        {
          "$ref": "referenced-schema-2.json#/definitions/minLengthObject"
        },
        {
          "$ref": "referenced-schema-2.json#/definitions/maxLengthObject"
        }
      ]
    },
    "property2": {
      "allOf": [
        {
          "$ref": "referenced-schema-1.json#/definitions/patternObject"
        }
      ]
    },
    "property3": {
      "allOf": [
        {
          "$ref": "referenced-schema-2.json#/definitions/enumObject"
        },
        {
          "$ref": "referenced-schema-2.json#/definitions/notPatternObject"
        }
      ]
    }
  },
  "allOf": [
    {
      "$ref": "referenced-schema-2.json#/definitions/dependenciesObject"
    },
    {
      "$ref": "referenced-schema-1.json#/definitions/requiredObject"
    },
    {
      "$ref": "referenced-schema-2.json#/definitions/misspelledOptionalProperties"
    }
  ]
}
`)

func validSchemaWithReferencesJsonBytes() ([]byte, error) {
	return _validSchemaWithReferencesJson, nil
}

func validSchemaWithReferencesJson() (*asset, error) {
	bytes, err := validSchemaWithReferencesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "valid-schema-with-references.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _validSchemaJson = []byte(`{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://raw.githubusercontent.com/arduino/arduino-lint/main/internal/rule/schema/testdata/input/valid-schema.json",
  "title": "Schema for use in unit tests",
  "type": "object",
  "properties": {
    "property1": {
      "minLength": 2
    }
  }
}
`)

func validSchemaJsonBytes() ([]byte, error) {
	return _validSchemaJson, nil
}

func validSchemaJson() (*asset, error) {
	bytes, err := validSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "valid-schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"invalid-schema.json":               invalidSchemaJson,
	"referenced-schema-1.json":          referencedSchema1Json,
	"referenced-schema-2.json":          referencedSchema2Json,
	"schema-without-id.json":            schemaWithoutIdJson,
	"valid-schema-with-references.json": validSchemaWithReferencesJson,
	"valid-schema.json":                 validSchemaJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"invalid-schema.json":               &bintree{invalidSchemaJson, map[string]*bintree{}},
	"referenced-schema-1.json":          &bintree{referencedSchema1Json, map[string]*bintree{}},
	"referenced-schema-2.json":          &bintree{referencedSchema2Json, map[string]*bintree{}},
	"schema-without-id.json":            &bintree{schemaWithoutIdJson, map[string]*bintree{}},
	"valid-schema-with-references.json": &bintree{validSchemaWithReferencesJson, map[string]*bintree{}},
	"valid-schema.json":                 &bintree{validSchemaJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
