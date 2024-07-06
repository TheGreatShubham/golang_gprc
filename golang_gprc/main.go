package main

import (
  "log"

  user "Documents\\golang_grpc\\pb" // Replace with your package path
  "Documents\\golang_grpc\\server" // Replace with your package path
)

func main() {
  // Replace with your actual user data access implementation
  repo := &server.InMemoryUserRepository{
    users: map[int]*user.User{
      // Add your sample user data here
      1: &user.User{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
      // ... add more users
    },
  }

  err := server.RunServer(50051, repo) // Replace 50051 with desired port
  if err != nil {
    log.Fatal(err)
  }
}
