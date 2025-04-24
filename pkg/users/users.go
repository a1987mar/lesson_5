package users

import (
	"lesson4/pkg/documentstore"
	"lesson4/pkg/err"
	"log/slog"
)

const (
	UserCollectionName = "name"
	UserCollectionKey  = "id"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *documentstore.Collection
}

func NewService(s *documentstore.Store) Service {
	s.CreateCollection(UserCollectionName, UserCollectionKey)
	collect, _ := s.GetCollection(UserCollectionName)
	return Service{
		coll: collect,
	}
}

func (s *Service) CreateUser(id, name string, doc *documentstore.Document) (*User, error) {
	if _, er := s.coll.Get(id); er == nil {
		return nil, err.ErrCreatedUser
	}
	if er := s.coll.Put(*doc); er != nil {
		slog.Error(err.ErrCreatedUser.Error())
		return nil, err.ErrCreatedUser
	}

	u := User{
		Id:   id,
		Name: name,
	}
	slog.Info("add user" + u.Id)
	return &u, nil
}

func (s *Service) ListUsers() ([]User, error) {
	tList := s.coll.List()
	if len(tList) > 0 {
		ulist := make([]User, 0, len(tList))
		for _, v := range tList {
			u := User{}
			er := documentstore.UnmarshalDocument(&v, &u)
			if er != nil {
				slog.Error(er.Error())
			}
			ulist = append(ulist, u)
		}
		return ulist, nil
	}
	return nil, err.ErrListEmpty
}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, er := s.coll.Get(userID)
	if er != nil {
		return nil, er
	}
	u := User{}
	er = documentstore.UnmarshalDocument(doc, &u)
	if er != nil {
		return nil, err.ErrCollectionAlreadyExists
	}
	return &u, nil

}

func (s *Service) DeleteUser(userID string) error {
	if ex := s.coll.Delete(userID); ex {
		slog.Info("delete user" + userID)
		return nil
	}
	return err.ErrNotFound
}
