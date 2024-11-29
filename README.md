# Uala-Challenge

#instrucciones

```
clonar el env
git clone https://github.com/david34corbalan/aula-challenge.git
cd uala-challenge
docker compose up -d
```

# tests

```
./test.sh
para la capa del repositorio se podria utilizar testcontainers, para levantar una base de datos en memoria (por falta de tiempo nose llego a implementar)
```

# db

```
tweets.db, cuenta con 3 tablas:
- tweets
- users
- follow

En caso productivo se deberia utilizar una base de datos relacional, como postgresql o dependiendo del servicio cloud queda a disposicion
de la empresa (aws - rds)
```

# arquitectura

```
_ Se utilizo el Framework gin (golang).
- Se utilizo una arquitectura hexagonal, con la finalidad de separar la logica de negocio de la logica de infraestructura.
- Se utilizo el patron de diseño repository, para separar la logica de negocio de la logica de persistencia.
- Se utilizo el patron de diseño service, para separar la logica de negocio de la logica de infraestructura.
-
```

# Kafka

```
se quizo implementar kafka, pero los contenedores no se podian comunicar entre si, pero la idea era para el uso de registro de usuario y para el futuro caso que se siga a algun usuario se podria consumir un servicio externo de notificacion
```

# Comentarios

```
en una terminal ingresar los comando para generar los mocks:
mockgen -destination=mocks/mock_users_port.go -package=mocks --source=pkg/users/ports/users.go
mockgen -destination=mocks/mock_tweets_port.go -package=mocks --source=pkg/tweets/ports/tweets.go
mockgen -destination=mocks/mock_follow_port.go -package=mocks --source=pkg/follow/ports/follow.go
```
