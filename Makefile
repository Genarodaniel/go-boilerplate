migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations up

migrate-up-force:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations -verbose force 0000001

test:
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out


TODO LIST:

1 - FIND A LOG LIBRARY
2 - gracefully shutdown
3 - get/put methods. 
4 - OTEL use open telemetry
5 - LEARN HOW TO TERRAFORM EC2/ECS AND RDS
6 - CI/CD
7 - README.MD