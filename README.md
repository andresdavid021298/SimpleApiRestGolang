# Proyecto ApiRest Golang

Api Rest sencilla usando Golang con la libreria Echo v4, usando como cache dos listas, una de carros(Car) y otra de motos(Motorcycle)

## Librerias

 - **github.com/labstack/echo/v4**: Facilita la creacion de servicios REST
 - **github.com/joho/godotenv**: Permite la carga de variables de entorno desde un archivo .env
 - **github.com/go-playground/validator**: Validador de paquetes implementa validaciones de valores para estructuras y campos individuales basados ​​en etiquetas
 - **github.com/google/uuid**: Genera e inspecciona UUID según RFC 9562 (Para los generar ids unicos)

## API Rest
Para su consumo, la app se expone en el puerto 8090 (Se puede cambiar en el archivo .env)
La API creada tiene los siguientes endpoints
- Car ["/car"]
  - ["" - **GET**] -> Retorna un texto como saludo
  - "[Path: "/all" - **GET**] -> Retorna todos los objetos Car almacenados en la cache
  - [Path: "/byBrand/:brand" - **GET**] -> Retorna los objetos Car que coincidan en su campo *brand* con el valor enviado en el Param
  - [Path: "/used" - **GET**] -> Retorna los objeto Car que tenga el campo *isUsed* como true
  - [Path: "" - **POST**] -> Inserta un nuevo objeto Car en la cache, debe enviarle el correspondiente request
  - [Path: "/:id" - **DELETE**] -> Eliminar un elemento de la cache segun el id enviado en el Param
- Motorcycle -> "/motorcycle"
  - ["" - **GET**] -> Retorna un texto como saludo
  - "[Path: "/all" - **GET**] -> Retorna todos los objetos Motorcycle almacenados en la cache
  - [Path: "/byBrand/:brand" - **GET**] -> Retorna los objetos Motorcycle que coincidan en su campo *brand* con el valor enviado en el Param
  - [Path: "/:id" - **GET**] -> Retorna el objeto Motorcycle que corresponda con el *id* enviado como Param
  - [Path: "/changeAmount" - **PATCH**] -> Actualiza el campo *amount* por el valor enviado en el QueryParam, segun el *id* enviado tambien en el QueryParam
  - [Path: "/:id" - **DELETE**] -> Eliminar un elemento de la cache segun el id enviado en el Param
