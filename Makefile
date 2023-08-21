.POSIX:
.SUFFIXES:
GO = go

# main targets
all: check

clean:
	@echo '> Remove dist directory'
	rm -rf ./dist

check:
	@echo '> Check project'	
	$(GO) test .
	$(GO) vet .
