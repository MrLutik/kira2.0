# Go Development Tools Setup

This guide provides instructions on setting up and using key development tools for Go, including linters and formatters.

## Linters

### golangci-lint

**golangci-lint** is a comprehensive linter that aggregates and configures many Go linters.

#### Installation

To install `golangci-lint`, run the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
```

*Note: Replace `v1.55.2` with the desired version.*

#### Usage

To run `golangci-lint` with a specific linter:

```bash
golangci-lint run ./... --disable-all --no-config -E <linter_name>
```

*Replace `<linter_name>` with the name of the linter you wish to use.*

To run `golangci-lint` with the configuration file `.golangci.yaml`:

```bash
golangci-lint run ./...
```

## Formatters

### gofumpt

**gofumpt** is a formatter that provides stricter formatting than `gofmt` and includes extra rules.

#### Installation

Install `gofumpt` using:

```bash
go install mvdan.cc/gofumpt@latest
```

#### Usage

To format your Go code with `gofumpt`:

```bash
gofumpt -w -extra [path]
```

*Replace `[path]` with the path to your Go files or directories.*

### goimports

**goimports** updates your Go import lines, adding missing ones and removing unreferenced ones.

#### Installation

Install `goimports` with the following command:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

#### Usage

To format and fix imports in your Go files:

```bash
goimports -w [path]
```

*Replace `[path]` with the path to your Go files or directories.*

---

Feel free to adapt this guide to your specific needs or project requirements.