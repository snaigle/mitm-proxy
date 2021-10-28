// +build release

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
)

func bindata_read(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resources_views_admin_index_gohtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5b\x6f\x13\x49\x16\x7e\xe7\x57\x1c\x6a\x91\x62\xb3\x49\x9b\x24\x8b\x36\xf2\xda\xd6\x4a\xb0\xab\xe5\x65\x89\x04\x1a\x69\x1e\x3b\x5d\xe5\xb8\x27\xed\xee\x56\x77\x39\x01\x99\x48\xc0\x04\x26\x09\x0e\xb1\x20\xc3\x25\x78\x94\x30\x13\xa2\x0c\x33\xc4\x8c\x60\xc0\x93\x0b\xf9\x33\x5d\xdd\xe6\x29\x7f\x61\x54\x5d\xbe\xc7\x6e\xb7\xc9\xd4\x93\xab\xea\x7c\xe7\x56\xe7\x7c\x55\xed\x7c\x1e\x30\x49\xab\x3a\x01\x24\xe3\xac\xaa\xc7\x54\x1d\x93\x1b\x08\xe6\xe7\xcf\x00\x00\x24\xce\x5e\xbe\x7a\xe9\xfa\xd7\x93\xff\x81\x0c\xcd\x6a\x29\xb1\xc6\x7f\x82\x26\xeb\xd3\x49\x44\x74\x24\x16\xc5\x06\x91\x71\x73\xca\x47\x3e\x4f\x49\xd6\xd4\x64\x4a\x00\xf1\x5d\x62\x21\x40\xde\xee\x4b\xaf\x78\x9f\x15\x1f\xb2\xd5\xb7\xa8\x66\xc8\xc7\xc7\xda\x15\x24\xa6\x0c\x7c\xb3\x5d\x5f\xc2\x26\x0a\x55\x0d\x1d\x14\x4d\xb6\xed\x24\xaa\x4d\x11\xa8\x38\x89\x64\xd3\x44\xed\xe2\x3e\x04\xab\xb3\x75\x71\xc5\xd0\xa9\xac\xea\xc4\xea\x22\x27\x22\x18\x4f\x39\x87\x47\xde\xda\x4e\xb5\xfc\xde\x7d\xf6\x30\x11\xcb\x8c\xf7\x90\x34\x2d\x22\x76\x84\x64\xdc\xff\x3d\x2a\x01\xdb\xfd\xd1\x5d\xfc\xe8\xfd\x52\xbe\x32\x09\xee\xb3\x72\x75\xfb\x8e\xbb\xfc\x98\x15\xee\x79\xeb\x0b\x6e\x69\x8f\x3d\xdf\x89\xb0\xed\xbb\xee\xd2\x03\xb7\xb4\x17\x33\x15\xef\xcd\x52\xd4\x5b\x5f\xb8\x32\xc9\x4a\x6f\xd9\x0f\xb7\x8f\x0f\x0a\xce\xfe\x43\x56\xac\x89\x3a\x95\x65\x77\xe3\x95\x57\x7a\xe0\x2b\x1f\x93\xc0\xd9\xff\xc9\x2b\xde\x17\x9a\x85\xda\xea\x61\xd9\xd9\xff\xe0\xab\x38\x3e\x28\xb0\xed\xbb\x4a\x46\xb6\x34\x62\xbb\x8b\x4f\xd2\x2a\xc6\x1a\xb1\xdc\xd5\xa2\xf3\xe9\x85\xb7\xbe\x20\xc0\xc2\x90\xaf\xd0\x7d\xbc\xe2\x1c\x96\xdc\x37\xaf\x3e\xbf\xde\x6a\xfa\x5f\x2c\xb3\xe5\x9d\x9a\x16\xef\xcd\x52\xab\x99\xe3\x83\x42\x75\xf7\xad\x73\xb8\x5a\xf7\xe3\xf8\xa0\xe0\x6e\x1e\xb0\x83\x55\xb6\x7d\x17\x46\x2f\x48\x63\xe3\x13\xd2\xd8\xb8\x34\x16\x9f\x98\x98\x98\xa8\x3b\x2d\x42\x68\x4d\x44\x23\x47\x2d\x7e\x8b\x94\xb8\x8b\x4f\x26\x2f\xf9\xa9\xfa\xd5\x2d\xed\xb5\x1b\x38\x3e\x58\x6f\x9a\xb8\xe8\x6b\x1f\x97\x80\x95\x76\x5a\xeb\xc9\xfd\xb8\xcf\x96\x37\x5b\x0e\x81\x3d\x2a\xd4\xbd\xf5\x21\xff\xf0\x21\x27\x8f\xc5\x5b\x5f\xa8\xee\x7e\xf2\x0e\x77\x2f\xff\xff\x5a\xa4\x66\x67\xf4\x9f\xd2\xe8\x78\xd4\x5d\x7c\xe2\x1c\xed\xba\x6b\x7f\xfc\xef\xea\xb5\xeb\x6d\x5b\xa0\x1b\x94\x48\x37\x8d\x1c\x96\x0d\x49\x31\xb2\x51\xdf\xc2\x45\x09\xaa\x47\xc5\xea\xcb\x42\xab\x11\xd9\x34\xd9\xca\x3b\xb6\x5a\x16\x99\x7f\xb7\xe3\x2e\xac\x36\x72\xde\x9a\x63\xf6\xa8\x13\xe7\x54\x56\xd8\xde\x1a\x2b\xed\xb0\x62\xc1\xa9\xdc\x76\x2a\xaf\x1b\xd5\x21\x52\xc7\xe5\x7f\x7f\x50\x2d\x7f\x6f\x2a\x4e\x65\xd9\x5b\x5f\x20\x1a\x51\xa8\x65\xe8\x8d\x34\xb8\x4f\x3f\xd4\x54\xac\x6e\x3b\x95\xdb\x3c\x4f\x75\x15\xc2\xd5\x56\x0f\xea\xe7\xe6\xbd\xd8\x65\x4b\x2b\xce\xfe\x3d\x77\xad\xec\x16\xee\x74\x06\x0b\xc3\x70\x5e\xd2\x6e\x48\x3a\xa1\x44\xb6\x89\xbf\x04\xe7\x25\xc5\x98\x33\xac\x99\xb6\xd5\xea\xd1\x73\xf6\xdd\xa6\x53\x79\xcd\x36\x36\x58\x71\xa5\x7b\x3f\xc5\x1a\x0d\xd5\xbe\x8c\xd5\xd9\xd3\xf7\x34\x5b\x7c\x5a\x7d\xb9\x13\xd0\xcd\x5d\xad\x34\x76\xa9\x3c\xa5\x91\xba\x3d\x31\x51\xed\x91\x8c\x31\x4b\x2c\x3e\xe9\x61\xb8\x09\x3f\xc9\x8b\xdd\xe5\xac\xfe\x42\x35\x85\xa9\x96\x0a\x4f\xc4\x68\x26\x3c\xb0\xde\x0c\x83\xa1\x04\x55\x84\xc3\x24\x62\xfd\x02\xe1\x7a\xfa\xa6\x24\x41\x4f\xd2\x7f\xb7\x91\xcf\x5b\xb2\x3e\x4d\xe0\xdc\xcc\xf0\xb9\x59\x88\x27\x41\x52\x29\xc9\xda\x2d\xf7\x4a\xb0\x95\x90\x49\x17\xc2\x38\x95\xcf\x9f\x9b\x99\x9f\x4f\xc4\x68\x88\x13\xed\xc0\xcd\x7e\x09\x2e\xb4\xb0\x0f\x98\xca\x51\xda\xbc\x1d\x6b\x33\xd5\x1e\xc1\x3c\x45\x16\x82\x7f\x2b\x9a\xaa\xcc\x24\x11\x26\x5a\x64\x48\x84\x32\x14\x45\x29\xb6\xb8\xf9\xf9\xf9\x56\x22\x26\x00\x03\x38\x18\x3a\x9c\xfe\x45\x01\xfe\x59\x12\x1d\x87\x38\xb9\x01\x5a\x05\xa7\x12\xaa\x6e\xe6\x68\x3d\x27\xfe\x04\xc1\xec\x48\xd6\xc0\x44\x4b\xa2\xb4\x65\x64\x11\x98\x9a\xac\x90\x8c\xa1\x61\x62\x25\x51\x4b\x6f\xa1\x58\x6a\x80\x18\xfb\x1a\xa3\x46\x87\xa9\x7a\x37\x0e\x6c\x27\xfc\x11\xf5\x2a\x09\x4d\xd5\x67\x9a\x05\x61\xe7\xa6\xb2\x2a\x8d\x44\x51\x4a\xdc\xa0\x83\xd5\x42\x38\xdf\x43\x11\x43\x70\xd3\x27\x62\x3e\xfd\xf6\x60\xf1\x1e\x97\x45\xf7\xe5\xd6\xa7\x69\xda\x30\x28\x6f\x0f\x49\xfc\xe8\xa8\xc0\x44\xac\xf6\xcc\x4c\x9d\xe9\x78\x8c\x2a\x96\x6a\xd2\x93\xaa\x67\x65\x0b\x64\xd3\x84\x24\xe8\x64\x0e\xbe\xca\x91\x48\xbe\xab\xc3\x44\x8b\xc3\xd0\xdf\x64\xd3\x1c\x1a\xee\xba\x8f\x65\x2a\xc7\xa1\x3b\x96\x0f\x5e\xbb\x71\x40\xa8\x3b\x9a\x0f\x6a\xf4\xde\x9f\xef\xbe\x9c\x25\x34\x63\x60\x3b\xc8\xae\x28\x96\x38\xa4\x73\xba\x78\x8d\x47\xa2\x01\xd2\x50\xcb\x08\xcd\xc8\x14\x92\x40\x33\xaa\xfd\xaf\xbe\xc2\x69\xc3\xca\x5e\x96\xa9\x5c\xcb\xe1\x7f\x6b\xd3\x48\x34\x10\xa9\xa6\x21\x72\x96\x1b\x90\x78\x6a\xe0\xd6\x2d\x10\x33\x6a\xf4\x73\x90\x0f\x59\x23\x16\x8d\x20\xa7\xb2\x52\xfd\xf6\xd0\xa9\xec\x79\x3f\xef\xa1\x68\xb0\xab\x7c\x58\x84\xe6\x2c\x3d\x58\x2e\x98\xd2\xea\xc1\x4a\xb2\x69\x12\x1d\x47\x04\x29\x0d\x43\x23\x92\xe0\xa8\x4f\xc0\xa9\x51\x07\x53\xa3\x4f\x00\x69\x42\x95\x4c\x04\xc5\xc4\xc7\x9f\x8c\xf1\x88\x6a\xa2\xe1\x10\xc9\x12\x65\x12\x07\x34\x79\xf5\xda\xf5\x80\x0a\xac\x0f\xc5\x22\x98\xe8\x54\x95\x35\x3b\x0e\xc8\x96\xb3\x64\xc4\xb0\xd4\x69\x55\x0f\x81\xe5\x9c\x10\x6f\x84\x19\x9c\xe8\xa8\x44\x33\x44\x8f\x34\x4b\xd3\x22\xb6\x19\xe6\xf4\xc5\x31\x02\x17\x97\xbe\xb1\x0d\xbd\x4f\xa9\x9d\x34\xc4\xbb\x35\x8c\x21\x5e\xa4\x5c\x56\x52\x0c\x4c\x20\x99\x84\x0b\x61\x50\x7c\xf0\x0e\x92\x74\x39\x4b\x20\x09\x08\xf5\xaf\xcc\x06\x26\x67\x69\x83\x42\x84\x73\xa1\x31\xb5\xde\xe1\x1f\x9b\x7b\x5b\xee\x62\x91\x2d\x6f\x84\xe9\x1d\x3e\xe6\x54\x1d\x1b\x73\x92\x66\x28\x32\x4f\xa4\x64\x11\xcd\x90\x71\x84\x5a\x39\x12\x7c\x04\x7c\xcc\x03\xd1\x6c\x32\x48\xfe\x4e\x13\x18\xdb\xfa\xad\xfa\xfe\x55\x1c\xc1\xdf\x7d\x6e\x96\xb2\xf6\x74\x08\x17\xfb\xd4\x51\x6f\x3f\x7a\x30\x34\x1f\x98\xdf\x1d\xcd\xda\x53\x71\x18\x06\xb6\x88\x9d\xd3\x38\x07\x2b\x86\x9e\x56\xad\x6c\x04\xb9\xcf\xca\xac\xb8\x2d\x9e\x7f\x3c\x2a\x15\xf7\x39\x34\x5e\xbc\x42\x4d\x98\xa2\xfd\x72\x26\x87\x00\x5a\x54\x71\x08\x6c\x1b\xaf\x61\xa2\x11\x4a\xc2\x52\x1b\x7c\x01\xbd\xc1\x29\x29\x0e\x06\xa3\x39\x38\x15\xd5\xc1\xe0\x74\xd7\xdd\x60\x58\xca\x83\x53\xd1\x1e\x9c\x96\x24\x60\x50\xa2\x80\xbf\xa4\xf1\xa1\x6f\xf3\x43\x1f\x02\x08\xd6\xd0\x7d\xe7\xe4\xea\x7c\xb4\xf3\x39\xdb\xf1\x6e\x4d\xc4\x9a\xaf\xee\x44\x4c\xfc\xd3\x5b\xfb\x14\xfb\x33\x00\x00\xff\xff\x9a\xbd\xd5\xd8\x20\x16\x00\x00"

func resources_views_admin_index_gohtml() ([]byte, error) {
	return bindata_read(
		_resources_views_admin_index_gohtml,
		"resources/views/admin/index.gohtml",
	)
}

var _resources_views_common_footer_gohtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\xaa\xad\xe5\x52\x50\x50\x50\xb0\x81\x70\xed\x20\x1c\x7d\x18\xaf\xba\x3a\x35\x2f\xa5\xb6\x16\x10\x00\x00\xff\xff\x53\x3a\x8e\x2c\x36\x00\x00\x00"

func resources_views_common_footer_gohtml() ([]byte, error) {
	return bindata_read(
		_resources_views_common_footer_gohtml,
		"resources/views/common/footer.gohtml",
	)
}

var _resources_views_common_header_gohtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x41\x4e\xc3\x30\x10\x45\xf7\x3d\x85\xe5\x35\xf1\x14\x89\x05\x42\x71\x96\x9c\x00\x0e\xe0\xda\xbf\xca\x04\xc7\x89\x3c\x43\x0a\x8a\x7c\x77\xd4\x16\x21\xb1\x9b\xf7\xa4\x79\xfa\xfb\x9e\x70\xe6\x02\x63\x47\x84\x84\x6a\x5b\x3b\xf4\x33\x34\x98\x38\x86\x2a\x50\x6f\xdf\xdf\x5e\xbb\x67\x3b\xfc\xea\x12\x66\x78\xbb\x31\x2e\xeb\x52\xd5\x9a\xb8\x14\x45\x51\x6f\x2f\x9c\x74\xf4\x09\x1b\x47\x74\x37\x78\x30\x5c\x58\x39\xe4\x4e\x62\xc8\xf0\x8f\xd7\x88\xb2\x66\x0c\xfb\xee\x5a\xeb\xe9\x0e\x87\x3e\x73\xf9\x30\x15\xd9\x5b\xd1\xef\x0c\x19\x01\xb5\x66\xac\x38\x7b\x3b\xaa\xae\xf2\x42\x14\x53\x99\xc4\xe5\x25\xb3\x2b\x50\x0a\x53\xf8\xa2\xcc\x27\x21\xbd\xb0\x2a\x6a\x77\x5a\x16\x15\xad\x61\xa5\x27\x77\x74\x47\x8a\x22\xf4\xe7\xdc\xcc\xc5\x45\x91\xeb\x04\x89\x95\x57\x35\x52\xe3\xbf\xba\x9b\x24\x21\xf3\x56\x6f\xfd\xb2\xce\xb4\x7d\x82\x12\x8b\x5e\x0f\x37\x89\x1d\x7a\xba\xff\x0e\x87\x7d\x47\x49\xad\xfd\x04\x00\x00\xff\xff\x52\x21\x59\x79\x3f\x01\x00\x00"

func resources_views_common_header_gohtml() ([]byte, error) {
	return bindata_read(
		_resources_views_common_header_gohtml,
		"resources/views/common/header.gohtml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"resources/views/admin/index.gohtml": resources_views_admin_index_gohtml,
	"resources/views/common/footer.gohtml": resources_views_common_footer_gohtml,
	"resources/views/common/header.gohtml": resources_views_common_header_gohtml,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resources": &_bintree_t{nil, map[string]*_bintree_t{
		"views": &_bintree_t{nil, map[string]*_bintree_t{
			"admin": &_bintree_t{nil, map[string]*_bintree_t{
				"index.gohtml": &_bintree_t{resources_views_admin_index_gohtml, map[string]*_bintree_t{
				}},
			}},
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"footer.gohtml": &_bintree_t{resources_views_common_footer_gohtml, map[string]*_bintree_t{
				}},
				"header.gohtml": &_bintree_t{resources_views_common_header_gohtml, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}