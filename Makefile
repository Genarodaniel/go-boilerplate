migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations up

migrate-up-force:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations -verbose force 0000001

test:
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out


TODO LIST:

1 - swagger
2 - JWT
2 - get/put/delete methods.
3 - OTEL use open telemetry / FIND A LOG LIBRARY
4 - LEARN HOW TO TERRAFORM EC2/ECS AND RDS
5 - CI/CD
6 - README.MD