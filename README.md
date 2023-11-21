# sortof

[![Go Reference](https://pkg.go.dev/badge/github.com/macie/sortof.svg)](https://pkg.go.dev/github.com/macie/sortof)
[![Quality check status](https://github.com/macie/sortof/actions/workflows/check.yml/badge.svg)](https://github.com/macie/sortof/actions/workflows/check.yml)

**sortof** provides implementations of ~~peculiar~~ _carefully selected_ sorting algorithms as a:

- CLI tool for sorting input, line by line (similar to _[POSIX sort](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sort.html)_),
- Go module with functions for sorting slices (similar to _[slices.Sort()](https://pkg.go.dev/slices#Sort)_).

Implemented algorithms:

- [bogosort](https://en.wikipedia.org/wiki/Bogosort)
- [miraclesort](https://en.wikipedia.org/wiki/Bogosort#miracle_sort) (currently module only)
- [slowsort](https://en.wikipedia.org/wiki/Slowsort)
- [stalinsort](https://mastodon.social/@mathew/100958177234287431).

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

Download [latest stable release from GitHub](https://github.com/macie/sortof/releases/latest) .

You can also build it manually with commands: `make && make build`.

## Development

Use `make` (GNU or BSD):

- `make` - install dependencies
- `make test` - runs test
- `make check` - static code analysis
- `make build` - compile binaries from latest commit
- `make dist` - compile binaries from latest commit for supported OSes
- `make clean` - removes compilation artifacts
- `make cli-release` - tag latest commit as a new release of CLI
- `make module-release` - tag latest commit as a new release of Go module
- `make info` - print system info (useful for debugging).

### Versioning

The repo contains CLI and Go module which can be developed with different pace.
Commits with versions are tagged with:
- `vX.X.X` (_[semantic versioning](https://semver.org/)_) - versions of Go module
- `cli/vYYYY.0M.MICRO` (_[calendar versioning](https://calver.org/)_) - versions of command-line utility.

## License

[MIT](./LICENSE)
