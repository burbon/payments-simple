build:
	go build payments-simple/cmd/ps

clean:
	rm ps

run:
	./ps

test: drop-test-database test-database run-test

run-test:
	go test payments-simple/internal/app/ps -v -count=1

database:
	sed s/__KEYSPACE__/paymentssimple/ hack/database-init.tcql | cqlsh

test-database:
	sed s/__KEYSPACE__/test_paymentssimple/ hack/database-init.tcql | cqlsh

drop-database:
	sed s/__KEYSPACE__/paymentssimple/ hack/database-drop.tcql | cqlsh

drop-test-database:
	sed s/__KEYSPACE__/test_paymentssimple/ hack/database-drop.tcql | cqlsh

swagger-validate:
	swagger validate ./api/swagger.yml

swagger-generate:
	mkdir -p ./internal/pkg/rest
	swagger generate server \
		--target=./internal/pkg/rest \
		--spec=./api/swagger.yml \
		--exclude-main
