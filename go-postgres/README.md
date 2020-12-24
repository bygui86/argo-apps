
# Go postgres

Example project to test PostgreSQL integration in Go

## Build
```shell script
make build
```

---

## Test
```shell script
make test
```

---

## Run

1. start PostgreSQL in a container
```shell script
make run-postgres
```

2. run application
```shell script
make run
```

3. play a bit with [Postman](https://www.postman.com/) loading the [prepared collection](postman/postman_collection.json)

---

## REST endpoints

- `GET /products` > Fetch a list of products in response to a valid 
- `GET /products/{id}` > Fetch a product in response to a valid 
- `POST /products` > Create a new product in response to a valid 
- `PUT /products/{id}` > Update a product in response to a valid 
- `DELETE /products/{id}` > Delete a product in response to a valid 

---

## Links
- https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
- https://errorsingo.com/github.com-lib-pq-err-ssl-not-supported/
- https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
