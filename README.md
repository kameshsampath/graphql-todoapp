# TODO App

A simple TODO app to practice [GraphQL](https://graphql.org/) using [go](https://golang.org). The app can be run with the following databases,

- Postgresql
- Mariadb/MySql
- Sqlite

## Required Tools

- Docker for Desktop
- Golang >= v1.18

## Application Configuration
| Variable Name	        | Description	                                                        | Default Value	      |
|-----------------------|---------------------------------------------------------------------|---------------------|
| BUNDEBUG              | The debug mode for bun                                              | 2 - log all queries |
| BUN_SCHEMA_GEN_MODE   | Schema Generation mode                                              | drop-and-create     |
| TODOS_DB_FILE         | The db file for sqlite                                              | $PWD/work/todos.db  |
| TODO_DB_TYPE          | The database to use.Valid values sqlite,postgresql,mariadb or mysql | sqlite              |
| PGPASSWORD            | Postgresql User password                                            | postgres            |
| PGUSER                | Postgresql User                                                     | postgres            |
| PGPORT                | Postgresql Port                                                     | 5432                |
| PGHOST                | Postrgesql Host                                                     | localhost           |
| PGDATABASE            | Postgresql database                                                 | postgres            |
| PGSSLMODE             | Use SSL                                                             | disable             |
| MARIADB_ROOT_PASSWORD | mariadb root password                                               | root                |
| MARIADB_HOST          | mariadb host                                                        | localhost           |
| MARIADB_PORT          | mariadb port                                                        | 3306                |
| MARIADB_USER          | mariadb user                                                        | root                |
| MARIADB_PASSWORD      | mariadb user password                                               | root                |
| MARIADB_DATABASE      | mariadb database                                                    | demodb              |

## Start Database Services

```shell
$SOURCE_ROOT/docker-compose up 
```

## Run the app

```shell
go run cmd/app
```

You can access the graphql console from url http://localhost:8080/

Some example queries/mutations are in ./graphql_ops.txt 
