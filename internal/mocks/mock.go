//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=mod -destination=userService.go -package=mocks -mock_names=UserService=UserService restapi/internal/domain UserService
//go:generate go run github.com/golang/mock/mockgen -destination=postService.go -package=mocks -mock_names=PostService=PostService restapi/internal/domain PostService
//go:generate go run github.com/golang/mock/mockgen -destination=DBService.go -package=mocks -mock_names=DBService=DBService restapi/internal/pkg/database DBService
//go:generate go run github.com/golang/mock/mockgen -destination=RatingService.go -package=mocks -mock_names=RatingService=RatingService restapi/internal/domain RatingService
//go:generate go run github.com/golang/mock/mockgen -destination=VoteService.go -package=mocks -mock_names=VoteService=VoteService restapi/internal/domain VoteService

//go:generate go run github.com/golang/mock/mockgen -destination=ImplementedAuthService.go -package=mocks -mock_names=ImplementedAuthService=ImplementedAuthService restapi/internal/infrastructure/inputports/grpc/handler ImplementedAuthService
//go:generate go run github.com/golang/mock/mockgen -destination=ImplementedUserService.go -package=mocks -mock_names=ImplementedUserService=ImplementedUserService restapi/internal/infrastructure/inputports/grpc/handler ImplementedUserService
//go:generate go run github.com/golang/mock/mockgen -destination=ImplementedPostService.go -package=mocks -mock_names=ImplementedPostService=ImplementedPostService restapi/internal/infrastructure/inputports/grpc/handler ImplementedPostService
//go:generate go run github.com/golang/mock/mockgen -destination=ImplementedRatingService.go -package=mocks -mock_names=ImplementedRatingService=ImplementedRatingService restapi/internal/infrastructure/inputports/grpc/handler ImplementedRatingService

package mocks
