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

var _migrations1688189195_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xcd\x6e\x9b\x4a\x14\xc7\xf7\x7e\x8a\xa3\x6c\x8c\x25\xeb\x92\x9b\x7b\x73\x17\x37\xaa\x2a\x12\x4f\x12\x54\x02\x29\x86\x7c\x74\x83\xb0\x21\x36\x12\x86\x68\x06\xd4\x2e\xad\x76\xd1\x45\x16\x59\x74\xd7\xa7\x88\x5a\x55\x4d\x5b\xc7\x79\x85\xc3\x2b\xf4\x49\xaa\x61\xfc\x81\x0d\x4e\xdc\xa8\xb3\x83\x39\xe7\x37\x67\xce\xf9\x9f\x33\x7b\x26\x51\x2c\x02\x96\xb2\xab\x11\x50\xf7\x41\x37\x2c\x20\x67\x6a\xdb\x6a\xc3\x85\xef\x7b\xac\x26\xd5\x00\x00\x02\x0f\x8a\xcb\xb6\xd5\x16\xac\x58\x9c\xa0\xdb\x9a\x06\x2d\xb2\xaf\xd8\x9a\x05\x3d\x3f\x72\xa8\x1b\x79\xf1\xc0\x49\xd3\xc0\x93\x1a\x70\x6c\xaa\x47\x8a\x79\x0e\x2f\xc8\x79\x33\xc7\x27\x41\x12\xfa\x73\xc4\x89\x62\xee\x1d\x2a\xa6\xb4\xb5\xbd\xdd\x58\x85\x17\x8e\x2c\x48\x7c\x27\xa5\xe1\x92\xe3\xdf\x9b\x5b\xff\x2e\x7b\x2e\x3a\xf2\xcb\x3d\xc9\x31\x74\x59\xe2\x04\x89\x3f\x70\xfa\x2e\xeb\xcf\x1c\xff\xd9\x2a\x45\x5a\xce\x44\xbd\x2e\x18\x3e\xa5\x31\x75\xba\x71\x1a\x25\xfc\x53\xd5\x2d\x72\x40\xcc\xf5\xb2\xb9\x29\x10\x5d\xea\xbb\x89\xef\x39\x6e\x4e\x00\x4b\x3d\x22\x6d\x4b\x39\x3a\x86\x53\xd5\x3a\x34\x6c\x2b\xff\x03\xaf\x0c\x9d\xcc\x59\x53\x44\x14\xbf\x96\x1a\x4b\xd7\x4a\x2f\xbd\x75\x79\xdc\xa9\xd6\xd8\xa9\xd5\x26\xda\x51\xf5\x16\x39\x5b\xd2\xce\x46\x2e\x1e\x67\x4a\x0d\x22\xcf\x7f\xb3\x01\x86\x2e\x44\x05\x52\xe1\xb8\x16\x69\xef\x35\x76\xa6\x30\x5b\x57\x5f\xda\x0f\x33\x69\xe8\xa4\x25\xe0\xb4\x9e\x3c\x2e\x55\x6f\x13\xd3\xe2\x69\x35\xc4\xb6\x94\xeb\xab\x39\x53\x4b\x73\x56\xfe\x46\xed\x44\xd1\x6c\xd2\x06\xa9\x7e\x10\xc3\xcf\xe1\x07\xc0\x8f\x38\xc6\x11\xde\xe3\x2d\xfe\xc0\xdb\x6c\x98\xbd\xc3\x2f\x38\xca\xae\xf0\x5b\x13\x70\x84\x77\x38\xc6\xcf\x38\xc6\x7b\x1c\x67\x6f\x71\x9c\xbd\xc7\x3b\xbe\x07\xd9\x35\x7e\xcd\xae\xf0\x3b\xe0\x7d\x36\xe4\x36\xd9\x10\x6f\x70\x84\x23\xce\xc0\x31\x7e\xc2\x1b\xbc\xc3\xdb\xec\xba\xde\x84\x7a\x3f\x49\x2e\xd9\xff\xb2\xdc\x77\x3b\xf4\xaf\x6e\x3c\x90\x69\x2a\xf7\xd3\x8e\xdc\x8b\xe5\x89\x42\xf8\xaa\x34\xa3\x8c\x4d\x4d\xdd\x30\x94\x9f\x5f\x84\xcf\x68\x5a\x2f\x54\xa3\xaa\x93\x53\xe6\xd3\x8a\x4e\xde\x55\x0f\x54\xdd\xaa\x94\xdd\x2a\x15\x96\x3a\xf7\x22\xa0\x2c\x71\x22\x77\xe0\xcf\x5a\xe1\xbf\x52\x07\xad\xc2\x15\x3a\x2a\x27\x3c\x19\xc1\x2f\x28\x08\x4f\x46\x04\xcc\xe9\xc4\x42\xfc\xbb\x86\xa1\xad\x93\x95\xf2\x45\xa2\x5e\xea\xf6\x78\x14\x62\x26\x3c\x16\x41\x19\x51\x68\xeb\x87\x5a\x70\xcd\x4e\x5e\x6b\x2a\x88\x28\x96\x7a\xba\x4a\x45\x2c\xed\xb0\x2e\x0d\x2e\x93\x20\x8e\x2a\xd4\x34\x79\x13\x66\x52\x29\xae\xdf\x78\x0a\x78\x29\x1d\x81\x9d\xe8\xb3\x62\x70\x8b\xfd\xa5\x13\xff\x40\x06\x8b\x29\x78\x68\x12\x2d\xa4\xc2\xc9\x43\x16\x13\x65\x3e\x97\x16\x4c\x40\x9a\x5c\xab\x39\x8d\x7f\x3e\xf1\xd6\x38\x40\xf8\xac\x42\xcf\x89\x35\x45\xb3\x88\x39\x29\x9e\x78\xbe\x79\x4a\x8c\x53\x9d\xff\x35\x80\x32\xd6\x89\x93\x9d\x05\x33\x31\x1b\x1e\x35\x5b\x2c\x7e\x95\xf9\xaf\x00\x00\x00\xff\xff\x43\x41\x25\x3a\x4a\x08\x00\x00")

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

	info := bindataFileInfo{name: "migrations/1688189195_init.up.sql", size: 2122, mode: os.FileMode(436), modTime: time.Unix(1689691540, 0)}
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
