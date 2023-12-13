# This Makefile intended to be POSIX-compliant (2018 edition with .PHONY target).
#
# .PHONY targets are used by:
#  - task definintions
#  - compilation of Go code (force usage of `go build` to changes detection).
#
# More info:
#  - docs: <https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html>
#  - .PHONY: <https://www.austingroupbugs.net/view.php?id=523>
#
.POSIX:
.SUFFIXES:


#
# INTERNAL MACROS
#

CLI_DIR = ./cmd/sortof


#
# DEVELOPMENT TASKS
#

.PHONY: all
all: install-dependencies

.PHONY: clean
clean:
	@echo '# Delete binaries: rm -rf ./dist' >&2
	@rm -rf ./dist

.PHONY: info
info:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@go version || true
	@echo '# Go environment variables:'
	@go env || true

.PHONY: check
check:
	@echo '# Static analysis: go vet' >&2
	@go vet

.PHONY: test
test:
	@echo '# Unit tests: go test .' >&2
	@go test .

.PHONY: e2e
e2e:
	@echo '# E2E tests of ./dist/sortof' >&2
	@printf '1\n2\n3\n' >test_case.sorted
	@printf '1\n3\n2\n' >test_case.unsorted
	@printf '1\n3\n' >test_case.stalinsorted
	./dist/sortof -v
	./dist/sortof -h
	./dist/sortof bogo <test_case.unsorted | diff test_case.sorted -
	./dist/sortof bogo -t 5s <test_case.unsorted | diff test_case.sorted -
	./dist/sortof miracle <test_case.sorted | diff test_case.sorted -
	./dist/sortof miracle -t 1ms <test_case.unsorted 2>&1 | grep '^sortof: '
	./dist/sortof slow <test_case.unsorted | diff test_case.sorted -
	./dist/sortof slow -t 100ms <test_case.unsorted | diff test_case.sorted -
	./dist/sortof stalin <test_case.unsorted | diff test_case.stalinsorted -
	./dist/sortof stalin -t 400000ns <test_case.unsorted | diff test_case.stalinsorted -

.PHONY: build
build:
	@echo '# Create release binary: ./dist/sortof' >&2
	@CURRENT_VER_TAG="$$(git tag --points-at HEAD | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		PREV_VER_TAG="$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		CURRENT_COMMIT_TAG="$$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')"; \
		PSEUDOVERSION="$${PREV_VER_TAG:-0001.01}-$$CURRENT_COMMIT_TAG"; \
		VERSION="$${CURRENT_VER_TAG:-$$PSEUDOVERSION}"; \
		go build -C $(CLI_DIR) -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../../dist/sortof'; \

.PHONY: dist
dist:
	@echo '# Create release binaries in ./dist' >&2
	@CURRENT_VER_TAG="$$(git tag --points-at HEAD | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		PREV_VER_TAG="$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		CURRENT_COMMIT_TAG="$$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')"; \
		PSEUDOVERSION="$${PREV_VER_TAG:-0001.01}-$$CURRENT_COMMIT_TAG"; \
		VERSION="$${CURRENT_VER_TAG:-$$PSEUDOVERSION}"; \
		GOOS=openbsd GOARCH=amd64 go build -C $(CLI_DIR) -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../../dist/sortof-openbsd_amd64'; \
		GOOS=linux GOARCH=amd64 go build -C $(CLI_DIR) -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../../dist/sortof-linux_amd64'; \
		GOOS=windows GOARCH=amd64 go build -C $(CLI_DIR) -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../../dist/sortof-windows_amd64.exe'; \

	@echo '# Create binaries checksum' >&2
	@sha256sum ./dist/* >./dist/sha256sum.txt

.PHONY: install-dependencies
install-dependencies:
	@echo '# Install CLI dependencies:' >&2
	@go get -v -x .

.PHONY: cli-release
cli-release: check test
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new CLI release tag' >&2
	@PREV_VER_TAG=$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1); \
		printf 'Choose new version number for CLI (calver; >%s): ' "$${PREV_VER_TAG:-2023.10}"
	@read -r VERSION; \
		git tag "cli/v$$VERSION"; \
		git push --tags

.PHONY: module-release
module-release: check test
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new Go module release tag' >&2
	@PREV_VER_TAG=$$(git tag | grep "^v" | sed 's/^v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1); \
		printf 'Choose new version number for module (semver; >%s): ' "$${PREV_VER_TAG:-0.0.0}"
	@read -r VERSION; \
		git tag "v$$VERSION"; \
		git push --tags
