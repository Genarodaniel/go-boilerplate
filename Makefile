migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations up

migrate-up-force:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations -verbose force 0000001

test:
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
