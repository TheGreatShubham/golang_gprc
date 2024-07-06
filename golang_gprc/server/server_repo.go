package server

import (
  "context"

  user "github.com/your-username/user-service/pb" // Replace with your package path
)

type UserRepository interface {
  GetUser(ctx context.Context, userID int) (*user.User, error)
  GetUsers(ctx context.Context, userIDs []int) ([]*user.User, error)
  SearchUsers(ctx context.Context, criteria string) ([]*user.User, error)
}

type InMemoryUserRepository struct {
  users map[int]*user.User
}

func (repo *InMemoryUserRepository) GetUser(ctx context.Context, userID int) (*user.User, error) {
  user, ok := repo.users[userID]
  if !ok {
    return nil, grpc.Errorf(codes.NotFound, "User with ID %d not found", userID)
  }
  return user, nil
}

func (repo *InMemoryUserRepository) GetUsers(ctx context.Context, userIDs []int) ([]*user.User, error) {
  var foundUsers []*user.User
  for _, userID := range userIDs {
    user, ok := repo.users[userID]
    if ok {
      foundUsers = append(foundUsers, user)
    }
  }
  return foundUsers, nil
}

func (repo *InMemoryUserRepository) SearchUsers(ctx context.Context, criteria string) ([]*user.User, error) {
  var matchingUsers []*user.User
  for _, user := range repo.users {
    // Implement search logic based on your criteria (e.g., city, phone number)
    if strings.Contains(user.GetCity(), criteria) || strings.Contains(fmt.Sprint(user.GetPhone()), criteria) {
      matchingUsers = append(matchingUsers, user)
    }
  }
  return matchingUsers, nil
}
