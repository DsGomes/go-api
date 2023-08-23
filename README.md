### Rest API

Project implementing a REST API using GO lang.

#### RUNNING LOCALLY GO API

Up database:
`docker-compose up -d`

Now you can access in browser a database management called **Adminer**:
> http://localhost:8080

Run `./script/create_api_db.sql` in your database management to create your database and api_todo table.

*If you have problems when try stop docker-compose services you can test the command `sudo aa-remove-unknown`

### RUNNING TESTS

You can run all tests using the following command on CLI:
`go test ./tests/...`

If you want to see more informations about:
`go test -v ./tests/...`