package documentstore

import (
	"lesson4/pkg/err"
	"log/slog"
)

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name, id string) (error, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil
	if _, exists := s.collections[name]; exists {
		return err.ErrCollectionAlreadyExists, nil
	}
	coll := &Collection{config: CollectionConfig{
		PrimaryKey: id,
	}}
	s.collections[name] = coll

	return nil, coll
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	if colect, ok := s.collections[name]; ok {
		return colect, nil
	}
	slog.Info("")

	return nil, err.ErrCollectionNotFound
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.collections[name]; ok {
		delete(s.collections, name)
		return true
	}
	return false
}
