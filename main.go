package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type PackageJSON struct {
	Version string `json:"version"`
}

type VerificationError struct {
	Message string
}

func (e *VerificationError) Error() string {
	return e.Message
}

func VerifyChangelog(packageJsonPath string, changelogPath string) error {
	pkgData, err := os.ReadFile(packageJsonPath)
	if err != nil {
		return fmt.Errorf("could not read package.json: %w", err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(pkgData, &pkg); err != nil {
		return fmt.Errorf("could not parse package.json: %w", err)
	}

	version := pkg.Version
	if version == "" {
		return &VerificationError{Message: "version not found in package.json"}
	}

	changelogData, err := os.ReadFile(changelogPath)
	if err != nil {
		return fmt.Errorf("could not read CHANGELOG.md: %w", err)
	}
	changelog := string(changelogData)

	titleRe := regexp.MustCompile(`(?m)^# Changelog\s*$`)
	if !titleRe.MatchString(changelog) {
		return &VerificationError{Message: "missing '# Changelog' title"}
	}

	versionRe := regexp.MustCompile(`(?m)^## (\d+\.\d+\.\d+)\s*$`)
	matches := versionRe.FindStringSubmatch(changelog)
	if len(matches) < 2 {
		return &VerificationError{Message: "missing version header in format '## x.y.z'"}
	}

	latestChangelogVersion := matches[1]
	if latestChangelogVersion != version {
		return &VerificationError{
			Message: fmt.Sprintf("changelog latest version (%s) does not match package.json version (%s)", latestChangelogVersion, version),
		}
	}

	allowedChangesSectionRe := regexp.MustCompile(`(?m)^### (Patch|Minor|Major) Changes\s*$`)
	if !allowedChangesSectionRe.MatchString(changelog) {
		return &VerificationError{
			Message: "missing changes section: expected one of '### Patch Changes', '### Minor Changes', or '### Major Changes'",
		}
	}

	listItemRe := regexp.MustCompile(`(?m)^-\s+.+`)
	if !listItemRe.MatchString(changelog) {
		return &VerificationError{Message: "missing changelog entries (bullet list items)"}
	}

	return nil
}

func main() {
	pkgPath := "package.json"
	changelogPath := "CHANGELOG.md"

	if len(os.Args) > 1 {
		pkgPath = os.Args[1]
	}
	if len(os.Args) > 2 {
		changelogPath = os.Args[2]
	}

	err := VerifyChangelog(pkgPath, changelogPath)
	if err != nil {
		if verreq, ok := err.(*VerificationError); ok {
			fmt.Printf("❌ CHANGELOG DOES NOT COMPLY:\n   %s\n", verreq.Message)
		} else {
			fmt.Printf("❌ ERROR: %s\n", err)
		}
		os.Exit(1)
	}

	fmt.Println("✅ Changelog is compliant with package.json version and format!")
}
