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

---

## 🔁 CI/CD & Deployment Workflow

This repository is integrated with the unified `ci-tools` workflows. 

### How it works
1. When changes are pushed to `main`, the `.github/workflows/ci.yml` in this repository triggers.
2. It delegates the execution to the **reusable workflow** located at `Volontariapp/ci-tools/.github/workflows/build-changelog-checker.yml`.
3. The reusable workflow tests the Go code and compiles static binaries for multiple platforms (Linux AMD64, macOS AMD64, macOS ARM64).
4. The workflow then clones the `Volontariapp/npm-packages` repository, copies the newly built binaries into `npm-packages/tools/`, and pushes a commit to update them.

### 🔑 Important: GitHub Actions Secrets

For the deployment step to work, this repository (`changelog-checker`) **MUST** have access to the `ORG_REPO_TOKEN` secret. 

Since GitHub Actions secrets are scoped per repository and are not automatically passed from the reusable workflow's parent repo (`ci-tools`), the token must be provided.

**How to configure the secret:**
1. Navigate to your the GitHub Organization settings (`Volontariapp`).
2. Go to **Settings > Secrets and variables > Actions**.
3. Ensure the `ORG_REPO_TOKEN` is created as an **Organization Secret**.
4. Configure the secret access to include `changelog-checker` (or set it to be available to all repositories).
