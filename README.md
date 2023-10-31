## gRPC User Service

- This is a sample gRPC User Service application written in Go. The service provides two endpoints to fetch user details:
  - Fetch single user details based on user id.
  - Fetch a list of user details based on a list of ids.

### Application Structure

```
/grpc-user-service      --> Root project directory
/cmd                     --> Contains application's entry points
    /grpc-server
        main.go          --> The server application's entry point
/api                     --> Protocol definition and generated source files
    /user
        user.proto       --> Protocol Buffers definition file
        user.pb.go       --> Generated from protobuf definitions
        user_grpc.pb.go  --> Generated from protobuf service definitions
/pkg                    --> Libraries and packages used by your applications which can also be used by other apps
    /user
        service.go       --> Contains actual implementation of user service
   	    repository.go    --> Repository implementing underlying data operations
        service_test.go  --> Test cases for service

Dockerfile              --> Dockerfile to containerize the application
README.md               --> Project description and instructions (You are here!)
go.mod, go.sum          --> Go modules file (describe dependencies)
```

### Design pattern used

The *Repository Pattern* is used here. The primary motivation to use this pattern here is for decoupling and abstraction of data access logic. Currently, we are maintaining all the users in an in-memory list. If in the future, you wish to switch to a database or a different data source for any reason, using this pattern will ensure that you only have to make the change in one place, i.e., the 'UserRepository'. None of your service logic would need to change since that would just call the methods provided by 'UserRepository'. This provides the flexibility to switch between data sources without affecting the business logic of your service. It's an extra layer of abstraction initializing your data—be it a slice, a database or something else doesn't matter—for your service code.

### Run Locally

Clone the project

```bash
  git clone https://github.com/Soumik43/grpc-user-service.git
```

Go to the project directory

```bash
  cd grpc-user-service
```

Install dependencies

```bash
  go mod download
```

Start the server

```bash
  go run cmd/grpc-server/main.go
```

### Run with Docker

```bash
  docker build -t grpc-user-service . && docker run -d -p 50051:50051 --env-file ./env.list --name grpc-user-service -dit grpc-user-service
```

### Calling the gRPC Request using grpcurl

To call the gRPC request using grpcurl, follow these steps:

- Install grpcurl by following the instructions in the [grpcurl GitHub repository](https://github.com/fullstorydev/grpcurl#installation).
- Start the gRPC server.
- In a separate terminal window, run the command `grpcurl -plaintext -d '{ "id": 11 }' localhost:50051 user.UserService/GetUser` (For single user).
- You should see the response from the server: `{"id":1,"fname":"Steve","city":"LA","phone":"1234567890","height":5.8,"married":true}`.
- Run the command `grpcurl -plaintext -d '{ "ids": [10,2,3] }' localhost:50051 user.UserService/GetUsers` (For multiple users).
- You should see the response from the server: `{"users":[{"id":10,"fname":"Laura","city":"Austin","phone":"4928394032","height":5.2},{"id":2,"fname":"Shimmy","city":"NYC","phone":"9876543210","height":5.2},{"id":3,"fname":"Jane","city":"HongKong","phone":"1234509876","height":5.1,"married":true}]}`.

### Running Tests

To run tests, run the following command

```bash
  go test ./...
```
