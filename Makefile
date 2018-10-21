build:
	go build payments-simple/cmd/ps

clean:
	rm ps

run:
	./ps

test:
	go test payments-simple/internal/app/ps -v -count=1
