.POSIX:
.SUFFIXES:

# MAIN TARGETS

all: install-dependencies

clean:
	@echo '# Delete binaries: rm -rf ./dist' >&2
	@rm -rf ./dist

info:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@go version || true
	@echo '# Go environment variables:'
	@go env || true

check:
	@echo '# Static analysis: go vet' >&2
	@go vet
	
test:
	@echo '# Unit tests: go test .' >&2
	@go test .

build: check test
	@echo '# Create release binaries in ./dist/sortof' >&2
	@CURRENT_VER_TAG="$$(git tag --points-at HEAD | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		PREV_VER_TAG="$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		CURRENT_COMMIT_TAG="$$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')"; \
		PSEUDOVERSION="$${PREV_VER_TAG:-0.0.0}-$$CURRENT_COMMIT_TAG"; \
		VERSION="$${CURRENT_VER_TAG:-$$PSEUDOVERSION}"; \
		go build -C cli/ -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../dist/sortof'; \

dist: check test
	@echo '# Create release binaries in ./dist' >&2
	@CURRENT_VER_TAG="$$(git tag --points-at HEAD | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		PREV_VER_TAG="$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1)"; \
		CURRENT_COMMIT_TAG="$$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')"; \
		PSEUDOVERSION="$${PREV_VER_TAG:-0.0.0}-$$CURRENT_COMMIT_TAG"; \
		VERSION="$${CURRENT_VER_TAG:-$$PSEUDOVERSION}"; \
		GOOS=openbsd GOARCH=amd64 go build -C cli/ -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../dist/sortof-openbsd_amd64'; \
		GOOS=linux GOARCH=amd64 go build -C cli/ -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../dist/sortof-linux_amd64'; \
		GOOS=windows GOARCH=amd64 go build -C cli/ -ldflags="-s -w -X main.AppVersion=$$VERSION" -o '../dist/sortof-windows_amd64.exe'; \

	@echo '# Create binaries checksum' >&2
	@sha256sum ./dist/* >./dist/sha256sum.txt

install-dependencies:
	@echo '# Install CLI dependencies:' >&2
	@go get -v -x .

cli-release:
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new CLI release tag' >&2
	@PREV_VER_TAG=$$(git tag | grep "^cli" | sed 's/^cli\/v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1); \
		printf 'Choose new version number for CLI (>%s): ' "$${PREV_VER_TAG:-0.0.0}"
	@read -r VERSION; \
		git tag "cli/v$$VERSION"; \
		git push --tags

module-release:
	@echo '# Update local branch' >&2
	@git pull --rebase
	@echo '# Create new Go module release tag' >&2
	@PREV_VER_TAG=$$(git tag | grep "^v" | sed 's/^v//' | sort -t. -k 1,1n -k 2,2n -k 3,3n | tail -1); \
		printf 'Choose new version number for module (>%s): ' "$${PREV_VER_TAG:-0.0.0}"
	@read -r VERSION; \
		git tag "v$$VERSION"; \
		git push --tags
