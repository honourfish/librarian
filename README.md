# librarian
librarian restful project

## Run the service

To run the service:

```
$> docker-compose -f service.yml up
```

## Generate Models

use swagger to first validate the config:

```
$> swagger validate ./swagger/swagger.yml
```

## Manual testing

Post:
```
curl --request POST --header "Content-Type: application/json" --data '{"title":"hp","author":"jk"}' localhost:8082/book
```

Get:
```
curl --request GET --header "Content-Type: application/json" localhost:8082/book?title=hp
```

then generate the server:

```
swagger generate server --target=./swagger --spec=./swagger/swagger.yml --exclude-main --name library
```