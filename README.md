# Pastebin API Project

One of the project of golang [list](https://github.com/diaszakir/Backend-Projects/)

## Description

Main idea of this project is creating messages and texts and sending to others by link to access.

You send request with JSON message and API generates link to access message.

```
{
    "content": "Hello"
}
```

After that generated code using [Base62-go](https://github.com/diaszakir/Base62-go)

Accessing -> GET link/:code -> content

## How to launch
Copy repository to your local machine using command:

```
git clone https://github.com/diaszakir/URL-Shortener-API-Go
```

In case if you do not have go.mod and go.sum

```
go get -u github.com/gin-gonic/gin
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/diaszakir/base62-go
```

After

```
go run cmd/app/main.go
```

Loading documentation in case if you changed structure

```
swag init -g cmd/app/main.go
```

To access swagger go to `localhost:8080/swagger/index.html`

## API Endpoints
- `GET /health` - checks API
- `GET /paste/:code` - Getting paste from code
- `GET /paste/raw/:code` - Getting string paste from code
- `POST /paste` - Creating paste URL
- `DELETE /paste/:code` - Deleting specific paste
