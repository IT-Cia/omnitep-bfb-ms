# Golang Microservice with Swagger and gRPC

Welcome to the Golang Microservice project utilizing Swagger for API documentation and gRPC for communication, with Gin as the web server. This project aims to simplify the setup process for integrating Swagger documentation into a microservice architecture in Golang. It includes a Task CRUD API as a demonstration of a minimal microservice system.

## Features

- **Swagger Integration**: Easily document your APIs with Swagger UI.
- **gRPC Communication**: Utilize efficient communication between client and server using gRPC protocol.
- **Gin Web Server**: Handle HTTP requests seamlessly with Gin framework.
- **Task CRUD API**: Implement basic CRUD operations for tasks.

## Project Structure

```
.
├── client                  # Client implementation with Gin
│   ├── main.go             # Entry point for the client
│   └── ...
├── protobuf/taskpb         # Protobuf definitions
│   ├── taskpb.proto        # Protobuf file for Task service
│   └── ...
├── server                  # Server implementation with gRPC
│   ├── main.go             # Entry point for the server
│   └── ...
└── ...
```

## Setup

1. **Clone the Repository**: 
   ```
   git clone https://github.com/aliashahi/go-swagger-microservice.git
   ```

2. **Install Dependencies**:
   - Make sure you have Go installed. If not, you can download it [here](https://golang.org/dl/).
   - Install necessary Go packages:
     ```
     go mod tidy
     ```

3. **Run the Server**:
   ```
   go run server/main.go
   ```

4. **Run the Client**:
   ```
   go run client/main.go
   ```
   or run with ```vscode run and debug```

5. **Access Swagger UI**:
   Open your web browser and navigate to `http://localhost:8002/api/swagger/index.html` to view Swagger documentation.

## Related Links

- [Swagger Documentation](https://swagger.io/)
- [gRPC Documentation](https://grpc.io/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

- [Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)


## Protobuf Setup

To install/update proto compiler we need to install/update these binaries :
- protoc
- protoc-gen-go-grpc
- protoc-gen-go

 thereby these steps are required:

1.remove already installed binaries :
 - find them with ```whereis <binary>``` and remove them with ```sudo rm -rf <binary>``` 

2.install protoc with these commands:

```shell
PROTOC_ZIP=protoc-25.2-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v25.2/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr 'include/*'
rm -f $PROTOC_ZIP
```

3. install `protoc-gen-go` and `protoc-gen-go-grpc`:
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### references:

[installing protoc](https://google.github.io/proto-lens/installing-protoc.html)

[gRPC Quick start](https://grpc.io/docs/languages/go/quickstart/)