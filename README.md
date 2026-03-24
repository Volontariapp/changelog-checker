# Changelog Checker

A lightweight, fast CLI tool written in Go to ensure that `CHANGELOG.md` files comply with the Volontariapp standards.

## 🎯 Purpose

This tool verifies release changelog compliance before a merge:
1. **Version Match:** The first changelog version header must exactly match the `version` field from `package.json`.
2. **Changelog Structure:** `CHANGELOG.md` must contain:
   - `# Changelog`
   - a version header in the format `## x.y.z`
   - `### Patch Changes`
   - at least one bullet item under the section (line starting with `- `)

It does **not** validate a footer PR link anymore.

Expected example:

```markdown
# Changelog

## 0.7.1

### Patch Changes

- [#25](https://github.com/Volontariapp/npm-packages/pull/25) [`fb3b3cf`](https://github.com/Volontariapp/npm-packages/commit/fb3b3cf44270654a85ff24517585f888bb78ea48) Thanks [@Fyorix](https://github.com/Fyorix)! - added new test snapshot for package contract

- test patch contracts
```

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
- `1`: Validation failed (mismatch, invalid format, or missing files).

## 🛠 Development

### Prerequisites
- Go 1.22+

### Run Tests
A full test suite is included to cover edge cases (valid files, missing sections, missing list entries, and version mismatches).
```bash
go test -v ./...
```

### Build Binary
```bash
go build -o changelog-checker main.go
```
