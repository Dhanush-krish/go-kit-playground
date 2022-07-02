### Greeter API using go kit

# local setup
*   clone the repo
*   download all the dependencies
*   run the main file  => 
```go run main.go```

## Request
GET    /greet

```
curl --location --request GET 'http://127.0.0.1:8080/greet' 
--header 'Content-Type: application/json'
--data-raw '{
    "Name":"Dhanush"
}'
```

## Response
```
{
    "message": "Welcome Dhanush!"
}
```