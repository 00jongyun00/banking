# Go micro web service And Auth service

## Database set up

```
docker-compose up --build -d
```

-   use mysql
-   default id: root
-   default password: password

## Start to project

#### Main server

-   go run main.go
    http://localhost:8080

#### Auth server

-   cd ./banking-auth && sh start.sh
    http://localhost:8181
