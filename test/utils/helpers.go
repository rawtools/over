package testutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/otiai10/copy"
	"github.com/stretchr/testify/require"

	over "raw.tools/over/pkg"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	TestRoot          = filepath.Dir(filepath.Dir(filename))
	TestData          = filepath.Join(TestRoot, "data")
)

func TestConfig(home string) *over.Config {
	return &over.Config{
		Home: home,
	}
}

func Data(path string) string {
	return filepath.Join(TestData, path)
}

func ForEachOverlay(t *testing.T, path string, test func(*testing.T, *over.Overlay)) {
	root := Data(path)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			name := file.Name()
			t.Run(name, func(t *testing.T) {
				defer func() {
					if err := recover(); err != nil {
						t.Error(err)
					}
				}()

				overlay := over.NewOverlay(name, filepath.Join(root, name))
				test(t, overlay)

			})
		}
	}
}

func WithTempDir(t *testing.T, name string, callback func(path string)) {
	require := require.New(t)
	tmp, err := ioutil.TempDir("", strings.ReplaceAll(name, string(os.PathSeparator), "_"))
	require.NoError(err)
	defer os.RemoveAll(tmp) // clean up
	callback(tmp)
}

func WithTempCopy(t *testing.T, src string, callback func(path string)) {
	WithTempDir(t, src, func(tmp string) {
		err := copy.Copy(src, tmp)
		if err != nil {
			t.Fatal(err)
		}
		callback(tmp)
	})
}
