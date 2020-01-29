GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

DOCKER_REPO:=beschmutzen/weathersource

# Generated Wire DI file
GENERATED_FILES += cmd/weather/deps/deps_wire_gen.go

-include .makefiles/Makefile
-include .makefiles/pkg/go/v1/Makefile
-include .makefiles/pkg/protobuf/v1/Makefile
-include .makefiles/pkg/docker/v1/Makefile

artifacts/bin/wire:
	@mkdir -p "$(@D)"
	GOBIN="$(shell pwd)/$(@D)" go install github.com/google/wire/cmd/wire

artifacts/bin/go-bindata:
	@mkdir -p "$(@D)"
	GOBIN="$(shell pwd)/$(@D)" go install github.com/go-bindata/go-bindata/go-bindata


DEPS := $(shell find cmd/weather/deps -type f \! -name "*_gen.go")
cmd/weather/deps/deps_wire_gen.go: $(DEPS) $(PROTO_FILES:.proto=.pb.go)| artifacts/bin/wire
	artifacts/bin/wire gen --output_file_prefix=deps_ "$(shell pwd)/$(@D)"

WWW_SRC_FILES := $(shell find www/weathersource/src -type f -name "*")
WWW_SRC_FILES += $(shell find www/weathersource/public -type f -name "*")
www/weathersource/build/index.html: $(WWW_SRC_FILES)
	npm --prefix $(shell pwd)/www/weathersource run-script build $(shell pwd)/www/weathersource

ASSETS_FILE := assets/assets.go
ASSETS_FILE_DIR = $(dir $(ASSETS_FILE))
.PHONY: assets
assets: www/weathersource/build/index.html | artifacts/bin/go-bindata
	artifacts/bin/go-bindata -o "$(ASSETS_FILE)" -pkg $(ASSETS_FILE_DIR:/=) \
		-prefix $(dir $<) \
		-modtime 1257894000000000000 \
		-mode 400 \
		-fs \
		"$(dir $<)..."
	gofmt -s -w "$(ASSETS_FILE)"

.PHONY: npm-start
npm-start:
	npm --prefix www/weathersource start

.PHONY: stack
stack:
	docker stack deploy -c docker-compose.yml weathersource

.PHONY: unstack
unstack:
	docker stack rm weathersource

.PHONY: mysql-schema
mysql-schema: storage/mysqlstorage/schema.sql
	docker run \
		--rm \
		--interactive \
		--tty \
		--network weathersource_backend \
		--volume "$(shell pwd):/mnt" \
		mariadb:latest \
		sh -c 'cat /mnt/storage/mysqlstorage/schema.sql | mysql -h db -u root --password=root'

.PHONY: mysql-client
mysql-client:
	docker run \
		--rm \
		--interactive \
		--tty \
		--network weathersource_backend \
		--volume "$(shell pwd):/mnt" \
		mariadb:latest \
		mysql -h db -u root --password=root

.PHONY: run
run: artifacts/build/debug/$(GOOS)/$(GOARCH)/weather
	WEATHER_DB_DSN="weathersource:weathersource-password@tcp(127.0.0.1:3306)/weathersource" \
	$< $(RUN_ARGS)


.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"


