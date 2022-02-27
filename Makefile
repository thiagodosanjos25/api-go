include .env

go:
	@go build -o api-go cmd/api-go/main.go
	@docker-compose -f deployments/docker-compose.yaml up --build -d
	@rm api-go
	@docker-compose -f deployments/docker-compose.yaml logs -f

stop:
	@docker-compose -f deployments/docker-compose.yaml stop

up:
	@docker-compose -f deployments/docker-compose.yaml up -d

build:
	@docker-compose -f deployments/docker-compose.yaml up --build -d

build-db:
	@docker-compose -f deployments/docker-compose.yaml up --build api-database

down:
	@docker-compose -f deployments/docker-compose.yaml down

logs:
	@docker-compose -f deployments/docker-compose.yaml logs -f

gen-bin:
	@go build -o api-go cmd/api-go/main.go

rm-bin:
	@rm api-go

run:
	@go run $(DIRCMD)main.go 

build-docker: login-docker build-push-docker

build-push-docker:
	@echo "Gerando binário..."
	@go build -o api-go cmd/api-go/main.go

	@echo "Gerando build do projeto docker..."
	@docker-compose -f deployments/docker-compose.yaml build

	@echo "Realizando push das imagens no docker.hub..."
	@docker-compose -f deployments/docker-compose.yaml push

	@echo "Removendo binário..."
	@rm api-go

	@echo "Fim."

login-docker:
	@docker login -u $DOCKER_USER -p $DOCKER_PASWD






