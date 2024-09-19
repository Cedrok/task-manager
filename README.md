The base of this project is a exercice from [Gophercises](https://gophercises.com/), using these packages:
- **sqlite** for the DB
- **GoFr** for the API
- **cobra** for the CLI

# Setup
For using *task-manager*, we need a .env file:
`API/configs/.env`
```
APP_NAME=task-manager
HTTP_PORT=9000
DB_NAME=../tasks.db
DB_DIALECT=sqlite
```

# Launching API
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