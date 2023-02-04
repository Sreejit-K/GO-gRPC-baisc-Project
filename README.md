GO-gRPC-baisc-Project

// command to create proto buffs 
    protoc --go_out=. --go-grpc_out=. proto/*.proto

// Tried using the makefile an use it to run the    command but was getting a few issues have to explore on it 

// The error shown is --
 make: *** No rule to make target 'proto_build'.  Stop.


// go mod tidy is kinda similat to npm install but it deletes the uneccessaty lib which are not used from the go.mod file 