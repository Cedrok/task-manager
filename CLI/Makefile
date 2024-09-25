APP_NAME = task

.PHONY: main run clean re

main:
	go build -o $(APP_NAME)

run: main
	./$(APP_NAME)

clean:
	rm $(APP_NAME)

re: clean main