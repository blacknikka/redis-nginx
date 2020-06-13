.PHONY: all
all: gen-client gen-server

.PHONY: gen-client
gen-client:
	rm -rf ./open-api/output/client
	docker run --rm -v ${PWD}/open-api:/local openapitools/openapi-generator-cli:v4.2.2 generate \
		-i  /local/src/openapi.yml \
		-g go \
		-o /local/output/client

gen-server:
	rm -rf ./open-api/output/server
	docker run --rm -v ${PWD}/open-api:/local openapitools/openapi-generator-cli:v4.2.2 generate \
		-i  /local/src/openapi.yml \
		-g go-server \
		-o /local/output/server
