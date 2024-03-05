//go:generate  protoc --go_out=. user.proto
//go:generate  protoc --go-grpc_out=. user.proto
//go:generate  protoc --go_out=. post.proto
//go:generate  protoc --go-grpc_out=. post.proto
//go:generate  protoc --go_out=. rating.proto
//go:generate  protoc --go-grpc_out=. rating.proto
//go:generate  protoc --go_out=. auth.proto
//go:generate  protoc --go-grpc_out=. auth.proto

package proto
