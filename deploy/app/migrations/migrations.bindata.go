// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/1688189195_init.down.sql
// migrations/1688189195_init.up.sql
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

var _migrations1688189195_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2f\xcd\x8c\x4f\xce\xcf\x4b\xcb\x4c\xb7\x06\x04\x00\x00\xff\xff\x49\xa7\x32\xcb\x20\x00\x00\x00")

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

	info := bindataFileInfo{name: "migrations/1688189195_init.down.sql", size: 32, mode: os.FileMode(436), modTime: time.Unix(1665214714, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1688189195_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcd\x6e\x9b\x4a\x14\xc7\xf7\x7e\x8a\xa3\x6c\x8c\x25\xeb\x92\x9b\x7b\xd3\x45\xa3\xaa\x22\xf1\x24\x41\x25\x90\x62\xc8\x47\x37\x08\x1b\x62\x23\x61\x88\x66\x40\xed\xd2\x6a\x17\x5d\x64\x91\x45\x77\x7d\x8a\xa8\x55\xd5\xb4\x75\x9c\x57\x38\xbc\x42\x9f\xa4\x1a\xc6\xdf\x60\xc7\x89\x3a\x3b\xc3\x39\xbf\x39\x73\xe6\x7f\xfe\x78\xcf\x24\x8a\x45\xc0\x52\x76\x35\x02\xea\x3e\xe8\x86\x05\xe4\x4c\x6d\x5a\x4d\xb8\xf0\x7d\x8f\x55\xa4\x0a\x00\x40\xe0\xc1\xec\xb2\x6d\xb5\x01\x4b\x16\x27\xe8\xb6\xa6\x41\x83\xec\x2b\xb6\x66\x41\xc7\x8f\x1c\xea\x46\x5e\xdc\x73\xd2\x34\xf0\xa4\x1a\x1c\x9b\xea\x91\x62\x9e\xc3\x2b\x72\x5e\xcf\xf1\x49\x90\x84\xfe\x14\x71\xa2\x98\x7b\x87\x8a\x29\x6d\x6d\x6f\xd7\x96\xe1\x45\x22\x0b\x12\xdf\x49\x69\xb8\x90\xf8\xef\xe6\xd6\xff\x8b\x99\xf3\x89\xfc\x70\x4f\x4a\x0c\x5d\x96\x38\x41\xe2\xf7\x9c\xae\xcb\xba\x93\xc4\xff\xb6\x0a\x95\x16\x3b\x51\xad\x0a\x86\x4f\x69\x4c\x9d\x76\x9c\x46\x09\xff\xa9\xea\x16\x39\x20\xe6\x7a\xdd\xdc\x14\x88\x36\xf5\xdd\xc4\xf7\x1c\x37\x27\x80\xa5\x1e\x91\xa6\xa5\x1c\x1d\xc3\xa9\x6a\x1d\x1a\xb6\x95\x3f\x81\x37\x86\x4e\xa6\xac\x31\x22\x8a\xdf\x4a\xb5\x85\x63\xa5\x97\xde\xba\x3c\x9e\x54\xa9\xed\x54\x2a\x23\xed\xa8\x7a\x83\x9c\x2d\x68\x67\x23\x17\x8f\x33\xa6\x06\x91\xe7\xbf\xdb\x00\x43\x17\xa2\x02\x69\x66\xbb\x06\x69\xee\xd5\x76\xc6\x30\x5b\x57\x5f\xdb\xab\x99\x34\x74\xd2\x02\x70\x7c\x9f\xbc\x2e\x55\x6f\x12\xd3\xe2\x6d\x35\xc4\x6b\x29\xd7\x57\x7d\xa2\x96\xfa\xe4\xfa\x6b\x95\x13\x45\xb3\x49\x13\xa4\xea\x41\x0c\xbf\xfb\x9f\x00\x3f\xe3\x10\x07\x78\x8f\xb7\xf8\x0b\x6f\xb3\x7e\xf6\x01\xbf\xe1\x20\xbb\xc2\x1f\x75\xc0\x01\xde\xe1\x10\xbf\xe2\x10\xef\x71\x98\xbd\xc7\x61\xf6\x11\xef\xf8\x3b\xc8\xae\xf1\x7b\x76\x85\x3f\x01\xef\xb3\x3e\x8f\xc9\xfa\x78\x83\x03\x1c\x70\x06\x0e\xf1\x0b\xde\xe0\x1d\xde\x66\xd7\xd5\x3a\x54\xbb\x49\x72\xc9\x9e\xcb\x72\xd7\x6d\xd1\x7f\xda\x71\x4f\xa6\xa9\xdc\x4d\x5b\x72\x27\x96\x47\x0a\xe1\xab\x34\x8c\x32\x36\x0e\x75\xc3\x50\x7e\x79\x11\xbe\xa0\x69\x75\xe6\x36\xca\x26\x39\x65\x3e\x2d\x99\xe4\x5d\xf5\x40\xd5\xad\x52\xd9\x2d\x53\x61\x61\x72\xdb\x5d\x37\x71\x04\xf3\xf1\xb8\xd1\x28\x06\x94\x25\x4e\xe4\xf6\xfc\xc9\x34\x3d\x2b\x0c\xe1\x6a\x44\x3e\x94\x39\xe1\xc9\x08\xde\x23\x41\x78\x32\x22\x60\x4e\x2b\x16\xf3\xb3\x6b\x18\xda\x3a\x9d\x28\x1e\x24\xea\xa4\x6e\x87\x57\x21\x6c\xe5\xa1\x0a\x8a\x88\x19\x67\x58\x35\xc5\x6b\x9a\xc1\x5a\xc6\x22\xaa\x58\xb0\x85\x32\x21\xb2\xb4\xc5\xda\x34\xb8\x4c\x82\x38\x2a\x11\xe4\xe8\xb3\x32\x51\xdb\xec\x7a\xc4\xd7\x84\x5f\xe5\xbc\x26\x4b\xbc\x5f\xbc\x5f\xd8\xf1\x2f\x74\x70\xb6\x05\xab\xcc\x6c\xae\x15\x4e\x5e\xb2\x30\xa5\xa9\xb5\xcd\x85\x80\x34\x3a\x56\x7d\x5c\xff\xd4\x34\xd7\xd8\x40\xe4\x2c\x43\x4f\x89\x15\x45\xb3\x88\x39\xba\x3c\xf1\x0f\x80\xb7\xc4\x38\xd5\xf9\x53\x03\x28\x63\xad\x38\xd9\x99\x0b\x13\xf6\xf2\x60\xd8\xfc\xe5\x97\x85\xff\x09\x00\x00\xff\xff\xc5\xda\x44\x2b\x8d\x08\x00\x00")

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

	info := bindataFileInfo{name: "migrations/1688189195_init.up.sql", size: 2189, mode: os.FileMode(436), modTime: time.Unix(1689505960, 0)}
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
	"migrations/1688189195_init.down.sql": migrations1688189195_initDownSql,
	"migrations/1688189195_init.up.sql":   migrations1688189195_initUpSql,
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
		"1688189195_init.down.sql": &bintree{migrations1688189195_initDownSql, map[string]*bintree{}},
		"1688189195_init.up.sql":   &bintree{migrations1688189195_initUpSql, map[string]*bintree{}},
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
