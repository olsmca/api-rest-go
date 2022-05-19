# api-rest-go

API en golang con Base de datos mongo

## Configuracion
En la carpeta `docker` se encuentra los archivos

`.env` => contiene la configuracion para autenticar en la base de datos de la aplicacion cargada en `actions.go` y de la imagen docker que contiene la base de datos mongo y mongo express

`mongo-docker-compose.yml` => definicion de las imagenes de docker para la base de datos de mongo y mongo express

## Docker
El proyecto incorpora una imagen de mongodb y mongo express para correr simulando el ambiente local y dev:  
#### Using Docker to simplify development (optional)
```
docker-compose -f docker/mongo-docker-compose.yml up -d
```
Para detenerlo y quitar del contenedor, ejecute:
```
docker-compose -f docker/mongo-docker-compose.yml down
```

## Como funciona la app
la app expone 3 servicios los cuales estan definidos en `router.go`

```GET - /peliculas``` 

```GET - /peliculas/{id}```

```POST - /peliculas```

### GET /peliculas
La peticion Get retorna una lista con los json de peliculas almacenadas en la base de datos

```
[
    {
        "name": "Lobo de wall stree",
        "year": 2015,
        "director": "Martin"
    },
    {
        "name": "matrix",
        "year": 2015,
        "director": "Martin"
    },
    {
        "name": "volver al futuro",
        "year": 2015,
        "director": "Martin"
    },
    {
        "name": "volver al futuro 2",
        "year": 2015,
        "director": "Martin"
    }
]
```

### POST /peliculas
la peticion post recibe un json con la informacion del struct definido en `movie.go`
```
{
    "name" : "volver al futuro 2",
    "year" : 2015,
    "director" : "Martin"
}
```

## Como correr la app

 Una vez se tenga corriendo la base de datos mongo basada en la imagen definida en la carpeta `docker` solo necesitas ejecutar el siguiente comando parado en la raiz del proyecto. 

`go run .`

si se configuro un base de datos diferente a la establecida en la imagen es necesario modificar los datos de acceso definidos en `docker/.env`

```
DB_USER=gouser
DB_PASSWD=gopass
DB_HOST=localhost
DB_PORT=27017
EXPRESS_PORT=8081
DB_NAME=godb
NODE_DOCKER_PORT=8080
MONGOEXPRESS_LOGIN=dev
MONGOEXPRESS_PASSWORD=dev
```