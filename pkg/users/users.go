package users

import (
	"fmt"
	"lesson4/pkg/documentstore"
	"lesson4/pkg/err"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	cfg  documentstore.CollectionConfig
}

type Service struct {
	coll  documentstore.Store
	users map[string]User
}

func NewService() *Service {
	return &Service{
		coll:  *documentstore.NewStore(),
		users: make(map[string]User),
	}
}

func (s *Service) CreateUser(id, name string, cfg documentstore.CollectionConfig, doc *documentstore.Document) (*User, error) {
	user := User{ID: id, Name: name, cfg: cfg}

	if _, exists := s.users[user.ID]; exists {
		return nil, err.ErrCreatedUser
	}
	s.users[user.ID] = user
	s.coll.CreateCollection(user.ID, &user.cfg)

	getCol, _ := s.coll.GetCollection(id)
	getCol.Put(*doc)

	fmt.Println(getCol.List())
	return &user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	sList := make([]User, 0, len(s.users))
	for _, v := range s.users {
		sList = append(sList, v)
	}
	if len(sList) > 0 {
		return sList, nil
	}
	return nil, err.ErrListEmpty
}

func (s *Service) GetUser(userID string) (*User, error) {
	if kUser, ok := s.users[userID]; ok {
		return &kUser, nil
	}
	return nil, err.ErrNotFound
}

func (s *Service) DeleteUser(userID string) error {
	if _, ok := s.users[userID]; ok {
		delete(s.users, userID)
		return nil
	}
	return err.ErrNotFound
}
