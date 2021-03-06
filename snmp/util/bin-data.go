// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// config/conf.ini
// config/dt.csv
package util

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

var _configConfIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x2e\x4e\x49\x2b\x4e\xe1\xe5\x72\x49\x2d\xcb\x4c\x4e\x0d\xae\xcc\x4b\x76\xce\xcf\x4b\xcb\x4c\x57\xb0\x55\x30\x30\x50\xd0\x82\x41\x40\x00\x00\x00\xff\xff\xae\x94\xdb\x95\x27\x00\x00\x00")

func configConfIniBytes() ([]byte, error) {
	return bindataRead(
		_configConfIni,
		"config/conf.ini",
	)
}

func configConfIni() (*asset, error) {
	bytes, err := configConfIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/conf.ini", size: 39, mode: os.FileMode(438), modTime: time.Unix(1450533414, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configDtCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x97\xd1\x4e\xe3\x3a\x10\x86\xef\x8f\x74\xde\x81\x07\xa8\xad\x99\xb1\x63\x8f\x2f\x39\x29\x54\xe8\xb4\x47\x55\xd3\x03\xdd\xab\x15\x62\x57\xbb\x95\x10\xac\x54\x10\xec\xdb\xaf\xdc\xa4\x49\x8a\x67\xc2\x0d\x42\xd5\xff\xe5\xf7\xfc\x1e\x4f\x1c\x8b\xd6\xd9\x60\xd1\x7a\x8b\x36\x59\xb4\xe8\xb0\x9a\xd5\xfb\xc3\xc3\xb3\xa9\xef\x5f\xee\x1f\x7f\x1f\x5e\x0c\xa5\x00\xf5\x7a\x6e\x78\xbd\x9d\xcd\xfe\xfe\x4b\x60\x42\x10\x19\xc3\xeb\xda\x2c\x35\x46\xf6\x31\xbc\x9d\x60\xbc\xce\x34\x1a\x13\x65\x06\x49\x5f\x1c\x45\x56\x43\xd8\x6a\x21\xa0\x1c\xc2\x22\xaf\x4e\x46\x62\x05\x12\x62\xd6\x8f\xaf\x07\x43\x7e\x59\x9b\xe5\x85\x46\x8a\xe9\x8d\xc8\x46\x23\x7d\x9a\x22\xd7\x53\x9e\x62\xfa\x23\x52\xf5\xac\x68\x8a\xdc\x4e\x79\x8a\xbb\x37\x22\xf5\x3a\xc5\x2d\x6c\x49\xcf\xeb\x66\x3b\x61\xea\x3e\x43\xf5\x4a\x71\x0a\x9d\xac\x54\x6c\x9f\x11\xa9\x78\x12\x3a\xb9\xc3\x77\x13\x47\xa9\xc2\xa2\x44\x17\x2b\xc8\x91\xca\xe7\xa8\x72\x45\xc3\x1d\x81\x1c\x87\x02\x14\xf5\x74\x0e\x1a\x90\x8a\xe8\x32\xb0\x30\x18\xb6\x73\x91\x08\x50\x34\x64\x4b\x78\x56\xaa\x08\x20\x96\xbd\xd0\xcb\x08\xa4\x78\xe4\xa4\x0c\xfe\xaf\xd8\x14\xfd\x7e\x82\x14\x1b\x0c\xbe\xac\x9e\xdb\x7c\x35\xa0\xf4\xe0\x36\x5f\x0d\x28\x6b\x6f\x1d\x34\x7d\x59\x79\x6b\xa0\x34\x70\x28\x0e\x5c\xf7\x7c\x39\x24\x8c\xe5\x64\xee\x0c\x14\x80\x63\x31\xba\x3a\x40\x09\x95\xb9\x98\xaf\x47\x00\x49\x06\x08\x83\x1c\x2a\xd2\x4e\x25\xe4\x54\xc9\xeb\x84\x9c\xab\x67\x9d\x28\x8f\x52\xe7\xa1\x24\x55\x38\xf8\x0a\x40\xd9\x66\x28\x0e\x76\x16\xef\x0c\x06\x0d\x28\x56\xd3\x02\x8e\x24\x80\xac\x3b\xfe\x94\xdc\x29\xda\xff\xbe\xbf\xbf\x1e\x4c\x04\x48\x53\xfa\x80\x1f\xf5\x28\x56\x70\xd2\xa3\xeb\x5f\x4f\x1d\x10\x21\x98\xe6\x6d\xff\xf2\xf0\x73\x92\x63\x3a\xe7\x92\x43\xc8\x83\xf3\x6a\x37\x8d\x7d\x58\x5f\x72\xc8\xf0\xe5\x53\x0c\x7d\x3c\xc3\xea\x08\x20\x1f\x28\xa2\x93\x41\x24\x08\xb7\xbb\x8d\xa8\xf2\xa7\xdd\x8b\x15\x44\x51\xc1\x7d\x87\xc6\xa0\x34\x81\xef\xe7\x55\x96\x28\x37\xa9\x50\xf1\x20\x92\x97\xcc\x7d\x6b\xc7\x90\xc3\x97\xe7\x3c\xa4\x41\x24\xb6\x80\xe5\xfe\x65\x93\x25\xda\x73\x68\x58\x0f\x3a\xb9\x57\x01\x4f\x9a\x04\x10\xcc\xbc\x36\xb7\x62\x93\x5a\xac\xfa\xb4\x2f\x9b\x4d\x02\x54\xce\x55\x3f\xa5\x8f\x2a\xd9\x95\xc8\xa5\xb1\x4a\xde\x16\x72\x09\x46\xaa\x24\x77\xb7\xc5\x08\xe9\x4c\xa6\xac\x1f\x30\x8e\x64\xff\x2a\x47\xb7\xdf\xc2\x7a\xd3\x18\xfe\xda\x98\x7f\x94\xfb\x6f\x7f\x31\x5d\x5d\x19\x17\xf2\xe9\x26\x7f\xdd\x98\x95\x66\x7e\xda\xf7\xd5\xdd\xc6\x50\xf2\x68\xe6\xca\x95\x97\xfa\xcf\x8b\x91\xd4\x5c\x8a\xe2\xe1\xea\x7e\xd7\x98\xda\x55\xc7\x21\x69\x9a\xd5\x8d\xdc\x9f\x78\xae\x0e\xdd\x45\xc0\x5c\x29\x37\xa0\x30\x96\xb7\x77\x99\x6d\xa3\x34\x1b\x31\x8d\xd4\xa1\x82\xa0\x3c\x96\xd8\x9d\x0b\x93\x22\xc4\xaa\xff\xc6\x60\x24\xe3\x16\xe6\xa6\x91\x8f\x38\x56\xfd\x7b\x2f\x2b\xeb\xfd\xf5\x5e\xd5\x72\x7f\x17\x64\x96\x7b\x98\x1d\x0f\x8a\x85\x22\x49\x83\xe4\x5a\x96\x78\x18\x24\xcd\xa6\x29\xbf\x89\x18\x62\x1e\x7a\x64\x11\x66\xcb\xfd\xd3\xeb\xfb\xec\x31\xff\xb5\xbf\x9e\x7e\x7c\x94\x3a\xc4\x5c\xe6\x71\x44\xe2\xec\x6e\xff\xf4\xed\xf9\xed\x02\x61\xf6\x76\xfc\xef\xd0\x21\x7f\x02\x00\x00\xff\xff\x7d\x1b\x96\x09\xab\x0e\x00\x00")

func configDtCsvBytes() ([]byte, error) {
	return bindataRead(
		_configDtCsv,
		"config/dt.csv",
	)
}

func configDtCsv() (*asset, error) {
	bytes, err := configDtCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/dt.csv", size: 3755, mode: os.FileMode(438), modTime: time.Unix(1450533414, 0)}
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
	"config/conf.ini": configConfIni,
	"config/dt.csv":   configDtCsv,
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
	"config": &bintree{nil, map[string]*bintree{
		"conf.ini": &bintree{configConfIni, map[string]*bintree{}},
		"dt.csv":   &bintree{configDtCsv, map[string]*bintree{}},
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
