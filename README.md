# Go Project Template

Go CLI starter: [Task](https://taskfile.dev/) for the dev loop, [GoReleaser](https://goreleaser.com/) for releases, Docker, and Claude Code permissions.

## Start a new repo

```bash
# 1. Create the repo from this template, clone it, then:
task init

# 2. Check it builds
task all
```

`task init` renames every `my-app` placeholder to the repo directory name and runs `go mod init` with the module path from your git remote. Override either:

```bash
task init APP=foo MODULE=github.com/you/foo
```

| Placeholder | Files |
|-------------|-------|
| `my-app` | `Taskfile.yml`, `Dockerfile`, `main.go`, `README.md` |
| `/path/to/main.go` | `Taskfile.yml` (`MAIN_FILE`), `.goreleaser.yml` (`main`) — only if you move `main.go` |

Then delete this section and write your project's docs. The `init` task is safe to leave in place.

## Build & Test

Install [Task](https://taskfile.dev/installation) and run `task --list`.

```bash
task all       # fmt, lint, vet, security, test
task run
task build     # -> dist/my-app
task build-docker
```

Release locally with [GoReleaser](https://goreleaser.com/install/):

```bash
goreleaser release --snapshot --clean
```

## Optional

- **Homebrew cask:** uncomment `homebrew_casks` in `.goreleaser.yml`. The default `GITHUB_TOKEN` can't push to a separate tap repo — add a `RELEASE_TOKEN` secret and use it in `.github/workflows/release.yml`. See [GoReleaser docs](https://goreleaser.com/errors/resource-not-accessible-by-integration/).
- **Distroless image via `ko`:** uncomment the `kos` block in `.goreleaser.yml`.
- **Docker publishing:** `.github/workflows/release.yml` pushes `ghcr.io/<owner>/<repo>` on tag.
