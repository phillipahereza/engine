export ROOT:=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export BIN:=$(ROOT)/bin
export GOPATH:=$(mktemp)
export GOBIN:=$(BIN)
export PATH:=$(BIN):$(PATH)
APP_NAME:=balloon
DEFAULT_PASS=bita123
GO=$(shell which go)
GIT=$(shell which git)
CURL:=$(shell which curl)
CHMOD=$(shell which chmod)
DB_PASS?=$(DEFAULT_PASS)
DB_USER?=$(APP_NAME)
DB_NAME?=$(APP_NAME)
WORK_DIR=$(ROOT)/tmp
LONG_HASH?=$(shell git log -n1 --pretty="format:%H" | cat)
SHORT_HASH?=$(shell git log -n1 --pretty="format:%h"| cat)
COMMIT_DATE?=$(shell git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")
COMMIT_COUNT?=$(shell git rev-list HEAD --count| cat)
BUILD_DATE=$(shell date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
FLAGS="-X pkg.version.hash=$(LONG_HASH) -X version.short=$(SHORT_HASH) -X version.date=$(COMMIT_DATE) -X version.count=$(COMMIT_COUNT) -X version.build=$(BUILD_DATE)"
LD_ARGS=-ldflags $(FLAGS)
GET=cd $(ROOT) && $(GO) get -u -v $(LD_ARGS)
BUILD=cd $(ROOT) && $(GO) build -v $(LD_ARGS)
INSTALL=cd $(ROOT) && $(GO) install -v $(LD_ARGS)
CG_SERVICES_POSTGRES_USER=$(DB_USER)
CG_SERVICES_POSTGRES_PASSWORD=$(DB_PASS)
CG_SERVICES_POSTGRES_DB=$(DB_NAME)
where-am-i = $(CURDIR)/$(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST))

# Include modules make file
include $(wildcard $(ROOT)/modules/*/module.mk)

need_root :
	@[ "$(shell id -u)" -eq "0" ] || exit 1

not_root :
	@[ "$(shell id -u)" != "0" ] || exit 1

database-setup: need_root
	sudo -u postgres psql -U postgres -d postgres -c "CREATE USER $(DB_USER) WITH PASSWORD '$(DB_PASS)';" || sudo -u postgres psql -U postgres -d postgres -c "ALTER USER $(DB_USER) WITH PASSWORD '$(DB_PASS)';"
	sudo -u postgres psql -U postgres -c "CREATE DATABASE $(DB_NAME);" || echo "Database $(DB_NAME) is already there?"
	sudo -u postgres psql -U postgres -c "GRANT ALL ON DATABASE $(DB_NAME) TO $(DB_USER);"

$(BIN)/prototool:
	$(CURL) -sSL https://github.com/uber/prototool/releases/download/v1.3.0/prototool-$(shell uname -s)-$(shell uname -m) -o $(BIN)/prototool
	$(CHMOD) +x $(BIN)/prototool

$(BIN)/protoc-gen-go:
	$(GET) github.com/golang/protobuf/protoc-gen-go

$(BIN)/protoc-gen-gogo:
	$(GET) github.com/gogo/protobuf/protoc-gen-gogo

$(BIN)/protoc-gen-grpc-gateway:
	$(GET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

$(BIN)/protoc-gen-swagger:
	$(GET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

$(BIN)/protoc-gen-grpchan:
	$(GET) github.com/fullstorydev/grpchan/cmd/protoc-gen-grpchan

$(BIN)/go-bindata:
	$(GET) github.com/shuLhan/go-bindata/cmd/go-bindata

swagger-to-go:
	$(INSTALL) ./cmd/swagger-to-go

generators:
	$(INSTALL) ./cmd/protoc-gen-wrapper
	$(INSTALL) ./cmd/protoc-gen-model


tools-migration: $(BIN)/go-bindata $(addsuffix -migration,$(dir $(wildcard $(ROOT)/modules/*/)))
	$(INSTALL) ./cmd/migration

mig-up: tools-migration
	$(BIN)/migration -action=up

mig-down: tools-migration
	$(BIN)/migration -action=down

proto: $(BIN)/prototool $(BIN)/protoc-gen-go $(BIN)/protoc-gen-grpc-gateway $(BIN)/protoc-gen-swagger $(BIN)/protoc-gen-grpchan $(BIN)/protoc-gen-gogo generators
	$(BIN)/prototool generate

swagger-ui: $(BIN)/go-bindata
	$(GIT) clone --depth 1 https://github.com/swagger-api/swagger-ui.git $(ROOT)/tmp/swagger-ui
	rm -rf $(ROOT)/third_party/swagger-ui
	mv $(ROOT)/tmp/swagger-ui/dist $(ROOT)/third_party/swagger-ui
	rm -rf $(ROOT)/tmp/swagger-ui
	sed -i -e 's/https:\/\/petstore.swagger.io\/v2\/swagger\.json/\/v1\/swagger\/index\.json/g' $(ROOT)/third_party/swagger-ui/index.html
	cd $(ROOT)/third_party/swagger-ui && $(BIN)/go-bindata -nometadata -o $(ROOT)/pkg/grpcgw/swagger.gen.go -nomemcopy=true -pkg=grpcgw ./...

swagger: swagger-to-go proto $(addsuffix -swagger,$(dir $(wildcard $(ROOT)/modules/*/)))

code-gen: swagger

LINTER:=$(BIN)/gometalinter.v2
LINTERCMD:=$(LINTER) -e ".*.pb.go" -e ".*_test.go" -e "$(ROOT)/vendor/.*" --cyclo-over=19 --line-length=120 --deadline=100s --disable-all --enable=structcheck --enable=deadcode --enable=gocyclo --enable=ineffassign --enable=golint --enable=goimports --enable=errcheck --enable=varcheck --enable=goconst --enable=gosimple --enable=staticcheck --enable=unused --enable=misspell

$(LINTER):
	$(GET) gopkg.in/alecthomas/gometalinter.v2
	$(LINTER) --install

lint: $(LINTER) $(addsuffix -lint,$(wildcard $(ROOT)/modules/*))
	$(LINTERCMD) $(ROOT)/cmd/...

build-server:
	@echo "Building server"
	$(INSTALL) ./cmd/server

run-server: code-gen build-server
	@echo "Running..."
	$(BIN)/server 2>&1

.PHONY: swagger-to-go proto swagger build-server run-server generate
