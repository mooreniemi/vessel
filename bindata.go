// Code generated by go-bindata.
// sources:
// data/vessel.csv
// data/vessel.yml
// DO NOT EDIT!

package vessel

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

var _dataVesselCsv = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\xd0\x31\xd0\x31\x04\x41\x5e\x2e\x43\x1d\x03\x30\x84\x32\x51\x44\xa1\x4c\x43\x18\x13\x10\x00\x00\xff\xff\xed\x89\x32\xc1\x37\x00\x00\x00")

func dataVesselCsvBytes() ([]byte, error) {
	return bindataRead(
		_dataVesselCsv,
		"data/vessel.csv",
	)
}

func dataVesselCsv() (*asset, error) {
	bytes, err := dataVesselCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vessel.csv", size: 55, mode: os.FileMode(420), modTime: time.Unix(1472437292, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataVesselYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xd3\xcd\x6e\x1a\x31\x10\x07\xf0\x3b\x4f\x31\xe2\xc2\x25\x54\x61\x81\x24\xe2\x52\x45\x6a\xd4\xf6\xd0\x1c\xaa\x5e\x7a\x34\xf6\x90\xb5\xb0\x77\x2c\x7f\x74\xcb\xdb\x77\xbc\x09\xa6\x89\x0d\x12\xd2\x5f\xcb\x9a\xf9\xed\xcc\xec\x72\xb9\x9c\xc9\x5e\xd8\x3d\xfa\xb0\x9b\x01\x2c\xf9\x0b\xa0\x30\xc8\x1d\xcc\x7f\xf5\x08\x6f\x3f\xc2\x28\x8c\x09\xb0\x4f\xfb\xbd\x41\xd0\x03\x04\x43\x23\x58\x8a\x9a\x86\x4f\xf3\xd7\x43\x44\xfe\xcb\x74\x50\x92\x51\xe5\xd2\xf4\xb7\xf9\xb3\x84\xd5\x94\xb4\xda\xc1\xed\x94\xfe\x9e\xc3\xe9\x35\xbc\x2f\x4e\x04\x4a\xf8\x23\x44\x82\x80\x08\x63\x2f\xe2\x0d\x03\x22\x04\xb2\x18\x7b\x3d\xbc\x80\x0e\xa0\xbc\x76\x8e\x73\x85\xb0\x8c\x38\xb5\x14\xb7\x25\x75\xc5\xb3\x3a\x7b\x56\x57\x3d\xdf\xe3\x22\x70\x69\x2e\x78\xca\xcf\xdf\xa3\x47\x38\x51\x82\x03\xa2\xc9\xc1\x33\x74\x78\x49\x08\x21\x6a\x79\x9c\xee\x99\xae\x5a\x4a\xb1\xaf\x5b\xe4\x51\x1c\x99\x7d\xbd\x4d\x39\xad\x4b\xda\x16\x6a\x57\x51\x57\x8d\xb9\x79\x64\xad\x00\x4f\x32\x57\xc9\x53\xd4\x1e\xc4\xa0\xf8\x9a\xf4\x7a\xff\xe6\xd1\x11\x6d\xb3\x70\xf7\x81\x3b\x0a\x6f\x5b\xd4\xae\xa4\x4d\x01\xae\xcf\xc0\xee\x2a\xf0\x37\xf7\xed\x38\xb0\x0d\xe8\x0f\xef\x16\x39\x1c\x20\xea\x21\x00\x1d\xe0\x40\xa4\x6e\x20\x38\x6d\x4c\x96\xab\x14\x22\x08\xe9\x29\x04\x88\xbc\x8f\x07\xc3\x82\xaa\x9d\xf3\xc5\x0f\xe4\x1b\xbe\xf1\x92\x2e\xe6\x2d\xe8\xba\xf0\x36\x67\xde\xfa\x2a\xef\x91\x67\x88\xe8\x20\xf4\xe2\x10\xc1\x27\x86\x29\x1a\x99\x48\xdc\x07\x35\xed\xdf\x38\xcd\x5f\x09\xeb\x1a\x96\xd5\x63\x1b\x71\xe9\xd6\xdd\x87\x6d\xcc\xb0\x6d\x35\xd8\xae\x01\x73\x42\x62\x48\x3a\xe6\xdd\x97\x3e\x59\x67\x50\x01\x0d\xff\xf5\x06\x9e\x89\xb7\xd3\xf0\x3b\x52\xd3\xa4\xa1\x80\xb1\x89\xdb\x96\x74\x5f\x48\x77\x15\x69\x5d\x93\x0c\x25\x05\x7d\xb2\x96\xc7\xc5\x8b\x9d\x2b\xe4\x77\x36\x83\x1c\x8d\x3c\xde\x90\x9c\x33\xa7\xcf\x15\x86\x9b\x39\xf0\x99\xa6\xe6\xd2\xa0\x87\xa2\xb9\xaf\x16\xab\xa1\x09\xe8\x35\x4e\x8b\x14\x46\x1d\x65\x8f\xa1\x31\xa0\x9f\x4f\x5f\x9f\x9e\xdb\x33\xba\x3c\xfc\x43\x55\x6e\x33\xfb\x17\x00\x00\xff\xff\x20\x01\x99\x34\x2f\x05\x00\x00")

func dataVesselYmlBytes() ([]byte, error) {
	return bindataRead(
		_dataVesselYml,
		"data/vessel.yml",
	)
}

func dataVesselYml() (*asset, error) {
	bytes, err := dataVesselYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vessel.yml", size: 1327, mode: os.FileMode(420), modTime: time.Unix(1472437226, 0)}
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
	"data/vessel.csv": dataVesselCsv,
	"data/vessel.yml": dataVesselYml,
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
	"data": &bintree{nil, map[string]*bintree{
		"vessel.csv": &bintree{dataVesselCsv, map[string]*bintree{}},
		"vessel.yml": &bintree{dataVesselYml, map[string]*bintree{}},
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
