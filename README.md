# sortof

**sortof** provides implementations of ~~peculiar~~ _carefully selected_ sorting algorithms as a:

- collection of CLI tools for sorting input, line by line (similar to _[POSIX sort](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sort.html)_),
- Go module with functions for sorting slices (similar to _[slices.Sort()](https://pkg.go.dev/slices#Sort)_).

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
Commits with versions are tagged using [semantic versioning](https://semver.org/):
- `vX.X.X` - versions of Go module
- `cli/vX.X.X` - versions of command-line utility.

## License

[MIT](./LICENSE)
