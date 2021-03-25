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

then generate the server:

```
swagger generate server --target=./swagger --spec=./swagger/swagger.yml --exclude-main --name library
```
