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
        --env NEO4J_AUTH=$(S_NEO4J_USER)/$(S_NEO4J_PASSWORD) \
        neo4j:latest
	# add the integration tests and unit tests here
	# go test -v -coverprofile=coverage.out ./internal/...
	# dont forget to set the envs for integration tests if required
	# go test -v ./internal/... --tags=integration
	make test_cleanup
# todo: write run:
# todo: write run_debug: that runs within docker with debug support

neo4j:
	make required
	docker run \
        --name neo4j \
        -p7474:7474 -p7687:7687 \
        -d \
        -v $(PWD)/.neo4j/data:/data \
        -v $(PWD)/.neo4j/logs:/logs \
        -v $(PWD)/.neo4j/import:/import \
        -v $(PWD)/.neo4j/plugins:/plugins \
        --env NEO4J_AUTH=neo4j/my-secure-password\
        neo4j:latest 2>/dev/null; true
        #todo: get this shit from environment variables

run_local:
	go run ./cmd/url-shorten/main.go --configFile=./cmd/url-shorten/config.yml
	#todo: makefile should run detached and have a make command to stop it later

stop_local:
	#todo: stop the execution of main here
	docker stop neo4j

required:
	make .set_envs
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
	make .set_envs
	cp "./cmd/url-shorten/config-example.yml" "./cmd/url-shorten/test-config.yml" 2>/dev/null
	mkdir -p ".neo4j-test"
	mkdir -p ".neo4j-test/data"
	mkdir -p ".neo4j-test/logs"
	mkdir -p ".neo4j-test/import"
	mkdir -p ".neo4j-test/plugins"

test_cleanup:
	# todo: add removal of volumes and container related to test
	docker stop neo4j-test
	docker rm neo4j-test
	docker volume rm neo4j-test-data
	docker volume rm neo4j-test-logs
	docker volume rm neo4j-test-import
	docker volume rm neo4j-test-plugins
	rm -fv "./cmd/url-shorten/test-config.yml"
	rm -rfv ".neo4j-test"

.set_envs:
	export S_NEO4J_SECURE=false
	export S_NEO4J_TARGET=bolt://localhost:7687
	export S_NEO4J_USER=neo4j
	export S_NEO4J_PASSWORD=my-secure-password
	export S_NEO4J_REALM=neo4j
