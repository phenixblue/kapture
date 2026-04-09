# Example Workflows

This page provides practical local and CI usage patterns for `kapture`.

## Local operator workflow

```bash
make build
./bin/kapture scan --output table
./bin/kapture scan --output json > report.json
```

## Namespace-scoped security review

```bash
./bin/kapture scan \
  --category security \
  --namespace tenant-a \
  --exclude-namespace tenant-a-shared \
  --output table
```

## Rego policy bundle validation

```bash
./bin/kapture scan \
  --engine rego \
  --policy-bundle ./policy/baseline \
  --output json
```

## CI gating workflow (shell)

```bash
set -euo pipefail

./bin/kapture scan --output json > report.json

# Exit code handling:
# 0 = pass, 2 = violations, 3 = partial/degraded
# Gate strictly on 0 in protected environments.
```

## CI gating workflow (GitHub Actions snippet)

```yaml
- name: Build scanner
  run: make build

- name: Run scan
  run: ./bin/kapture scan --output json > report.json

- name: Upload report artifact
  uses: actions/upload-artifact@v4
  with:
    name: kapture-report
    path: report.json
```

## Exception-managed workflow with waivers

```bash
./bin/kapture scan \
  --output json \
  --waiver-file ./waivers.yaml > report.json
```

Use this mode when temporary exceptions are approved and tracked with owner + expiry metadata.
