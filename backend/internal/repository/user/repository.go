package user

import (
  "gorm.io/gorm"
)

type UserRepository interface {
  GetUser(id string) (*User, error)
  CreateUser(id string, name string, password string, email string, hash string) (*User, error)
  UpdateUser(id string, name string, email string) (*User, error)
}

type userRepository struct {
  db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
  return &userRepository{
    db: db,
  }
}

func (u *userRepository) GetUser(id string) (*User, error) {
  user := &User{}
  user.ID = id
  if err := u.db.First(user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

func (u *userRepository) CreateUser(id string, name string, password string, email string, hash string) (*User, error) {
  user := &User{ID: id, Name: name, Password: password, Email: email, Hash: hash}
  if err := u.db.Create(user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

func (u *userRepository) UpdateUser(id string, name string, email string) (*User, error) {
  user := &User{ID: id}
  if err := u.db.First(user).Error; err != nil {
    return nil, err
  }
  user.Name = name
  user.Email = email
  if err := u.db.Save(user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

