package utils

import (
	"github.com/google/uuid"
	"strings"
)

var DocId = NewDocumentId()

func NewDocumentId() *docId {
	return &docId{}
}

type docId struct{}

// get new id
func (f *docId) Create() string {
	uuid2 := uuid.New()
	uuidWithoutDash := strings.Replace(uuid2.String(), "-", "", -1)
	return 	uuidWithoutDash
}

