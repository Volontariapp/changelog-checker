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

	headerRe := regexp.MustCompile(`(?m)^## \[([^\]]+)\]`)
	matches := headerRe.FindStringSubmatch(changelog)
	if len(matches) < 2 {
		return &VerificationError{Message: "no version header found in CHANGELOG.md"}
	}

	latestChangelogVersion := matches[1]
	if latestChangelogVersion != version {
		return &VerificationError{
			Message: fmt.Sprintf("changelog latest version (%s) does not match package.json version (%s)", latestChangelogVersion, version),
		}
	}

	linkPattern := fmt.Sprintf(`(?m)^\[%s\]:\s*https?://github\.com/Volontariapp/[A-Za-z0-9_\-\.]+/pull/(PLACEHOLDER|\d+)`, regexp.QuoteMeta(version))
	linkRe := regexp.MustCompile(linkPattern)
	
	if !linkRe.MatchString(changelog) {
		return &VerificationError{
			Message: fmt.Sprintf("no valid PR link found for version [%s] in the footer. Expected format: [%s]: https://github.com/Volontariapp/repo/pull/PLACEHOLDER (or a number)", version, version),
		}
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

	fmt.Println("✅ Changelog is compliant with package.json!")
}
