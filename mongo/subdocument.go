package mongo

import (
	"github.com/ungerik/go-start/mgo/bson"
	"github.com/ungerik/go-start/model"
)

///////////////////////////////////////////////////////////////////////////////
// SubDocument

type SubDocument interface {
	Init(collection *Collection, selector string, embeddingStruct interface{})
	Collection() *Collection
	RootDocumentObjectId() bson.ObjectId
	RootDocumentSetObjectId(id bson.ObjectId)
	Iterator() model.Iterator
	Save() error
	RemoveRootDocument() error
}
