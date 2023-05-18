unittest:
	@echo "==> Testing identity package"
	# create test db container
	docker-compose -f ./test/docker-compose.yml --env-file=./test/.env up -d
	go test ./... -v
	# tear down
	docker-compose -f ./test/docker-compose.yml down -v
	# remove the test container
	docker-compose -f ./test/docker-compose.yml rm