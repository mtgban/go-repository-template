# go-repository-template

The starting point for a new Go repository in the mtgban org. It carries the
org's CI gates, lint rules and Dependabot setup, so a new repository starts
out matching the others.

## After "Use this template"

1. Rename the module in `go.mod` from
   `github.com/mtgban/go-repository-template`, and the binary
   `/go-repository-template` in `.gitignore`.
2. Replace `main.go` and `main_test.go` with the real code.
3. If the default branch is not `main`, change the `push` branch in
   `.github/workflows/ci.yml`.
4. Rewrite this README.
5. In the new repository's settings, protect the default branch and require
   the `ci` check. Templates do not copy settings.

## What every pull request runs

`.github/workflows/ci.yml`, under the Go version that `go.mod` names
(`go 1.26.0`, `toolchain go1.26.8`):

| Gate | Command |
|---|---|
| Formatting | `gofmt -s -l .` must print nothing |
| Vet | `go vet ./...` |
| Lint | `revive@v1.13.0` with `.revive.toml` |
| Static analysis | `staticcheck@2026.2.1` |
| Vulnerabilities | `govulncheck@latest`, which covers the standard library too |
| Build | `go build ./...` |
| Tests | `go test -race ./...` |

revive and staticcheck are pinned, so a new release cannot fail a build that
changed nothing. govulncheck is not pinned, because its verdict follows the
vulnerability database. Actions stay on version tags rather than SHA pins.

`.revive.toml` is the org-wide rule set. A repository keeps its own
`unhandled-error` exemptions, each with a line saying why, and may add or
drop a rule with a comment. `reflect` is blocked: compare with
`slices`/`maps` or with a comparison written for the type.

`.github/dependabot.yml` proposes Go modules and GitHub Actions weekly, one
grouped pull request each, after a seven-day cooldown. Security updates skip
the cooldown, and `github.com/mtgban/*` modules are proposed without waiting.

## Running the gates locally

Run them under the same toolchain as CI:

```bash
export GOTOOLCHAIN=go1.26.8
"$(go env GOROOT)/bin/gofmt" -s -l .
go vet ./...
go run github.com/mgechev/revive@v1.13.0 -set_exit_status -config .revive.toml ./...
go run honnef.co/go/tools/cmd/staticcheck@2026.2.1 ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
go build ./...
go test -race ./...
```

`GOTOOLCHAIN` switches the `go` command but not a `gofmt` on your `PATH`. A
newer `gofmt` can lay out a file differently from CI's, so call the
toolchain's own binary as above.
