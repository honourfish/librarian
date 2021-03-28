# librarian
librarian restful project

## Run the service

To run the service:

```
$> docker-compose -f service.yml up
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

## Automated testing:

first goto `acceptance_test` directory then:

```
godog -c 0
```

N.B. we have to turn concurrent scenario execution off.

## Manual testing

Post:
```
curl --request POST --header "Content-Type: application/json" --data '{"title":"hp","author":"jk"}' localhost:8082/book
```

Get:
```
curl --request GET --header "Content-Type: application/json" localhost:8082/book?title=hp
```
```
curl --request GET --header "Content-Type: application/json" localhost:8082/book/{title}
```

Put:
```
curl --request PUT --header "Content-Type: application/json" --data '{"title":"hp","author":"jkrowling"}' localhost:8082/book/{title}
```

Delete:
```
curl --request DELETE --header "Content-Type: application/json" localhost:8082/book/{title}
```

