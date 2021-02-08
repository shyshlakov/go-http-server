// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// migration/schema/1_init_db_structure.down.sql
// migration/schema/1_init_db_structure.up.sql
// migration/schema/2_add_tag_in_articles.down.sql
// migration/schema/2_add_tag_in_articles.up.sql
package main

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

var __1_init_db_structureDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1_init_db_structureDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_init_db_structureDownSql,
		"1_init_db_structure.down.sql",
	)
}

func _1_init_db_structureDownSql() (*asset, error) {
	bytes, err := _1_init_db_structureDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_init_db_structure.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1611841354, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_init_db_structureUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x52\xcb\x6e\xc2\x30\x10\x3c\x37\x5f\xb1\xe2\x44\xa4\xf6\x0b\x38\x45\xc9\x22\x45\x0d\x0e\x4a\x1c\x15\x7a\x89\x4c\x6c\x52\xb7\x8e\x8d\xfc\xa8\xca\xdf\x57\x60\xa8\xa8\xd4\x4b\x6f\x5c\x2c\xcd\xec\xec\x7a\x76\xec\xbc\xc1\x8c\x22\xe0\x86\x22\x69\xcb\x9a\x40\xb9\x04\x52\x53\xc0\x4d\xd9\xd2\x16\x66\x21\x48\xfe\x64\x9c\x3b\xcc\x16\xc9\x60\x05\xf3\x02\x3c\xdb\xa9\xd3\x39\x3a\x98\x27\x0f\x92\x43\xd7\x95\x05\xac\x9b\x72\x95\x35\x5b\x78\xc6\xed\x63\x02\x00\x10\xd5\xbc\x67\x1e\x68\xb9\xc2\x96\x66\xab\x35\x7d\x3d\x0f\x27\x5d\x55\x41\x81\xcb\xac\xab\x28\x90\xfa\x65\x9e\xc6\x96\x70\xe0\xff\x6d\xd1\x6c\x12\xe0\xc5\x97\xff\x51\x25\xe9\x22\xf9\x6d\x95\x05\xff\x66\xec\xfd\xb9\xd5\x41\xa9\xc8\x5b\x31\x4a\xe7\x85\xed\x8d\x86\xd3\xd0\xc8\xca\x89\x8d\x51\xfe\xc7\x4e\xd6\xcb\x41\x89\x7b\x58\xca\xa9\x30\x9e\x5d\x46\xe8\xa5\x57\xe2\x06\x73\xe1\x06\x2b\x0f\x5e\x1a\x7d\xc3\xee\x0c\x3f\xde\xc0\x3d\xfb\x34\x56\x7a\xe1\xfa\xc1\x04\xed\x41\xea\x4b\x61\x30\xd3\x24\xb4\x87\x77\x67\xf4\xee\x72\xdf\x60\xac\x80\xbd\x32\xec\xa2\x51\xf2\x43\xf0\x3e\x38\x61\x5d\x0c\x97\x59\xcb\x8e\xb1\x16\xdf\xbe\xbf\x46\xd4\xe0\x12\x1b\x24\x39\xb6\xd7\x5f\x31\x97\x3c\x85\x9a\x40\x81\x15\x52\x84\x3c\x6b\xf3\xac\xc0\x24\x5d\x7c\x07\x00\x00\xff\xff\x1e\xbc\x66\xbd\x1a\x03\x00\x00")

func _1_init_db_structureUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_init_db_structureUpSql,
		"1_init_db_structure.up.sql",
	)
}

func _1_init_db_structureUpSql() (*asset, error) {
	bytes, err := _1_init_db_structureUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_init_db_structure.up.sql", size: 794, mode: os.FileMode(420), modTime: time.Unix(1611841440, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_tag_in_articlesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\x28\x49\x4c\x8f\xcf\xcc\x8b\x4f\x2c\x2a\xc9\x4c\xce\x49\x2d\xb6\x06\x04\x00\x00\xff\xff\xef\x6c\x3b\xdb\x1b\x00\x00\x00")

func _2_add_tag_in_articlesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_tag_in_articlesDownSql,
		"2_add_tag_in_articles.down.sql",
	)
}

func _2_add_tag_in_articlesDownSql() (*asset, error) {
	bytes, err := _2_add_tag_in_articlesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_tag_in_articles.down.sql", size: 27, mode: os.FileMode(420), modTime: time.Unix(1611841491, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_tag_in_articlesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcf\xb1\xaa\x83\x30\x14\xc6\xf1\xf9\xfa\x14\xdf\xa8\x70\xdf\xa0\x53\x30\x47\x90\xc6\x28\x31\xa1\xd8\x45\xd2\x24\x94\x80\x94\xa2\xe9\xfb\x17\xab\x5d\x4a\x97\x2e\x67\x38\xf0\x83\xef\xef\xe6\x60\x53\x40\xb2\x97\x69\xbd\xd7\x31\xde\x46\x3b\xa7\xe8\xa6\xb0\x20\xcf\xfe\xa2\x87\x31\x35\x47\xa7\xea\x86\xa9\x01\x47\x1a\xfe\x33\x00\xd8\xa0\x1f\x6d\x82\xae\x1b\xea\x35\x6b\x3a\x7d\x86\x6c\x35\xa4\x11\x02\x9c\x2a\x66\x84\x86\x6c\x4f\x79\xb1\x91\xc7\xdd\xff\x4a\x5e\x8b\xf6\x09\x8a\x2a\x52\x24\x4b\xea\xd7\xf7\x92\x47\x5f\xa0\x95\xe0\x24\x48\x13\x9c\x5d\x9c\xf5\x61\x63\x7b\xc1\x37\xfa\x8e\xfb\xe0\x25\xeb\x4b\xc6\x29\x2b\x0e\xcf\x00\x00\x00\xff\xff\x92\x6f\xde\xd9\x12\x01\x00\x00")

func _2_add_tag_in_articlesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_tag_in_articlesUpSql,
		"2_add_tag_in_articles.up.sql",
	)
}

func _2_add_tag_in_articlesUpSql() (*asset, error) {
	bytes, err := _2_add_tag_in_articlesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_tag_in_articles.up.sql", size: 274, mode: os.FileMode(420), modTime: time.Unix(1611841434, 0)}
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
	"1_init_db_structure.down.sql":   _1_init_db_structureDownSql,
	"1_init_db_structure.up.sql":     _1_init_db_structureUpSql,
	"2_add_tag_in_articles.down.sql": _2_add_tag_in_articlesDownSql,
	"2_add_tag_in_articles.up.sql":   _2_add_tag_in_articlesUpSql,
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
	"1_init_db_structure.down.sql":   &bintree{_1_init_db_structureDownSql, map[string]*bintree{}},
	"1_init_db_structure.up.sql":     &bintree{_1_init_db_structureUpSql, map[string]*bintree{}},
	"2_add_tag_in_articles.down.sql": &bintree{_2_add_tag_in_articlesDownSql, map[string]*bintree{}},
	"2_add_tag_in_articles.up.sql":   &bintree{_2_add_tag_in_articlesUpSql, map[string]*bintree{}},
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
