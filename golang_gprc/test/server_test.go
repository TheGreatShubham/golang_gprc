package server_test

import (
  "context"
  "testing"

  "github.com/stretchr/testify/assert" // Consider adding a testing framework like testify

  user "github.com/your-username/user-service/pb" // Replace with your package path
  "github.com/your-username/user-service/server" // Replace with your package path
)

func TestGetUser(t *testing.T) {
  // Replace with your user data access implementation
  repo := &server.InMemoryUserRepository{
    users: map[int]*user.User{
      1: &user.User{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
    },
  }
  s := server.NewServer(repo)

  req := &user.GetUserIDRequest{Id:
