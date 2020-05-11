run:
	cd deployments/stage &&	\
	docker-compose up -d --build && \
	docker image prune -f && \
	cd ../../migrations && \
	go build -o goose *.go && \
	./goose --host=127.0.0.1 --port=5438 --user=postgres --password=password --dbname=otus --sslmode=disable up

integration_test:
	cd deployments/test && \
	docker-compose up -d postgres && \
	sleep 3 && \
	cd ../../migrations && \
	go build -o goose *.go && \
	./goose --host=127.0.0.1 --port=5450 --user=postgres --password=password --dbname=otus --sslmode=disable up \
	&& cd .. \
	&& go test -v ./... -count 1 \
	&& rm migrations/goose \
	&& docker rm -f postgres_test

shutdown:
	cd deployments/stage && docker-compose down && docker system prune --force --volumes && cd ../../migrations && rm goose