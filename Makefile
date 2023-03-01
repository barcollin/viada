build-dev:
	docker build -t xaria -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/composers/dc.dev.yml down

run-dev:
	docker-compose -f containers/composers/dc.dev.yml up
