guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable '$*' not set"; \
		exit 1; \
	fi

test:
	go test -v -cover ./...

build:
	go build -v ./...

coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run:
	go run web/main.go

create: guard-name guard-type guard-style
	scripts/create.sh -n $(name) -t $(type) -s $(style)

list:
	scripts/list.sh

get: guard-id
	scripts/get.sh $(id)

update: guard-id guard-name guard-type guard-style
	scripts/update.sh -i $(id) -n $(name) -t $(type) -s $(style)

delete: guard-id
	scripts/delete.sh $(id)

container:
	docker build . -t gobeer
	docker run -d -p 4000:4000 gobeer
