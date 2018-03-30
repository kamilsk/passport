// Code generated by go-bindata.
// sources:
// static/migrations/1_initial.sql
// static/scripts/passport.min.js
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticMigrations1_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd0\x31\x6f\xb3\x30\x10\x06\xe0\xdd\xbf\xe2\xe4\x09\xf4\x85\xe1\x93\xaa\x2e\x4c\x26\x5c\xd2\x53\xc1\x50\x63\x57\xc9\x54\xa1\xe2\x46\xa8\x0d\x20\x07\x9a\xbf\x5f\x91\xc4\x55\x23\x71\xf3\xf3\xbe\xf6\x5d\x14\xc1\xbf\x63\x7b\x70\xf5\x68\xc1\x0c\x8c\xad\x15\x0a\x8d\x80\x3b\x8d\xb2\xa2\x42\x02\x6d\x40\x16\x1a\x70\x47\x95\xae\x80\x4f\x53\xdb\x44\xfd\xe9\x34\xf0\xf8\x17\x6b\x91\x64\x08\xfc\xa3\xed\x0e\xd6\x0d\xae\xed\x46\x0e\x01\x03\xe0\x6d\xc3\xc1\x4f\x42\xdb\x0a\x15\x89\x0c\x4a\x45\xb9\x50\x7b\x78\xc6\xfd\x6a\x56\xc7\xda\x7d\x5a\x77\x95\xc6\x50\xea\x13\xf3\xb3\xd2\x64\xd9\x05\x7d\xd7\x5f\x93\xbd\xb5\xbd\x0a\xb5\x7e\x12\x2a\x78\x7c\x08\xef\xd1\x7b\x3f\x75\xe3\xad\x2a\xa1\x2d\x49\x7d\xdf\x04\x29\x6e\x84\xc9\x34\xfc\xbf\x72\x67\xeb\xd1\x36\x6f\xf5\xc8\x41\x53\x8e\x95\x16\x79\xb9\xc4\xbb\xfe\x1c\x84\x97\xc8\x34\x34\xcb\x91\x99\xcf\xe3\x23\xfe\x4f\x46\xd2\x8b\x41\x08\xfc\x96\x2b\xbf\x4a\xc8\xc2\x98\xb1\xbf\xe7\x4f\xfb\x73\xc7\x58\xaa\x8a\x72\xe9\xa2\x31\xfb\x09\x00\x00\xff\xff\xc9\xeb\xcc\xee\xab\x01\x00\x00")

func staticMigrations1_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations1_initialSql,
		"static/migrations/1_initial.sql",
	)
}

func staticMigrations1_initialSql() (*asset, error) {
	bytes, err := staticMigrations1_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/1_initial.sql", size: 427, mode: os.FileMode(420), modTime: time.Unix(1522409989, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticScriptsPassportMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xcf\x6e\x1b\x21\x10\xc6\xef\x7d\x0a\xcc\x61\x05\x32\xd9\xd8\x3d\x55\xb6\x68\xa5\xfe\x93\x5a\x45\x4d\xda\xa4\xea\x99\xb0\xb3\x36\xd1\x1a\x56\x30\x1b\xd7\xc2\xbc\x7b\x85\xed\x5d\xaf\xab\xe6\xd2\x1b\x33\x0c\x33\x1f\xbf\x0f\x26\x75\x67\x35\x1a\x67\x19\x0a\x2b\xbc\x00\xe1\x84\xe2\x91\x76\x01\x48\x40\x6f\x34\xd2\x25\x96\x77\x2a\x84\xd6\x79\x94\xb1\x36\x76\x05\xbe\xf5\xc6\xe2\xe2\xd9\x99\x8a\xcc\xc4\x06\x50\x55\x0a\x55\x1f\x6b\xd7\x59\x04\xdf\x87\x61\x67\x35\x54\xa7\x28\x2d\x9f\x95\x27\xe6\x3f\xfa\x24\xa1\xe5\x64\x2e\x82\x9c\xcc\x97\xbd\x66\x52\x31\xe4\x51\x15\x85\x63\xb4\x3d\x49\x5c\x10\x3a\x45\x9e\x86\x92\x8e\x79\xe1\x78\xac\x98\x9b\x52\xb2\x56\x81\x04\xb0\x48\x14\xb1\x0e\x4d\x6d\xb4\xca\x45\x94\x1f\x64\x29\xc9\xe8\x35\x95\x52\xda\x32\x74\x8f\x01\x3d\xbb\x9a\xf3\x77\x43\x30\x13\xb6\x6c\xc0\x5e\xcd\xf9\xc2\xf2\x29\x8d\xb1\xfc\x64\xab\xd6\x19\x8b\x29\xd1\x25\xb0\x88\xbb\x16\x16\xf4\xee\xf6\xfe\x81\x8a\xce\x37\x0b\x25\x0e\xd7\xf9\x7a\x7f\xfb\xad\xcc\x2c\xed\xca\xd4\x3b\x66\xb8\xd0\xce\x22\x58\x7c\x38\xd4\xab\xb6\x6d\x4e\x3a\xae\x9f\x82\xb3\x4b\xa2\xd7\xca\x07\x40\xd9\x61\x7d\xf5\x86\x8a\xd0\x69\x0d\x21\x2c\x06\xa7\x78\xd4\x72\x32\x13\xba\x01\xe5\xbf\x64\x44\xcf\xaa\x61\x9e\x8b\xb3\x4f\xe5\x08\xaf\x34\xe3\x68\x5c\xd4\x03\x97\x66\x58\x8e\xb7\x4f\xfc\xa5\xe9\x57\xa2\x62\x34\x80\xad\xc0\x1f\x41\x1e\x8c\x25\x8a\xb4\x6a\xd7\x38\x55\x51\x2e\x8e\x94\x4d\x20\x95\xb3\x40\x79\x12\xda\x6d\xda\x06\x10\xc6\xe2\xb3\x85\x29\x8d\x2c\x6a\x19\x8f\x1e\xb0\xf3\x76\xa2\x8b\x62\x18\xf7\x36\xc6\xf2\xc6\x6c\x0c\xa6\xf4\x97\xc3\xda\x1b\x34\x5a\x35\x8b\x7e\x34\xd9\xaa\x90\x1d\x25\x8f\x70\xf0\x97\x72\xa1\xf7\xfb\x7f\x75\x4a\xd9\xe8\x5a\x06\xc0\x81\xdc\x48\xda\x24\x14\x05\xb3\xb0\x25\x9e\x97\x2b\x40\x36\xfe\x1c\x3c\x9a\x9a\xe1\x44\x5e\xf2\x2c\x0a\x76\x11\x4b\x14\x67\x9a\xd2\x8a\x41\x83\x9c\x65\x7c\xda\x79\x0f\x1a\xdd\x91\xe0\x46\x55\x40\x14\x39\x25\x0f\x0f\x91\x9f\x4f\x4c\xa7\xa2\x65\x9c\x1f\xc9\x90\x4b\xb7\x6b\x2e\x0e\xdf\xe2\xa2\xe5\xc0\x7d\x79\xbe\xb9\x8c\xb1\x7c\x58\x7b\x08\x6b\xd7\x54\x19\x24\x0b\xf9\xed\x74\xac\x16\xe7\x93\x94\xf3\xc4\x93\x88\xb1\xfc\x70\x4c\xa5\xc4\x45\xf3\x12\x24\x53\xb3\x17\x75\x35\x67\x5d\x5b\x85\x7a\x0d\x63\x55\x61\xbf\xef\xa7\x37\xa2\xdf\xcf\xb3\xf3\xe4\x5f\x39\xcc\xcf\x82\x6d\x8d\xad\xdc\x56\xe4\x0f\xf6\x5e\x05\xf8\xf9\xe3\x26\x25\x2a\x8e\xd9\xf2\xf3\x19\xf5\xeb\x3e\xf7\xf4\xbd\x03\xbf\x2b\xd5\x93\xfa\xdd\xa7\xb4\xb3\xc1\x35\x50\x36\x6e\x95\x9b\x7f\x84\xc7\x6e\x95\x12\x5f\xbe\xfa\x13\x00\x00\xff\xff\xc1\x67\xe8\xaa\xf0\x04\x00\x00")

func staticScriptsPassportMinJsBytes() ([]byte, error) {
	return bindataRead(
		_staticScriptsPassportMinJs,
		"static/scripts/passport.min.js",
	)
}

func staticScriptsPassportMinJs() (*asset, error) {
	bytes, err := staticScriptsPassportMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/scripts/passport.min.js", size: 1264, mode: os.FileMode(420), modTime: time.Unix(1522422579, 0)}
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
	"static/migrations/1_initial.sql": staticMigrations1_initialSql,
	"static/scripts/passport.min.js": staticScriptsPassportMinJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial.sql": &bintree{staticMigrations1_initialSql, map[string]*bintree{}},
		}},
		"scripts": &bintree{nil, map[string]*bintree{
			"passport.min.js": &bintree{staticScriptsPassportMinJs, map[string]*bintree{}},
		}},
	}},
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

