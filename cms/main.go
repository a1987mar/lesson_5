package main

import (
	"fmt"
	"lesson4/pkg/documentstore"
	"lesson4/pkg/users"
)

func main() {

	//marshalExample()
	unmarshalExample()
	//lesson5()
}

func marshalExample() {
	s := &documentstore.MyStruct{X: 15}
	doc, err := documentstore.MarshalDocument(s)
	if err != nil {
		fmt.Printf("failed to marshal document: %+v\n", err)
		return
	}
	fmt.Printf("marshaled document: %+v\n", doc)
}

func unmarshalExample() {
	doc := &documentstore.Document{Fields: map[string]documentstore.DocumentField{}}
	doc.Fields["X"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeNumber,
		Value: int(32),
	}

	s := &documentstore.MyStruct{}
	err := documentstore.UnmarshalDocument(doc, s)
	if err != nil {
		fmt.Printf("failed to unmarshal document: %+v\n", err)
		return
	}
	fmt.Printf("unmarshaled document: %+v\n", s)
}

func lesson5() {
	s := users.NewService()
	d1 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d1.Fields["id-1"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "setup.exe",
	}

	cfg1 := documentstore.CollectionConfig{PrimaryKey: "id-1"}
	_, err := s.CreateUser("id-1", "UserTest-1", cfg1, &d1)
	if err != nil {
		fmt.Println(err)
	}

	d2 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d2.Fields["id-2"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "main.go",
	}
	cfg2 := documentstore.CollectionConfig{PrimaryKey: "id-2"}
	_, err = s.CreateUser("id-2", "UserTest-2", cfg2, &d2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.ListUsers())
}
