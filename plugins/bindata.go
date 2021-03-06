// Code generated by go-bindata.
// sources:
// plugins/plugins.toml
// DO NOT EDIT!

package plugins

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

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x97\xdf\x73\xda\x38\x10\xc7\xdf\xf9\x2b\x76\xc8\xcb\xdd\x4d\x09\xd7\xde\xb4\x0f\x9d\xb9\x07\x8e\x90\x36\xb9\xfc\x60\x30\xcd\x4c\x27\x93\x49\x17\x5b\x36\xba\xd8\x96\x46\x92\x69\xf8\xef\x4f\x2b\x7e\x38\x80\xd2\x5a\xd4\xbc\x18\x76\xa5\xdd\xfd\x7c\x77\x2d\x9b\x13\x18\x0a\xb9\x54\x3c\x9b\x1b\xf8\x2d\xfe\x1d\xde\xfd\xf9\xf6\x2f\xe8\xd1\xe5\x03\xcc\x72\x8c\x9f\x8c\x90\x70\x29\xf4\xbc\x42\xb8\x46\x5e\xb2\x37\x30\xc8\x73\x98\xd0\x06\x0d\x13\xa6\x99\x5a\xb0\xe4\xb4\x73\x02\x11\x63\x70\x75\x31\x1c\xdd\x44\x23\x48\x85\x82\x9c\xc7\xac\xd4\x0c\x78\x69\x7f\x15\x68\xb8\x28\x4f\x3b\x9d\x93\x76\x3e\x36\xdf\xf8\xea\xcb\xa7\x8b\x1b\x5b\x7e\x99\xf2\xac\x52\x2e\x01\x84\xc7\x69\xa9\x9e\x8e\xe1\x26\x67\xf0\x37\x74\xaf\x91\xc8\x61\x9c\x57\x19\x2f\x77\xcb\xd3\xdd\x4e\xe7\xfe\x5e\x3a\xcf\xc3\x43\x07\x80\x95\x38\xcb\x59\x62\xb7\xa5\x98\x6b\x66\x2d\x25\x16\x2e\x4a\xa9\x55\xde\xb5\xbf\x13\xa6\x63\xc5\xa5\x83\xb3\xe6\x9b\x68\x72\x05\x67\x68\x70\x86\x56\xda\xcf\xa8\xe7\x56\x77\x54\xf1\x9c\xd6\xc6\x68\x58\x26\xd4\x92\x16\xf2\xd2\x30\x17\x80\x17\x98\xb9\x88\x85\xab\xab\xbf\x09\xac\x98\x14\x9a\x9b\xf5\xf2\xb9\x31\x52\x7f\xec\xf7\x33\x6e\xe6\xd5\xec\x34\x16\x45\x7f\xb5\x9e\x8b\xf5\x97\x1e\x6d\x3c\xb5\x7e\xda\x3c\xab\x78\xbe\xad\x1a\xec\xc7\xda\x0a\xbe\xaa\x7c\x6e\x8b\xa2\x35\x74\x35\x4b\xc9\xb4\x35\xde\xdb\xf4\xc9\xfb\x2e\x3c\xbc\xca\x6f\x54\xf5\x12\x7f\xc1\x55\xa5\x8d\x30\xe8\x13\xe1\x8e\x9c\x53\x72\xda\x49\x4d\x79\x6e\x53\xe8\x18\x4b\xc0\x32\x71\x59\x21\x17\xe2\xa9\x92\x8d\x25\xd9\x4d\x16\x2c\x4c\xbd\xdd\x2f\x8f\x35\x54\x32\x17\x58\x5b\xe8\x66\x31\x1b\xa3\x11\xf0\x8d\xe8\xbf\x01\x4f\x61\x29\x2a\xf8\x8e\xa5\x21\xeb\xda\xaf\xb1\x90\x76\xb2\xac\xe1\x45\x22\x5b\x87\x0d\x8b\x92\x3f\x31\x57\xa6\x63\x2d\x28\x43\xb7\x86\x6f\xd0\x91\x37\xd0\xd5\x73\x7c\xbb\xbe\xbe\x7b\xff\x81\x7a\x44\x7d\x59\xd0\x9a\xee\xf5\x80\xee\xe8\xc7\xbb\xe9\xe3\x60\x7c\xd1\x6d\xdc\x3d\x1b\x2a\x11\xdf\x7b\xee\x6c\x50\x9e\x06\x46\xce\x1f\x39\xb7\x6d\xe1\x31\x3d\x3b\x48\x11\xdc\xb6\x9d\x08\xaf\x0c\x76\x98\x88\x3f\x18\xef\xfd\xdb\xdb\x30\x2c\x7a\xf1\xb2\x50\x95\x47\x9e\xa9\x75\x0e\xc9\x77\xa4\x36\xbb\xc1\x83\x85\xa9\xb7\xbf\x76\xbb\x07\xcb\x72\x30\x51\xd3\xcf\x8f\x5f\xa2\xd1\x84\xd6\xd4\x96\x7f\x47\x5f\x9b\xcf\x18\xdd\xf7\xf4\x68\xf1\xe8\xa7\x75\xc2\x98\xec\x4f\x27\xfc\xac\xcf\x9e\x79\x6a\x84\xc8\xf7\xc5\x2b\x98\xc1\xc4\x9e\xa3\x3e\xfd\x5e\x86\x0e\x56\x6f\xb3\xf9\xe7\x13\xf5\xc7\xeb\x4f\x83\x3d\xd6\x25\x2a\xf4\x70\x7e\x1d\x4c\x06\x10\xd9\x93\x6f\x1f\x0e\x17\x3e\xac\x4d\x94\x60\x24\xda\xf8\x4b\x38\xfb\xd3\x8f\x0b\xd4\xc6\x03\x34\x20\x3b\x0c\x4a\xc3\xdd\x11\xdf\x0c\x6b\x1b\x2c\x98\x0b\x17\x44\x75\xf2\x71\x1b\xa1\x8d\x5e\xe1\x22\xf3\x91\xdd\x7d\x0a\xe7\xca\x7e\x91\x2a\x6b\x8b\x69\xc6\x4d\xc2\x52\x56\x26\xde\xd3\xfc\x9f\xda\x1b\xca\xb8\x17\xf8\x68\xd6\xbd\x38\x6d\x30\xc7\x39\x16\xab\xa2\xf7\x70\x87\xd6\x31\xb8\x6b\xc6\x57\x07\x39\x1a\xad\x0e\xd1\x0a\x95\x28\x44\xe2\x3b\x33\x87\xce\x11\xda\xbf\x3a\xdc\xf1\x7c\xdb\x10\x6d\xf0\xa5\x52\x09\xdf\xc9\x72\xde\x1b\x4f\x6e\xa7\xa1\x78\xdb\x68\x47\xd3\x6d\x23\x6c\xe0\xd6\xf5\x1e\x71\x68\x4a\xe6\xe1\x1a\x8f\xec\x3b\x02\x3d\xdf\xe8\xfd\xd0\x28\x4e\xf5\x4b\xa1\x0c\x05\x01\xf6\xcc\xe2\xca\x7d\x3d\xe0\xb5\x2e\x1f\xf0\x2a\x47\x30\xad\x64\x3f\x7f\x3c\xa0\x94\x76\xad\xfb\x1b\xd4\x7f\xee\x25\x42\x53\x75\xcd\xfb\x9a\x0b\xad\x7d\x7d\xe5\x8a\x8d\x96\xf6\x8f\x27\xce\x34\xdc\xce\xd2\x4a\x13\x65\x02\x91\xd5\xa2\xcc\x20\x12\xf9\xfa\xf5\xb0\x01\xfc\x36\x47\xf8\x33\x9f\x76\xb6\x2e\xc1\xfe\x00\x88\x34\xb5\xc9\x3c\x22\xdc\x3a\xc7\xe1\x20\xdc\x5e\x8d\xfa\x93\xe9\x39\x24\x22\xae\x0a\x56\x9a\x83\x29\xd8\x38\x7c\x6a\xd4\xd9\x82\xe5\x58\x6d\x6d\xf5\x8d\x41\x26\xa9\x6f\xfa\xcf\xce\x0f\xa9\xc9\x78\x14\xf1\x3a\x47\xf8\xf4\x27\x69\x58\xef\x5d\xa2\xa6\xe4\xff\xd1\xdb\x89\x83\xf6\x08\x70\xb9\x75\x1e\xea\x70\x19\xc1\xca\x15\x24\xc2\x6e\xba\x60\x2d\xea\xed\x61\x92\xbc\x4c\xfb\x7f\x00\x00\x00\xff\xff\xbd\x41\xcd\x1c\x93\x12\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 4755, mode: os.FileMode(420), modTime: time.Unix(1471840236, 0)}
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
	"plugins/plugins.toml": pluginsPluginsToml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
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

