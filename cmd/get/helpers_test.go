package get

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"testing/fstest"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/gmeghnag/omc/vars"
)

func TestKindGroupNamespacedFromCrds_LazyInitsAliasToCrd(t *testing.T) {
	root := t.TempDir()
	crdsDir := filepath.Join(root, "cluster-scoped-resources", "apiextensions.k8s.io", "customresourcedefinitions")
	if err := os.MkdirAll(crdsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	saved := vars.MustGatherRootPath
	savedMap := vars.AliasToCrd
	t.Cleanup(func() {
		vars.MustGatherRootPath = saved
		vars.AliasToCrd = savedMap
	})
	vars.MustGatherRootPath = root
	vars.AliasToCrd = nil

	// The call will return an error (no CRDs found), which is expected.
	_, _, _, _, _ = kindGroupNamespacedFromCrds("nonexistent")

	if vars.AliasToCrd == nil {
		t.Fatal("expected AliasToCrd to be initialized, got nil")
	}
	// Confirm the type is correct by assigning a value.
	vars.AliasToCrd["test"] = apiextensionsv1.CustomResourceDefinition{}
}

func TestReadDirForResources(t *testing.T) {
	tests := []struct {
		name     string
		in       fstest.MapFS
		expected []string
	}{
		{
			name: "read correct resource files/dirs",
			in: fstest.MapFS{
				"resource-file-1.yaml":            {Data: []byte("abc")},
				"resource.yaml":                   {Data: []byte("abc")},
				"1.yaml":                          {Data: []byte("abc")},
				"resource.with.dot.filename.yaml": {Data: []byte("abc")},
				"resource-directory-name":         {Data: []byte("abc"), Mode: fs.ModeDir},
				"resource.directory.with.dot":     {Data: []byte("abc"), Mode: fs.ModeDir},
			},
			expected: []string{
				"resource-file-1.yaml",
				"resource.yaml",
				"1.yaml",
				"resource.with.dot.filename.yaml",
				"resource-directory-name",
				"resource.directory.with.dot",
			},
		},
		{
			name: "read only resource files/dirs matching the expected name convention",
			in: fstest.MapFS{
				"resource-file-1.yaml":             {Data: []byte("abc")},
				"._faulthy-resource-filename.yaml": {Data: []byte("abc")}, // e.g. AppleDouble encoded Macintosh file
				".resource-filename.yaml.swp":      {Data: []byte("abc")},
				"-resource-filename.yaml":          {Data: []byte("abc")},
			},
			expected: []string{"resource-file-1.yaml"},
		},
		{
			name: "read only resource files/dir with size > 0",
			in: fstest.MapFS{
				"resource-file-1.yaml":         {Data: []byte("abc")},
				"empty-resource-filename.yaml": {},
			},
			expected: []string{"resource-file-1.yaml"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := readDirForResources(tc.in)
			if len(got) != len(tc.expected) {
				t.Errorf("Got: %v \n", got)
				t.Errorf("Want: %v \n", tc.expected)
			} else {
				gotNames := make([]string, 0)
				for _, dir := range got {
					gotNames = append(gotNames, dir.Name())
				}
				sort.Slice(gotNames, func(i, j int) bool {
					return gotNames[i] > gotNames[j]
				})
				sort.Slice(tc.expected, func(i, j int) bool {
					return tc.expected[i] > tc.expected[j]
				})
				if !reflect.DeepEqual(gotNames, tc.expected) {
					t.Error("Got:")
					for _, g := range gotNames {
						t.Error(g)
					}
					t.Error("Want:")
					for _, te := range tc.expected {
						t.Error(te)
					}
				}
			}
		})
	}
}
