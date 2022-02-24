# configuracion_api

API CRUD para la gestión de aplicativos, procesos y perfiles de aplicaciones.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
CONFIGURACION_API_HTTP_PORT=[Descripción]
CONFIGURACION_API_DB_USER=[Descripción]
CONFIGURACION_API_DB_PASS=[Descripción]
CONFIGURACION_API_DB_URL=[Descripción]
CONFIGURACION_API_DB_NAME=[Descripción]
CONFIGURACION_API_DB_SCHEMA=[Descripción]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con CONFIGURACION*API*...

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/configuracion_api

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/configuracion_api

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CONFIGURACION_API__PORT=8080 CONFIGURACION_API__DB_HOST=127.0.0.1:27017 CONFIGURACION_API_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
docker build --tag=configuracion_api . --no-cache
docker run -p 80:80 configuracion_api
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/configuracion_api

#2. Moverse a la carpeta del repositorio
cd configuracion_api

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias

```shell
# Not Data
```

## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/configuracion_api/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/configuracion_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/configuracion_api/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/configuracion_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/configuracion_api/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/configuracion_api) |

## Modelo de Datos

[SVG](database/configuracion.svg) - [PGModeler](database/configuracion.dbm)

## Licencia

This file is part of configuracion_api.

configuracion_api is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

configuracion_api is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with configuracion_api. If not, see https://www.gnu.org/licenses/.
