package server

import (
  "context"
  "fmt"
  "log"
  "net" // Make sure this line is present
  "google.golang.org/grpc" // Make sure this line is present

  user "github.com/TheGreatShubham/golang_gprc/tree/main/golang_gprc/pb"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type server struct {
  repo user.UserRepository
}

func NewServer(repo user.UserRepository) *server {
  return &server{repo: repo}
}

func (s *server) GetUser(ctx context.Context, req *user.GetUserIDRequest) (*user.User, error) {
  if req.GetId() <= 0 {
    return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID")
  }
  user, err := s.repo.GetUser(ctx, req.GetId())
  if err != nil {
    return nil, err
  }
  return user, nil
}

func (s *server) GetUsers(ctx context.Context, req *user.GetUserIDsRequest, stream user.UserService_GetUsersServer) error {
  for _, id := range req.GetIds() {
    if id <= 0 {
      return status.Errorf(codes.InvalidArgument, "Invalid user ID: %d", id)
    }
    user, err := s.repo.GetUser(ctx, id)
    if err != nil {
      return err
    }
    if err := stream.Send(user); err != nil {
      return err
    }
  }
  return nil
}

func (s *server) SearchUsers(ctx context.Context, req *user.SearchRequest, stream user.UserService_SearchUsersServer) error {
  criteria := req.GetCriteria()
  if criteria == "" {
    return nil, status.Errorf(codes.InvalidArgument, "Empty search criteria")
  }
  stream, err := stream.New(&stream.ServerStream{Decider: stream.Strip})
  if err != nil {
    return nil, err
  }
  matchingUsers, err := s.repo.SearchUsers(ctx, criteria)
  if err != nil {
    return nil, err
  }
  for _, user := range matchingUsers {
    if err := stream.SendMsg(user); err != nil {
      return nil, err
    }
  }
  return stream, nil
}

func RunServer(port int, repo user.UserRepository) error {
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  grpcServer := grpc.NewServer()
  user.RegisterUserServiceServer(grpcServer, NewServer(repo))
  log.Printf("server listening at %v", lis.Addr())
  return grpcServer.Serve(lis)
}
