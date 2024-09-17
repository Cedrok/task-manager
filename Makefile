.PHONY: main run clean fclean re

main:
	@go build

run: main
	./task

clean:
	rm task

fclean: clean
	rm tasks.db

re: clean main