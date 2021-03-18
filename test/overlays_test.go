package test

import (
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	over "raw.tools/over/pkg"
	tu "raw.tools/over/test/utils"
)

func TestOverlays(t *testing.T) {
	tu.ForEachOverlay(t, "overlays", func(t *testing.T, overlay *over.Overlay) {
		name := overlay.Name
		tu.WithTempDir(t, name, func(target string) {
			require := require.New(t)
			assert := assert.New(t)
			err := overlay.Apply(target)
			require.NoError(err)

			expectations := filepath.Join(tu.TestData, "expected", name)
			// Ensure all files are present in the right state
			require.NoError(filepath.Walk(expectations, func(path string, info fs.FileInfo, err error) error {
				require.NoErrorf(err, "missing expectations for %s", name)

				base, err := filepath.Rel(expectations, path)
				require.NoError(err)
				fullTarget := filepath.Join(target, base)

				if info.IsDir() {
					assert.DirExists(fullTarget)
				} else {
					assert.FileExists(fullTarget)
				}
				return nil
			}))
			// Ensure there is not extra file
			require.NoError(filepath.Walk(target, func(path string, info fs.FileInfo, err error) error {
				require.NoErrorf(err, "missing expectations for %s", name)

				base, err := filepath.Rel(target, path)
				require.NoError(err)
				src := filepath.Join(target, base)

				if info.IsDir() {
					assert.DirExists(src)
				} else {
					assert.FileExists(src)
				}
				return nil
			}))
		})
	})
}

func TestOverlayErrors(t *testing.T) {
	tu.ForEachOverlay(t, "errors", func(t *testing.T, overlay *over.Overlay) {
		name := overlay.Name
		tpl := filepath.Join(tu.TestData, "targets", name)
		tu.WithTempCopy(t, tpl, func(target string) {
			require := require.New(t)

			require.Error(overlay.Apply(target))
		})
	})
}

// func TestExpressions(t *testing.T) {
// 	root := "./testData/expressions"
// 	env := tu.TestEnv(root)
// 	tu.GlobTemplateTests(t, root, env)
// }

// func TestFilters(t *testing.T) {
// 	root := "./testData/filters"
// 	env := tu.TestEnv(root)
// 	tu.GlobTemplateTests(t, root, env)
// }

// func TestFunctions(t *testing.T) {
// 	root := "./testData/functions"
// 	env := tu.TestEnv(root)
// 	tu.GlobTemplateTests(t, root, env)
// }

// func TestTests(t *testing.T) {
// 	root := "./testData/tests"
// 	env := tu.TestEnv(root)
// 	tu.GlobTemplateTests(t, root, env)
// }

// func TestStatements(t *testing.T) {
// 	root := "./testData/statements"
// 	env := tu.TestEnv(root)
// 	tu.GlobTemplateTests(t, root, env)
// }

// // func TestCompilationErrors(t *testing.T) {
// // 	tu.GlobErrorTests(t, "./testData/errors/compilation")
// // }

// // func TestExecutionErrors(t *testing.T) {
// // 	tu.GlobErrorTests(t, "./testData/errors/execution")
// // }
