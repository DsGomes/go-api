### Rest API

Project implementing a REST API using GO lang.

#### RUNNING LOCALLY GO API

Up database:
`docker-compose up -d`

Now you can access in browser a database management called **PgAdmin**:
> http://localhost:54321

Run `./script/create_api_db.sql` in your database management to create your database and api_todo table.

*If you have problems when try stop docker-compose services you can test the command `sudo aa-remove-unknown`

Then you can run the project:

`go run main.go`

### RUNNING TESTS

You can run all tests using the following command on CLI:
`go test ./tests/...`

If you want to see more informations about:
`go test -v ./tests/...`


### FOLDER STRUCTURE

```bash
├── configs
├── internal
│   ├── core
│   │   ├── domain
│   │   ├── ports
│   │   └── usecases
│   ├── handlers
│   └── infra
│       ├── db
│       └── repositories
├── script
└── tests
    └── unit
        ├── mocks
        └── usecases
```