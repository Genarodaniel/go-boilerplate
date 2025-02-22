migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations up

migrate-up-force:
	migrate -database ${POSTGRESQL_URL} -path internal/infra/database/migrations -verbose force ${MIGRATION}

test:
	go test -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

generate-certificate:
	@dir="certs"; \
	if [[ ! -e $$dir ]]; then \
		mkdir -p $$dir; \
	fi; \
	cd $$dir && openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048

extract-public-key:
	@dir="certs"; \
	if [[ ! -e $$dir ]]; then \
		echo "run generate-certificate first"1>&2; \
	fi; \
	cd $$dir && openssl rsa -in private.pem -pubout -out public.pem
