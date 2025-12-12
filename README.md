# My Go Project Template

> Other references
>
> - https://github.com/thockin/go-build-template/tree/master
> - https://peter.bourgon.org/go-best-practices-2016/

## TODO

- [ ] Update your app name in:
  - [ ] [Taskfile.yml](./Taskfile.yml)
  - [ ] [Dockerfile](./Dockerfile)
  - [ ] [main.go](./main.go)
- [ ] Update the `main.go` file location in:
  - [ ] [Taskfile.yml](./Taskfile.yml)
  - [ ] [.goreleaser.yml](./.goreleaser.yml)
- [ ] Initialize your Go module:
  ```bash
  go mod init YOUR_MODULE_NAME
  go mod tidy
  ```

## Build & Test

- Using [Taskfile](https://taskfile.dev/)

_Install Taskfile: [Installation Guide](https://taskfile.dev/docs/installation)_

```bash
# Available tasks
task --list
task: Available tasks for this project:
* all:                   Run comprehensive checks: format, lint, security and test
* build:                 Build the application binary for the current platform
* build-docker:          Build Docker image
* build-platforms:       Build the application binaries for multiple platforms and architectures
* fmt:                   Formats all Go source files
* install:               Install required tools and dependencies
* lint:                  Run static analysis and code linting using golangci-lint
* run:                   Runs the main application
* security:              Run security vulnerability scan
* test:                  Runs all tests in the project      (aliases: tests)
* vet:                   Examines Go source code and reports suspicious constructs
```

- Build with [goreleaser](https://goreleaser.com/)

_Install GoReleaser: [Installation Guide](https://goreleaser.com/install/)_

```bash
# Build locally
goreleaser release --snapshot --clean
...
```
