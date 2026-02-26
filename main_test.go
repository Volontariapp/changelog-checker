package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVerifyChangelog_Valid(t *testing.T) {
	tmpDir := t.TempDir()
	
	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "1.0.1"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## [1.0.1] - 2026-02-24
### Fixed
- A bug.

[1.0.1]: https://github.com/Volontariapp/ms-user/pull/123
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifyChangelog_ValidPlaceholder(t *testing.T) {
	tmpDir := t.TempDir()
	
	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "0.1.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## [0.1.0] - 2026-02-24

[0.1.0]: https://github.com/Volontariapp/npm-packages/pull/PLACEHOLDER
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifyChangelog_VersionMismatch(t *testing.T) {
	tmpDir := t.TempDir()
	
	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "2.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## [1.0.0] - 2023-01-01
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to mismatch")
	}

	if !strings.Contains(err.Error(), "does not match package.json version") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVerifyChangelog_MissingPRLink(t *testing.T) {
	tmpDir := t.TempDir()
	
	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "1.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## [1.0.0] - 2023-01-01
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to missing link")
	}

	if !strings.Contains(err.Error(), "no valid PR link found") {
		t.Errorf("unexpected error message: %v", err)
	}
}
