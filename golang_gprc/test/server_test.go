package server_test

import (
  "context"
  "testing"

  "github.com/stretchr/testify/assert" // Consider adding a testing framework like testify

  user "github.com/TheGreatShubham/golang_gprc/tree/main/golang_gprc/pb"
  "github.com/TheGreatShubham/golang_gprc/tree/main/golang_gprc/server"
)

func TestGetUser(t *testing.T) {
  // Replace with your user data access implementation
  repo := &server.InMemoryUserRepository{
    users: map[int]*user.User{
      1: &user.User{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
    },
  }
  s := server.NewServer(repo)

  req := &user.GetUserIDRequest{Id: 1} // Set the user ID to fetch

  // Call the server method under test
  resp, err := s.GetUser(context.Background(), req)

  // Assertions to verify the response and error handling
  if err != nil {
    t.Errorf("Unexpected error: %v", err)
  } else {
    assert.Equal(t, req.GetId(), resp.GetId())  // Check if fetched ID matches request
    assert.Equal(t, "Steve", resp.GetFname()) // Check for expected user data (modify as needed)
    // Add more assertions to verify other user data fields as needed
  }
}

