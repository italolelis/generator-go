masterNO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: all clean deps test build
all: clean deps test build

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure --vendor-only

# Builds the project
build: 
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X cmd.version=${VERSION}" -o "dist/<%=projectName%>" <%=projectRoot%>

test: lint format vet
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	@go test -v -cover -covermode=atomic ./...

format:
	@gofmt -l -s cmd pkg | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	@go vet ./...

lint: tools.golint
	@echo "$(OK_COLOR)==> checking code style with 'golint' tool$(NO_COLOR)"
	@go list ./... | xargs -n 1 golint -set_exit_status

#---------------
#-- tools
#---------------

.PHONY: tools tools.dep tools.golint
tools: tools.dep tools.golint

tools.golint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get github.com/golang/lint/golint; \
	fi

tools.dep:
	@command -v dep >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing dep"; \
		@go get -u github.com/golang/dep/cmd/dep; \
	fi
