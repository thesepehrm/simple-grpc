# Simple GRPC
A simple GRPC server to learn GRPC on iOS and Android.
## Service Features

- Login
- Logout
- Server promotions stream

## How To Run
Run the server using `make run-server`.

### Test Client
Also there's a basic client for testing purposes. Run these commands to run the client:

- Login: `go run cmd/grpc-client/main.go login`
- Logout `go run cmd/grpc-client/main.go logout`
- Connect to Live Updates stream: `go run cmd/grpc-client/main.go update`

## Protocol File

```protobuf
// Update Service
service Update {
  // Login method
  rpc Login(LoginRequest) returns (LoginResponse) {}
  // Logout method
  rpc Logout(LogoutRequest) returns (google.protobuf.Empty) {}
  // Server promotions stream
  rpc ServerPromotions(google.protobuf.Empty)
      returns (stream UpdateStreamResponse) {}
}

message LoginRequest {
  string username = 1;
  string password = 2;
}

message LoginResponse { string token = 1; }

message LogoutRequest { string token = 1; }

message ChatStreamRequest { string message = 1; }

message UpdateStreamResponse {
  google.protobuf.Timestamp timestamp = 1;
  Update update = 2;

  message Update {
    string type = 1;
    string text = 2;
  }
}

```
