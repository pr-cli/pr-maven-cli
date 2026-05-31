package prmaven

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestAnalyzeDoesNotRequireProviderEnvironment(t *testing.T) {
	for _, name := range []string{
		"GITHUB_TOKEN",
		"GH_TOKEN",
		"GITHUB_ACTIONS",
		"GITLAB_TOKEN",
		"GITLAB_PRIVATE_TOKEN",
		"CI_JOB_TOKEN",
	} {
		t.Setenv(name, "")
	}

	report, err := Analyze(Options{ProjectDir: "../../demo/multi-module-failure"})
	if err != nil {
		t.Fatal(err)
	}
	if report.Summary.FindingCount != 2 {
		t.Fatalf("finding count = %d, want 2", report.Summary.FindingCount)
	}
}

func TestCorePackageDoesNotImportNetworkOrProviderClients(t *testing.T) {
	forbiddenImports := map[string]bool{
		"net":                                 true,
		"net/http":                            true,
		"net/rpc":                             true,
		"golang.org/x/oauth2":                 true,
		"github.com/google/go-github":         true,
		"github.com/xanzy/go-gitlab":          true,
		"gitlab.com/gitlab-org/api/client-go": true,
	}

	forEachProductionFile(t, func(path string, file *ast.File) {
		for _, spec := range file.Imports {
			importPath, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				t.Fatalf("unquote import in %s: %v", path, err)
			}
			if forbiddenImports[importPath] {
				t.Fatalf("%s imports %q; core analysis must stay provider-agnostic and no-network", path, importPath)
			}
		}
	})
}

func TestCorePackageDoesNotReadProviderEnvironment(t *testing.T) {
	forbiddenSelectors := map[string]bool{
		"Getenv":    true,
		"LookupEnv": true,
		"Environ":   true,
		"ExpandEnv": true,
	}

	forEachProductionFile(t, func(path string, file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			selector, ok := node.(*ast.SelectorExpr)
			if !ok || !forbiddenSelectors[selector.Sel.Name] {
				return true
			}
			ident, ok := selector.X.(*ast.Ident)
			if ok && ident.Name == "os" {
				t.Fatalf("%s reads environment through os.%s; core analysis must not require provider tokens", path, selector.Sel.Name)
			}
			return true
		})
	})
}

func forEachProductionFile(t *testing.T, visit func(path string, file *ast.File)) {
	t.Helper()

	fileSet := token.NewFileSet()
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(fileSet, path, nil, 0)
		if err != nil {
			return err
		}
		visit(path, file)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
