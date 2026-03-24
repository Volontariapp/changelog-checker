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

	pkgContent := `{"version": "0.7.1"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 0.7.1

### Patch Changes

- [#25](https://github.com/Volontariapp/npm-packages/pull/25) Thanks [@Fyorix](https://github.com/Fyorix)! - added new test snapshot for package contract

- test patch contracts
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifyChangelog_MissingTitle(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "0.7.1"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `## 0.7.1

### Patch Changes

- test patch contracts
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to missing title")
	}

	if !strings.Contains(err.Error(), "missing '# Changelog' title") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVerifyChangelog_MissingVersionHeader(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "0.7.1"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog

### Patch Changes

- test patch contracts
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to missing version header")
	}

	if !strings.Contains(err.Error(), "missing version header") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVerifyChangelog_MissingChangesSection(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "1.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 1.0.0

- test patch contracts
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to missing changes section")
	}

	if !strings.Contains(err.Error(), "missing changes section") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVerifyChangelog_ValidMinorChangesSection(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "1.2.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 1.2.0

### Minor Changes

- add a small feature
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifyChangelog_ValidMajorChangesSection(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "2.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 2.0.0

### Major Changes

- introduce a breaking API change
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestVerifyChangelog_MissingListItems(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "1.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 1.0.0

### Patch Changes
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to missing list items")
	}

	if !strings.Contains(err.Error(), "missing changelog entries") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestVerifyChangelog_VersionMismatch(t *testing.T) {
	tmpDir := t.TempDir()

	pkgPath := filepath.Join(tmpDir, "package.json")
	changelogPath := filepath.Join(tmpDir, "CHANGELOG.md")

	pkgContent := `{"version": "2.0.0"}`
	os.WriteFile(pkgPath, []byte(pkgContent), 0644)

	changelogContent := `# Changelog
## 1.0.0

### Patch Changes

- test patch contracts
`
	os.WriteFile(changelogPath, []byte(changelogContent), 0644)

	err := VerifyChangelog(pkgPath, changelogPath)
	if err == nil {
		t.Fatal("expected error due to version mismatch")
	}

	if !strings.Contains(err.Error(), "does not match package.json version") {
		t.Errorf("unexpected error message: %v", err)
	}
}
