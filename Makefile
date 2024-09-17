build:
	@go build .

run: build
	./task

clean:
	rm task
	rm tasks.db

re: clean main