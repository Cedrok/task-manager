The base of this project is a exercice from [Gophercises](https://gophercises.com/), using these packages:
- **sqlite** for the DB
- **GoFr** for the API
- **cobra** for the CLI

# Setup
For using *task-manager*, we need a .env file:
`API/configs/.env`
```properties
APP_NAME=task-manager
HTTP_PORT=9000

# DO NOT MODIFY BELOW
DB_NAME=./db/tasks.db
DB_DIALECT=sqlite
```

# Launching API
## With docker compose
```
docker compose up --build -d
```
A volume *task-manager_db* will be created to keep your data when recreating the container.

## With local directory
```
cd API
make run
```

# Using CLI
```
cd CLI
make
./task 
```
This CLI is not persistent. You'll have to run `./task [cmd]` for each command.  
e.g:
```
./task add write an example
Added "write an example" to your task list at index 0.

./task list
You have the following tasks to do:
        0: write an example

./task do 0
Task marked as complete
```