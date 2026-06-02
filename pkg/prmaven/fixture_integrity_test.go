package prmaven

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFixtureIntegrity(t *testing.T) {
	expectedFixtureFiles := []fixtureFile{
		{path: "../../demo/multi-module-failure/pom.xml", description: "multi-module failure demo root POM"},
		{path: "../../demo/multi-module-failure/payment-api/pom.xml", description: "payment-api module POM"},
		{
			path:        "../../demo/multi-module-failure/payment-api/target/failsafe-reports/TEST-dev.prmaven.demo.PaymentApiIT.xml",
			description: "Failsafe demo report",
		},
		{path: "../../demo/multi-module-failure/payment-core/pom.xml", description: "payment-core module POM"},
		{
			path:        "../../demo/multi-module-failure/payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml",
			description: "Surefire demo report",
		},
		{path: "../../demo/no-failure/pom.xml", description: "no-failure demo root POM"},
		{path: "../../demo/no-failure/inventory-core/pom.xml", description: "inventory-core module POM"},
		{
			path:        "../../demo/no-failure/inventory-core/target/surefire-reports/TEST-dev.prmaven.demo.InventoryServiceTest.xml",
			description: "no-failure Surefire report",
		},
		{path: "testdata/checkstyle-project/pom.xml", description: "Checkstyle fixture root POM"},
		{path: "testdata/checkstyle-project/service-core/pom.xml", description: "Checkstyle fixture module POM"},
		{path: "testdata/checkstyle-project/service-core/target/checkstyle-result.xml", description: "Checkstyle report fixture"},
		{path: "testdata/spotbugs-project/pom.xml", description: "SpotBugs fixture root POM"},
		{path: "testdata/spotbugs-project/service-core/pom.xml", description: "SpotBugs fixture module POM"},
		{path: "testdata/spotbugs-project/service-core/target/spotbugsXml.xml", description: "SpotBugs report fixture"},
		{path: "testdata/enforcer-project/pom.xml", description: "Enforcer fixture root POM"},
		{path: "testdata/enforcer-project/service-core/pom.xml", description: "Enforcer fixture module POM"},
		{path: "testdata/enforcer-project/service-core/target/maven-enforcer.log", description: "Enforcer log fixture"},
		{path: "testdata/jacoco-project/pom.xml", description: "JaCoCo fixture root POM"},
		{path: "testdata/jacoco-project/service-core/pom.xml", description: "JaCoCo fixture module POM"},
		{path: "testdata/jacoco-project/service-core/target/jacoco.log", description: "JaCoCo log fixture"},
		{path: "testdata/nested-module-project/pom.xml", description: "nested fixture root POM"},
		{path: "testdata/nested-module-project/platform/pom.xml", description: "nested platform POM"},
		{path: "testdata/nested-module-project/platform/service-core/pom.xml", description: "nested service-core POM"},
		{
			path:        "testdata/nested-module-project/platform/service-core/target/surefire-reports/TEST-dev.prmaven.demo.NestedPaymentTest.xml",
			description: "nested Surefire report",
		},
		{path: "testdata/golden/multi-module-failure.txt", description: "multi-module text golden file"},
		{path: "testdata/golden/no-failure.txt", description: "no-failure text golden file"},
	}

	for _, file := range expectedFixtureFiles {
		assertFixtureFile(t, file.path, file.description)
	}

	expectedTargetFiles := map[string]bool{}
	for _, file := range expectedFixtureFiles {
		if containsTargetSegment(file.path) {
			expectedTargetFiles[fixturePathKey(file.path)] = true
		}
	}

	seenTargetFiles := map[string]bool{}
	for _, root := range []string{"../../demo", "testdata"} {
		walkFixtureRoot(t, root, func(path string, entry fs.DirEntry) {
			if entry.Type()&fs.ModeSymlink != 0 {
				t.Fatalf("fixture %s is a symlink; fixtures must be self-contained committed files", fixturePathKey(path))
			}
			if entry.IsDir() || !containsTargetSegment(path) {
				return
			}

			key := fixturePathKey(path)
			if !expectedTargetFiles[key] {
				t.Fatalf("unexpected committed target artifact %s; add it to TestFixtureIntegrity if it is an intentional sanitized fixture", key)
			}
			seenTargetFiles[key] = true
		})
	}

	for path := range expectedTargetFiles {
		if !seenTargetFiles[path] {
			t.Fatalf("expected committed target fixture %s was not found during fixture scan", path)
		}
	}
}

type fixtureFile struct {
	path        string
	description string
}

func assertFixtureFile(t *testing.T, path, description string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("missing %s at %s: %v", description, fixturePathKey(path), err)
	}
	if !info.Mode().IsRegular() {
		t.Fatalf("%s at %s is not a regular file", description, fixturePathKey(path))
	}
	if info.Size() == 0 {
		t.Fatalf("%s at %s is empty", description, fixturePathKey(path))
	}
}

func walkFixtureRoot(t *testing.T, root string, visit func(path string, entry fs.DirEntry)) {
	t.Helper()

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		visit(path, entry)
		return nil
	})
	if err != nil {
		t.Fatalf("walk fixture root %s: %v", fixturePathKey(root), err)
	}
}

func containsTargetSegment(path string) bool {
	normalized := filepath.ToSlash(filepath.Clean(path))
	return normalized == "target" ||
		strings.HasPrefix(normalized, "target/") ||
		strings.Contains(normalized, "/target/")
}

func fixturePathKey(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}
