include .env

run-dev:
	@go run $(DIRCMD)main.go 

run-docker-dev:
	@go build -o api-go cmd/api-go/main.go
	@docker-compose -f deployments/docker-compose.desenvolvimento.yaml up --build -d
	@rm api-go
	@docker-compose -f deployments/docker-compose.desenvolvimento.yaml logs -f

build-push-docker:
	@echo "Gerando binário..."
	@go build -o api-go cmd/api-go/main.go

	@echo "Gerando build do projeto docker..."
	@docker-compose -f deployments/docker-compose.desenvolvimento.yaml build

	@echo "Realizando push das imagens no docker.hub..."
	@docker-compose -f deployments/docker-compose.desenvolvimento.yaml push

	@echo "Removendo binário..."
	@rm api-go

	@echo "Fim."

login-docker:
	@docker login -u $DOCKER_USER -p $DOCKER_PASWD



