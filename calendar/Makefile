run:
	cd deployments/stage &&	\
	docker-compose up -d --build && \
	docker image prune -f && \
	cd ../../migrations && \
	go build -o goose *.go && \
	./goose --host=127.0.0.1 --port=5438 --user=postgres --password=password --dbname=otus --sslmode=disable up

shutdown:
	cd deployments/stage && \
	docker-compose down && \
	docker system prune --force --volumes && \
	rm  ../../migrations/goose