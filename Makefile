.PHONY: run shutdown bdd

run:
	cd deployments/stage &&	\
	docker-compose up -d --build && \
	docker image prune -f && \
	cd ../../calendar/migrations && \
	go build -o goose *.go && \
	./goose --host=127.0.0.1 --port=5432 --user=postgres --password=password --dbname=otus --sslmode=disable up

shutdown:
	cd deployments/stage && \
	docker-compose down && \
	docker system prune --force --volumes && \
	rm  ../../calendar/migrations/goose

bdd:
	set -e ;\
	cd tests/docker && docker-compose -f docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	sleep 3 ;\
	docker-compose -f docker-compose.test.yml run migrator_test ./migrator --dir=/calendar/migrations --host=postgres_test --port=5432 --user=postgres --password=password --dbname=otus --sslmode=disable up ;\
	docker-compose -f docker-compose.test.yml run calendar_integration_tests go test -tags integration ./... || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml run scheduler_integration_tests go test -tags integration ./... || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml down ;\
	docker system prune --force --volumes ;\
	exit $$test_status_code ;\
