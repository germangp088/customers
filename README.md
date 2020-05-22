# customers

CRUD customers example on golang.

# Run docker
```sh
$ docker-compose build
$ docker-compose up
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
