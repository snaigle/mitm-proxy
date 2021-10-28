// +build debug

package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// resources_views_admin_index_gohtml reads file data from disk. It returns an error on failure.
func resources_views_admin_index_gohtml() ([]byte, error) {
	return bindata_read(
		"/Users/snaigle/workspace/mework/goproxy2/resources/views/admin/index.gohtml",
		"resources/views/admin/index.gohtml",
	)
}

// resources_views_common_footer_gohtml reads file data from disk. It returns an error on failure.
func resources_views_common_footer_gohtml() ([]byte, error) {
	return bindata_read(
		"/Users/snaigle/workspace/mework/goproxy2/resources/views/common/footer.gohtml",
		"resources/views/common/footer.gohtml",
	)
}

// resources_views_common_header_gohtml reads file data from disk. It returns an error on failure.
func resources_views_common_header_gohtml() ([]byte, error) {
	return bindata_read(
		"/Users/snaigle/workspace/mework/goproxy2/resources/views/common/header.gohtml",
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
