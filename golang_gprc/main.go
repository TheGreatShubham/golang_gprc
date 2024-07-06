package main

import (
  "log"

  user "github.com/TheGreatShubham/golang_gprc/tree/main/golang_gprc/pb"
  "github.com/TheGreatShubham/golang_gprc/tree/main/golang_gprc/server"
)

func main() {
  repo := &server.InMemoryUserRepository{
    users: map[int]*user.User{
      // Add your sample user data here
      1: &user.User{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
      // ... add more users
    },
  }

  err := server.RunServer(50051, repo)
  if err != nil {
    log.Fatal(err)
  }
}
