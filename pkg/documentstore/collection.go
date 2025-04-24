package documentstore

import (
	"fmt"
	"lesson4/pkg/err"
	"log/slog"
)

type Collection struct {
	documents map[string]Document
	config    CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) error {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`

	keyFilds, ok := doc.Fields[s.config.PrimaryKey]
	if !ok {
		fmt.Println("Error: Document must contain a key field")
		return err.ErrUnsupportedDocumentField
	}

	if keyFilds.Type != DocumentFieldTypeString {
		fmt.Println("Error: Key field must be of type string")
		return err.ErrUnsupportedDocumentField
	}
	keyValue, ok := keyFilds.Value.(string)
	if !ok {
		slog.Error("Error: Key field value is not a string")
		return err.ErrUnsupportedDocumentField
	}

	if s.documents == nil {
		s.documents = map[string]Document{}
	}
	s.documents[keyValue] = doc
	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, exists := s.documents[key]; exists {
		return &doc, nil
	}
	return nil, err.ErrDocumentNotFound
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.documents[key]; exists {
		delete(s.documents, key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.documents))
	for _, v := range s.documents {
		sList = append(sList, v)
	}
	return sList
}
