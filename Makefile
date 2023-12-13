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
# PUBLIC MACROS
#

CLI     = sortof
DESTDIR = ./dist
GO      = go
GOFLAGS = 
LDFLAGS = -ldflags "-s -w -X main.AppVersion=$(CLI_VERSION)"


#
# INTERNAL MACROS
#

CLI_DIR               = ./cmd/sortof
CLI_CURRENT_VER_TAG   = $$(git tag --points-at HEAD | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)
CLI_LATEST_VERSION    = $$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)
CLI_PSEUDOVERSION     = $$(VER="$(CLI_LATEST_VERSION)"; echo "$${VER:-0001.01}")-$$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')
CLI_VERSION           = $$(VER="$(CLI_CURRENT_VER_TAG)"; echo "$${VER:-$(CLI_PSEUDOVERSION)}")
MODULE_LATEST_VERSION = $$(git tag | grep "^v" | sed 's/^v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)


#
# DEVELOPMENT TASKS
#

.PHONY: all
all: install-dependencies

.PHONY: clean
clean:
	@echo '# Delete bulid directory' >&2
	rm -rf ./dist

.PHONY: info
info:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@$(GO) version || true
	@echo '# Go environment variables:'
	@$(GO) env || true

.PHONY: check
check:
	@echo '# Static analysis' >&2
	$(GO) vet

.PHONY: test
test:
	@echo '# Unit tests' >&2
	@$(GO) test .

.PHONY: e2e
e2e:
	@echo '# E2E tests of $(DESTDIR)/$(CLI)' >&2
	@printf '1\n2\n3\n' >test_case.sorted
	@printf '1\n3\n2\n' >test_case.unsorted
	@printf '1\n3\n' >test_case.stalinsorted
	$(DESTDIR)/$(CLI) -v
	$(DESTDIR)/$(CLI) -h
	$(DESTDIR)/$(CLI) bogo <test_case.unsorted | diff test_case.sorted -
	$(DESTDIR)/$(CLI) bogo -t 5s <test_case.unsorted | diff test_case.sorted -
	$(DESTDIR)/$(CLI) miracle <test_case.sorted | diff test_case.sorted -
	$(DESTDIR)/$(CLI) miracle -t 1ms <test_case.unsorted 2>&1 | grep '^sortof: '
	$(DESTDIR)/$(CLI) slow <test_case.unsorted | diff test_case.sorted -
	$(DESTDIR)/$(CLI) slow -t 100ms <test_case.unsorted | diff test_case.sorted -
	$(DESTDIR)/$(CLI) stalin <test_case.unsorted | diff test_case.stalinsorted -
	$(DESTDIR)/$(CLI) stalin -t 400000ns <test_case.unsorted | diff test_case.stalinsorted -

.PHONY: build
build:
	@echo '# Build CLI executable: $(DESTDIR)/$(CLI)' >&2
	$(GO) build -C $(CLI_DIR) $(GOFLAGS) $(LDFLAGS) -o '../../$(DESTDIR)/$(CLI)'
	@echo '# Add executable checksum to: $(DESTDIR)/sha256sum.txt' >&2
	cd $(DESTDIR); sha256sum $(CLI) >> sha256sum.txt

.PHONY: dist
dist: sortof-linux_amd64 sortof-openbsd_amd64 sortof-windows_amd64.exe

.PHONY: install-dependencies
install-dependencies:
	@echo '# Install CLI dependencies:' >&2
	@GOFLAGS='-v -x' $(GO) get -C $(CLI_DIR) $(GOFLAGS) .

.PHONY: cli-release
cli-release: check test
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new CLI release tag' >&2
	@VER="$(CLI_LATEST_VERSION)"; printf 'Choose new version number for CLI (calver; >%s): ' "$${VER:-2023.10}"
	@read -r NEW_VERSION; \
		git tag "cli/v$$NEW_VERSION"; \
		git push --tags

.PHONY: module-release
module-release: check test
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new Go module release tag' >&2
	@VER="$(MODULE_LATEST_VERSION)"; printf 'Choose new version number for module (semver; >%s): ' "$${VER:-0.0.0}"
	@read -r NEW_VERSION; \
		git tag "v$$NEW_VERSION"; \
		git push --tags


#
# SUPPORTED EXECUTABLES
#

# this forces using `go build` for changes detection in Go related files (instead of `make`)
.PHONY: sortof-linux_amd64 sortof-openbsd_amd64 sortof-windows_amd64.exe

sortof-linux_amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) CLI=$@ build

sortof-openbsd_amd64:
	GOOS=openbsd GOARCH=amd64 $(MAKE) CLI=$@ build

sortof-windows_amd64.exe:
	GOOS=windows GOARCH=amd64 $(MAKE) CLI=$@ build
