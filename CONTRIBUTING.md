# Contributing to sortof

Thank you for your interest in **sortof**! Contributions are welcome from everyone, no matter your experience level with programming or open source.

By contributing, you agree that your work will be licensed under the same terms as the rest of the project (see [LICENSE](./LICENSE) for details).

By following these guidelines, you’ll benefit from:

- smoother collaboration and less back-and-forth
- faster reviews and merges
- opportunities to learn Go best practices, testing, and open source workflows.

## Ways to contribute

### For everyone

- Ask **questions** about algorithms or code in [discussions](https://github.com/macie/sortof/discussions)
- Report **bugs** using _[Bug report](https://github.com/macie/sortof/issues)_ issue template
- Suggest **new algorithms** (even without implementation) with _[Feature suggestion](https://github.com/macie/sortof/issues)_ issue template
- Participate in **discussions** to help others
- Create **examples** demonstrating the library or CLI usage
- Fix **typos** in documentation and code comments (see below)
- Test on **different OSes** and report any issues
- Improve **documentation** for clarity and completeness (see below)
- Help **organize issues** by labeling and triaging bug reports

### For programmers

- Fix **bugs** in existing implementations
- Write additional **tests** to improve coverage
- Expand **platform support** for different operating systems
- Implement **new sorting algorithms** (see detailed guide below)

When reporting **build issues**, include `make info` output with your report.

If you're looking for a place to start, check issues labeled `good first issue` or `help wanted`.

## Small improvements

For **typo fixes** or **documentation improvements**:

1. Edit the file directly on GitHub (click the pencil icon)
2. Describe your changes in the commit message
3. Create a pull request

## Code changes

Before contributing code, you'll need:

- `git` - for version control ([cheat sheet](https://git-scm.com/book/en/v2/Getting-Started-Git-Basics))
- `go` - the programming language ([installation](https://go.dev/doc/install))
- `make` - for running build commands (usually pre-installed on Linux/macOS), such as:
    - `make check` - runs code quality checks (formatting, linting)
    - `make test` - runs unit tests for all algorithms
    - `make e2e` - runs end-to-end tests (tests the CLI tool)
    - `make info` - shows system and build information

For **bug fixes** or **small improvements** use the standard GitHub workflow:

1. Fork and clone the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Test: `make check && make test && make e2e`
5. Commit and submit a pull request

### Adding a new sorting algorithm

Each PR with a new algorithm must include:

1. A new `{algorithm}.go` file with implementation:

    ```go
    func {Algorithm}[T constraints.Ordered](ctx context.Context, data []T) error {
        // implementation must support context cancellation (check ctx.Done() periodically)
    }
    ```

2. A new `{algorithm}_test.go` file with tests (see `bogosort_test.go` as example)

3. Changes to `cmd/sortof/cli.go` file (to enable the algorithm in the CLI tool)

4. Updates to the `Makefile` file (to add E2E tests for the new algorithm)

5. Updates to the `README.md` file (with a brief description of the new algorithm)

#### Quality requirements

Use this PR template as a checklist:

```markdown
## Description

Brief algorithm description and motivation:

- **Complexity**: O(?)
- **Unique behavior**: What makes it interesting?

## Checklist

- [ ] Algorithm implementation with context support
- [ ] Comprehensive tests (various inputs, edge cases, cancellation)
- [ ] CLI integration added (`cmd/sortof/cli.go`)
- [ ] E2E CLI tests added (`Makefile`)
- [ ] `README.md` updated
```
