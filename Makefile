.PHONY:  clean test secruity build

APP_NAME = simpe_auth_backend
BUILD_DIR = $(PWD)/build
MIGRATION_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

secruity:
	gosec -quite ./...

test: secruity
	go test -v -timeout 30s -coverprofile=cover.out -cover ./simple-auth-backend/
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BUILD_DIR}/${APP_NAME} main.go

run: build
	${BUILD_DIR}/${APP_NAME}

migrate.up:
	migrate -path  ${MIGRATION_FOLDER} -database "${DATABASE_URL}" up 

migrate.down:
	migrate -path  ${MIGRATION_FOLDER} -database "${DATABASE_URL}" down 

migrate.force:
	migrate -path  ${MIGRATION_FOLDER} -database "${DATABASE_URL}" force $(version)


docker.run: docker.network docker.postgres docker.fiber migrate.up
	

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name dev-fiber \
		--network dev-network \
		-p 8080:8080 \
		fiber

docker.postgres:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.stop: docker.stop.fiber docker.stop.postgres

docker.stop.fiber:
	docker stop dev-fiber

docker.stop.postgres:
	docker stop dev-postgres

docker.purge: migrate.down docker.stop 
