# customers

CRUD customers example on golang.

# Run docker
```sh
$ docker build --tag customers:1.0 .
$ docker run --publish 8000:8080 --detach --name bb customers:1.0
```

###### ENDPOINTS

-- GET /{URL}/customer

-- GET /{URL}/customer/:id

-- POST /{URL}/customer

-- PUT /{URL}/customer:id

-- DELETE /{URL}/customer:id

**Transaction Body Example**
```json
{
	"Firstname": "Mathias",
	"Lastname": "Pencil"
}
```

## Swagger
-- /{URL}/swaggerui/
