// Code generated for package mysqlstorage by go-bindata DO NOT EDIT. (@generated)
// sources:
// storage/mysqlstorage/schema.sql
package mysqlstorage

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

var _schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x96\x51\x6f\xda\x3a\x14\xc7\xdf\xf3\x29\xfc\x72\x05\x5c\x41\x75\xef\xb4\x3e\xa1\x49\x4b\x89\xa1\xd1\x42\x82\x12\xd3\xad\x4f\x91\x49\x3c\xb0\x96\xd8\x96\xed\x94\x76\x9f\x7e\x22\x25\xab\x1b\xe2\x00\x5b\x1e\x4f\xce\xff\xe7\x63\xfb\xf8\xaf\xe3\xc5\xd1\x0a\x24\xb3\x7b\xb8\x74\x81\x3f\x07\xf0\x9b\x9f\xa0\x04\xec\x09\xd6\x3b\x22\x15\xaf\x64\x46\xa6\xce\x2c\x86\x2e\x82\x4d\x5a\xeb\xe7\x3a\x81\xed\xd0\x22\x76\x43\x04\x12\x18\xc0\x19\x1a\xfb\x61\x02\x63\x34\x5e\xaf\x3c\x17\xc1\xb1\x07\x03\x88\x20\x88\xc2\xf7\x9a\x9b\x7f\x01\x8a\xc0\xe0\x5d\x6c\xf0\x79\xf0\xcf\x00\xf8\x1e\x0c\x91\x3f\xf7\xa1\x07\xee\x1e\x5b\x19\x13\x81\x95\xda\x73\x99\x0f\xa6\x8e\x33\x99\x38\x93\x49\x43\x4d\x73\xac\x31\xd0\x78\x53\x10\xa0\x34\x97\x44\x81\x92\x68\xc2\x0b\xbe\xa5\x19\x2e\x40\xfd\x7b\x88\x33\x5d\xe1\x02\x60\x96\x83\xef\x5c\x92\x0c\x2b\x0d\x88\xd2\xb4\xc4\x9a\xa8\xd1\xcd\x01\x79\xdc\x3a\x72\xef\x02\x78\x38\xa0\x30\x42\xad\x43\x7a\x5d\x6a\xe8\x00\x00\x6a\x6c\xaa\x69\x49\x40\xf7\xe7\x87\x08\x2e\x60\x5c\x53\xc2\x75\x10\x8c\x0d\xd5\x8b\xb0\xa9\x1e\xdc\x78\x76\xef\xc6\xc3\x0f\xb7\xb7\xa3\x57\xc5\x61\x89\x9f\x9c\xd9\x04\x1d\x8a\x02\x6b\xaa\xab\xdc\xaa\xf0\xe0\xcc\x5f\xba\x01\x18\xfe\xff\xdf\xf8\x63\xa3\xe1\x6c\xdb\x27\xea\xd2\x60\x21\xb0\x24\x4c\xa7\x9a\x94\x82\x48\xac\x2b\x49\xd2\x1d\xdd\xee\xfe\x44\x73\x3c\xca\x3b\x7f\xe1\x87\xa8\x27\xb7\xe0\xfb\xab\x6b\x2a\xf8\xbe\xb9\xa9\xb3\xfc\x12\x3f\x5f\xcd\x2f\xf1\xf3\xe5\x7c\xca\xae\xe7\x53\xd6\xc5\x37\x52\xfe\xea\xce\xce\x69\xb2\x82\x57\x79\x9a\xf1\x27\x22\x2f\x5e\x27\x27\xfb\x54\x70\xca\xf4\x15\xfd\xb4\xab\x4a\x9a\x53\xfd\x62\x91\x74\x6a\x68\xc6\x99\x2d\x1f\x74\xbe\x8e\x92\x73\x96\x8a\x1d\x56\x96\x56\xef\x5a\x85\xf7\x3d\x40\x8b\x46\x48\x92\x51\x91\x52\xa6\x09\x53\x1d\xbb\xba\x44\x63\x76\xe3\x35\x1a\xd3\x97\xcc\x8e\x39\xe6\x0a\xc9\x37\x78\x43\x8b\x93\xaa\x7a\xf8\x56\xcf\x3a\x3d\x61\x21\x89\x52\xb6\xa6\xb4\xac\xa2\xaa\xb2\xc4\xd2\x7a\xf5\x1d\xab\xa8\x8a\x49\xaa\x88\xd5\x82\xcd\x7d\xab\x8a\x29\xa2\xed\x6e\x6d\x79\x55\xa6\x99\xf5\x56\x6f\x31\xb3\xb3\x7c\xc3\xcc\x2e\xe6\x1b\x66\x76\x96\xdf\x6a\x9f\x8b\xf8\xad\xf6\xe9\xe7\xd3\xd3\xc7\x77\x96\xff\x66\x66\x9d\xfc\xea\x29\xa5\x2c\x27\xa7\x85\x1f\xbf\x93\x44\xdb\xb5\x9a\xd0\x27\xaa\x68\x57\xc3\xf7\x15\xbd\xa7\x2c\x4f\x37\x04\x4b\xca\xb6\xbd\x85\xd4\x89\xdb\x4a\x5d\xe3\x75\xbf\x35\x96\xea\xcd\xe2\xeb\x5c\x25\x08\xc9\x2f\xe3\xd7\xa2\x59\x14\x26\x28\x76\xfd\x10\x01\xf1\x23\x7d\x37\xbf\xac\x62\x7f\xe9\xc6\x8f\xe0\x0b\x7c\x3c\xce\x32\xc0\x9c\x67\xc6\xad\xd0\x8b\x30\x42\xcd\x34\xf2\x16\x69\xa6\x0d\x23\xd2\xcc\x12\x75\x64\xe4\x8c\x40\x1c\x7d\x4d\xe7\x51\xbc\x74\xd1\xa7\x59\xb4\x5c\xc5\x30\x49\xa0\x37\x75\x1c\x67\x1e\xac\x93\xfb\x43\x41\x0f\x7e\x00\x17\x30\x99\x3a\xbf\x02\x00\x00\xff\xff\x93\x91\x12\x06\xa5\x0a\x00\x00")

func schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSql,
		"schema.sql",
	)
}

func schemaSql() (*asset, error) {
	bytes, err := schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.sql", size: 2725, mode: os.FileMode(400), modTime: time.Unix(1257894000000000000, 0)}
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
	"schema.sql": schemaSql,
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
	"schema.sql": {schemaSql, map[string]*bintree{}},
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
