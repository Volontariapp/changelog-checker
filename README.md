# Changelog Checker

A lightweight, fast CLI tool written in Go to ensure that `CHANGELOG.md` files comply with the Volontariapp standards.

## 🎯 Purpose

This tool verifies two main things before a release can be merged:
1. **Version Match:** The latest version header in `CHANGELOG.md` (e.g., `## [1.0.0] - ...`) must exactly match the `version` field from `package.json`.
2. **Pull Request Link:** The changelog must contain a valid footer link to the Pull Request that introduces these changes.
   - Format: `[1.0.0]: https://github.com/Volontariapp/<repo>/pull/<number or PLACEHOLDER>`

## 🚀 Usage

You can run the checker by providing the paths to both the `package.json` and the `CHANGELOG.md` files. 

```bash
# Default usage (assumes files are in the current directory)
./changelog-checker

# Specify custom paths
./changelog-checker path/to/package.json path/to/CHANGELOG.md
```

**Exit Codes:**
- `0`: Validation successful.
- `1`: Validation failed (mismatch, missing link, or missing files). 

## 🛠 Development

### Prerequisites
- Go 1.22+

### Run Tests
A full test suite is included to cover various edge cases (valid files, missing PR links, version mismatches, etc.).
```bash
go test -v ./...
```

### Build Binary
```bash
go build -o changelog-checker main.go
```
