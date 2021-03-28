# librarian
librarian restful project

## Run the service

To run the service:

1. start the mongo database service:
```
$> docker-compose -f mongo.yml up
```

2. start the server:
```
$> ./librarian
```

## Generate Server

use swagger to first validate the config:

```
$> swagger validate ./swagger/swagger.yml
```

then generate the server:

```
swagger generate server --target=./swagger --spec=./swagger/swagger.yml --exclude-main --name library
```

## Generate Client

to generate the client run:

```
swagger generate client --target=./client --spec=./swagger/swagger.yml --name library
```

## Acceptance testing:

### Setup

1. start the mongo database:
```
$> docker-compose -f mongo.yml up
```

2. Then the service:
```
$> ./librarian
```

3. Due to not implementing any Librarian management endpoints, we have to create a couple of librarians first:
    a. head to the mongo db admin: `localhost:8081`
    b. create a database (if one doesn't exist) called `library`
    c. create a collection (if one doesn't exist) called `librarians`
    d. create the librarians with the following attributes:
| role | username |
|------|----------|
| Senior | Janice |
| Librarian | Jason |

4. Now in a different terminal get the godog executable:
```
$> go get github.com/cucumber/godog/cmd/godog
```

## Run the tests

first goto `acceptance_test` directory then:

```
godog -c 0 --random
```

N.B. we have to turn concurrent scenario execution off (-c 0).
N.B. --random is used to ensure random scenario execution order.

