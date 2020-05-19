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