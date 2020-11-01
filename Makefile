test:
	make test_required
	docker run \
        --name neo4j-test \
        -p7474:7474 -p7687:7687 \
        -d \
        -v $(PWD)/.neo4j-test/data:/data \
        -v $(PWD)/.neo4j-test/logs:/logs \
        -v $(PWD)/.neo4j-test/import:/import \
        -v $(PWD)/.neo4j-test/plugins:/plugins \
        --env NEO4J_AUTH=neo4j/my-neo4j-password \
        neo4j:latest
    #add the integration tests and unit tests here
    #go test -v -coverprofile=coverage.out ./internal/...
    #dont forget to set the envs for integration tests if required
    #go test -v ./internal/... --tags=integration
	make test_cleanup


run_local:
	make required
	docker run \
        --name neo4j \
        -p7474:7474 -p7687:7687 \
        -d \
        -v $(PWD)/.neo4j/data:/data \
        -v $(PWD)/.neo4j/logs:/logs \
        -v $(PWD)/.neo4j/import:/import \
        -v $(PWD)/.neo4j/plugins:/plugins \
        --env NEO4J_AUTH=neo4j/my-neo4j-password \
        neo4j:latest
	go run ./cmd/url-shorten/main.go --configFile=./cmd/url-shorten/config.yml

required:
	cp "./cmd/url-shorten/config-example.yml" "./cmd/url-shorten/config.yml" 2>/dev/null
	mkdir -p ".neo4j"
	mkdir -p ".neo4j/data"
	mkdir -p ".neo4j/logs"
	mkdir -p ".neo4j/import"
	mkdir -p ".neo4j/plugins"

cleanup:
	rm -rfv "./cmd/url-shorten/config.yml"
	rm -rfv ".neo4j"
	rm -rfv ".build"

test_required:
	cp "./cmd/url-shorten/config-example.yml" "./cmd/url-shorten/test-config.yml" 2>/dev/null
	mkdir -p ".neo4j-test"
	mkdir -p ".neo4j-test/data"
	mkdir -p ".neo4j-test/logs"
	mkdir -p ".neo4j-test/import"
	mkdir -p ".neo4j-test/plugins"

test_cleanup:
	rm -fv "./cmd/url-shorten/test-config.yml"
	rm -rfv ".neo4j-test"
