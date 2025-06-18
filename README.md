# sortof

[![Go Reference](https://pkg.go.dev/badge/github.com/macie/sortof.svg)](https://pkg.go.dev/github.com/macie/sortof)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8153/badge)](https://www.bestpractices.dev/projects/8153)
[![Quality check status](https://github.com/macie/sortof/actions/workflows/check.yml/badge.svg)](https://github.com/macie/sortof/actions/workflows/check.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/macie/sortof)](https://goreportcard.com/report/github.com/macie/sortof)

**sortof** provides implementations of ~~peculiar~~ _carefully selected_ sorting algorithms as a:

- CLI tool for sorting input, line by line (similar to _[POSIX sort](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sort.html)_),
- Go module with functions for sorting slices (similar to _[slices.Sort()](https://pkg.go.dev/slices#Sort)_).

Implemented algorithms:

- impractical due to poor time complexity:
    - [bogosort](https://en.wikipedia.org/wiki/Bogosort)
    - [miraclesort](https://en.wikipedia.org/wiki/Bogosort#miracle_sort)
    - [slowsort](https://en.wikipedia.org/wiki/Slowsort)

- impractical due to destructive behavior:
    - [hitlersort](https://devrant.com/rants/2066915/hitler-sort-delete-every-odd-element)
    - [stalinsort](https://mastodon.social/@mathew/100958177234287431)

## Usage

```sh
$ cat letters.txt
c
a
b
$ sortof bogo -t 10s letters.txt
a
b
c
```

## Installation

Download the [latest stable release from GitHub](https://github.com/macie/sortof/releases/latest).

You can also build it manually with: `make && make build`.

## Development

For detailed contribution guidelines, see [CONTRIBUTING.md](./CONTRIBUTING.md).

Use `make` (GNU or BSD):

- `make` - install dependencies
- `make test` - run tests
- `make e2e` - run end-to-end tests for CLI
- `make check` - run static code analysis
- `make build` - compile binaries from latest commit
- `make dist` - compile binaries from latest commit for supported OSes
- `make clean` - remove compilation artifacts
- `make cli-release` - tag latest commit as a new release of CLI
- `make module-release` - tag latest commit as a new release of Go module
- `make info` - print system info (useful for debugging).

### Versioning

This repository contains both a CLI and Go module which can be developed at different paces.
Commits with versions are tagged with:
- `vX.X.X` (_[semantic versioning](https://semver.org/)_) - versions of Go module
- `cli/vYYYY.0M.MICRO` (_[calendar versioning](https://calver.org/)_) - versions of command-line utility.

### Security hardening

On modern Linux distributions and OpenBSD, the CLI application has restricted access to kernel
calls with [seccomp](https://en.wikipedia.org/wiki/Seccomp) and [pledge](https://man.openbsd.org/pledge.2).

## License

[MIT](./LICENSE)
