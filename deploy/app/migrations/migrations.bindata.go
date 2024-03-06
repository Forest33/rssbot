// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/1688189195_init.down.sql
// migrations/1688189195_init.up.sql
// migrations/1688189196_items.down.sql
// migrations/1688189196_items.up.sql
package migrations

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

var _migrations1688189195_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrations1688189195_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189195_initDownSql,
		"migrations/1688189195_init.down.sql",
	)
}

func migrations1688189195_initDownSql() (*asset, error) {
	bytes, err := migrations1688189195_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189195_init.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1709700192, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1688189195_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xcd\x6e\xd3\x40\x10\xc7\xef\x7e\x8a\x51\x2f\x71\xa4\x08\x97\x42\x39\x50\x21\xe4\x36\xdb\xd6\x22\xb5\x8b\x63\xf7\x83\x8b\xe5\xc4\x6e\x62\xc9\xb1\xab\x5d\x5b\x70\x8c\xe0\xc0\xa1\x87\x1e\xb8\xf1\x14\x15\x08\x51\x20\x4d\x5f\x61\xfc\x0a\x3c\x09\x5a\x6f\x3e\x9c\xc4\x69\x43\x61\x6f\xb1\x67\x7e\x9e\x9d\xf9\xff\x27\x3b\x26\x51\x2d\x02\x96\xba\xdd\x20\xa0\xed\x82\x6e\x58\x40\x4e\xb4\xa6\xd5\x84\x33\xdf\xf7\x98\x24\x4b\x00\x00\x81\x07\xc5\x63\xdb\x5a\x1d\x96\x1c\x4e\xd0\xed\x46\x03\xea\x64\x57\xb5\x1b\x16\x74\xfc\xc8\xa1\x6e\xe4\xc5\x3d\x27\x4d\x03\x4f\xae\xc2\xa1\xa9\x1d\xa8\xe6\x29\xbc\x22\xa7\xb5\x1c\x9f\x04\x49\xe8\x4f\x11\x47\xaa\xb9\xb3\xaf\x9a\xf2\xc6\xe6\x66\x75\x19\x5e\x24\xb2\x20\xf1\x9d\x94\x86\x73\x89\x8f\xd7\x37\x9e\xce\x67\xce\x26\xf2\xcb\x3d\x28\x31\x74\x59\xe2\x04\x89\xdf\x73\xba\x2e\xeb\x4e\x12\x9f\x6c\x2c\x54\xba\xd8\x89\x4a\x45\x30\x7c\x4a\x63\xea\xb4\xe3\x34\x4a\xf8\x4f\x4d\xb7\xc8\x1e\x31\x57\xeb\xe6\xba\x40\xb4\xa9\xef\x26\xbe\xe7\xb8\x39\x01\x2c\xed\x80\x34\x2d\xf5\xe0\x10\x8e\x35\x6b\xdf\xb0\xad\xfc\x09\xbc\x31\x74\x32\x65\x8d\x11\x51\xfc\x56\xae\xce\x5d\x2b\x3d\xf7\x56\xe5\xf1\x24\xa9\xba\x25\x49\x23\xed\x68\x7a\x9d\x9c\xcc\x69\x67\x2d\x17\x8f\x33\xa6\x06\x91\xe7\xbf\x5b\x03\x43\x17\xa2\x02\xb9\xf0\xb9\x3a\x69\xee\x54\xb7\xc6\x30\x5b\xd7\x5e\xdb\x77\x33\x69\xe8\xa4\x0b\xc0\xf1\x3c\x79\x5d\x9a\xde\x24\xa6\xc5\xdb\x6a\x88\xd7\x72\xae\xaf\xda\x44\x2d\xb5\xc9\xf8\xab\xd2\x91\xda\xb0\x49\x13\xe4\xca\x5e\x0c\xbf\xfb\x9f\x00\x3f\xe3\x10\x07\x78\x8b\xd7\xf8\x0b\xaf\xb3\x7e\xf6\x01\xbf\xe1\x20\xbb\xc0\x1f\x35\xc0\x01\xde\xe0\x10\xbf\xe2\x10\x6f\x71\x98\xbd\xc7\x61\xf6\x11\x6f\xf8\x3b\xc8\x2e\xf1\x7b\x76\x81\x3f\x01\x6f\xb3\x3e\x8f\xc9\xfa\x78\x85\x03\x1c\x70\x06\x0e\xf1\x0b\x5e\xe1\x0d\x5e\x67\x97\x95\x1a\x54\xba\x49\x72\xce\x9e\x2b\x4a\xd7\x6d\xd1\x47\xed\xb8\xa7\xd0\x54\xe9\xa6\x2d\xa5\x13\x2b\x23\x85\xf0\x53\x1a\x46\x19\x1b\x87\xba\x61\xa8\xbc\x3c\x0b\x5f\xd0\xb4\x52\x98\x46\x99\x93\x53\xe6\xd3\x12\x27\x6f\x6b\x7b\x9a\x6e\x95\xca\x6e\x99\x0a\x17\x9c\x7b\x16\x50\x96\x38\x91\xdb\xf3\x27\x56\x78\xb6\xe0\xa0\x65\xb8\x82\xa3\x72\xc2\x83\x11\xfc\x82\x82\xf0\x0f\x55\x44\x9d\xd4\xed\x70\x84\x30\xf4\x7d\xe9\x8b\x88\x82\x27\xef\xf2\xcf\x8a\x36\x5c\xc9\xd2\xa2\x8a\x39\x43\x96\x49\x80\xa5\x2d\xd6\xa6\xc1\x79\x12\xc4\x51\x89\x14\x46\x0b\x7d\x32\xe7\xe2\xf9\x8b\x3d\xce\xe7\xe0\x08\xec\x48\x5c\x25\x5b\x57\xbc\x9f\xfb\xe2\x7f\xe8\x60\xb1\x05\x77\xad\x91\x99\x56\x38\x79\xc9\x62\x1d\x4c\x97\xca\x4c\x08\xc8\xa3\x6b\xd5\xc6\xf5\x4f\xd7\xd5\x0a\x1f\x10\x39\xcb\xd0\x53\xa2\xa4\x36\x2c\x62\x8e\x86\x27\xfe\x7b\x79\x4b\x8c\x63\x9d\x3f\x35\x80\x32\xd6\x8a\x93\xad\x99\x30\x61\xec\x7b\xc3\x66\x87\x5f\x16\xfe\x27\x00\x00\xff\xff\x35\x59\x61\x95\x07\x08\x00\x00")

func migrations1688189195_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189195_initUpSql,
		"migrations/1688189195_init.up.sql",
	)
}

func migrations1688189195_initUpSql() (*asset, error) {
	bytes, err := migrations1688189195_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189195_init.up.sql", size: 2055, mode: os.FileMode(436), modTime: time.Unix(1689691654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1688189196_itemsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrations1688189196_itemsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189196_itemsDownSql,
		"migrations/1688189196_items.down.sql",
	)
}

func migrations1688189196_itemsDownSql() (*asset, error) {
	bytes, err := migrations1688189196_itemsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189196_items.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1709700192, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1688189196_itemsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xbd\x4e\xc3\x30\x14\x85\x77\x3f\xc5\x51\xa7\x44\xea\x04\x63\x27\xd3\xde\xaa\x16\xa9\x5d\x1c\x9b\xb6\x2c\x56\xc0\x86\x66\x68\x22\xe5\x47\xf0\xf8\x28\x69\x42\xaa\x0a\x84\x07\xcb\x3f\xd7\x9f\xcf\x39\x77\xa9\x89\x1b\x82\xe1\x0f\x09\x41\xac\x21\x95\x01\x1d\x44\x6a\x52\xbc\x87\xe0\x5d\xde\x84\x73\xcd\x22\x06\x00\xb9\xc7\x38\xac\x15\xab\x71\xdd\x3d\x91\x36\x49\xc6\xfd\x8a\xd6\xdc\x26\x06\x1f\xa1\x70\x55\x56\xf8\xf2\xec\xda\x36\xf7\x51\x8c\x9d\x16\x5b\xae\x8f\x78\xa4\xe3\xbc\x27\x5e\xbe\xf0\x7f\x11\x2f\x45\x9d\x04\x77\xca\xea\x13\xf0\xcc\xf5\x72\xc3\x75\x74\x7f\x17\xdf\x14\xbd\x55\x21\x6b\x82\x77\x59\x03\x23\xb6\x94\x1a\xbe\xdd\x61\x2f\xcc\x46\x59\xd3\x9f\xe0\x45\x49\xfa\xd1\x56\x94\x9f\xd1\x84\x60\xf1\x82\xb1\x21\x09\x2b\xc5\x93\x25\x08\xb9\xa2\xc3\x4d\x20\xb3\x29\x91\x7e\x76\x6d\x5e\xf8\xf0\x35\x83\x92\x57\x61\x21\x1a\x5c\xcd\x27\xe5\xf1\x62\xc4\xff\xc7\x1d\x7d\xfc\x4e\x9e\x5c\x76\x8a\x79\x62\x48\x0f\xad\xbb\x6a\x56\x17\x87\xda\xcb\xee\x4a\xa1\xaa\xeb\xd7\xb2\x59\xb0\xef\x00\x00\x00\xff\xff\x49\x99\x7a\x84\xe8\x01\x00\x00")

func migrations1688189196_itemsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189196_itemsUpSql,
		"migrations/1688189196_items.up.sql",
	)
}

func migrations1688189196_itemsUpSql() (*asset, error) {
	bytes, err := migrations1688189196_itemsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189196_items.up.sql", size: 488, mode: os.FileMode(436), modTime: time.Unix(1709712742, 0)}
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
	"migrations/1688189195_init.down.sql":  migrations1688189195_initDownSql,
	"migrations/1688189195_init.up.sql":    migrations1688189195_initUpSql,
	"migrations/1688189196_items.down.sql": migrations1688189196_itemsDownSql,
	"migrations/1688189196_items.up.sql":   migrations1688189196_itemsUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"migrations": &bintree{nil, map[string]*bintree{
		"1688189195_init.down.sql":  &bintree{migrations1688189195_initDownSql, map[string]*bintree{}},
		"1688189195_init.up.sql":    &bintree{migrations1688189195_initUpSql, map[string]*bintree{}},
		"1688189196_items.down.sql": &bintree{migrations1688189196_itemsDownSql, map[string]*bintree{}},
		"1688189196_items.up.sql":   &bintree{migrations1688189196_itemsUpSql, map[string]*bintree{}},
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
