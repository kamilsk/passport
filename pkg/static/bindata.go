// Code generated by go-bindata.
// sources:
// pkg/static/migrations/1_prepare.sql
// pkg/static/migrations/2_domain.sql
// pkg/static/scripts/passport.min.js
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

var _staticMigrations1_prepareSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x4f\x83\x30\x18\x86\xef\xfd\x15\x6f\x96\x1d\x74\x66\xfb\x01\x36\x1e\xd8\xfa\x81\x4d\xb0\x98\xb6\xc4\xdd\x16\x22\x0d\x21\x19\x0c\xa1\xc4\xbf\x6f\x80\x98\x4d\xdc\x0e\x7a\xec\xd7\xa7\x6f\x9f\xbe\x5d\xaf\xf1\x50\x95\x45\x9b\x79\x87\xb4\x61\x6c\xa7\x29\xb0\x04\xda\x5b\x52\x46\x26\x0a\x32\x84\x4a\x2c\x68\x2f\x8d\x35\x58\xf4\x7d\x99\xaf\x4f\x5d\xd7\x2c\x38\x63\x97\x87\x8d\xcf\xbc\xab\x5c\xed\xb7\xae\x28\xeb\xef\x9c\x30\x55\x3b\x3b\xc4\xf4\x4d\x9e\x79\x77\xf0\x65\xe5\x3a\x9f\x55\xcd\xdd\x3d\x03\x34\xd9\x54\x2b\x03\xab\x65\x14\x91\x46\x60\xb0\x9c\x83\x4b\xb6\xa5\x48\x2a\x86\xd1\x84\xde\x36\x2b\x48\x03\x21\x8d\x95\x6a\x67\x11\xea\xe4\x05\x49\x2c\x36\x2b\x06\xd8\x67\x1a\x40\x8c\xdc\x14\x94\x1f\x32\x8f\xc7\x27\xbc\xf7\x6d\xeb\x6a\x7f\xce\xe5\x23\x38\x19\x0c\xfc\xb0\xa6\xd8\xd0\xe5\x38\x89\xc5\x38\x56\x02\x32\xe4\x8c\x94\xe0\xec\x8a\x60\x1c\xa8\x28\x0d\x22\x42\x73\x6c\x8a\xee\xe3\xc8\xaf\x17\x43\x75\xfe\xa7\xca\xca\xa2\x3e\xb5\xee\x30\xdd\x77\xb3\xaf\x1f\xd4\xb9\xac\xcb\x17\x4c\xde\x33\x10\xf8\xaf\xb6\x38\x7d\xd6\x8c\x09\x9d\xbc\xde\x74\xe5\xf3\xfd\xdf\xdf\xcf\xd9\x57\x00\x00\x00\xff\xff\x1f\x2f\xa5\xa8\x7b\x02\x00\x00")

func staticMigrations1_prepareSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations1_prepareSql,
		"static/migrations/1_prepare.sql",
	)
}

func staticMigrations1_prepareSql() (*asset, error) {
	bytes, err := staticMigrations1_prepareSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/1_prepare.sql", size: 635, mode: os.FileMode(420), modTime: time.Unix(1546590942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMigrations2_domainSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x31\x6b\xc3\x30\x10\x85\xf7\xfb\x15\x87\x26\x9b\x26\x43\xa1\x74\xc9\x24\xc7\x6a\x7a\x54\x76\x5c\x45\x2a\x64\x2a\xa2\x56\x83\x69\x23\x1b\x21\x37\x7f\xbf\x38\x89\x20\x01\xdf\xfc\x7d\x8f\x7b\x6f\xb9\xc4\x87\x63\x77\x08\x36\x3a\x34\x03\xc0\x5a\x09\xae\x05\x6a\x5e\x48\x81\xec\xbb\xf3\x07\x17\x86\xd0\xf9\xc8\x20\x03\x44\xd6\xb5\x0c\xd3\x15\xb4\xd9\x09\x45\x5c\x62\xa3\xa8\xe2\x6a\x8f\x6f\x62\xbf\x98\xa8\xa3\x0d\x3f\x2e\x5c\x48\x63\xa8\x4c\x46\xbd\xd5\x58\x1b\x29\xcf\xd0\x9f\xfd\x1d\xdd\x35\xed\x83\xab\xf5\x2b\x57\xd9\xf3\x53\x7e\x0f\x7d\xf5\xa3\x8f\xd7\xa8\x82\x36\x54\xeb\xfb\x24\x2c\xc5\x0b\x37\x52\xe3\xe3\x05\x0f\xce\x46\xd7\x7e\xda\xc8\x50\x53\x25\x76\x9a\x57\xcd\x1c\xee\xfb\x53\x96\x9f\x95\x71\x68\xe7\x95\x09\x9f\x2e\x29\xe9\x27\x53\xd3\xbb\x11\x98\xa5\x96\x8b\x54\x25\x87\x7c\x05\x70\xbb\x68\xd9\x9f\x3c\x40\xa9\xb6\xcd\xdc\xa2\x2b\xf8\x0f\x00\x00\xff\xff\x31\xa9\x1f\xc6\x7e\x01\x00\x00")

func staticMigrations2_domainSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations2_domainSql,
		"static/migrations/2_domain.sql",
	)
}

func staticMigrations2_domainSql() (*asset, error) {
	bytes, err := staticMigrations2_domainSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/2_domain.sql", size: 382, mode: os.FileMode(420), modTime: time.Unix(1546590937, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticScriptsPassportMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x4d\x8f\x23\x27\x10\xbd\xe7\x57\x60\x0e\x2d\x90\x2b\xc4\x93\x53\xd4\x16\x7b\xd9\x64\xa5\x44\x51\x76\xa3\x1d\x29\x67\x06\xaa\x6d\x46\x0c\xb4\xa0\x7a\x3c\x16\xe6\xbf\x47\xdd\xfe\x18\xcf\x21\x39\x75\xd5\x13\x45\xbf\xf7\xaa\x8a\xd5\x30\x45\x4b\x3e\x45\x41\x10\x01\xc1\x43\x96\x95\x4f\x05\x59\xa1\xec\x2d\xf1\xed\xab\xc9\xcc\x68\x52\xdf\x4c\x29\x63\xca\x74\x17\x9e\x4e\xd5\xe1\xd3\xb4\xeb\xb3\x5a\xbe\x0d\x92\xae\x83\x8f\x3b\xcc\x63\xf6\x91\xfa\xd7\xe4\x1d\xdb\xc0\x0b\x92\x71\x86\xcc\x25\x6f\x60\xf5\x06\x9c\x5e\x3d\x40\xd1\xab\x87\xed\x95\x02\x1b\x05\xc9\x6a\xce\x77\x75\x9d\x17\x59\x8d\x19\x07\xff\xb6\x26\xd9\x6e\x87\x86\x99\xa9\xac\x7e\x10\x2b\xd7\x75\xf6\x53\x56\xc1\xbf\x78\xfa\x70\x9e\xdb\xec\xc9\x5b\x13\x7a\x36\x9a\x63\x48\xc6\xb1\x83\x29\x2c\x26\x62\x4f\xc8\x0a\x46\xe2\x12\xdc\xe9\x74\xab\x96\x19\x69\xca\x91\xd9\x80\x26\xff\x1e\x09\xf3\xab\x09\x82\x24\x2c\x8c\x47\x11\xd7\x9c\xf9\xc2\x5c\x8a\xc8\xe5\xb6\x9c\x4e\xa2\xe8\xd5\x06\xec\x7a\x0d\x28\x2a\x1d\x47\xec\xf9\xb7\xaf\xdf\x1f\x39\x4c\x39\xf4\x59\x61\x74\x63\xf2\x91\x60\x91\xfd\xc7\xf7\xaf\x7f\xa9\xd9\xcf\xb8\xf3\xc3\x51\x24\x09\x36\x45\xc2\x48\x8f\x4b\xa1\x19\xc7\xe0\xad\x99\xd5\xfd\xf4\x5c\x52\xdc\x32\xbb\x37\xb9\x20\xe9\x89\x86\x1f\x7f\xe1\xf0\xb6\xcf\x5f\x3c\x06\x57\xfa\x7a\xf0\xb4\xff\x9c\xd1\x61\x24\x6f\x42\xe9\x57\x9b\x06\x65\xb2\x16\x4b\xe9\x6f\xad\x94\xd5\xcd\xec\x8c\xba\xeb\x85\x4e\xf7\x19\x18\x75\x6d\x8a\x4e\xb7\x10\x46\xc1\x0b\x46\x87\x99\xed\x4d\x61\xe5\x18\x2d\x3a\x66\xae\x1e\x72\xd9\xc0\xa6\x97\x31\x20\xe1\xfd\xcf\xe6\x26\xc2\xd9\xa3\xa5\x0c\x23\x31\x33\x9b\xed\x87\x8b\x2c\x46\x89\xf1\xf5\xbb\x2d\xb2\x35\x29\xdb\x3c\x59\x01\x26\x5d\x90\x6e\x96\x8b\xa0\x1f\x40\x44\x3c\xb0\x28\xd5\x0e\x49\xdc\xcf\xa7\xac\x1f\x44\x68\x82\x77\xee\x3a\x36\x09\x77\xa4\xfe\xfb\x0e\x5a\xe9\x8f\x66\x74\x9d\xf8\xbf\x7b\x21\xe8\xcd\xec\x8c\x4d\x39\xa3\xa5\x74\x36\xe7\xc5\x38\x64\x86\x5d\x40\x9f\x22\x97\x12\xd6\xeb\xf0\x49\x67\x45\xfb\x8c\x65\x9f\x82\xeb\xba\x41\x4c\xf0\x5e\xc9\x65\x93\x4d\x42\x56\x17\x44\xc2\xee\x83\xfa\x3b\xfe\x83\xd8\x01\x3f\x18\xb2\x7b\x9c\xcb\x20\xab\x25\x91\x4d\x1c\x7c\x74\xe9\x00\xe7\x8f\xfa\xf2\xce\xfb\xe7\x2b\xf6\xfc\xf7\x84\xf9\xa8\xcc\xb3\x79\xbb\x42\x36\xc5\x92\x02\xaa\x90\x76\x50\xaf\x6d\xe8\x79\xad\xea\xb7\x4b\xd2\x1a\x87\xf3\xfa\xf4\x7c\xbc\x2c\x78\xcf\x38\x2c\xfb\xd1\xd7\xaa\xfe\x9c\x83\xd6\xe0\xa6\x6e\x06\x1f\xaf\x49\x9b\x67\x63\x11\x35\xc3\x9f\xcf\x61\x6b\xb0\xb0\x9e\xa1\x7f\xe6\xa0\x35\x38\xbf\x18\xb5\xaa\x5f\x97\x27\xa3\x35\xb9\xfd\xe1\xdf\x00\x00\x00\xff\xff\x7f\xa7\xf5\x89\x86\x04\x00\x00")

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

	info := bindataFileInfo{name: "static/scripts/passport.min.js", size: 1158, mode: os.FileMode(420), modTime: time.Unix(1522527314, 0)}
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
	"static/migrations/1_prepare.sql": staticMigrations1_prepareSql,
	"static/migrations/2_domain.sql": staticMigrations2_domainSql,
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
			"1_prepare.sql": &bintree{staticMigrations1_prepareSql, map[string]*bintree{}},
			"2_domain.sql": &bintree{staticMigrations2_domainSql, map[string]*bintree{}},
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
