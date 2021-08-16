Requirement:

- Golang installed
- Mysql server

 How to build server:

- Prepare the database in "database.sql"
- "private.key" and "public.crt" is for https config, can replace if needed.
- Config server's information in "properties.ini"
- Parse the "dist" folder after building from frontend
- In "dist/index.html", there is a variable named "host". Change it to "https://<your_server>/api"
- For testing, run cmd: go run main.go
- For building, run cmd: go build